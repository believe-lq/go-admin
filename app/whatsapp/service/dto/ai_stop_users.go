package dto

import (

	"go-admin/app/whatsapp/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type AiStopUsersGetPageReq struct {
	dto.Pagination     `search:"-"`
    From string `form:"from"  search:"type:contains;column:from;table:ai_stop_users" comment:"发送者"`
    To string `form:"to"  search:"type:contains;column:to;table:ai_stop_users" comment:"接收者"`
    CountryCode string `form:"countryCode"  search:"type:exact;column:country_code;table:ai_stop_users" comment:"国家代码"`
    StopType int64 `form:"stopType"  search:"type:exact;column:stop_type;table:ai_stop_users" comment:"停止类型"`
    AiStopUsersOrder
}

type AiStopUsersOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:ai_stop_users"`
    From string `form:"fromOrder"  search:"type:order;column:from;table:ai_stop_users"`
    To string `form:"toOrder"  search:"type:order;column:to;table:ai_stop_users"`
    CountryCode string `form:"countryCodeOrder"  search:"type:order;column:country_code;table:ai_stop_users"`
    StopType string `form:"stopTypeOrder"  search:"type:order;column:stop_type;table:ai_stop_users"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ai_stop_users"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ai_stop_users"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ai_stop_users"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:ai_stop_users"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:ai_stop_users"`
    
}

func (m *AiStopUsersGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AiStopUsersInsertReq struct {
    Id int `json:"-" comment:""` // 
    From string `json:"from" comment:"发送者"`
    To string `json:"to" comment:"接收者"`
    CountryCode string `json:"countryCode" comment:"国家代码"`
    StopType int64 `json:"stopType" comment:"停止类型"`
    common.ControlBy
}

func (s *AiStopUsersInsertReq) Generate(model *models.AiStopUsers)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.From = s.From
    model.To = s.To
    model.CountryCode = s.CountryCode
    model.StopType = s.StopType
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *AiStopUsersInsertReq) GetId() interface{} {
	return s.Id
}

type AiStopUsersUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    From string `json:"from" comment:"发送者"`
    To string `json:"to" comment:"接收者"`
    CountryCode string `json:"countryCode" comment:"国家代码"`
    StopType int64 `json:"stopType" comment:"停止类型"`
    common.ControlBy
}

func (s *AiStopUsersUpdateReq) Generate(model *models.AiStopUsers)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.From = s.From
    model.To = s.To
    model.CountryCode = s.CountryCode
    model.StopType = s.StopType
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *AiStopUsersUpdateReq) GetId() interface{} {
	return s.Id
}

// AiStopUsersGetReq 功能获取请求参数
type AiStopUsersGetReq struct {
     Id int `uri:"id"`
}
func (s *AiStopUsersGetReq) GetId() interface{} {
	return s.Id
}

// AiStopUsersDeleteReq 功能删除请求参数
type AiStopUsersDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *AiStopUsersDeleteReq) GetId() interface{} {
	return s.Ids
}
