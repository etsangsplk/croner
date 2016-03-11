//************************************************************************//
// API "croner": Application Media Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/rightscale/croner
// --design=github.com/rightscale/croner/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"time"

	"github.com/goadesign/goa"
)

// Execution media type.
//
// Identifier: application/vnd.rightscale.croner.execution+json
type Execution struct {
	// Execution exit status if finished
	ExitStatus *int `json:"exit_status,omitempty" xml:"exit_status,omitempty"`
	// Execution finished at timestamp if finished
	FinishedAt *time.Time `json:"finished_at,omitempty" xml:"finished_at,omitempty"`
	// Execution OS pid
	Pid int `json:"pid" xml:"pid"`
	// Execution started at timestamp
	StartedAt time.Time `json:"started_at" xml:"started_at"`
	// Execution stderr output if finished and if not empty
	Stderr *string `json:"stderr,omitempty" xml:"stderr,omitempty"`
}

// Validate validates the Execution media type instance.
func (mt *Execution) Validate() (err error) {

	return
}

// ExecutionCollection media type is a collection of Execution.
//
// Identifier: application/vnd.rightscale.croner.execution+json; type=collection
type ExecutionCollection []*Execution

// Validate validates the ExecutionCollection media type instance.
func (mt ExecutionCollection) Validate() (err error) {
	return
}

// Job media type.
//
// Identifier: application/vnd.rightscale.croner.job+json
type Job struct {
	// scheduled command
	Cmd string `json:"cmd" xml:"cmd"`
	// last execution
	Last *Execution `json:"last,omitempty" xml:"last,omitempty"`
	// currently running executions if any
	Running ExecutionCollection `json:"running,omitempty" xml:"running,omitempty"`
	// cron schedule spec
	Schedule string `json:"schedule" xml:"schedule"`
}

// Validate validates the Job media type instance.
func (mt *Job) Validate() (err error) {
	if mt.Cmd == "" {
		err = goa.MissingAttributeError(`response`, "cmd", err)
	}
	if mt.Schedule == "" {
		err = goa.MissingAttributeError(`response`, "schedule", err)
	}
	return
}
