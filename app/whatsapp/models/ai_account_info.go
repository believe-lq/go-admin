package models

import (

	"go-admin/common/models"

)

type AiAccountInfo struct {
    models.Model
    
    MobilePhone string `json:"mobilePhone" gorm:"type:varchar(255);comment:手机号"` 
    Mark string `json:"mark" gorm:"type:varchar(255);comment:备注"` 
    Email string `json:"email" gorm:"type:varchar(255);comment:邮箱"` 
    Openid string `json:"openid" gorm:"type:varchar(255);comment:openid"` 
    IsStopAi int64 `json:"isStopAi" gorm:"type:int;comment:禁用AI"` 
    models.ModelTime
    models.ControlBy
}

func (AiAccountInfo) TableName() string {
    return "ai_account_info"
}

func (e *AiAccountInfo) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *AiAccountInfo) GetId() interface{} {
	return e.Id
}