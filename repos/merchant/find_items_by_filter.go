package merchant

import (
	"context"

	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/mewsql"
)

// AdminListMerchant implements MerchantRepository.
func (lol *merchantRepoImpl) FindMerchantItemsByFilter(ctx context.Context,
	filter MerchantItemListFilter,
) (res []MerchantItem, total int64, err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	var conditions []mewsql.Condition

	if filter.MerchantItemId != nil {
		conditions = append(
			conditions,
			mewsql.WithCondition("merchant_item_id = ?", *filter.MerchantItemId),
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

	ctx, sess, err := lol.dbRizz.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	countRows, err := sess.Ext.QueryxContext(ctx, sqlCount, args...)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}
	defer countRows.Close()

	for countRows.Next() {
		var cur Total
		if err = countRows.StructScan(&cur); err != nil {
			err = errorutil.AddCurrentContext(err)
			return
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
		err = errorutil.AddCurrentContext(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var cur MerchantItem
		if err = rows.StructScan(&cur); err != nil {
			err = errorutil.AddCurrentContext(err)
			return
		}

		res = append(res, cur)
	}

	return
}
