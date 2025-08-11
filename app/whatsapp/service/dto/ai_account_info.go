package dto

import (

	"go-admin/app/whatsapp/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type AiAccountInfoGetPageReq struct {
	dto.Pagination     `search:"-"`
    MobilePhone string `form:"mobilePhone"  search:"type:contains;column:mobile_phone;table:ai_account_info" comment:"手机号"`
    Mark string `form:"mark"  search:"type:contains;column:mark;table:ai_account_info" comment:"备注"`
    Email string `form:"email"  search:"type:contains;column:email;table:ai_account_info" comment:"邮箱"`
    Openid string `form:"openid"  search:"type:contains;column:openid;table:ai_account_info" comment:"openid"`
    IsStopAi int64 `form:"isStopAi"  search:"type:exact;column:is_stop_ai;table:ai_account_info" comment:"禁用AI"`
    AiAccountInfoOrder
}

type AiAccountInfoOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:ai_account_info"`
    MobilePhone string `form:"mobilePhoneOrder"  search:"type:order;column:mobile_phone;table:ai_account_info"`
    Mark string `form:"markOrder"  search:"type:order;column:mark;table:ai_account_info"`
    Email string `form:"emailOrder"  search:"type:order;column:email;table:ai_account_info"`
    Openid string `form:"openidOrder"  search:"type:order;column:openid;table:ai_account_info"`
    IsStopAi string `form:"isStopAiOrder"  search:"type:order;column:is_stop_ai;table:ai_account_info"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:ai_account_info"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:ai_account_info"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:ai_account_info"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:ai_account_info"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:ai_account_info"`
    
}

func (m *AiAccountInfoGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AiAccountInfoInsertReq struct {
    Id int `json:"-" comment:""` // 
    MobilePhone string `json:"mobilePhone" comment:"手机号"`
    Mark string `json:"mark" comment:"备注"`
    Email string `json:"email" comment:"邮箱"`
    Openid string `json:"openid" comment:"openid"`
    IsStopAi int64 `json:"isStopAi" comment:"禁用AI"`
    common.ControlBy
}

func (s *AiAccountInfoInsertReq) Generate(model *models.AiAccountInfo)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.MobilePhone = s.MobilePhone
    model.Mark = s.Mark
    model.Email = s.Email
    model.Openid = s.Openid
    model.IsStopAi = s.IsStopAi
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *AiAccountInfoInsertReq) GetId() interface{} {
	return s.Id
}

type AiAccountInfoUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    MobilePhone string `json:"mobilePhone" comment:"手机号"`
    Mark string `json:"mark" comment:"备注"`
    Email string `json:"email" comment:"邮箱"`
    Openid string `json:"openid" comment:"openid"`
    IsStopAi int64 `json:"isStopAi" comment:"禁用AI"`
    common.ControlBy
}

func (s *AiAccountInfoUpdateReq) Generate(model *models.AiAccountInfo)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.MobilePhone = s.MobilePhone
    model.Mark = s.Mark
    model.Email = s.Email
    model.Openid = s.Openid
    model.IsStopAi = s.IsStopAi
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *AiAccountInfoUpdateReq) GetId() interface{} {
	return s.Id
}

// AiAccountInfoGetReq 功能获取请求参数
type AiAccountInfoGetReq struct {
     Id int `uri:"id"`
}
func (s *AiAccountInfoGetReq) GetId() interface{} {
	return s.Id
}

// AiAccountInfoDeleteReq 功能删除请求参数
type AiAccountInfoDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *AiAccountInfoDeleteReq) GetId() interface{} {
	return s.Ids
}
