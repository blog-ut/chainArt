// Package data
// Time : 2023/8/16 23:32
// Author : PTJ
// File : user
// Software: GoLand
package data

import (
	"chainArt/internal/biz"
	"chainArt/internal/pkg/util"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	PassWordCost = 12 // 密码加密难度
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

func (r *userRepo) Create(ctx context.Context, u *biz.User) (*biz.User, error) {
	user := BuildUser(u)
	user.SetPassword(u.Password)
	if err := r.data.db.Model(&user).Create(&user).Error; err != nil {
		return nil, err
	}
	return BuildUserToBiz(user), nil
}
func (r *userRepo) GetUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	var user User
	db := r.data.db.Model(&user)
	if u.ID != 0 {
		db.Where("id = ?", u.ID)
	}
	if u.UserName != "" {
		db.Where("user_name = ?", u.UserName)
	}
	if u.Phone != "" {
		db.Where("phone = ?", u.Phone)
	}
	result := db.First(&user)
	if result.RowsAffected != 1 {
		return nil, errors.NotFound("USER_NOT_FOUND", "用户不存在")
	}
	return BuildUserToBiz(&user), nil
}
func (r *userRepo) ListUser(ctx context.Context, pageNum int64, pageSize int64, u *biz.User) ([]*biz.User, int64, error) {
	user := make([]*User, 0)
	var total int64
	db := r.data.db.Model(&user)
	// db.Where
	if err := db.Count(&total).Scopes(util.Paginate(pageNum, pageSize)).Find(&user).Error; err != nil {
		return nil, 0, err
	}
	ru := make([]*biz.User, 0)
	for _, us := range user {
		ru = append(ru, BuildUserToBiz(us))
	}
	return ru, total, nil
}
func (r *userRepo) Update(ctx context.Context, u *biz.User) (bool, error) {
	if u.ID == 0 {
		return false, errors.New(500, "INVALID_PARAM", "ID为空")
	}
	userInfo := BuildUser(u)
	if u.Password != "" {
		userInfo.SetPassword(u.Password)
	}
	if err := r.data.db.Model(User{}).Where("id = ?", u.ID).Updates(userInfo).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *userRepo) Delete(ctx context.Context, u *biz.User) (bool, error) {
	if u.ID == 0 {
		return false, errors.New(500, "INVALID_PARAM", "ID为空")
	}
	if err := r.data.db.Delete(User{}, u.ID).Error; err != nil {
		return false, err
	}
	return true, nil
	return true, nil
}

// CheckPassword 校验密码
func (r *userRepo) CheckPassword(ctx context.Context, password string, encryptedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
	return err == nil
}

// SetPassword 密码加密
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func BuildUser(user *biz.User) *User {
	return &User{
		UserName: user.UserName,
		NickName: user.NickName,
		Intro:    user.Intro,
		Avatar:   user.Avatar,
		Phone:    user.Phone,
		Email:    user.Email,
		Address:  user.Address,
		IsEnable: user.IsEnable,
		IsAdmin:  user.IsAdmin,
	}
}

func BuildUserToBiz(user *User) *biz.User {
	return &biz.User{
		ID:       int64(user.ID),
		UserName: user.UserName,
		NickName: user.NickName,
		Password: user.Password,
		Intro:    user.Intro,
		Avatar:   user.Avatar,
		Phone:    user.Phone,
		Email:    user.Email,
		Address:  user.Address,
		IsEnable: user.IsEnable,
		IsAdmin:  user.IsAdmin,
	}
}
