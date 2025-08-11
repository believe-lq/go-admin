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

type AiStopUsers struct {
	service.Service
}

// GetPage 获取AiStopUsers列表
func (e *AiStopUsers) GetPage(c *dto.AiStopUsersGetPageReq, p *actions.DataPermission, list *[]models.AiStopUsers, count *int64) error {
	var err error
	var data models.AiStopUsers

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("AiStopUsersService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取AiStopUsers对象
func (e *AiStopUsers) Get(d *dto.AiStopUsersGetReq, p *actions.DataPermission, model *models.AiStopUsers) error {
	var data models.AiStopUsers

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetAiStopUsers error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建AiStopUsers对象
func (e *AiStopUsers) Insert(c *dto.AiStopUsersInsertReq) error {
	var err error
	var data models.AiStopUsers
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("AiStopUsersService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改AiStopUsers对象
func (e *AiStopUsers) Update(c *dto.AiStopUsersUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.AiStopUsers{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("AiStopUsersService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除AiStopUsers
func (e *AiStopUsers) Remove(d *dto.AiStopUsersDeleteReq, p *actions.DataPermission) error {
	var data models.AiStopUsers

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Unscoped().Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveAiStopUsers error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
