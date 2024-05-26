package merchant

import (
	"context"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/lib/pq"
)

func (repo *merchantRepoImpl) ListItemsByIds(
	ctx context.Context,
	ids []string,
) (res []MerchantItem, err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	ctx, sess, err := repo.dbRizz.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	rows, err := sess.
		Stmt(ctx, repo.statements.listItemsByIds).
		QueryxContext(ctx, pq.StringArray(ids))

	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var item MerchantItem

		err = rows.StructScan(&item)
		if err != nil {
			err = errorutil.AddCurrentContext(err)
			return
		}

		res = append(res, item)
	}

	return
}
