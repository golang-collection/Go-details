package dao

import "github.com/jinzhu/gorm"

/**
* @Author: super
* @Date: 2020-10-07 14:01
* @Description:
**/

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}