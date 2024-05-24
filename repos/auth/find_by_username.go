package auth

import (
	"context"

	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
)

func (repo *authRepositoryImpl) FindUserByUsername(
	ctx context.Context,
	username string,
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
		Stmt(ctx, repo.statements.findByUsername).
		QueryxContext(ctx, username)

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
		err = ErrUsernameNotFound
		return
	}

	return
}
