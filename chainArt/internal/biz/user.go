// Package biz
// Time : 2023/8/16 23:32
// Author : PTJ
// File : user
// Software: GoLand
package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

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
	Create(context.Context, *User) (*User, error)
	GetUser(context.Context, *User) (*User, error)
	ListUser(context.Context, int64, int64, *User) ([]*User, int64, error)
	Update(context.Context, *User) (bool, error)
	Delete(context.Context, *User) (bool, error)
	CheckPassword(context.Context, string, string) bool
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

func (r *UserUsecase) Create(ctx context.Context, u *User) (*User, error) {
	return r.repo.Create(ctx, u)
}
func (r *UserUsecase) GetUser(ctx context.Context, u *User) (*User, error) {
	return r.repo.GetUser(ctx, u)
}
func (r *UserUsecase) ListUser(ctx context.Context, pageNum int64, pageSize int64, u *User) ([]*User, int64, error) {
	return r.repo.ListUser(ctx, pageNum, pageSize, u)
}
func (r *UserUsecase) Update(ctx context.Context, u *User) (bool, error) {
	return r.repo.Update(ctx, u)
}
func (r *UserUsecase) Delete(ctx context.Context, u *User) (bool, error) {
	return r.repo.Delete(ctx, u)
}
func (r *UserUsecase) CheckPassword(ctx context.Context, password string, encryptedPassword string) bool {
	return r.repo.CheckPassword(ctx, password, encryptedPassword)
}
