// Package biz
// Time : 2023/8/23 10:03
// Author : PTJ
// File : article
// Software: GoLand
package biz

import (
	"chainArt/internal/domain"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ArticleRepo interface {
	Create(context.Context, *domain.Article) (*domain.Article, error)
	GetArticle(context.Context, *domain.Article) (*domain.Article, error)
	ListArticle(context.Context, int64, int64, *domain.Article) ([]*domain.Article, int64, error)
	Update(context.Context, *domain.Article) (bool, error)
	Delete(context.Context, *domain.Article) (bool, error)
}

type ArticleUsecase struct {
	repo ArticleRepo
	log  *log.Helper
}

func NewArticleUsecase(repo ArticleRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (u *ArticleUsecase) Create(ctx context.Context, article *domain.Article) (*domain.Article, error) {
	return u.repo.Create(ctx, article)
}

func (u *ArticleUsecase) GetArticle(ctx context.Context, article *domain.Article) (*domain.Article, error) {
	return u.repo.Create(ctx, article)
}
func (u *ArticleUsecase) ListArticle(ctx context.Context, pageNum int64, pageSize int64, article *domain.Article) ([]*domain.Article, int64, error) {
	return u.repo.ListArticle(ctx, pageNum, pageSize, article)
}
func (u *ArticleUsecase) Update(ctx context.Context, article *domain.Article) (bool, error) {
	return u.repo.Update(ctx, article)
}
func (u *ArticleUsecase) Delete(ctx context.Context, article *domain.Article) (bool, error) {
	return u.repo.Delete(ctx, article)
}
