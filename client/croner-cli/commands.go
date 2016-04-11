package main

import (
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	"github.com/rightscale/croner/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"log"
	"os"
)

type (
	// DoHealthCheckCommand is the command line data structure for the do action of health_check
	DoHealthCheckCommand struct {
	}
	// ShowJobCommand is the command line data structure for the show action of job
	ShowJobCommand struct {
	}
)

// Run makes the HTTP request corresponding to the DoHealthCheckCommand command.
func (cmd *DoHealthCheckCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/health-check"
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.UseLogger(context.Background(), logger)
	resp, err := c.DoHealthCheck(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DoHealthCheckCommand) RegisterFlags(cc *cobra.Command) {
}

// Run makes the HTTP request corresponding to the ShowJobCommand command.
func (cmd *ShowJobCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/job"
	}
	logger := goa.NewStdLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.UseLogger(context.Background(), logger)
	resp, err := c.ShowJob(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowJobCommand) RegisterFlags(cc *cobra.Command) {
}
