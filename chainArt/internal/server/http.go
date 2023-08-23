package server

import (
	article "chainArt/api/article/v1"
	user "chainArt/api/user/v1"
	"chainArt/internal/conf"
	"chainArt/internal/pkg/middleware"
	"chainArt/internal/pkg/util"
	"chainArt/internal/service"
	"github.com/go-kratos/grpc-gateway/v2/protoc-gen-openapiv2/generator"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, userService *service.UserService, articleService *service.ArticleService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			selector.Server(middleware.Auth()).Match(middleware.NewWhiteListMatcher()).Build(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	uploadFile(srv)
	openAPIhandler := openapiv2.NewHandler(
		openapiv2.WithGeneratorOptions(
			generator.UseJSONNamesForFields(false),
			generator.EnumsAsInts(true)))

	srv.HandlePrefix("/q/", openAPIhandler)
	user.RegisterUserHTTPServer(srv, userService)
	article.RegisterArticleHTTPServer(srv, articleService)
	return srv
}

func uploadFile(srv *http.Server) {
	route := srv.Route("/v1")
	route.POST("/upload", func(ctx http.Context) error {
		req := ctx.Request()
		f, h, err := req.FormFile("file")
		if err != nil {
			return err
		}
		defer f.Close()
		file, err := h.Open()

		path, err := util.UploadFile(file, h.Size)
		if err != nil {
			return err
		}
		return ctx.JSON(200, path)

	})
}
