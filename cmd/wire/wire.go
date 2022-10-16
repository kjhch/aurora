//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/kjhch/aurora/internal/api"
	"github.com/kjhch/aurora/internal/config"
	"github.com/kjhch/aurora/internal/core"
)

var confSet = wire.NewSet(config.NewLoggerFactory)

var coreSet = wire.NewSet(core.NewApp, core.NewHttpServer)

var apiSet = wire.NewSet(api.NewBaseApi)

func InitApp() *core.App {
	panic(wire.Build(confSet, coreSet, apiSet))
}
