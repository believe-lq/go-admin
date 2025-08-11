package dto

import (

	"go-admin/app/whatsapp/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type AiYinliuUsersGetPageReq struct {
	dto.Pagination     `search:"-"`
    From string `form:"from"  search:"type:contains;column:from;table:ai_yinliu_users" comment:"引流客户"`
    To string `form:"to"  search:"type:contains;column:to;table:ai_yinliu_users" comment:"主播账号"`
    Nickname string `form:"nickname"  search:"type:contains;column:nickname;table:ai_yinliu_users" comment:"用户昵称"`
    Mark string `form:"mark"  search:"type:contains;column:mark;table:ai_yinliu_users" comment:"备注信息"`
    CountryCode string `form:"countryCode"  search:"type:exact;column:country_code;table:ai_yinliu_users" comment:"国家代码"`
    UserType int64 `form:"userType"  search:"type:exact;column:user_type;table:ai_yinliu_users" comment:"用户类型"`
    AiYinliuUsersOrder
}

type AiYinliuUsersOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:ai_yinliu_users"`
    From string `form:"fromOrder"  search:"type:order;column:from;table:ai_yinliu_users"`
    To string `form:"toOrder"  search:"type:order;column:to;table:ai_yinliu_users"`
    Nickname string `form:"nicknameOrder"  search:"type:order;column:nickname;table:ai_yinliu_users"`
    Mark string `form:"markOrder"  search:"type:order;column:mark;table:ai_yinliu_users"`
    CountryCode string `form:"countryCodeOrder"  search:"type:order;column:country_code;table:ai_yinliu_users"`
    UserType string `form:"userTypeOrder"  search:"type:order;column:user_type;table:ai_yinliu_users"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ai_yinliu_users"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ai_yinliu_users"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ai_yinliu_users"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:ai_yinliu_users"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:ai_yinliu_users"`
    
}

func (m *AiYinliuUsersGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AiYinliuUsersInsertReq struct {
    Id int `json:"-" comment:""` // 
    From string `json:"from" comment:"引流客户"`
    To string `json:"to" comment:"主播账号"`
    Nickname string `json:"nickname" comment:"用户昵称"`
    Mark string `json:"mark" comment:"备注信息"`
    CountryCode string `json:"countryCode" comment:"国家代码"`
    UserType int64 `json:"userType" comment:"用户类型"`
    common.ControlBy
}

func (s *AiYinliuUsersInsertReq) Generate(model *models.AiYinliuUsers)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.From = s.From
    model.To = s.To
    model.Nickname = s.Nickname
    model.Mark = s.Mark
    model.CountryCode = s.CountryCode
    model.UserType = s.UserType
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *AiYinliuUsersInsertReq) GetId() interface{} {
	return s.Id
}

type AiYinliuUsersUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    From string `json:"from" comment:"引流客户"`
    To string `json:"to" comment:"主播账号"`
    Nickname string `json:"nickname" comment:"用户昵称"`
    Mark string `json:"mark" comment:"备注信息"`
    CountryCode string `json:"countryCode" comment:"国家代码"`
    UserType int64 `json:"userType" comment:"用户类型"`
    common.ControlBy
}

func (s *AiYinliuUsersUpdateReq) Generate(model *models.AiYinliuUsers)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.From = s.From
    model.To = s.To
    model.Nickname = s.Nickname
    model.Mark = s.Mark
    model.CountryCode = s.CountryCode
    model.UserType = s.UserType
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *AiYinliuUsersUpdateReq) GetId() interface{} {
	return s.Id
}

// AiYinliuUsersGetReq 功能获取请求参数
type AiYinliuUsersGetReq struct {
     Id int `uri:"id"`
}
func (s *AiYinliuUsersGetReq) GetId() interface{} {
	return s.Id
}

// AiYinliuUsersDeleteReq 功能删除请求参数
type AiYinliuUsersDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *AiYinliuUsersDeleteReq) GetId() interface{} {
	return s.Ids
}
