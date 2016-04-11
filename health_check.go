package main

import (
	"github.com/goadesign/goa"
	"github.com/rightscale/croner/app"
)

// HealthCheckController implements the health_check resource.
type HealthCheckController struct {
	*goa.Controller
}

// NewHealthCheckController creates a health_check controller.
func NewHealthCheckController(service *goa.Service) app.HealthCheckController {
	return &HealthCheckController{Controller: service.NewController("HealthCheckController")}
}

// Do runs the do action.
func (c *HealthCheckController) Do(ctx *app.DoHealthCheckContext) error {
	return ctx.OK([]byte("OK"))
}
