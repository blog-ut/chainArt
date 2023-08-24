// Package data
// Time : 2023/8/23 9:50
// Author : PTJ
// File : article
// Software: GoLand
package data

import (
	"chainArt/internal/biz"
	"chainArt/internal/domain"
	"chainArt/internal/pkg/util"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	UserId        int64     `gorm:"not null;Index;comment:用户ID"`
	CategoryId    int64     `gorm:"Index;comment:分类ID"`
	Title         string    `gorm:"Index;comment:文章标题"`
	Intro         string    `gorm:"type:text;comment:简介"`
	Cover         string    `gorm:"comment:封面"`
	Content       string    `gorm:"type:text;comment:文章内容"`
	ContentMd     string    `gorm:"type:text;comment:文章内容-markdown"`
	Code          string    `gorm:"not null;comment:文章编码"`
	KeyWords      util.Strs `gorm:"type:text;comment:关键字"`
	Sort          int64     `gorm:"comment:排序"`
	IsElite       int64     `gorm:"comment:是否推荐"`
	Hits          int64     `gorm:"comment:浏览量"`
	ArticleStatus int64     `gorm:"comment:文章状态0-发布1-草稿;"`
}

func (*Article) TableName() string {
	return "biz_article"
}

type articleRepo struct {
	data *Data
	log  *log.Helper
}

func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u *articleRepo) Create(ctx context.Context, a *domain.Article) (*domain.Article, error) {
	var article = BuildArticle(a)
	if err := u.data.db.Create(&article).Error; err != nil {
		return nil, err
	}
	return BuildArticleByBiz(article), nil
}

func (u *articleRepo) GetArticle(ctx context.Context, a *domain.Article) (*domain.Article, error) {
	var article Article
	db := u.data.db.Model(&article)
	if a.Id != 0 {
		db.Where("id = ?", a.Id)
	}
	if a.Code != "" {
		db.Where("code = ?", a.Code)
	}
	if err := db.First(&article).Error; err != nil {
		return nil, err
	}
	return BuildArticleByBiz(&article), nil
}
func (u *articleRepo) ListArticle(ctx context.Context, pageNum int64, pageSize int64, a *domain.Article) ([]*domain.Article, int64, error) {
	var article = make([]*Article, 0)
	var total int64
	db := u.data.db.Model(&Article{})
	if a.CategoryId != 0 {
		db.Where("category_id = ?", a.CategoryId)
	}
	if a.Title != "" {
		db.Where("title like ?", "%"+a.Title+"%")
	}
	if err := db.Count(&total).Scopes(util.Paginate(pageNum, pageSize)).Find(&article).Error; err != nil {
		return nil, 0, err
	}
	var da = make([]*domain.Article, 0)
	for _, ar := range article {
		da = append(da, BuildArticleByBiz(ar))
	}
	return da, total, nil

}
func (u *articleRepo) Update(ctx context.Context, a *domain.Article) (bool, error) {
	return true, nil
}
func (u *articleRepo) Delete(ctx context.Context, a *domain.Article) (bool, error) {
	return true, nil
}

func BuildArticle(article *domain.Article) *Article {
	return &Article{
		UserId:        article.UserId,
		CategoryId:    article.CategoryId,
		Title:         article.Title,
		Intro:         article.Intro,
		Cover:         article.Cover,
		Content:       article.Content,
		ContentMd:     article.ContentMd,
		Code:          article.Code,
		KeyWords:      article.KeyWords,
		Sort:          article.Sort,
		IsElite:       article.IsElite,
		Hits:          article.Hits,
		ArticleStatus: article.ArticleStatus,
	}
}

func BuildArticleByBiz(article *Article) *domain.Article {
	return &domain.Article{
		Id:            int64(article.ID),
		UserId:        article.UserId,
		CategoryId:    article.CategoryId,
		Title:         article.Title,
		Intro:         article.Intro,
		Cover:         article.Cover,
		Content:       article.Content,
		ContentMd:     article.ContentMd,
		Code:          article.Code,
		KeyWords:      article.KeyWords,
		IsElite:       article.IsElite,
		Hits:          article.Hits,
		ArticleStatus: article.ArticleStatus,
		CreateAt:      article.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdateAt:      article.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
