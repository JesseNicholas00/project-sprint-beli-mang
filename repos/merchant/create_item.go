package merchant

import (
	"context"

	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
)

func (repo *merchantRepoImpl) CreateMerchantItem(
	ctx context.Context,
	mi MerchantItem,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := repo.dbRizz.GetOrNoTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	_, err = sess.NamedStmt(ctx, repo.statements.createItem).Exec(mi)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return nil
}
