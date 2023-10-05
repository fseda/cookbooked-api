package routes

import (
	"github.com/fseda/cookbooked-api/internal/infra/config"
)

func AddRoutes(ctx *config.AppContext) {
	addUserRoutes(ctx.App, ctx.DB)
	addAuthRoutes(ctx.App, ctx.DB, ctx.Env)
}
