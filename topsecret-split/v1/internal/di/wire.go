//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/ocasti/fuego-quazar/topsecret-split/v1/internal/ctx"
)

func Initialize() (*ctx.Handler, error) {
	wire.Build(stdSet)

	return &ctx.Handler{}, nil
}
