package main

import (
	"github.com/rightscale/croner/client"
	"github.com/spf13/cobra"
)

type (
	// IndexJobsCommand is the command line data structure for the index action of jobs
	IndexJobsCommand struct {
	}
)

// Run makes the HTTP request corresponding to the IndexJobsCommand command.
func (cmd *IndexJobsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/jobs"
	}
	resp, err := c.IndexJobs(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *IndexJobsCommand) RegisterFlags(cc *cobra.Command) {
}
