// Package data
// Time : 2023/8/16 23:32
// Author : PTJ
// File : user
// Software: GoLand
package data

import (
	"chainArt/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"not null;type:varchar(255);uniqueIndex;comment:用户名" json:"username"`
	NickName string `gorm:"not null;comment:昵称" json:"nickName"`
	Password string `gorm:"not null;comment:密码" json:"password"`
	Intro    string `gorm:"comment:简介" json:"intro"`
	Avatar   string `gorm:"comment:头像" json:"avatar"`
	Phone    string `gorm:"comment:电话号码" json:"phone"`
	Email    string `gorm:"comment:邮箱" json:"email"`
	Address  string `gorm:"comment:地址" json:"address"`
	IsEnable uint32 `gorm:"not null;type:tinyint(1);comment:是否禁用 0无效 1有效" json:"isEnable"`
	IsAdmin  uint32 `gorm:"not null;type:tinyint(1);comment:是否是管理员 0无效 1有效" json:"isAdmin"`
}

func (*User) TableName() string {
	return "sys_user"
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
