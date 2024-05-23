package auth

import (
	"context"

	"github.com/KerakTelor86/GoBoiler/utils/errorutil"
	"github.com/jmoiron/sqlx"
)

func (repo *authRepostioryImpl) CreateStaff(
	ctx context.Context,
	staff Staff,
) (res Staff, err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	query := `
		INSERT INTO staffs(
			staff_id,
			staff_name,
			staff_phone_number,
			staff_password
		) VALUES (
			:staff_id,
			:staff_name,
			:staff_phone_number,
			:staff_password
		) RETURNING
			staff_id,
			staff_name,
			staff_phone_number,
			staff_password,
			created_at,
			updated_at
	`
	ctx, sess, err := repo.dbRizzer.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	rows, err := sqlx.NamedQueryContext(ctx, sess.Ext, query, staff)
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
