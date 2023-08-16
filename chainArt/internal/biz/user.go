// Package biz
// Time : 2023/8/16 23:32
// Author : PTJ
// File : user
// Software: GoLand
package biz

import "github.com/go-kratos/kratos/v2/log"

type User struct {
	ID       int64
	UserName string
	NickName string
	Password string
	Intro    string
	Avatar   string
	Phone    string
	Email    string
	Address  string
	IsEnable uint32
	IsAdmin  uint32
}

type UserRepo interface {
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}
