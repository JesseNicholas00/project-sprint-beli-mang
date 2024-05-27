package merchant

import (
	"context"

	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/mewsql"
	"github.com/lib/pq"
)

// ListAllMerchant implements MerchantRepository.
func (lol *merchantRepoImpl) ListAllMerchant(ctx context.Context, filter MerchantListAllFilter) (res []Merchant, err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	var conditions []mewsql.Condition

	if filter.MerchantId != nil {
		conditions = append(
			conditions,
			mewsql.WithCondition("merchant_id = ?", *filter.MerchantId),
		)
	}

	if len(filter.MerchantIds) != 0 {
		conditions = append(
			conditions,
			mewsql.WithCondition("merchant_id = ANY (?) ", pq.StringArray(filter.MerchantIds)),
		)
	}

	if filter.Name != nil {
		conditions = append(
			conditions,
			mewsql.WithCondition("name ILIKE ?", "%"+*filter.Name+"%"),
		)
	}

	if filter.MerchantCategory != nil {
		conditions = append(
			conditions,
			mewsql.WithCondition("category = ?", *filter.MerchantCategory),
		)
	}

	options := []mewsql.SelectOption{
		mewsql.WithWhere(conditions...),
	}

	ctx, sess, err := lol.dbRizz.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	sql, args := mewsql.Select(
		"merchant_id, name, category, image_url, created_at, ST_X(location::geometry) AS longitude, ST_Y(location::geometry) AS latitude",
		"merchants",
		options...,
	)

	rows, err := sess.Ext.QueryxContext(ctx, sql, args...)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var cur Merchant
		if err = rows.StructScan(&cur); err != nil {
			err = errorutil.AddCurrentContext(err)
			return
		}

		res = append(res, cur)
	}

	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	return
}
