package main

import (
	"github.com/goadesign/goa"
	"github.com/rightscale/croner/app"
	"github.com/rightscale/croner/cron"
)

// JobController implements the jobs resource.
type JobController struct {
	*goa.Controller
	job *cron.Job
}

// NewJobController creates a jobs controller.
func NewJobController(service *goa.Service, job *cron.Job) app.JobController {
	return &JobController{Controller: service.NewController("jobs"), job: job}
}

// Show returns information about the cron job.
func (c *JobController) Show(ctx *app.ShowJobContext) error {
	return ctx.OK(c.job.ToMediaType())
}
