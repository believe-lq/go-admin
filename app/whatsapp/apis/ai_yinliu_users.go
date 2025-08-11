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

type AiYinliuUsers struct {
	api.Api
}

// GetPage 获取引流用户列表列表
// @Summary 获取引流用户列表列表
// @Description 获取引流用户列表列表
// @Tags 引流用户列表
// @Param from query string false "引流客户"
// @Param to query string false "主播账号"
// @Param nickname query string false "用户昵称"
// @Param mark query string false "备注信息"
// @Param countryCode query string false "国家代码"
// @Param userType query int64 false "用户类型"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.AiYinliuUsers}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ai-yinliu-users [get]
// @Security Bearer
func (e AiYinliuUsers) GetPage(c *gin.Context) {
    req := dto.AiYinliuUsersGetPageReq{}
    s := service.AiYinliuUsers{}
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
	list := make([]models.AiYinliuUsers, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取引流用户列表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取引流用户列表
// @Summary 获取引流用户列表
// @Description 获取引流用户列表
// @Tags 引流用户列表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.AiYinliuUsers} "{"code": 200, "data": [...]}"
// @Router /api/v1/ai-yinliu-users/{id} [get]
// @Security Bearer
func (e AiYinliuUsers) Get(c *gin.Context) {
	req := dto.AiYinliuUsersGetReq{}
	s := service.AiYinliuUsers{}
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
	var object models.AiYinliuUsers

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取引流用户列表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建引流用户列表
// @Summary 创建引流用户列表
// @Description 创建引流用户列表
// @Tags 引流用户列表
// @Accept application/json
// @Product application/json
// @Param data body dto.AiYinliuUsersInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/ai-yinliu-users [post]
// @Security Bearer
func (e AiYinliuUsers) Insert(c *gin.Context) {
    req := dto.AiYinliuUsersInsertReq{}
    s := service.AiYinliuUsers{}
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
		e.Error(500, err, fmt.Sprintf("创建引流用户列表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改引流用户列表
// @Summary 修改引流用户列表
// @Description 修改引流用户列表
// @Tags 引流用户列表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.AiYinliuUsersUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/ai-yinliu-users/{id} [put]
// @Security Bearer
func (e AiYinliuUsers) Update(c *gin.Context) {
    req := dto.AiYinliuUsersUpdateReq{}
    s := service.AiYinliuUsers{}
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
		e.Error(500, err, fmt.Sprintf("修改引流用户列表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除引流用户列表
// @Summary 删除引流用户列表
// @Description 删除引流用户列表
// @Tags 引流用户列表
// @Param data body dto.AiYinliuUsersDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/ai-yinliu-users [delete]
// @Security Bearer
func (e AiYinliuUsers) Delete(c *gin.Context) {
    s := service.AiYinliuUsers{}
    req := dto.AiYinliuUsersDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除引流用户列表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
