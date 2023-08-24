// Package data
// Time : 2023/8/23 9:55
// Author : PTJ
// File : category
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

type Category struct {
	gorm.Model
	Code     string `gorm:"comment:'分类编码'"`
	Name     string `gorm:"comment:'分类名称'"`
	ParentId string `gorm:"comment:'父级ID'"`
}

func (*Category) TableName() string {
	return "biz_category"
}

type categoryRepo struct {
	data *Data
	log  *log.Helper
}

func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	return &categoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (cr *categoryRepo) Create(ctx context.Context, c *domain.Category) (*domain.Category, error) {
	category := BuildCategory(c)
	if err := cr.data.db.Create(&category).Error; err != nil {
		return nil, err
	}
	return &domain.Category{
		Id:       int64(category.ID),
		Code:     category.Code,
		Name:     category.Name,
		ParentId: category.ParentId,
	}, nil
}

func (cr *categoryRepo) GetCategory(ctx context.Context, c *domain.Category) (*domain.Category, error) {
	var category Category
	db := cr.data.db.Model(&category)
	if c.Id != 0 {
		db.Where("id = ?", c.Id)
	}
	if c.Code != "" {
		db.Where("code = ?", c.Code)
	}
	if err := db.First(&category).Error; err != nil {
		return nil, err
	}
	return &domain.Category{
		Id:       int64(category.ID),
		Code:     category.Code,
		Name:     category.Name,
		ParentId: category.ParentId,
	}, nil
}

func (cr *categoryRepo) ListCategory(ctx context.Context, pageNum int64, pageSize int64, c *domain.Category) ([]*domain.Category, int64, error) {
	var total int64
	var category = make([]*Category, 0)
	if err := cr.data.db.Count(&total).Scopes(util.Paginate(pageNum, pageSize)).Find(&category).Error; err != nil {
		return nil, 0, err
	}
	var dc = make([]*domain.Category, 0)
	for _, cat := range category {
		dc = append(dc, &domain.Category{
			Id:       int64(cat.ID),
			Code:     cat.Code,
			ParentId: cat.ParentId,
		})
	}
	return dc, total, nil
}

func BuildCategory(category *domain.Category) *Category {
	return &Category{
		Code:     category.Code,
		Name:     category.Name,
		ParentId: category.ParentId,
	}
}
