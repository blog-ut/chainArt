package service

import (
	"chainArt/internal/biz"
	"chainArt/internal/domain"
	"context"
	"github.com/go-kratos/kratos/v2/log"

	pb "chainArt/api/category/v1"
)

type CategoryService struct {
	pb.UnimplementedCategoryServer
	cu  *biz.CategoryUseCase
	log *log.Helper
}

func NewCategoryService(cu *biz.CategoryUseCase, logger log.Logger) *CategoryService {
	return &CategoryService{cu: cu, log: log.NewHelper(logger)}
}

func (s *CategoryService) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CreateCategoryReply, error) {
	_, err := s.cu.Create(ctx, &domain.Category{
		Code:     req.GetCode(),
		Name:     req.GetName(),
		ParentId: req.GetParentId(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateCategoryReply{}, nil
}
func (s *CategoryService) GetCategory(ctx context.Context, req *pb.GetCategoryRequest) (*pb.GetCategoryReply, error) {
	category, err := s.cu.GetCategory(ctx, &domain.Category{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.GetCategoryReply{
		Info: BuildCategory(category),
	}, nil
}
func (s *CategoryService) ListCategory(ctx context.Context, req *pb.ListCategoryRequest) (*pb.ListCategoryReply, error) {
	categories, total, err := s.cu.ListCategory(ctx, req.GetPageNum(), req.GetPageSize(), &domain.Category{})
	if err != nil {
		return nil, err
	}
	var list = make([]*pb.CategoryResp, 0)
	for _, category := range categories {
		list = append(list, BuildCategory(category))
	}
	return &pb.ListCategoryReply{
		Total: total,
		Info:  list,
	}, nil
}

func BuildCategory(category *domain.Category) *pb.CategoryResp {
	return &pb.CategoryResp{
		Id:       category.Id,
		Code:     category.Code,
		Name:     category.Name,
		ParentId: category.ParentId,
	}
}
