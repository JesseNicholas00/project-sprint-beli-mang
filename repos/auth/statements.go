package auth

import (
	"github.com/JesseNicholas00/BeliMang/utils/statementutil"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	create                     *sqlx.NamedStmt
	findByUsername             *sqlx.Stmt
	findFirstByEmailAndIsAdmin *sqlx.Stmt
}

func prepareStatements() statements {
	return statements{
		create: statementutil.MustPrepareNamed(`
			INSERT INTO users(
				user_id,
				username,
				email,
				password,
				is_admin	
			) VALUES (
				:user_id,
				:username,
				:email,
				:password,
				:is_admin
			) RETURNING
				user_id,
				username,
				email,
				password,
				is_admin
		`),
		findByUsername: statementutil.MustPrepare(`
			SELECT
				*
			FROM
				users
			WHERE
				username LIKE $1
		`),
		findFirstByEmailAndIsAdmin: statementutil.MustPrepare(`
			SELECT
				*
			FROM
				users
			WHERE
				email = $1 AND
				is_admin= $2
			LIMIT 1
		`),
	}
}
