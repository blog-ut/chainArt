// Package domain
// Time : 2023/8/23 9:59
// Author : PTJ
// File : category
// Software: GoLand
package domain

type Category struct {
	Id       int64
	Code     string
	Name     string
	ParentId string
}
