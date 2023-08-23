// Package data
// Time : 2023/8/23 9:55
// Author : PTJ
// File : category
// Software: GoLand
package data

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Code     string `gorm:"comment:'分类编码'"`
	Name     string `gorm:"comment:'分类名称'"`
	ParentId string `gorm:"comment:'父级ID'"`
}

func (*Category) TableName() string {
	return "biz_category"
}
