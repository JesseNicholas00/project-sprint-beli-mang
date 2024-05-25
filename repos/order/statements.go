package order

import (
	"github.com/JesseNicholas00/BeliMang/utils/statementutil"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	createEstimate                *sqlx.NamedStmt
	findEstimateById              *sqlx.Stmt
	findEstimateItemsByEstimateId *sqlx.Stmt
}

func prepareStatements() statements {
	return statements{
		createEstimate: statementutil.MustPrepareNamed(`
			INSERT INTO estimates(
			    estimate_id,
				user_id
			) VALUES (
				:estimate_id,
				:user_id
			)
		`),

		findEstimateById: statementutil.MustPrepare(`
			SELECT
				*
			FROM
				estimates
			WHERE
				estimate_id = $1
		`),

		findEstimateItemsByEstimateId: statementutil.MustPrepare(`
			SELECT
				merchant_id,
				merchant_item_id,
				quantity
			FROM
				estimates JOIN estimate_items ON
					estimates.estimate_id = estimate_items.estimate_id
			WHERE
				estimates.estimate_id = $1
		`),
	}
}
