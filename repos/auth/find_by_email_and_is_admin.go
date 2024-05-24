package auth

import (
	"context"

	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/jmoiron/sqlx"
)

func (repo *authRepositoryImpl) FindUserByEmailAndIsAdmin(
	ctx context.Context,
	email string,
	isAdmin bool,
) (res User, err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	query := `
		SELECT
			*
		FROM
			users
		WHERE
			email= :email AND
			is_admin= :isAdmin
		LIMIT 1
	`
	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	rows, err := sqlx.NamedQueryContext(
		ctx,
		sess.Ext,
		query,
		map[string]interface{}{
			"email":    email,
			"is_admin": isAdmin,
		},
	)

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

	if res.Id == "" {
		err = ErrEmailAndIsAdminNotFound
		return
	}

	return
}
