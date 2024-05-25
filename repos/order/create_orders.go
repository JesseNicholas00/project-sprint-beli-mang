package order

import (
	"context"
	"github.com/JesseNicholas00/BeliMang/utils/transaction"

	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
)

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

		_, err = sess.
			NamedStmt(ctx, repo.statements.createEstimateItems).
			Exec(estItems)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}

		return nil
	})
}
