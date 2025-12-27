package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type GiftPack struct {
	api.Api
}

// GetPage 获取兑换码管理列表
// @Summary 获取兑换码管理列表
// @Description 获取兑换码管理列表
// @Tags 兑换码管理
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.GiftPack}} "{"code": 200, "data": [...]}"
// @Router /api/v1/gift-pack [get]
// @Security Bearer
func (e GiftPack) GetPage(c *gin.Context) {
	req := dto.GiftPackGetPageReq{}
	s := service.GiftPack{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Query).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.GiftPack, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取兑换码管理失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取兑换码管理
// @Summary 获取兑换码管理
// @Description 获取兑换码管理
// @Tags 兑换码管理
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.GiftPack} "{"code": 200, "data": [...]}"
// @Router /api/v1/gift-pack/{id} [get]
// @Security Bearer
func (e GiftPack) Get(c *gin.Context) {
	req := dto.GiftPackGetReq{}
	s := service.GiftPack{}
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
	var object models.GiftPack

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取兑换码管理失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建兑换码管理
// @Summary 创建兑换码管理
// @Description 创建兑换码管理
// @Tags 兑换码管理
// @Accept application/json
// @Product application/json
// @Param data body dto.GiftPackInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/gift-pack [post]
// @Security Bearer
func (e GiftPack) Insert(c *gin.Context) {
	req := dto.GiftPackInsertReq{}
	s := service.GiftPack{}
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
	req.CreatedBy = user.GetUserName(c)

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建兑换码管理失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改兑换码管理
// @Summary 修改兑换码管理
// @Description 修改兑换码管理
// @Tags 兑换码管理
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.GiftPackUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/gift-pack/{id} [put]
// @Security Bearer
func (e GiftPack) Update(c *gin.Context) {
	req := dto.GiftPackUpdateReq{}
	s := service.GiftPack{}
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
	req.SetOperator(user.GetUserName(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改兑换码管理失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

func (e GiftPack) UpdateStatus(c *gin.Context) {
	req := dto.GiftPackUpdateReq{}
	s := service.GiftPack{}
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

	err = s.UpdateStatus(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改兑换码管理失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除兑换码管理
// @Summary 删除兑换码管理
// @Description 删除兑换码管理
// @Tags 兑换码管理
// @Param data body dto.GiftPackDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/gift-pack [delete]
// @Security Bearer
func (e GiftPack) Delete(c *gin.Context) {
	s := service.GiftPack{}
	req := dto.GiftPackDeleteReq{}
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

	// req.SetOperator(user.GetUserName(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除兑换码管理失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
