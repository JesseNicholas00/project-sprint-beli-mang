package order

import (
	"context"

	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/transaction"
)

// CreateOrder implements OrderRepository.
func (repo *orderRepositoryImpl) CreateOrder(
	ctx context.Context,
	order Order,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, sess, err := repo.dbRizzer.GetOrAppendTx(ctx)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	return transaction.RunWithAutoCommit(&sess, func() error {
		_, err := sess.
			NamedStmt(ctx, repo.statements.createOrder).
			Exec(order)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		return nil
	})
}
