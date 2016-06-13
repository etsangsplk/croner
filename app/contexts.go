//************************************************************************//
// API "croner": Application Contexts
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --design=github.com/rightscale/croner/design
// --out=$(GOPATH)/src/github.com/rightscale/croner
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
)

// DoHealthCheckContext provides the health_check do action context.
type DoHealthCheckContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
}

// NewDoHealthCheckContext parses the incoming request URL and body, performs validations and creates the
// context used by the health_check controller do action.
func NewDoHealthCheckContext(ctx context.Context, service *goa.Service) (*DoHealthCheckContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := DoHealthCheckContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *DoHealthCheckContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// ShowJobContext provides the job show action context.
type ShowJobContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
}

// NewShowJobContext parses the incoming request URL and body, performs validations and creates the
// context used by the job controller show action.
func NewShowJobContext(ctx context.Context, service *goa.Service) (*ShowJobContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := ShowJobContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowJobContext) OK(r *Job) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.rightscale.croner.job+json")
	return ctx.Service.Send(ctx.Context, 200, r)
}
