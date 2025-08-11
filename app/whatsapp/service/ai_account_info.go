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

type AiAccountInfo struct {
	service.Service
}

// GetPage 获取AiAccountInfo列表
func (e *AiAccountInfo) GetPage(c *dto.AiAccountInfoGetPageReq, p *actions.DataPermission, list *[]models.AiAccountInfo, count *int64) error {
	var err error
	var data models.AiAccountInfo

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("AiAccountInfoService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取AiAccountInfo对象
func (e *AiAccountInfo) Get(d *dto.AiAccountInfoGetReq, p *actions.DataPermission, model *models.AiAccountInfo) error {
	var data models.AiAccountInfo

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetAiAccountInfo error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建AiAccountInfo对象
func (e *AiAccountInfo) Insert(c *dto.AiAccountInfoInsertReq) error {
	var err error
	var data models.AiAccountInfo
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("AiAccountInfoService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改AiAccountInfo对象
func (e *AiAccountInfo) Update(c *dto.AiAccountInfoUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.AiAccountInfo{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("AiAccountInfoService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除AiAccountInfo
func (e *AiAccountInfo) Remove(d *dto.AiAccountInfoDeleteReq, p *actions.DataPermission) error {
	var data models.AiAccountInfo

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Unscoped().Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveAiAccountInfo error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
