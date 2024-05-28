package mewsql_test

import (
	"testing"

	"github.com/JesseNicholas00/BeliMang/utils/helper"

	"github.com/JesseNicholas00/BeliMang/utils/mewsql"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSelect(t *testing.T) {
	Convey("When constructing a query", t, func() {
		sql, vars := mewsql.Select(
			"product_id, product_name",
			"products",
			mewsql.WithWhere(
				mewsql.WithCondition("product_id = ?", "turbo-gyatt-001"),
				mewsql.WithCondition("product_name ILIKE ?", "%amogus%"),
				mewsql.WithCondition("product_sku = ?", "sku deez nuts"),
				// not an actual postgis function but eh
				mewsql.WithConditionMultiArgs(
					"DIST(product_pos, Point(?, ?)) < ?",
					float32(1.2),
					float32(2.3),
					float32(25),
				),
			),
			mewsql.WithOrderBy("created_at", "asc"),
			mewsql.WithLimit(5),
			mewsql.WithOffset(0),
		)
		Convey("Should construct the query string correctly", func() {
			So(
				sql,
				ShouldEqual,
				"SELECT product_id, product_name FROM products  WHERE (product_id = $1 AND product_name ILIKE $2 AND product_sku = $3 AND DIST(product_pos, Point($4, $5)) < $6) ORDER BY created_at ASC LIMIT 5 OFFSET 0",
			)
			So(vars, ShouldHaveLength, 6)
			So(vars[0], ShouldEqual, "turbo-gyatt-001")
			So(vars[1], ShouldEqual, "%amogus%")
			So(vars[2], ShouldEqual, "sku deez nuts")
			So(helper.IsEqualFloat(vars[3].(float32), 1.2), ShouldBeTrue)
			So(helper.IsEqualFloat(vars[4].(float32), 2.3), ShouldBeTrue)
			So(helper.IsEqualFloat(vars[5].(float32), 25.0), ShouldBeTrue)
		})
	})
	Convey("When constructing a query with AND and OR joiners", t, func() {
		sql, vars := mewsql.Select(
			"*",
			"users",
			mewsql.WithWhere(
				mewsql.Or(
					mewsql.WithCondition("user_id = ?", "turbo-gyatt-001"),
					mewsql.WithCondition("user_name = ?", "fanum tax"),
				),
				mewsql.Or(
					mewsql.WithCondition("user_is_gamer = ?", "true"),
					mewsql.WithCondition(`user_name = "epic"`, nil),
				),
			),
			mewsql.WithOrderBy("created_at", "asc"),
			mewsql.WithLimit(5),
			mewsql.WithOffset(0),
		)
		Convey("Should construct the query string correctly", func() {
			So(
				sql,
				ShouldEqual,
				`SELECT * FROM users  WHERE ((user_id = $1 OR user_name = $2) AND (user_is_gamer = $3 OR user_name = "epic")) ORDER BY created_at ASC LIMIT 5 OFFSET 0`,
			)
			So(vars, ShouldHaveLength, 3)
			So(vars[0], ShouldEqual, "turbo-gyatt-001")
			So(vars[1], ShouldEqual, "fanum tax")
			So(vars[2], ShouldEqual, "true")
		})
	})
}
