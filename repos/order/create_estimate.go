package order

import (
	"context"
	"github.com/JesseNicholas00/BeliMang/utils/transaction"
	"github.com/jmoiron/sqlx"

	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
)

type dbEstimateItem struct {
	EstimateItem
	Id      int64  `db:"estimate_item_id"`
	OrderId string `db:"estimate_id"`
}

func (repo *orderRepositoryImpl) CreateEstimate(
	ctx context.Context,
	est Estimate,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	var estItems []dbEstimateItem
	for _, item := range est.Items {
		estItems = append(estItems, dbEstimateItem{
			EstimateItem: item,
			OrderId:      est.Id,
		})
	}

	ctx, sess, err := repo.dbRizzer.GetOrAppendTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return transaction.RunWithAutoCommit(&sess, func() error {
		_, err := sess.
			NamedStmt(ctx, repo.statements.createEstimate).
			Exec(est)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		_, err = sqlx.NamedExecContext(
			ctx,
			sess.Ext,
			`INSERT INTO estimate_items(
				estimate_id,
				merchant_id,
				merchant_item_id,
				quantity
			) VALUES (
				:estimate_id,
				:merchant_id,
				:merchant_item_id,
				:quantity
			)`,
			estItems,
		)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		return nil
	})
}
