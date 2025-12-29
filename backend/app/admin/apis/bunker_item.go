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

type Item struct {
	api.Api
}

// GetPage 获取道具表列表
// @Summary 获取道具表列表
// @Description 获取道具表列表
// @Tags 道具表
// @Param itemId query int64 false "道具id"
// @Param type query int64 false "道具类型"
// @Param name query string false "道具名称"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Item}} "{"code": 200, "data": [...]}"
// @Router /api/v1/item [get]
// @Security Bearer
func (e Item) GetPage(c *gin.Context) {
	req := dto.ItemGetPageReq{}
	s := service.Item{}
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
	list := make([]models.Item, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取道具表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

func (e Item) GetOption(c *gin.Context) {
	s := service.Item{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	list, err := s.GetOption()
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取道具表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(list, "查询成功")
}

// Get 获取道具表
// @Summary 获取道具表
// @Description 获取道具表
// @Tags 道具表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Item} "{"code": 200, "data": [...]}"
// @Router /api/v1/item/{id} [get]
// @Security Bearer
func (e Item) Get(c *gin.Context) {
	req := dto.ItemGetReq{}
	s := service.Item{}
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
	var object models.Item

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取道具表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建道具表
// @Summary 创建道具表
// @Description 创建道具表
// @Tags 道具表
// @Accept application/json
// @Product application/json
// @Param data body dto.ItemInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/item [post]
// @Security Bearer
func (e Item) Insert(c *gin.Context) {
	req := dto.ItemInsertReq{}
	s := service.Item{}
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
	req.Operator = user.GetUserName(c)

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建道具表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改道具表
// @Summary 修改道具表
// @Description 修改道具表
// @Tags 道具表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.ItemUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/item/{id} [put]
// @Security Bearer
func (e Item) Update(c *gin.Context) {
	req := dto.ItemUpdateReq{}
	s := service.Item{}
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
		e.Error(500, err, fmt.Sprintf("修改道具表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

func (e Item) UpdateStatus(c *gin.Context) {
	req := dto.ItemUpdateReq{}
	s := service.Item{}
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

	err = s.UpdateStatus(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改道具表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除道具表
// @Summary 删除道具表
// @Description 删除道具表
// @Tags 道具表
// @Param data body dto.ItemDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/item [delete]
// @Security Bearer
func (e Item) Delete(c *gin.Context) {
	s := service.Item{}
	req := dto.ItemDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除道具表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
