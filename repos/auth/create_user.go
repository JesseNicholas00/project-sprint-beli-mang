package auth

import (
	"context"

	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/jmoiron/sqlx"
)

func (repo *authRepositoryImpl) CreateUser(
	ctx context.Context,
	user User,
) (res User, err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	query := `
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
	`
	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	rows, err := sqlx.NamedQueryContext(ctx, sess.Ext, query, user)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(&res)
		if err != nil {
			err = errorutil.AddCurrentContext(err)
			return
		}
	}

	return
}
