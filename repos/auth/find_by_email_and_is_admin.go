package auth

import (
	"context"

	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
)

func (repo *authRepositoryImpl) FindUserByEmailAndIsAdmin(
	ctx context.Context,
	email string,
	isAdmin bool,
) (res User, err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	rows, err := sess.
		Stmt(ctx, repo.statements.findFirstByEmailAndIsAdmin).
		QueryxContext(ctx, email, isAdmin)

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
