package components

import (
	"github.com/go-raptor/raptor/v4"
	"github.com/go-raptor/raptor/v4/core"
	"github.com/h00s/lotoweb3/app/middlewares"
)

func Middlewares() raptor.Middlewares {
	return raptor.Middlewares{
		core.Use(middlewares.NewCORSMiddleware(middlewares.CORSConfig{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST"},
			AllowHeaders:     []string{"Authorization", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           3600, // 1 hour
		})),
	}
}
