package controller

import (
	"admin-system/common/response"
	"admin-system/service"

	"github.com/gin-gonic/gin"
)

// SystemController 系统配置控制器
type SystemController struct {
	systemService *service.SystemService
}

// NewSystemController 创建系统配置控制器
func NewSystemController() *SystemController {
	return &SystemController{
		systemService: &service.SystemService{},
	}
}

func (c *SystemController) SystemList(ctx *gin.Context) {
	category := ctx.Param("category")
	if category == "" {
		response.BadRequest(ctx, "no param")
		return
	}

	// 调用服务层
	resp, err := c.systemService.SystemList(category)
	if err != nil {
		response.Unauthorized(ctx, err.Error())
		return
	}

	response.Data(ctx, resp)
}

func (c *SystemController) SystemCreate(ctx *gin.Context) {
	category := ctx.Param("category")
	if category == "" {
		response.BadRequest(ctx, "no param")
		return
	}

	var system service.SystemCreateReq
	if err := ctx.ShouldBind(&system); err != nil {
		response.BadRequest(ctx, "no body")
		return
	}

	if err := c.systemService.SystemCreate(category, system); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	response.OK(ctx)
}
