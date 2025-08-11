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

type AiYinliuUsers struct {
	service.Service
}

// GetPage 获取AiYinliuUsers列表
func (e *AiYinliuUsers) GetPage(c *dto.AiYinliuUsersGetPageReq, p *actions.DataPermission, list *[]models.AiYinliuUsers, count *int64) error {
	var err error
	var data models.AiYinliuUsers

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("AiYinliuUsersService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取AiYinliuUsers对象
func (e *AiYinliuUsers) Get(d *dto.AiYinliuUsersGetReq, p *actions.DataPermission, model *models.AiYinliuUsers) error {
	var data models.AiYinliuUsers

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetAiYinliuUsers error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建AiYinliuUsers对象
func (e *AiYinliuUsers) Insert(c *dto.AiYinliuUsersInsertReq) error {
    var err error
    var data models.AiYinliuUsers
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("AiYinliuUsersService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改AiYinliuUsers对象
func (e *AiYinliuUsers) Update(c *dto.AiYinliuUsersUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.AiYinliuUsers{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("AiYinliuUsersService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除AiYinliuUsers
func (e *AiYinliuUsers) Remove(d *dto.AiYinliuUsersDeleteReq, p *actions.DataPermission) error {
	var data models.AiYinliuUsers

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveAiYinliuUsers error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
