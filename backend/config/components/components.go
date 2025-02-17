package components

import (
	"github.com/go-raptor/raptor/v3"
)

func New(c *raptor.Config) *raptor.Components {
	return &raptor.Components{
		Controllers: Controllers(),
		Services:    Services(),
		Middlewares: Middlewares(),
	}
}
