package service

import (
	"chainArt/internal/biz"
	"chainArt/internal/domain"
	"context"
	"github.com/go-kratos/kratos/v2/log"

	pb "chainArt/api/article/v1"
)

type ArticleService struct {
	pb.UnimplementedArticleServer
	ar  *biz.ArticleUsecase
	log *log.Helper
}

func NewArticleService(ar *biz.ArticleUsecase, logger log.Logger) *ArticleService {
	return &ArticleService{ar: ar, log: log.NewHelper(logger)}
}

func (s *ArticleService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleReply, error) {
	_, err := s.ar.Create(ctx, &domain.Article{
		UserId:        req.GetUserId(),
		CategoryId:    req.GetCategoryId(),
		Title:         req.GetTitle(),
		Intro:         req.GetIntro(),
		Cover:         req.GetCover(),
		Content:       req.GetContent(),
		ContentMd:     req.GetContentMd(),
		Code:          req.GetCode(),
		KeyWords:      req.GetKeyWords(),
		Sort:          req.GetSort(),
		IsElite:       req.GetIsElite(),
		Hits:          req.GetHits(),
		ArticleStatus: req.GetArticleStatus(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateArticleReply{}, nil
}
func (s *ArticleService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleReply, error) {
	return &pb.UpdateArticleReply{}, nil
}
func (s *ArticleService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleReply, error) {
	return &pb.DeleteArticleReply{}, nil
}
func (s *ArticleService) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.GetArticleReply, error) {
	article, err := s.ar.GetArticle(ctx, &domain.Article{
		Id:   req.GetId(),
		Code: req.GetCode(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.GetArticleReply{
		Info: BuildArticle(article),
	}, nil
}
func (s *ArticleService) ListArticle(ctx context.Context, req *pb.ListArticleRequest) (*pb.ListArticleReply, error) {
	listArticle, total, err := s.ar.ListArticle(ctx, req.GetPageNum(), req.GetPageSize(), &domain.Article{
		Title:      req.GetTitle(),
		CategoryId: req.GetCategoryId(),
		KeyWords:   req.GetKeyWords(),
	})
	if err != nil {
		return nil, err
	}
	var list = make([]*pb.ArticleResp, 0)
	for _, la := range listArticle {
		list = append(list, BuildArticle(la))
	}
	return &pb.ListArticleReply{
		Total: total,
		Info:  list,
	}, nil
}

func BuildArticle(article *domain.Article) *pb.ArticleResp {
	return &pb.ArticleResp{
		Id:            article.Id,
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
		CreateAt:      article.CreateAt,
		UpdateAt:      article.UpdateAt,
	}
}
