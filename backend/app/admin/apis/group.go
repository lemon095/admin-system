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

type Group struct {
	api.Api
}

// GetPage 获取战团列表
// @Summary 获取战团列表
// @Description 获取战团列表
// @Tags 战团
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Group}} "{"code": 200, "data": [...]}"
// @Router /api/v1/group [get]
// @Security Bearer
func (e Group) GetPage(c *gin.Context) {
	req := dto.GroupGetPageReq{}
	s := service.Group{}
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
	list := make([]models.Group, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取战团失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取战团
// @Summary 获取战团
// @Description 获取战团
// @Tags 战团
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Group} "{"code": 200, "data": [...]}"
// @Router /api/v1/group/{id} [get]
// @Security Bearer
func (e Group) Get(c *gin.Context) {
	req := dto.GroupGetReq{}
	s := service.Group{}
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
	var object models.Group

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取战团失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建战团
// @Summary 创建战团
// @Description 创建战团
// @Tags 战团
// @Accept application/json
// @Product application/json
// @Param data body dto.GroupInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/group [post]
// @Security Bearer
func (e Group) Insert(c *gin.Context) {
	req := dto.GroupInsertReq{}
	s := service.Group{}
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
		e.Error(500, err, fmt.Sprintf("创建战团失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改战团
// @Summary 修改战团
// @Description 修改战团
// @Tags 战团
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.GroupUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/group/{id} [put]
// @Security Bearer
func (e Group) Update(c *gin.Context) {
	req := dto.GroupUpdateReq{}
	s := service.Group{}
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
	//req.Operator = user.GetUserName(c)
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改战团失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除战团
// @Summary 删除战团
// @Description 删除战团
// @Tags 战团
// @Param data body dto.GroupDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/group [delete]
// @Security Bearer
func (e Group) Delete(c *gin.Context) {
	s := service.Group{}
	req := dto.GroupDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除战团失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
