package order

import (
	"context"
	"github.com/JesseNicholas00/BeliMang/services/auth"
	"github.com/JesseNicholas00/BeliMang/services/order"
	"github.com/JesseNicholas00/BeliMang/utils/helper"
	"github.com/JesseNicholas00/BeliMang/utils/unittesting"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

type itemReq struct {
	ItemId   string `json:"ItemId"`
	Quantity int    `json:"Quantity"`
}

type orderReq struct {
	MerchantId      string    `json:"MerchantId"`
	IsStartingPoint bool      `json:"IsStartingPoint"`
	Items           []itemReq `json:"Items"`
}

func TestEstimateOrder(t *testing.T) {
	Convey("When given a valid request", t, func() {
		mockCtrl, controller, mockedService := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		lat := 6.9
		long := 42.0

		orders := []orderReq{
			{
				MerchantId:      "bread",
				IsStartingPoint: false,
				Items: []itemReq{
					{
						ItemId:   "towa",
						Quantity: 1,
					},
					{
						ItemId:   "fubuki",
						Quantity: 2,
					},
				},
			},
			{
				MerchantId:      "boogie",
				IsStartingPoint: false,
				Items: []itemReq{
					{
						ItemId:   "iofi",
						Quantity: 3,
					},
					{
						ItemId:   "okayu",
						Quantity: 4,
					},
				},
			},
			{
				MerchantId:      "bruh",
				IsStartingPoint: true,
				Items: []itemReq{
					{
						ItemId:   "juan",
						Quantity: 69,
					},
				},
			},
			{
				MerchantId:      "mewing",
				IsStartingPoint: false,
				Items: []itemReq{
					{
						ItemId:   "original-juan",
						Quantity: 69,
					},
					{
						ItemId:   "fake-juan",
						Quantity: 42,
					},
					{
						ItemId:   "gyatttt",
						Quantity: 361,
					},
				},
			},
		}

		callWithOrders := func() *httptest.ResponseRecorder {
			rec := httptest.NewRecorder()
			ctx := unittesting.CreateEchoContextFromRequest(
				http.MethodPost,
				"/users/estimate",
				rec,
				unittesting.WithJsonPayload(map[string]interface{}{
					"userLocation": map[string]interface{}{
						"lat":  lat,
						"long": long,
					},
					"orders": orders,
				}),
			)
			ctx.Set("session", auth.GetSessionFromTokenRes{
				UserId:  "henry",
				IsAdmin: false,
			})
			unittesting.CallController(ctx, controller.estimateOrder)
			return rec
		}

		Convey("On invalid number of start points", func() {
			Convey("When there are no start points", func() {
				orders[2].IsStartingPoint = false

				Convey("Should return http code 400", func() {
					rec := callWithOrders()
					So(rec.Code, ShouldEqual, http.StatusBadRequest)
				})
			})
			Convey("When there are multiple start points", func() {
				orders[0].IsStartingPoint = true
				orders[3].IsStartingPoint = true

				Convey("Should return http code 400", func() {
					rec := callWithOrders()
					So(rec.Code, ShouldEqual, http.StatusBadRequest)
				})
			})
		})

		Convey("On valid number of start points", func() {
			callWithResponse := func(
				returnedErr error,
				returnedRes *order.EstimateOrderRes,
			) *httptest.ResponseRecorder {
				mockedService.
					EXPECT().
					EstimateOrder(gomock.Any(), gomock.Any(), gomock.Any()).
					Do(
						func(
							_ context.Context,
							req order.EstimateOrderReq,
							res *order.EstimateOrderRes,
						) {
							for i, ord := range req.Orders {
								So(
									ord.MerchantId,
									ShouldEqual,
									orders[i].MerchantId,
								)
								So(
									*ord.IsStartingPoint,
									ShouldEqual,
									orders[i].IsStartingPoint,
								)
								for j, it := range ord.Items {
									So(
										it.ItemId,
										ShouldEqual,
										orders[i].Items[j].ItemId,
									)
									So(
										it.Quantity,
										ShouldEqual,
										orders[i].Items[j].Quantity,
									)
								}
							}
							if returnedRes != nil {
								*res = *returnedRes
							}
						},
					).
					Return(returnedErr).
					Times(1)
				return callWithOrders()
			}

			Convey("When traveled distance is too far", func() {
				Convey("Should return http code 400", func() {
					rec := callWithResponse(order.ErrTooFar, nil)
					So(rec.Code, ShouldEqual, http.StatusBadRequest)
				})
			})
			Convey("When merchant is not found", func() {
				Convey("Should return http code 404", func() {
					rec := callWithResponse(order.ErrMerchantNotFound, nil)
					So(rec.Code, ShouldEqual, http.StatusNotFound)
				})
			})
			Convey("When item is not found", func() {
				Convey("Should return http code 404", func() {
					rec := callWithResponse(order.ErrItemNotFound, nil)
					So(rec.Code, ShouldEqual, http.StatusNotFound)
				})
			})
			Convey("When all is good", func() {
				Convey(
					"Should return the response from the service layer",
					func() {
						expectedRes := order.EstimateOrderRes{
							TotalPrice:               69420,
							EstimatedDeliveryTimeMin: 1337,
							Id:                       "turbo-mewing-gyatt-69",
						}
						rec := callWithResponse(nil, &expectedRes)
						So(rec.Code, ShouldEqual, http.StatusOK)
						So(
							rec.Body.String(),
							ShouldEqualJSON,
							string(helper.MustMarshalJson(expectedRes)),
						)
					},
				)
			})
		})
	})
}
