// Package util
// Time : 2023/8/17 21:02
// Author : PTJ
// File : page
// Software: GoLand
package util

import "gorm.io/gorm"

// Paginate 设置分页参数
func Paginate(pageNum, pageSize int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNum == 0 {
			pageNum = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (pageNum - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}
