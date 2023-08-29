//go:build wireinject
// +build wireinject

package wireinject

import (
	"github.com/google/wire"
	"layout/internal/handler/timer"
	_ "layout/pkg/pprof"
)

var JobSet = wire.NewSet(timer.NewJob)

func NewApp() (*timer.Job, func(), error) {
	panic(wire.Build(
		JobSet,
	))
}
