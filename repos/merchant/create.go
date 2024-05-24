package merchant

import (
	"context"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
)

func (repo *merchantRepoImpl) CreateMerchant(
	ctx context.Context,
	m Merchant,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := repo.dbRizz.GetOrNoTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	_, err = sess.NamedStmt(ctx, repo.statements.create).Exec(m)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return nil
}
