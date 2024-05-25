package order

import (
	"github.com/JesseNicholas00/BeliMang/utils/statementutil"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	createEstimate      *sqlx.NamedStmt
	createEstimateItems *sqlx.NamedStmt
}

func prepareStatements() statements {
	return statements{
		createEstimate: statementutil.MustPrepareNamed(`
			INSERT INTO estimates(
			    estimate_id,
				user_id
			) VALUES (
				:order_id,
				:user_id
			)
		`),

		createEstimateItems: statementutil.MustPrepareNamed(`
			INSERT INTO estimate_items(
				estimate_id,
				merchant_id,
				merchant_item_id,
				quantity
			) VALUES (
				:estimate_id,
				:merchant_id,
				:merchant_item_id,
				:quantity
			)
		`),
	}
}
