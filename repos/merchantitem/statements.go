package merchantitem

import (
	"github.com/JesseNicholas00/BeliMang/utils/statementutil"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	create *sqlx.NamedStmt
}

func prepareStatements() statements {
	return statements{
		create: statementutil.MustPrepareNamed(`
			INSERT INTO merchant_items(
				merchant_item_id,
				merchant_id,
				name,
				category,
				price,
				image_url
			) VALUES (
				:merchant_item_id,
				:merchant_id,
				:name,
				:category,
				:price,
				:image_url
			)
		`),
	}
}
