package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Category struct {
	BaseModel
	Name    string `gorm:"type:varchar(64);not null;comment:'分类名称'" json:"name"`
	IsValid int    `gorm:"type:int;not null;default:1;comment:'是否可用：1-是 2-否'" json:"is_valid"`
}

// TableName 自定义表名
func (Category) TableName() string {
	return "category"
}

//BeforeCreate 在创建之前，先把创建时间赋值
func (cate *Category) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now().Unix())
	return nil
}

//BeforeUpdate 在更新之前，先把更新时间赋值
func (cate *Category) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now().Unix())
	return nil
}

func (cate *Category) SetName(name string) {
	cate.Name = name
}

func (cate *Category) SetIsValid(isValid int) {
	cate.IsValid = isValid
}
