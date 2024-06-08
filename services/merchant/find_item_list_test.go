package merchant

import (
	"context"
	"testing"
	"time"

	"github.com/JesseNicholas00/BeliMang/repos/merchant"
	"github.com/JesseNicholas00/BeliMang/types/pagination"
	"github.com/JesseNicholas00/BeliMang/utils/unittesting"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMerchantItemList(t *testing.T) {
	Convey("When listing merchant items for admin", t, func() {
		mockCtrl, service, mockedRepo := NewWithMockedRepo(t)
		defer mockCtrl.Finish()
		merchantId := "merchantId"
		merchantItemId := "merchantItemId"
		limit := 5
		offset := 0
		name := "epic"
		category := "category"
		price := int64(100)
		createdAtSort := "asc"

		Convey("When called with a valid request", func() {
			req := MerchantItemListReq{
				MerchantItemId: &merchantItemId,
				Name:           &name,
				Category:       &category,
				CreatedAtSort:  &createdAtSort,
				Limit:          &limit,
				Offset:         &offset,
			}
			var res MerchantItemListRes

			expectedTotal := int64(1)

			Convey("Should pass the request to the repo layer", func() {
				expectedMeta := pagination.Page{
					Offset: &offset,
					Limit:  &limit,
					Total:  &expectedTotal,
				}

				m := merchant.MerchantItemListFilter{
					MerchantItemId: req.MerchantItemId,
					Name:           req.Name,
					Limit:          *req.Limit,
					Offset:         *req.Offset,
					Category:       req.Category,
					CreatedAtSort:  req.CreatedAtSort,
				}

				mockedRepo.
					EXPECT().
					FindMerchantItemsByFilter(gomock.Any(), m).
					Return([]merchant.MerchantItem{
						{
							Id:        merchantId,
							Name:      name,
							Category:  category,
							Price:     price,
							ImageUrl:  "https://images.com/image.jpg",
							CreatedAt: time.Now(),
						},
					}, expectedTotal, nil).
					Times(1)

				Convey("And return the expected id", func() {
					unittesting.FixNextUuid()
					err := service.FindMerchantItemList(context.TODO(), req, &res)
					So(err, ShouldBeNil)
					So(res.Meta, ShouldEqual, expectedMeta)
				})
			})
		})
	})
}
