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

type ConfigureBullet struct {
	api.Api
}

// GetPage 获取ConfigureBullet列表
// @Summary 获取ConfigureBullet列表
// @Description 获取ConfigureBullet列表
// @Tags ConfigureBullet
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.ConfigureBullet}} "{"code": 200, "data": [...]}"
// @Router /api/v1/configure-bullet [get]
// @Security Bearer
func (e ConfigureBullet) GetPage(c *gin.Context) {
	req := dto.ConfigureBulletGetPageReq{}
	s := service.ConfigureBullet{}
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
	list := make([]models.ConfigureBullet, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取ConfigureBullet失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取ConfigureBullet
// @Summary 获取ConfigureBullet
// @Description 获取ConfigureBullet
// @Tags ConfigureBullet
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.ConfigureBullet} "{"code": 200, "data": [...]}"
// @Router /api/v1/configure-bullet/{id} [get]
// @Security Bearer
func (e ConfigureBullet) Get(c *gin.Context) {
	req := dto.ConfigureBulletGetReq{}
	s := service.ConfigureBullet{}
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
	var object models.ConfigureBullet

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取ConfigureBullet失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建ConfigureBullet
// @Summary 创建ConfigureBullet
// @Description 创建ConfigureBullet
// @Tags ConfigureBullet
// @Accept application/json
// @Product application/json
// @Param data body dto.ConfigureBulletInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/configure-bullet [post]
// @Security Bearer
func (e ConfigureBullet) Insert(c *gin.Context) {
	req := dto.ConfigureBulletInsertReq{}
	s := service.ConfigureBullet{}
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
		e.Error(500, err, fmt.Sprintf("创建ConfigureBullet失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改ConfigureBullet
// @Summary 修改ConfigureBullet
// @Description 修改ConfigureBullet
// @Tags ConfigureBullet
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.ConfigureBulletUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/configure-bullet/{id} [put]
// @Security Bearer
func (e ConfigureBullet) Update(c *gin.Context) {
	req := dto.ConfigureBulletUpdateReq{}
	s := service.ConfigureBullet{}
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
		e.Error(500, err, fmt.Sprintf("修改ConfigureBullet失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除ConfigureBullet
// @Summary 删除ConfigureBullet
// @Description 删除ConfigureBullet
// @Tags ConfigureBullet
// @Param data body dto.ConfigureBulletDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/configure-bullet [delete]
// @Security Bearer
func (e ConfigureBullet) Delete(c *gin.Context) {
	s := service.ConfigureBullet{}
	req := dto.ConfigureBulletDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除ConfigureBullet失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
