package apis

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/whatsapp/models"
	"go-admin/app/whatsapp/service"
	"go-admin/app/whatsapp/service/dto"
	"go-admin/common/actions"
)

type AiStopUsers struct {
	api.Api
}

// GetPage 获取已停止AI回复客户列表
// @Summary 获取已停止AI回复客户列表
// @Description 获取已停止AI回复客户列表
// @Tags 已停止AI回复客户
// @Param from query string false "发送者"
// @Param to query string false "接收者"
// @Param countryCode query string false "国家代码"
// @Param stopType query int64 false "停止类型"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.AiStopUsers}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ai-stop-users [get]
// @Security Bearer
func (e AiStopUsers) GetPage(c *gin.Context) {
	req := dto.AiStopUsersGetPageReq{}
	s := service.AiStopUsers{}
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
	list := make([]models.AiStopUsers, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取已停止AI回复客户失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取已停止AI回复客户
// @Summary 获取已停止AI回复客户
// @Description 获取已停止AI回复客户
// @Tags 已停止AI回复客户
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.AiStopUsers} "{"code": 200, "data": [...]}"
// @Router /api/v1/ai-stop-users/{id} [get]
// @Security Bearer
func (e AiStopUsers) Get(c *gin.Context) {
	req := dto.AiStopUsersGetReq{}
	s := service.AiStopUsers{}
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
	var object models.AiStopUsers

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取已停止AI回复客户失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建已停止AI回复客户
// @Summary 创建已停止AI回复客户
// @Description 创建已停止AI回复客户
// @Tags 已停止AI回复客户
// @Accept application/json
// @Product application/json
// @Param data body dto.AiStopUsersInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/ai-stop-users [post]
// @Security Bearer
func (e AiStopUsers) Insert(c *gin.Context) {
	req := dto.AiStopUsersInsertReq{}
	s := service.AiStopUsers{}
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
		// 检查各种可能的重复键错误
		errMsg := err.Error()
		if errors.Is(err, gorm.ErrDuplicatedKey) ||
			strings.Contains(errMsg, "1062") ||
			strings.Contains(errMsg, "Duplicate entry") {
			e.Error(400, err, "记录已存在")
			return
		}

		e.Error(500, err, fmt.Sprintf("创建已停止AI回复客户失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改已停止AI回复客户
// @Summary 修改已停止AI回复客户
// @Description 修改已停止AI回复客户
// @Tags 已停止AI回复客户
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.AiStopUsersUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/ai-stop-users/{id} [put]
// @Security Bearer
func (e AiStopUsers) Update(c *gin.Context) {
	req := dto.AiStopUsersUpdateReq{}
	s := service.AiStopUsers{}
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
		// 检查各种可能的重复键错误
		errMsg := err.Error()
		if errors.Is(err, gorm.ErrDuplicatedKey) ||
			strings.Contains(errMsg, "1062") ||
			strings.Contains(errMsg, "Duplicate entry") {
			e.Error(400, err, "记录已存在")
			return
		}

		e.Error(500, err, fmt.Sprintf("修改已停止AI回复客户失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除已停止AI回复客户
// @Summary 删除已停止AI回复客户
// @Description 删除已停止AI回复客户
// @Tags 已停止AI回复客户
// @Param data body dto.AiStopUsersDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/ai-stop-users [delete]
// @Security Bearer
func (e AiStopUsers) Delete(c *gin.Context) {
	s := service.AiStopUsers{}
	req := dto.AiStopUsersDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除已停止AI回复客户失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
