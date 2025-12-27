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

type Workshop struct {
	api.Api
}

// GetPage 获取工坊列表
// @Summary 获取工坊列表
// @Description 获取工坊列表
// @Tags 工坊
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Workshop}} "{"code": 200, "data": [...]}"
// @Router /api/v1/workshop [get]
// @Security Bearer
func (e Workshop) GetPage(c *gin.Context) {
	req := dto.WorkshopGetPageReq{}
	s := service.Workshop{}
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
	list := make([]models.Workshop, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取工坊失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取工坊
// @Summary 获取工坊
// @Description 获取工坊
// @Tags 工坊
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Workshop} "{"code": 200, "data": [...]}"
// @Router /api/v1/workshop/{id} [get]
// @Security Bearer
func (e Workshop) Get(c *gin.Context) {
	req := dto.WorkshopGetReq{}
	s := service.Workshop{}
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
	var object models.Workshop

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取工坊失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建工坊
// @Summary 创建工坊
// @Description 创建工坊
// @Tags 工坊
// @Accept application/json
// @Product application/json
// @Param data body dto.WorkshopInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/workshop [post]
// @Security Bearer
func (e Workshop) Insert(c *gin.Context) {
	req := dto.WorkshopInsertReq{}
	s := service.Workshop{}
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
		e.Error(500, err, fmt.Sprintf("创建工坊失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改工坊
// @Summary 修改工坊
// @Description 修改工坊
// @Tags 工坊
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.WorkshopUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/workshop/{id} [put]
// @Security Bearer
func (e Workshop) Update(c *gin.Context) {
	req := dto.WorkshopUpdateReq{}
	s := service.Workshop{}
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
		e.Error(500, err, fmt.Sprintf("修改工坊失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除工坊
// @Summary 删除工坊
// @Description 删除工坊
// @Tags 工坊
// @Param data body dto.WorkshopDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/workshop [delete]
// @Security Bearer
func (e Workshop) Delete(c *gin.Context) {
	s := service.Workshop{}
	req := dto.WorkshopDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除工坊失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
