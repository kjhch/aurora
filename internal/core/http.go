package core

import (
	"github.com/gin-gonic/gin"
	"github.com/kjhch/aurora/internal/api"
	"net/http"
)

func NewHttpServer(baseApi *api.BaseApi) *gin.Engine {
	engine := gin.Default()
	engine.Handle(http.MethodGet, "/", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
	apiV1Group := engine.Group("/api/v1")
	{
		apiV1Group.GET("author", baseApi.Author)
	}
	return engine
}
