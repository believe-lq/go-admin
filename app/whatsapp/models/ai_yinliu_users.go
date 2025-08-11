package models

import (

	"go-admin/common/models"

)

type AiYinliuUsers struct {
    models.Model
    
    From string `json:"from" gorm:"type:varchar(255);comment:引流客户"` 
    To string `json:"to" gorm:"type:varchar(255);comment:主播账号"` 
    Nickname string `json:"nickname" gorm:"type:varchar(255);comment:用户昵称"` 
    Mark string `json:"mark" gorm:"type:varchar(255);comment:备注信息"` 
    CountryCode string `json:"countryCode" gorm:"type:varchar(255);comment:国家代码"` 
    UserType int64 `json:"userType" gorm:"type:int;comment:用户类型"` 
    models.ModelTime
    models.ControlBy
}

func (AiYinliuUsers) TableName() string {
    return "ai_yinliu_users"
}

func (e *AiYinliuUsers) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *AiYinliuUsers) GetId() interface{} {
	return e.Id
}