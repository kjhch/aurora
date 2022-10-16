package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kjhch/aurora/internal/config"
	"go.uber.org/zap"
	"net/http"
)

type BaseApi struct {
	logger *zap.SugaredLogger
}

func NewBaseApi(factory *config.LoggerFactory) *BaseApi {
	return &BaseApi{logger: factory.GetDefaultLogger().Sugar()}
}

func (a *BaseApi) Author(ctx *gin.Context) {
	a.logger.Debugf("get author")
	ctx.JSON(http.StatusOK, gin.H{"author": config.Get().App.Author})
}
