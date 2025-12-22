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

type Auction struct {
	api.Api
}

// GetPage 获取Auction列表
// @Summary 获取Auction列表
// @Description 获取Auction列表
// @Tags Auction
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Auction}} "{"code": 200, "data": [...]}"
// @Router /api/v1/auction [get]
// @Security Bearer
func (e Auction) GetPage(c *gin.Context) {
	req := dto.AuctionGetPageReq{}
	s := service.Auction{}
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
	list := make([]models.Auction, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Auction失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Auction
// @Summary 获取Auction
// @Description 获取Auction
// @Tags Auction
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Auction} "{"code": 200, "data": [...]}"
// @Router /api/v1/auction/{id} [get]
// @Security Bearer
func (e Auction) Get(c *gin.Context) {
	req := dto.AuctionGetReq{}
	s := service.Auction{}
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
	var object models.Auction

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Auction失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建Auction
// @Summary 创建Auction
// @Description 创建Auction
// @Tags Auction
// @Accept application/json
// @Product application/json
// @Param data body dto.AuctionInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/auction [post]
// @Security Bearer
func (e Auction) Insert(c *gin.Context) {
	req := dto.AuctionInsertReq{}
	s := service.Auction{}
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
		e.Error(500, err, fmt.Sprintf("创建Auction失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Auction
// @Summary 修改Auction
// @Description 修改Auction
// @Tags Auction
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.AuctionUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/auction/{id} [put]
// @Security Bearer
func (e Auction) Update(c *gin.Context) {
	req := dto.AuctionUpdateReq{}
	s := service.Auction{}
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
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改Auction失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除Auction
// @Summary 删除Auction
// @Description 删除Auction
// @Tags Auction
// @Param data body dto.AuctionDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/auction [delete]
// @Security Bearer
func (e Auction) Delete(c *gin.Context) {
	s := service.Auction{}
	req := dto.AuctionDeleteReq{}
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

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除Auction失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
