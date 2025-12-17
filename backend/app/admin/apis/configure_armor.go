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

type ConfigureArmor struct {
	api.Api
}

// GetPage 获取ConfigureArmor列表
// @Summary 获取ConfigureArmor列表
// @Description 获取ConfigureArmor列表
// @Tags ConfigureArmor
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.ConfigureArmor}} "{"code": 200, "data": [...]}"
// @Router /api/v1/configure-armor [get]
// @Security Bearer
func (e ConfigureArmor) GetPage(c *gin.Context) {
	req := dto.ConfigureArmorGetPageReq{}
	s := service.ConfigureArmor{}
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
	list := make([]models.ConfigureArmor, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取ConfigureArmor失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取ConfigureArmor
// @Summary 获取ConfigureArmor
// @Description 获取ConfigureArmor
// @Tags ConfigureArmor
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.ConfigureArmor} "{"code": 200, "data": [...]}"
// @Router /api/v1/configure-armor/{id} [get]
// @Security Bearer
func (e ConfigureArmor) Get(c *gin.Context) {
	req := dto.ConfigureArmorGetReq{}
	s := service.ConfigureArmor{}
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
	var object models.ConfigureArmor

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取ConfigureArmor失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建ConfigureArmor
// @Summary 创建ConfigureArmor
// @Description 创建ConfigureArmor
// @Tags ConfigureArmor
// @Accept application/json
// @Product application/json
// @Param data body dto.ConfigureArmorInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/configure-armor [post]
// @Security Bearer
func (e ConfigureArmor) Insert(c *gin.Context) {
	req := dto.ConfigureArmorInsertReq{}
	s := service.ConfigureArmor{}
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
		e.Error(500, err, fmt.Sprintf("创建ConfigureArmor失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改ConfigureArmor
// @Summary 修改ConfigureArmor
// @Description 修改ConfigureArmor
// @Tags ConfigureArmor
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.ConfigureArmorUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/configure-armor/{id} [put]
// @Security Bearer
func (e ConfigureArmor) Update(c *gin.Context) {
	req := dto.ConfigureArmorUpdateReq{}
	s := service.ConfigureArmor{}
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
		e.Error(500, err, fmt.Sprintf("修改ConfigureArmor失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除ConfigureArmor
// @Summary 删除ConfigureArmor
// @Description 删除ConfigureArmor
// @Tags ConfigureArmor
// @Param data body dto.ConfigureArmorDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/configure-armor [delete]
// @Security Bearer
func (e ConfigureArmor) Delete(c *gin.Context) {
	s := service.ConfigureArmor{}
	req := dto.ConfigureArmorDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除ConfigureArmor失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
