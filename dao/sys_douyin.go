package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"zeropiece/common"
)

type DouyinUser struct {
	gorm.Model
	UserName string `gorm:"column:user_name" json:"userName"`
	UID      string `gorm:"column:uid;unique_index" json:"uid"`
	Status   string `gorm:"column:status;default:'1'" json:"status"`
	Remark   string `gorm:"column:remark" json:"remark"`
}

type DouyinWorks struct {
	gorm.Model
	UID       string `gorm:"column:uid" json:"uid"`
	WorkID    string `gorm:"column:work_id;unique_index" json:"workId"`
	Title     string `gorm:"column:title" json:"title"`
	Name      string `gorm:"column:name" json:"name"`
	CommCount int    `gorm:"column:comm_count" json:"commCount"`
}

// Create 创建数据
func (userWorks *DouyinWorks) Create() (err error) {
	return common.DB.Create(&userWorks).Error
}

// Save 更新数据
func (userWorks *DouyinWorks) Save() (err error) {
	return common.DB.Save(&userWorks).Error
}

// FindAll 查询所有数据
func (userWorks *DouyinWorks) FindAll() (outModels []DouyinWorks, err error) {
	err = common.DB.Find(&outModels).Error
	return
}

// Filter 根据条件进行查询
func (userWorks *DouyinWorks) Filter() (outModels []DouyinWorks, err error) {
	err = common.DB.Where(&userWorks).Find(&outModels).Error
	return
}

// Delete 删除数据
func (userWorks *DouyinWorks) Delete() (err error) {
	return common.DB.Delete(&userWorks).Error
}

// FindWorksIdList 查询用户的作品id列表
func (userWorks *DouyinWorks) FindWorksIdList() (outModel []interface{}, err error) {
	var workList []DouyinWorks
	err = common.DB.Where(&userWorks).Find(&workList).Error
	for _, v := range workList {
		outModel = append(outModel, v.WorkID)
	}
	return
}

func (userWorks *DouyinWorks) CheckWorkIdInList(workId string, workList []interface{}) (err error) {
	for _, v := range workList {
		if v.(string) == workId {
			return nil
		}

	}
	return errors.New("not found")
}

// Create 创建数据
func (userName *DouyinUser) Create() (err error) {
	return common.DB.Create(&userName).Error
}

// Save 更新数据
func (userName *DouyinUser) Save() (err error) {
	return common.DB.Save(&userName).Error
}

// FindAll 查询所有数据
func (userName *DouyinUser) FindAll() (outModels []DouyinUser, err error) {
	err = common.DB.Find(&outModels).Error
	return
}

// Filter 根据条件进行查询
func (userName *DouyinUser) Filter() (outModels []DouyinUser, err error) {
	err = common.DB.Where(&userName).Find(&outModels).Error
	return
}

// Delete 删除数据
func (userName *DouyinUser) Delete() (err error) {
	return common.DB.Delete(&userName).Error
}

// FindAllUserIdList 查询所有用户id列表
func (userName *DouyinUser) FindAllUserIdList() (userList []DouyinUser, err error) {
	err = common.DB.Find(&userList).Error
	return
}
