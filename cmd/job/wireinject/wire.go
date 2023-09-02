//go:build wireinject
// +build wireinject

package wireinject

import (
	"github.com/google/wire"
	"layout/application/timer_handler"
	_ "layout/pkg/pprof"
)

var JobSet = wire.NewSet(timer_handler.NewJob)

func NewApp() (*timer.Job, func(), error) {
	panic(wire.Build(
		JobSet,
	))
}
