package models

import (

	"go-admin/common/models"

)

type AiStopUsers struct {
    models.Model
    
    From string `json:"from" gorm:"type:varchar(255);comment:发送者"` 
    To string `json:"to" gorm:"type:varchar(255);comment:接收者"` 
    CountryCode string `json:"countryCode" gorm:"type:varchar(255);comment:国家代码"` 
    StopType int64 `json:"stopType" gorm:"type:int;comment:停止类型"` 
    models.ModelTime
    models.ControlBy
}

func (AiStopUsers) TableName() string {
    return "ai_stop_users"
}

func (e *AiStopUsers) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *AiStopUsers) GetId() interface{} {
	return e.Id
}