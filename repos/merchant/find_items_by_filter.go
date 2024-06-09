package merchant

import (
	"context"

	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/mewsql"
	"github.com/JesseNicholas00/BeliMang/utils/transaction"
)

// AdminListMerchant implements MerchantRepository.
func (repo *merchantRepoImpl) FindMerchantItemsByFilter(ctx context.Context,
	filter MerchantItemListFilter,
) (res []MerchantItem, total int64, err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	ctx, sess, err := repo.dbRizz.GetOrAppendTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	err = transaction.RunWithAutoCommit(&sess, func() error {
		var conditions []mewsql.Condition
		if filter.MerchantId != nil {
			conditions = append(
				conditions,
				mewsql.WithCondition("merchant_id = ?", *filter.MerchantId),
			)
		}

		if filter.MerchantItemId != nil {
			conditions = append(
				conditions,
				mewsql.WithCondition(
					"merchant_item_id = ?",
					*filter.MerchantItemId,
				),
			)
		}
		if filter.Name != nil {
			conditions = append(
				conditions,
				mewsql.WithCondition("name ILIKE ?", "%"+*filter.Name+"%"),
			)
		}
		if filter.Category != nil {
			conditions = append(
				conditions,
				mewsql.WithCondition("category = ?", *filter.Category),
			)
		}
		options := []mewsql.SelectOption{
			mewsql.WithWhere(conditions...),
		}

		//get count before adding pagination
		sqlCount, args := mewsql.Select(
			"count(*) as total",
			"merchant_items",
			options...,
		)

		countRows, err := sess.Ext.QueryxContext(ctx, sqlCount, args...)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}
		defer countRows.Close()

		for countRows.Next() {
			var cur Total
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
		}

		sql, args := mewsql.Select(
			"merchant_item_id, name, category, price, image_url, created_at",
			"merchant_items",
			options...,
		)

		rows, err := sess.Ext.QueryxContext(ctx, sql, args...)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}
		defer rows.Close()

		for rows.Next() {
			var cur MerchantItem
			if err = rows.StructScan(&cur); err != nil {
				return errorutil.AddCurrentContext(err)
			}

			res = append(res, cur)
		}
		return nil
	})
	return
}
