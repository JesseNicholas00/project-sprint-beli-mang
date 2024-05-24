package merchant

import (
	"context"

	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
)

func (repo *merchantRepoImpl) FindMerchantById(
	ctx context.Context,
	merchant_id string,
) (res Merchant, err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	ctx, sess, err := repo.dbRizz.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	rows, err := sess.
		Stmt(ctx, repo.statements.findById).
		QueryxContext(ctx, merchant_id)

	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(&res)
		if err != nil {
			err = errorutil.AddCurrentContext(err)
			return
		}
	}

	if res.Id == "" {
		err = ErrMerchantNotFound
		return
	}

	return
}
