package models

import (

	"go-admin/common/models"

)

type AiAlertUrl struct {
    models.Model
    
    AlertName string `json:"alertName" gorm:"type:varchar(255);comment:告警昵称"` 
    AlertUrl string `json:"alertUrl" gorm:"type:varchar(255);comment:告警地址"` 
    IsUse string `json:"isUse" gorm:"type:int;comment:使用状态"` 
    models.ModelTime
    models.ControlBy
}

func (AiAlertUrl) TableName() string {
    return "ai_alert_url"
}

func (e *AiAlertUrl) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *AiAlertUrl) GetId() interface{} {
	return e.Id
}