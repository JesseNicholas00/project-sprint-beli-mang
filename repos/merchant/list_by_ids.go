package merchant

import (
	"context"
	"github.com/lib/pq"

	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
)

func (repo *merchantRepoImpl) ListMerchantsByIds(
	ctx context.Context,
	merchantIds []string,
) (res []Merchant, err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	ctx, sess, err := repo.dbRizz.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	rows, err := sess.
		Stmt(ctx, repo.statements.listByIds).
		QueryxContext(ctx, pq.StringArray(merchantIds))

	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var merchant Merchant

		err = rows.StructScan(&merchant)
		if err != nil {
			err = errorutil.AddCurrentContext(err)
			return
		}

		res = append(res, merchant)
	}

	return
}
