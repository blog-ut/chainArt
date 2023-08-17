package service

import (
	"chainArt/internal/biz"
	"chainArt/internal/pkg/token"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"

	pb "chainArt/api/user/v1"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	user, err := s.uc.GetUser(ctx, &biz.User{UserName: req.UserName})
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.NotFound("USER_NOT_FOUND", "账户不存在")
	}
	b := s.uc.CheckPassword(ctx, req.Password, user.Password)
	if b {
		return nil, errors.New(500, "PASSWORD_ERROR", "密码错误")
	}
	tokenStr, err := token.GenToken(user.UserName, uint(user.ID))
	return &pb.LoginReply{
		Token: fmt.Sprintf("Beaber %s", tokenStr),
	}, nil
}
func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	s.uc.Create(ctx, &biz.User{
		UserName: req.GetUserName(),
		Password: req.GetPassword(),
		NickName: req.GetNickName(),
		Intro:    req.GetIntro(),
		Avatar:   req.GetAvatar(),
		Phone:    req.GetPhone(),
		Email:    req.GetEmail(),
		Address:  req.GetAddress(),
		IsEnable: req.GetIsEnable(),
		IsAdmin:  req.GetIsAdmin(),
	})
	return &pb.RegisterReply{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	b, err := s.uc.Update(ctx, &biz.User{
		ID:       req.GetUserId(),
		UserName: req.GetUserName(),
		Password: req.GetPassword(),
		NickName: req.GetNickName(),
		Intro:    req.GetIntro(),
		Avatar:   req.GetAvatar(),
		Phone:    req.GetPhone(),
		Email:    req.GetEmail(),
		Address:  req.GetAddress(),
		IsEnable: req.GetIsEnable(),
		IsAdmin:  req.GetIsAdmin(),
	})
	if err != nil || !b {
		return nil, err
	}
	return &pb.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	b, err := s.uc.Update(ctx, &biz.User{
		ID:    req.GetUserId(),
		Phone: req.GetPhone(),
	})
	if err != nil || !b {
		return nil, err
	}
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	user, err := s.uc.GetUser(ctx, &biz.User{
		ID:       req.GetUserId(),
		Phone:    req.GetPhone(),
		UserName: req.GetUserName(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.GetUserReply{
		Info: BuildUser(user),
	}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	user, total, err := s.uc.ListUser(ctx, req.GetPageNum(), req.GetPageSize(), &biz.User{})
	if err != nil {
		return nil, err
	}
	us := make([]*pb.UserResp, 0)
	for _, u := range user {
		us = append(us, BuildUser(u))
	}
	return &pb.ListUserReply{
		Total: total,
		Info:  us,
	}, nil
}

func BuildUser(user *biz.User) *pb.UserResp {
	return &pb.UserResp{
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
