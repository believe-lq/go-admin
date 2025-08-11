package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/whatsapp/models"
	"go-admin/app/whatsapp/service"
	"go-admin/app/whatsapp/service/dto"
	"go-admin/common/actions"
)

type AiAlertUrl struct {
	api.Api
}

// GetPage 获取告警地址管理列表
// @Summary 获取告警地址管理列表
// @Description 获取告警地址管理列表
// @Tags 告警地址管理
// @Param alertName query string false "告警昵称"
// @Param alertUrl query string false "告警地址"
// @Param isUse query string false "使用状态"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.AiAlertUrl}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ai-alert-url [get]
// @Security Bearer
func (e AiAlertUrl) GetPage(c *gin.Context) {
    req := dto.AiAlertUrlGetPageReq{}
    s := service.AiAlertUrl{}
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
	list := make([]models.AiAlertUrl, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取告警地址管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取告警地址管理
// @Summary 获取告警地址管理
// @Description 获取告警地址管理
// @Tags 告警地址管理
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.AiAlertUrl} "{"code": 200, "data": [...]}"
// @Router /api/v1/ai-alert-url/{id} [get]
// @Security Bearer
func (e AiAlertUrl) Get(c *gin.Context) {
	req := dto.AiAlertUrlGetReq{}
	s := service.AiAlertUrl{}
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
	var object models.AiAlertUrl

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取告警地址管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建告警地址管理
// @Summary 创建告警地址管理
// @Description 创建告警地址管理
// @Tags 告警地址管理
// @Accept application/json
// @Product application/json
// @Param data body dto.AiAlertUrlInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/ai-alert-url [post]
// @Security Bearer
func (e AiAlertUrl) Insert(c *gin.Context) {
    req := dto.AiAlertUrlInsertReq{}
    s := service.AiAlertUrl{}
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
		e.Error(500, err, fmt.Sprintf("创建告警地址管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改告警地址管理
// @Summary 修改告警地址管理
// @Description 修改告警地址管理
// @Tags 告警地址管理
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.AiAlertUrlUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/ai-alert-url/{id} [put]
// @Security Bearer
func (e AiAlertUrl) Update(c *gin.Context) {
    req := dto.AiAlertUrlUpdateReq{}
    s := service.AiAlertUrl{}
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
		e.Error(500, err, fmt.Sprintf("修改告警地址管理失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除告警地址管理
// @Summary 删除告警地址管理
// @Description 删除告警地址管理
// @Tags 告警地址管理
// @Param data body dto.AiAlertUrlDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/ai-alert-url [delete]
// @Security Bearer
func (e AiAlertUrl) Delete(c *gin.Context) {
    s := service.AiAlertUrl{}
    req := dto.AiAlertUrlDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除告警地址管理失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
