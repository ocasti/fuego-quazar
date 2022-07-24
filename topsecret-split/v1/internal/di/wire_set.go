//go:build wireinject
// +build wireinject

package di

import "github.com/google/wire"

var stdSet = wire.NewSet(

	providerHandler,
)
