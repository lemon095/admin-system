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

type ConfigureWeapon struct {
	api.Api
}

// GetPage 获取武器配置列表
// @Summary 获取武器配置列表
// @Description 获取武器配置列表
// @Tags 武器配置
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.ConfigureWeapon}} "{"code": 200, "data": [...]}"
// @Router /api/v1/configure-weapon [get]
// @Security Bearer
func (e ConfigureWeapon) GetPage(c *gin.Context) {
	req := dto.ConfigureWeaponGetPageReq{}
	s := service.ConfigureWeapon{}
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
	list := make([]models.ConfigureWeapon, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取武器配置失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取武器配置
// @Summary 获取武器配置
// @Description 获取武器配置
// @Tags 武器配置
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.ConfigureWeapon} "{"code": 200, "data": [...]}"
// @Router /api/v1/configure-weapon/{id} [get]
// @Security Bearer
func (e ConfigureWeapon) Get(c *gin.Context) {
	req := dto.ConfigureWeaponGetReq{}
	s := service.ConfigureWeapon{}
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
	var object models.ConfigureWeapon

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取武器配置失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建武器配置
// @Summary 创建武器配置
// @Description 创建武器配置
// @Tags 武器配置
// @Accept application/json
// @Product application/json
// @Param data body dto.ConfigureWeaponInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/configure-weapon [post]
// @Security Bearer
func (e ConfigureWeapon) Insert(c *gin.Context) {
	req := dto.ConfigureWeaponInsertReq{}
	s := service.ConfigureWeapon{}
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
		e.Error(500, err, fmt.Sprintf("创建武器配置失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改武器配置
// @Summary 修改武器配置
// @Description 修改武器配置
// @Tags 武器配置
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.ConfigureWeaponUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/configure-weapon/{id} [put]
// @Security Bearer
func (e ConfigureWeapon) Update(c *gin.Context) {
	req := dto.ConfigureWeaponUpdateReq{}
	s := service.ConfigureWeapon{}
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
		e.Error(500, err, fmt.Sprintf("修改武器配置失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除武器配置
// @Summary 删除武器配置
// @Description 删除武器配置
// @Tags 武器配置
// @Param data body dto.ConfigureWeaponDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/configure-weapon [delete]
// @Security Bearer
func (e ConfigureWeapon) Delete(c *gin.Context) {
	s := service.ConfigureWeapon{}
	req := dto.ConfigureWeaponDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除武器配置失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
