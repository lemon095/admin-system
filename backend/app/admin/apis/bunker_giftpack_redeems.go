package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type GiftpackRedeems struct {
	api.Api
}

// GetPage 获取兑换记录列表
// @Summary 获取兑换记录列表
// @Description 获取兑换记录列表
// @Tags 兑换记录
// @Param redeemerUnionId query string false "兑换人微信union_id"
// @Param redeemCode query string false "兑换码"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.GiftpackRedeems}} "{"code": 200, "data": [...]}"
// @Router /api/v1/giftpack-redeems [get]
// @Security Bearer
func (e GiftpackRedeems) GetPage(c *gin.Context) {
	req := dto.GiftpackRedeemsGetPageReq{}
	s := service.GiftpackRedeems{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]dto.GiftpackRedeemsGetPageResp, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取兑换记录失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取兑换记录
// @Summary 获取兑换记录
// @Description 获取兑换记录
// @Tags 兑换记录
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.GiftpackRedeems} "{"code": 200, "data": [...]}"
// @Router /api/v1/giftpack-redeems/{id} [get]
// @Security Bearer
func (e GiftpackRedeems) Get(c *gin.Context) {
	req := dto.GiftpackRedeemsGetReq{}
	s := service.GiftpackRedeems{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.GiftpackRedeems

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取兑换记录失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建兑换记录
// @Summary 创建兑换记录
// @Description 创建兑换记录
// @Tags 兑换记录
// @Accept application/json
// @Product application/json
// @Param data body dto.GiftpackRedeemsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/giftpack-redeems [post]
// @Security Bearer
func (e GiftpackRedeems) Insert(c *gin.Context) {
	req := dto.GiftpackRedeemsInsertReq{}
	s := service.GiftpackRedeems{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建兑换记录失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改兑换记录
// @Summary 修改兑换记录
// @Description 修改兑换记录
// @Tags 兑换记录
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.GiftpackRedeemsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/giftpack-redeems/{id} [put]
// @Security Bearer
func (e GiftpackRedeems) Update(c *gin.Context) {
	req := dto.GiftpackRedeemsUpdateReq{}
	s := service.GiftpackRedeems{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.Operator = user.GetUserName(c)
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改兑换记录失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除兑换记录
// @Summary 删除兑换记录
// @Description 删除兑换记录
// @Tags 兑换记录
// @Param data body dto.GiftpackRedeemsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/giftpack-redeems [delete]
// @Security Bearer
func (e GiftpackRedeems) Delete(c *gin.Context) {
	s := service.GiftpackRedeems{}
	req := dto.GiftpackRedeemsDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.Operator = user.GetUserName(c)
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除兑换记录失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
