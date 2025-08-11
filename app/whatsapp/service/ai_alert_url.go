package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/whatsapp/models"
	"go-admin/app/whatsapp/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type AiAlertUrl struct {
	service.Service
}

// GetPage 获取AiAlertUrl列表
func (e *AiAlertUrl) GetPage(c *dto.AiAlertUrlGetPageReq, p *actions.DataPermission, list *[]models.AiAlertUrl, count *int64) error {
	var err error
	var data models.AiAlertUrl

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("AiAlertUrlService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取AiAlertUrl对象
func (e *AiAlertUrl) Get(d *dto.AiAlertUrlGetReq, p *actions.DataPermission, model *models.AiAlertUrl) error {
	var data models.AiAlertUrl

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetAiAlertUrl error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建AiAlertUrl对象
func (e *AiAlertUrl) Insert(c *dto.AiAlertUrlInsertReq) error {
	var err error
	var data models.AiAlertUrl
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("AiAlertUrlService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改AiAlertUrl对象
func (e *AiAlertUrl) Update(c *dto.AiAlertUrlUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.AiAlertUrl{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("AiAlertUrlService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除AiAlertUrl
func (e *AiAlertUrl) Remove(d *dto.AiAlertUrlDeleteReq, p *actions.DataPermission) error {
	var data models.AiAlertUrl

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Unscoped().Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveAiAlertUrl error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
