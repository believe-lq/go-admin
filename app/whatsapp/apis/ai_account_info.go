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

type AiAccountInfo struct {
	api.Api
}

// GetPage 获取主播账号列表列表
// @Summary 获取主播账号列表列表
// @Description 获取主播账号列表列表
// @Tags 主播账号列表
// @Param mobilePhone query string false "手机号"
// @Param mark query string false "备注"
// @Param email query string false "邮箱"
// @Param openid query string false "openid"
// @Param isStopAi query int64 false "禁用AI"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.AiAccountInfo}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ai-account-info [get]
// @Security Bearer
func (e AiAccountInfo) GetPage(c *gin.Context) {
	req := dto.AiAccountInfoGetPageReq{}
	s := service.AiAccountInfo{}
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
	list := make([]models.AiAccountInfo, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取主播账号列表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取主播账号列表
// @Summary 获取主播账号列表
// @Description 获取主播账号列表
// @Tags 主播账号列表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.AiAccountInfo} "{"code": 200, "data": [...]}"
// @Router /api/v1/ai-account-info/{id} [get]
// @Security Bearer
func (e AiAccountInfo) Get(c *gin.Context) {
	req := dto.AiAccountInfoGetReq{}
	s := service.AiAccountInfo{}
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
	var object models.AiAccountInfo

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取主播账号列表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建主播账号列表
// @Summary 创建主播账号列表
// @Description 创建主播账号列表
// @Tags 主播账号列表
// @Accept application/json
// @Product application/json
// @Param data body dto.AiAccountInfoInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/ai-account-info [post]
// @Security Bearer
func (e AiAccountInfo) Insert(c *gin.Context) {
	req := dto.AiAccountInfoInsertReq{}
	s := service.AiAccountInfo{}
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
			e.Error(400, err, "账号已存在")
			return
		}

		// 其他错误
		e.Error(500, err, fmt.Sprintf("创建主播账号列表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改主播账号列表
// @Summary 修改主播账号列表
// @Description 修改主播账号列表
// @Tags 主播账号列表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.AiAccountInfoUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/ai-account-info/{id} [put]
// @Security Bearer
func (e AiAccountInfo) Update(c *gin.Context) {
	req := dto.AiAccountInfoUpdateReq{}
	s := service.AiAccountInfo{}
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
			e.Error(400, err, "账号已存在")
			return
		}

		e.Error(500, err, fmt.Sprintf("修改主播账号列表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除主播账号列表
// @Summary 删除主播账号列表
// @Description 删除主播账号列表
// @Tags 主播账号列表
// @Param data body dto.AiAccountInfoDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/ai-account-info [delete]
// @Security Bearer
func (e AiAccountInfo) Delete(c *gin.Context) {
	s := service.AiAccountInfo{}
	req := dto.AiAccountInfoDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除主播账号列表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
