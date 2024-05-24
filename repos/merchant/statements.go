package merchant

import (
	"github.com/JesseNicholas00/BeliMang/utils/statementutil"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	create   *sqlx.NamedStmt
	findById *sqlx.Stmt
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
		findById: statementutil.MustPrepare(`
			SELECT
				* 
			FROM
				merchants
			WHERE
				merchant_id LIKE $1
		`),
	}
}
