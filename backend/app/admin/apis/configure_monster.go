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

type ConfigureMonster struct {
	api.Api
}

// GetPage 获取ConfigureMonster列表
// @Summary 获取ConfigureMonster列表
// @Description 获取ConfigureMonster列表
// @Tags ConfigureMonster
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.ConfigureMonster}} "{"code": 200, "data": [...]}"
// @Router /api/v1/configure-monster [get]
// @Security Bearer
func (e ConfigureMonster) GetPage(c *gin.Context) {
	req := dto.ConfigureMonsterGetPageReq{}
	s := service.ConfigureMonster{}
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
	list := make([]models.ConfigureMonster, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取ConfigureMonster失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取ConfigureMonster
// @Summary 获取ConfigureMonster
// @Description 获取ConfigureMonster
// @Tags ConfigureMonster
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.ConfigureMonster} "{"code": 200, "data": [...]}"
// @Router /api/v1/configure-monster/{id} [get]
// @Security Bearer
func (e ConfigureMonster) Get(c *gin.Context) {
	req := dto.ConfigureMonsterGetReq{}
	s := service.ConfigureMonster{}
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
	var object models.ConfigureMonster

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取ConfigureMonster失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建ConfigureMonster
// @Summary 创建ConfigureMonster
// @Description 创建ConfigureMonster
// @Tags ConfigureMonster
// @Accept application/json
// @Product application/json
// @Param data body dto.ConfigureMonsterInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/configure-monster [post]
// @Security Bearer
func (e ConfigureMonster) Insert(c *gin.Context) {
	req := dto.ConfigureMonsterInsertReq{}
	s := service.ConfigureMonster{}
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
		e.Error(500, err, fmt.Sprintf("创建ConfigureMonster失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改ConfigureMonster
// @Summary 修改ConfigureMonster
// @Description 修改ConfigureMonster
// @Tags ConfigureMonster
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.ConfigureMonsterUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/configure-monster/{id} [put]
// @Security Bearer
func (e ConfigureMonster) Update(c *gin.Context) {
	req := dto.ConfigureMonsterUpdateReq{}
	s := service.ConfigureMonster{}
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
		e.Error(500, err, fmt.Sprintf("修改ConfigureMonster失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除ConfigureMonster
// @Summary 删除ConfigureMonster
// @Description 删除ConfigureMonster
// @Tags ConfigureMonster
// @Param data body dto.ConfigureMonsterDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/configure-monster [delete]
// @Security Bearer
func (e ConfigureMonster) Delete(c *gin.Context) {
	s := service.ConfigureMonster{}
	req := dto.ConfigureMonsterDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除ConfigureMonster失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
