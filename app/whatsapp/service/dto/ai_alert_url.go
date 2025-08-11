package dto

import (

	"go-admin/app/whatsapp/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type AiAlertUrlGetPageReq struct {
	dto.Pagination     `search:"-"`
    AlertName string `form:"alertName"  search:"type:contains;column:alert_name;table:ai_alert_url" comment:"告警昵称"`
    AlertUrl string `form:"alertUrl"  search:"type:contains;column:alert_url;table:ai_alert_url" comment:"告警地址"`
    IsUse string `form:"isUse"  search:"type:exact;column:is_use;table:ai_alert_url" comment:"使用状态"`
    AiAlertUrlOrder
}

type AiAlertUrlOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:ai_alert_url"`
    AlertName string `form:"alertNameOrder"  search:"type:order;column:alert_name;table:ai_alert_url"`
    AlertUrl string `form:"alertUrlOrder"  search:"type:order;column:alert_url;table:ai_alert_url"`
    IsUse string `form:"isUseOrder"  search:"type:order;column:is_use;table:ai_alert_url"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ai_alert_url"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ai_alert_url"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ai_alert_url"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:ai_alert_url"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:ai_alert_url"`
    
}

func (m *AiAlertUrlGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AiAlertUrlInsertReq struct {
    Id int `json:"-" comment:""` // 
    AlertName string `json:"alertName" comment:"告警昵称"`
    AlertUrl string `json:"alertUrl" comment:"告警地址"`
    IsUse string `json:"isUse" comment:"使用状态"`
    common.ControlBy
}

func (s *AiAlertUrlInsertReq) Generate(model *models.AiAlertUrl)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.AlertName = s.AlertName
    model.AlertUrl = s.AlertUrl
    model.IsUse = s.IsUse
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *AiAlertUrlInsertReq) GetId() interface{} {
	return s.Id
}

type AiAlertUrlUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    AlertName string `json:"alertName" comment:"告警昵称"`
    AlertUrl string `json:"alertUrl" comment:"告警地址"`
    IsUse string `json:"isUse" comment:"使用状态"`
    common.ControlBy
}

func (s *AiAlertUrlUpdateReq) Generate(model *models.AiAlertUrl)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.AlertName = s.AlertName
    model.AlertUrl = s.AlertUrl
    model.IsUse = s.IsUse
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *AiAlertUrlUpdateReq) GetId() interface{} {
	return s.Id
}

// AiAlertUrlGetReq 功能获取请求参数
type AiAlertUrlGetReq struct {
     Id int `uri:"id"`
}
func (s *AiAlertUrlGetReq) GetId() interface{} {
	return s.Id
}

// AiAlertUrlDeleteReq 功能删除请求参数
type AiAlertUrlDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *AiAlertUrlDeleteReq) GetId() interface{} {
	return s.Ids
}
