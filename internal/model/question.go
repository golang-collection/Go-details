package model

import "github.com/jinzhu/gorm"

/**
* @Author: super
* @Date: 2020-10-07 14:22
* @Description:
**/

type Question struct {
	*Model
	Source   string `gorm:"column:source" json:"source"`
	State    uint8  `gorm:"column:state" json:"state"`
	Question string `gorm:"column:question" json:"question"`
}

// TableName sets the insert table name for this struct type
func (q Question) TableName() string {
	return "questions"
}

////用于swagger的内容展示
//type BedtimeStorySwagger struct {
//	List  []*Question
//	Pager *app.Pager
//}

//以下内容是数据库的CRUD操作
func (q Question) Create(db *gorm.DB) (*Question, error) {
	if err := db.Create(&q).Error; err != nil {
		return nil, err
	}

	return &q, nil
}

func (q Question) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&q).Where("id = ? AND is_del = ?", q.ID, 0).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

func (q Question) Get(db *gorm.DB) (Question, error) {
	var question Question
	db = db.Where("id = ? AND state = ? AND is_del = ?", q.ID, q.State, 0)
	err := db.First(&question).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return question, err
	}

	return question, nil
}

func (q Question) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? AND is_del = ?", q.Model.ID, 0).Delete(&q).Error; err != nil {
		return err
	}

	return nil
}
