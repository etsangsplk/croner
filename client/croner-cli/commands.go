package main

import (
	"github.com/rightscale/croner/client"
	"github.com/spf13/cobra"
)

type (
	// ShowJobCommand is the command line data structure for the show action of job
	ShowJobCommand struct {
	}
)

// Run makes the HTTP request corresponding to the ShowJobCommand command.
func (cmd *ShowJobCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/job"
	}
	resp, err := c.ShowJob(path)
	if err != nil {
		return err
	}
	HandleResponse(c, resp)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowJobCommand) RegisterFlags(cc *cobra.Command) {
}
