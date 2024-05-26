package merchant

import (
	"context"

	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/mewsql"
)

func (repo *merchantRepoImpl) FindMerchantByFilter(
	ctx context.Context,
	filter MerchantFilter,
) (res []Merchant, err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	var conditions []mewsql.Condition

	lol := func() bool {
		if filter.Location.Latitude != nil && filter.Location.Longitude != nil {
			return true
		}
		return false
	}()

	if filter.MerchantId != nil {
		conditions = append(conditions,
			mewsql.WithCondition("m.merchant_id = ?", *filter.MerchantId),
		)
	}

	if filter.Name != nil {
		conditions = append(
			conditions,
			mewsql.WithCondition("m.name ILIKE ?", "%"+*filter.Name+"%"),
		)
	}

	if filter.MerchantCategory != nil {
		conditions = append(conditions,
			mewsql.WithCondition("m.category = ?", *filter.MerchantCategory),
		)
	}

	options := []mewsql.SelectOption{
		mewsql.WithLimit(filter.Limit),
		mewsql.WithOffset(filter.Offset),
		mewsql.WithWhere(conditions...),
		mewsql.WithJoin("JOIN", "merchant_items mi", "m.merchant_id = mi.merchant_id"),
	}

	if lol {
		options = append(options, mewsql.WithOrderByNearestLocation("location", *filter.Location.Latitude, *filter.Location.Longitude))
	}

	sql, args := mewsql.Select(
		`
            m.merchant_id as merchant_id, 
            m.name as merchant_name, 
            m.category as merchant_category, 
            m.image_url as merchant_url, 
            ST_Y(m.location::geometry) as latitude, 
            ST_X(m.location::geometry) as longitude,
            m.created_at merchant_created_at, 
            mi.merchant_item_id as item_id, 
            mi.name as item_name, 
            mi.category as item_category, 
            mi.price as item_price, 
            mi.image_url as item_url, 
            mi.created_at as item_created_at
        `,
		"merchants m",
		options...,
	)

	ctx, sess, err := repo.dbRizz.GetOrNoTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	rows, err := sess.Ext.QueryxContext(ctx, sql, args...)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}
	defer rows.Close()

	merchantMap := make(map[string]*Merchant)
	for rows.Next() {
		var merchantID string
		var item MerchantItem
		var merchant Merchant

		err := rows.Scan(
			&merchantID,
			&merchant.Name,
			&merchant.Category,
			&merchant.ImageUrl,
			&merchant.Latitude,
			&merchant.Longitude,
			&merchant.CreatedAt,
			&item.Id,
			&item.Name,
			&item.Category,
			&item.Price,
			&item.ImageUrl,
			&item.CreatedAt,
		)
		if err != nil {
			err = errorutil.AddCurrentContext(err)
		}

		if _, exists := merchantMap[merchantID]; !exists {
			merchant.Id = merchantID
			merchantMap[merchantID] = &merchant
		}

		if item.Id != "" {
			merchantMap[merchantID].Items = append(merchantMap[merchantID].Items, item)
		}
	}

	for _, merchant := range merchantMap {
		res = append(res, *merchant)
	}
	return
}
