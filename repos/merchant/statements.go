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
				CAST(Point(:longitude, :latitude) AS geometry)
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
				merchant_id,
				name,
				category,
				image_url,
				ST_X(location::geometry) AS longitude,
				ST_Y(location::geometry) AS latitude,
				created_at 
			FROM
				merchants
			WHERE
				merchant_id = $1
		`),
		listByIds: statementutil.MustPrepare(`
			SELECT
				merchant_id,
				name,
				category,
				image_url,
				ST_X(location::geometry) AS longitude,
				ST_Y(location::geometry) AS latitude,
				created_at 
			FROM
				merchants
			WHERE
				merchant_id = ANY ($1)
		`),
		listItemsByIds: statementutil.MustPrepare(`
			SELECT
				* 
			FROM
				merchant_items
			WHERE
				merchant_item_id = ANY ($1)
		`),
	}
}
