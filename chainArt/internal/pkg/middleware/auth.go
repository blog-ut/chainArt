// Package middleware
// Time : 2023/8/17 23:10
// Author : PTJ
// File : auth
// Software: GoLand
package middleware

import (
	"chainArt/internal/pkg/token"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"strconv"
)

func Auth() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				tokenStr := tr.RequestHeader().Get("Authorization")
				if tokenStr == "" {
					return nil, errors.New(500, "no auth", "未授权")
				}
				userClaims, err := token.ParseToken(tokenStr)
				if err != nil {
					return nil, errors.New(500, "no auth", "未授权")
				}
				id := strconv.Itoa(int(userClaims.UserId))
				ctx = metadata.NewServerContext(ctx, metadata.New(map[string][]string{
					"userId":   {id},
					"username": {userClaims.Username},
				}))
			}
			return handler(ctx, req)
		}
	}
}

func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/api.user.v1.User/Login"] = struct{}{}
	whiteList["/api.user.v1.User/Register"] = struct{}{}
	whiteList["/api.article.v1.Article/ListArticle"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
