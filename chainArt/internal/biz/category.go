// Package biz
// Time : 2023/8/24 9:21
// Author : PTJ
// File : category
// Software: GoLand
package biz

import (
	"chainArt/internal/domain"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type CategoryRepo interface {
	Create(context.Context, *domain.Category) (*domain.Category, error)
	GetCategory(context.Context, *domain.Category) (*domain.Category, error)
	ListCategory(context.Context, int64, int64, *domain.Category) ([]*domain.Category, int64, error)
}

type CategoryUseCase struct {
	repo CategoryRepo
	log  *log.Helper
}

func NewCategoryUseCase(repo CategoryRepo, logger log.Logger) *CategoryUseCase {
	return &CategoryUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (u *CategoryUseCase) Create(ctx context.Context, c *domain.Category) (*domain.Category, error) {
	return u.repo.Create(ctx, c)
}

func (u *CategoryUseCase) GetCategory(ctx context.Context, c *domain.Category) (*domain.Category, error) {
	return u.repo.GetCategory(ctx, c)
}

func (u *CategoryUseCase) ListCategory(ctx context.Context, pageNum int64, pageSize int64, c *domain.Category) ([]*domain.Category, int64, error) {
	return u.repo.ListCategory(ctx, pageNum, pageSize, c)
}
