package order

import (
	"context"

	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/mewsql"
	"github.com/JesseNicholas00/BeliMang/utils/transaction"
	"github.com/lib/pq"
)

// ListOrderSummary implements OrderRepository.
func (lol *orderRepositoryImpl) ListOrderSummary(
	ctx context.Context,
	filter OrderSummaryListFilter,
) (res []OrderSummaryView, err error) {
	if err = ctx.Err(); err != nil {
		return
	}

	var conditions []mewsql.Condition

	if len(filter.MerchantIds) != 0 && filter.ItemName != nil {
		conditions = append(
			conditions,
			mewsql.WithConditionMultiArgs("merchant_id = ANY (?) or item_name ILIKE ?", pq.StringArray(filter.MerchantIds), "%"+*filter.ItemName+"%"),
		)
	} else if len(filter.MerchantIds) > 0 {
		conditions = append(
			conditions,
			mewsql.WithCondition("merchant_id = ANY (?)", pq.StringArray(filter.MerchantIds)),
		)
	} else if filter.ItemName != nil {
		conditions = append(
			conditions,
			mewsql.WithCondition("item_name ILIKE ?", "%"+*filter.ItemName+"%"),
		)
	}

	conditions = append(
		conditions,
		mewsql.WithCondition("user_id = ?", filter.UserId),
	)

	ctx, sess, err := lol.dbRizzer.GetOrAppendTx(ctx)
	if err != nil {
		err = errorutil.AddCurrentContext(err)
		return
	}

	options := []mewsql.SelectOption{
		mewsql.WithWhere(conditions...),
	}

	err = transaction.RunWithAutoCommit(&sess, func() error {
		options = append(options, mewsql.WithLimit(filter.Limit),
			mewsql.WithOffset(filter.Offset))

		sql, args := mewsql.Select(
			"*",
			"order_summary_view",
			options...,
		)

		rows, err := sess.Ext.QueryxContext(ctx, sql, args...)
		if err != nil {
			return errorutil.AddCurrentContext(err)
		}
		defer rows.Close()

		for rows.Next() {
			var cur OrderSummaryView
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
