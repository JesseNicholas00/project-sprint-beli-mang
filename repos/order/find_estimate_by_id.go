package order

import (
	"context"
	"database/sql"
	"errors"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/transaction"
)

type dbEstimateItemRet struct {
	MerchantId     string `db:"merchant_id"`
	MerchantItemId string `db:"merchant_item_id"`
	Quantity       int    `db:"quantity"`
}

func (repo *orderRepositoryImpl) FindEstimateById(
	ctx context.Context,
	id string,
) (res Estimate, err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	ctx, sess, err := repo.dbRizzer.GetOrAppendTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	err = transaction.RunWithAutoCommit(&sess, func() error {
		err := sess.
			Stmt(ctx, repo.statements.findEstimateById).
			QueryRowxContext(ctx, id).
			StructScan(&res)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return ErrEstimateNotFound
			}
			return errorutil.AddCurrentContext(err)
		}

		rows, err := sess.
			Stmt(ctx, repo.statements.findEstimateById).
			QueryxContext(ctx, id)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}
		defer rows.Close()

		for rows.Next() {
			var cur dbEstimateItemRet
			if err := rows.StructScan(&cur); err != nil {
				return errorutil.AddCurrentContext(err)
			}

			res.Items = append(res.Items, EstimateItem{
				MerchantId: cur.MerchantId,
				ItemId:     cur.MerchantItemId,
				Quantity:   cur.Quantity,
			})
		}

		return nil
	})

	return
}
