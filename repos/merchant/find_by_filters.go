package merchant

import (
	"context"
	"time"

	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/mewsql"
)

type MerchantItem struct {
	Id         string    `db:"item_id"`
	MerchantId string    `db:"item_merchant_id"`
	Name       string    `db:"item_name"`
	Category   string    `db:"item_category"`
	Price      int64     `db:"item_price"`
	ImageUrl   string    `db:"item_image_url"`
	CreatedAt  time.Time `db:"item_created_at"`
}

type MerchantWithItems struct {
	Id        string    `db:"merchant_id"`
	Name      string    `db:"merchant_name"`
	Category  string    `db:"merchant_category"`
	ImageUrl  string    `db:"merchant_image_url"`
	Latitude  float64   `db:"latitude"`
	Longitude float64   `db:"longitude"`
	CreatedAt time.Time `db:"merchant_created_at"`
	Items     []MerchantItem
}

func (repo *merchantRepoImpl) FindMerchantByFilter(
	ctx context.Context,
	filter MerchantFilter,
) (res []MerchantWithItems, err error) {
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
            m.name as name, 
            m.category as category, 
            m.image_url as image_url, 
            ST_Y(m.location::geometry) as latitude, 
            ST_X(m.location::geometry) as longitude,
            m.created_at created_at, 
            mi.merchant_item_id as item_id, 
            mi.name as item_name, 
            mi.category as item_category, 
            mi.price as item_price, 
            mi.image_url as item_image_url, 
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

	merchantMap := make(map[string]MerchantWithItems)
	for rows.Next() {

		var dbRes struct {
			Merchant
			MerchantItem
		}
		err := rows.StructScan(&dbRes)
		if err != nil {
			err = errorutil.AddCurrentContext(err)
		}

		var merchant MerchantWithItems
		if _, exists := merchantMap[dbRes.Merchant.Id]; !exists {
			merchantMap[dbRes.Merchant.Id] = MerchantWithItems{
				Id:        dbRes.Merchant.Id,
				Name:      dbRes.Merchant.Name,
				Category:  dbRes.Merchant.Category,
				ImageUrl:  dbRes.Merchant.ImageUrl,
				Latitude:  dbRes.Merchant.Latitude,
				Longitude: dbRes.Merchant.Longitude,
				CreatedAt: dbRes.Merchant.CreatedAt,
				Items:     []MerchantItem{},
			}
		}

		if dbRes.MerchantItem.Id != "" {
			merchant = merchantMap[dbRes.Merchant.Id]
			merchant.Items = append(merchant.Items, dbRes.MerchantItem)
			merchantMap[merchant.Id] = merchant
		}
	}

	for _, merchant := range merchantMap {
		res = append(res, merchant)
	}
	return
}
