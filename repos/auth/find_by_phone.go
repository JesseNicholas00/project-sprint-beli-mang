package auth

import (
	"context"

	"github.com/KerakTelor86/GoBoiler/utils/errorutil"
	"github.com/jmoiron/sqlx"
)

func (repo *authRepostioryImpl) FindStaffByPhone(
	ctx context.Context,
	phone string,
) (res Staff, err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	query := `
		SELECT
			*
		FROM
			staffs
		WHERE
			staff_phone_number = :phone_number
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
			"phone_number": phone,
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
		err = ErrPhoneNumberNotFound
		return
	}

	return
}
