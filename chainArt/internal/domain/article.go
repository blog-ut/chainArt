// Package domain
// Time : 2023/8/23 9:57
// Author : PTJ
// File : article
// Software: GoLand
package domain

type Article struct {
	Id            int64
	UserId        int64
	CategoryId    int64
	Title         string
	Intro         string
	Cover         string
	Content       string
	ContentMd     string
	Code          string
	KeyWords      []string
	Sort          int64
	IsElite       int64
	Hits          int64
	ArticleStatus int64
	CreateAt      int64
	UpdateAt      int64
}
