package di

import "github.com/ocasti/fuego-quazar/topsecret-split/v1/internal/ctx"

func providerHandler() *ctx.Handler {
	return ctx.NewHandler()
}
