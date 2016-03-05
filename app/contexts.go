//************************************************************************//
// API "croner": Application Contexts
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
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
)

// ShowJobContext provides the job show action context.
type ShowJobContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewShowJobContext parses the incoming request URL and body, performs validations and creates the
// context used by the job controller show action.
func NewShowJobContext(ctx context.Context) (*ShowJobContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := ShowJobContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowJobContext) OK(r *Job) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.rightscale.croner.job+json")
	return ctx.ResponseData.Send(ctx.Context, 200, r)
}
