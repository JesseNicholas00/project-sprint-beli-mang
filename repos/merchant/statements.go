package merchant

import (
	"github.com/JesseNicholas00/BeliMang/utils/statementutil"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	create         *sqlx.NamedStmt
	createItem     *sqlx.NamedStmt
	findById       *sqlx.Stmt
	listByIds      *sqlx.Stmt
	listItemsByIds *sqlx.Stmt
}

func prepareStatements() statements {
	return statements{
		create: statementutil.MustPrepareNamed(`
			INSERT INTO merchants(
				merchant_id,
				name,
				category,
				image_url,
				location
			) VALUES (
				:merchant_id,
				:name,
				:category,
				:image_url,
				CAST(Point(:latitude, :longitude) AS geometry)
			)
		`),
		createItem: statementutil.MustPrepareNamed(`
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
		findById: statementutil.MustPrepare(`
			SELECT
				* 
			FROM
				merchants
			WHERE
				merchant_id = $1
		`),
		listByIds: statementutil.MustPrepare(`
			SELECT
				* 
			FROM
				merchants
			WHERE
				merchant_id IN $1
		`),
		listItemsByIds: statementutil.MustPrepare(`
			SELECT
				* 
			FROM
				merchant_items
			WHERE
				merchant_item_id IN $1
		`),
	}
}
