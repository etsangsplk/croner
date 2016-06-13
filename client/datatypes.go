//************************************************************************//
// User Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --design=github.com/rightscale/croner/design
// --out=$(GOPATH)/src/github.com/rightscale/croner
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import (
	"net/http"
	"time"
)

// A cron job together with information on the last execution
type Job struct {
	// scheduled command
	Cmd string `json:"cmd" xml:"cmd" form:"cmd"`
	// last execution
	Last *Execution `json:"last,omitempty" xml:"last,omitempty" form:"last,omitempty"`
	// currently running executions if any
	Running ExecutionCollection `json:"running,omitempty" xml:"running,omitempty" form:"running,omitempty"`
	// cron schedule spec
	Schedule string `json:"schedule" xml:"schedule" form:"schedule"`
}

// DecodeJob decodes the Job instance encoded in resp body.
func (c *Client) DecodeJob(resp *http.Response) (*Job, error) {
	var decoded Job
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A job execution
type Execution struct {
	// Execution exit status if finished
	ExitStatus *int `json:"exit_status,omitempty" xml:"exit_status,omitempty" form:"exit_status,omitempty"`
	// Execution finished at timestamp if finished
	FinishedAt *time.Time `json:"finished_at,omitempty" xml:"finished_at,omitempty" form:"finished_at,omitempty"`
	// Execution OS pid
	Pid int `json:"pid" xml:"pid" form:"pid"`
	// Execution started at timestamp
	StartedAt time.Time `json:"started_at" xml:"started_at" form:"started_at"`
	// Execution stderr output if finished and if not empty
	Stderr *string `json:"stderr,omitempty" xml:"stderr,omitempty" form:"stderr,omitempty"`
}

// ExecutionCollection media type is a collection of Execution.
type ExecutionCollection []*Execution
