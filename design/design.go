package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("croner", func() {
	Title("The cron service")
	Description(`croner is a simple cron as a service implementation.
croner accepts cron-like schedule specifications on the command line.`)
	Host("localhost:8080")
	Scheme("http")
})

var _ = Resource("job", func() {
	Action("show", func() {
		Routing(GET("job"))
		Description("Show information about cron job")
		Response(OK, Job)
	})
})

var _ = Resource("health_check", func() {
	Action("do", func() {
		Description("Health check")
		Routing(
			GET("/health-check"),
			GET("//health-check"),
		)
		Response(OK)
	})
})

// Job media type
var Job = MediaType("application/vnd.rightscale.croner.job+json", func() {
	Description("A cron job together with information on the last execution")
	TypeName("Job")
	Attributes(func() {
		Attribute("cmd", String, "scheduled command", func() {
			Example("bundle exec rake db:snapshot")
		})
		Attribute("schedule", String, "cron schedule spec", func() {
			Example("0 20 * * *")
		})
		Attribute("last", Execution, "last execution")
		Attribute("running", CollectionOf(Execution), "currently running executions if any")
		Required("cmd", "schedule")
	})
	View("default", func() {
		Attribute("cmd")
		Attribute("schedule")
		Attribute("last")
		Attribute("running")
	})
})

// Execution media type
var Execution = MediaType("application/vnd.rightscale.croner.execution+json", func() {
	Description("A job execution")
	TypeName("Execution")
	Attributes(func() {
		Attribute("pid", Integer, "Execution OS pid")
		Attribute("started_at", DateTime, "Execution started at timestamp")
		Attribute("finished_at", DateTime, "Execution finished at timestamp if finished")
		Attribute("stderr", String, "Execution stderr output if finished and if not empty")
		Attribute("exit_status", Integer, "Execution exit status if finished", func() {
			Example(0)
		})
		Required("pid", "started_at")
	})
	View("default", func() {
		Attribute("pid")
		Attribute("started_at")
		Attribute("finished_at")
		Attribute("stderr")
		Attribute("exit_status")
	})
})
