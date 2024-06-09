package merchant

import (
	"net/http"

	"github.com/JesseNicholas00/BeliMang/services/merchant"
	"github.com/JesseNicholas00/BeliMang/utils/errorutil"
	"github.com/JesseNicholas00/BeliMang/utils/helper"
	"github.com/JesseNicholas00/BeliMang/utils/request"
	"github.com/labstack/echo/v4"
)

func (mc *merchantController) findMerchantItemList(c echo.Context) error {
	var req merchant.MerchantItemListReq
	if err := request.BindAndValidate(c, &req); err != nil {
		return err
	}

	// desc sort should be default
	if req.CreatedAtSort == nil ||
		(*req.CreatedAtSort != "asc" && *req.CreatedAtSort != "desc") {
		req.CreatedAtSort = helper.ToPointer("desc")
	}
	if req.Limit == nil {
		req.Limit = helper.ToPointer(5)
	}
	if req.Offset == nil {
		req.Offset = helper.ToPointer(0)
	}

	var res merchant.MerchantItemListRes
	err := mc.service.FindMerchantItemList(c.Request().Context(), req, &res)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	if res.Data == nil {
		res.Data = make([]merchant.ListMerchantItemResData, 0)
	}

	return c.JSON(http.StatusOK, res)
}
