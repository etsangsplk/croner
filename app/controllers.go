//************************************************************************//
// API "croner": Application Controllers
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
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// HealthCheckController is the controller interface for the HealthCheck actions.
type HealthCheckController interface {
	goa.Muxer
	Do(*DoHealthCheckContext) error
}

// MountHealthCheckController "mounts" a HealthCheck resource controller on the given service.
func MountHealthCheckController(service *goa.Service, ctrl HealthCheckController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDoHealthCheckContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Do(rctx)
	}
	service.Mux.Handle("GET", "/health-check", ctrl.MuxHandler("Do", h, nil))
	service.LogInfo("mount", "ctrl", "HealthCheck", "action", "Do", "route", "GET /health-check")
}

// JobController is the controller interface for the Job actions.
type JobController interface {
	goa.Muxer
	Show(*ShowJobContext) error
}

// MountJobController "mounts" a Job resource controller on the given service.
func MountJobController(service *goa.Service, ctrl JobController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowJobContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/job", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Job", "action", "Show", "route", "GET /job")
}
