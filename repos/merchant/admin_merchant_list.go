package merchant

import (
	"context"

	"github.com/JesseNicholas00/BeliMang/types/pagination"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/mewsql"
	"github.com/JesseNicholas00/BeliMang/utils/transaction"
)

// AdminListMerchant implements MerchantRepository.
func (lol *merchantRepoImpl) AdminListMerchant(ctx context.Context,
	filter AdminMerchantListFilter,
) (res []Merchant, total int64, err error) {
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

	ctx, sess, err := lol.dbRizz.GetOrAppendTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	err = transaction.RunWithAutoCommit(&sess, func() error {
		//get count before adding pagination
		sqlCount, args := mewsql.Select(
			"count(*) as total",
			"merchants",
			options...,
		)

		countRows, err := sess.Ext.QueryxContext(ctx, sqlCount, args...)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}
		defer countRows.Close()

		for countRows.Next() {
			var cur pagination.Total
			if err = countRows.StructScan(&cur); err != nil {
				return errorutil.AddCurrentContext(err)
			}

			total = cur.Total
		}

		options = append(options, mewsql.WithLimit(filter.Limit),
			mewsql.WithOffset(filter.Offset))

		if filter.CreatedAtSort != nil {
			options = append(
				options,
				mewsql.WithOrderBy("created_at", *filter.CreatedAtSort),
			)
		} else {
			options = append(
				options,
				mewsql.WithOrderBy("created_at", "desc"),
			)
		}

		sql, args := mewsql.Select(
			"merchant_id, name, category, image_url, created_at, ST_X(location::geometry) AS longitude, ST_Y(location::geometry) AS latitude",
			"merchants",
			options...,
		)

		rows, err := sess.Ext.QueryxContext(ctx, sql, args...)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}
		defer rows.Close()

		for rows.Next() {
			var cur Merchant
			if err = rows.StructScan(&cur); err != nil {
				return errorutil.AddCurrentContext(err)
			}

			res = append(res, cur)
		}
		return nil
	})
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	return
}
