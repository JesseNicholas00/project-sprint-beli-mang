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

func TestAdminListMerchant(t *testing.T) {
	Convey("When listing merchants for admin", t, func() {
		mockCtrl, service, mockedRepo := NewWithMockedRepo(t)
		defer mockCtrl.Finish()
		merchantId := "220132512342"
		limit := 5
		offset := 0
		name := "epic"
		merchantCategory := "goodCategory"
		lat := float64(1.1)
		lon := float64(2.2)
		createdAtSort := "asc"

		Convey("When called with a valid request", func() {
			req := AdminListMerchantReq{
				MerchantId:       &merchantId,
				Name:             &name,
				MerchantCategory: &merchantCategory,
				CreatedAtSort:    &createdAtSort,
				Limit:            &limit,
				Offset:           &offset,
			}
			var res AdminListMerchantRes

			expectedTotal := int64(1)

			Convey("Should pass the request to the repo layer", func() {
				expectedMeta := pagination.Page{
					Offset: &offset,
					Limit:  &limit,
					Total:  &expectedTotal,
				}

				m := merchant.AdminMerchantListFilter{
					MerchantId:       req.MerchantId,
					Name:             req.Name,
					Limit:            *req.Limit,
					Offset:           *req.Offset,
					MerchantCategory: req.MerchantCategory,
					CreatedAtSort:    req.CreatedAtSort,
				}

				mockedRepo.
					EXPECT().
					AdminListMerchant(gomock.Any(), m).
					Return([]merchant.Merchant{
						{
							Id:        merchantId,
							Name:      name,
							Category:  merchantCategory,
							ImageUrl:  "https://images.com",
							Latitude:  lat,
							Longitude: lon,
							CreatedAt: time.Now(),
						},
					}, expectedTotal, nil).
					Times(1)

				Convey("And return the expected id", func() {
					unittesting.FixNextUuid()
					err := service.AdminListMerchant(context.TODO(), req, &res)
					So(err, ShouldBeNil)
					So(res.Meta, ShouldEqual, expectedMeta)
				})
			})
		})
	})
}
