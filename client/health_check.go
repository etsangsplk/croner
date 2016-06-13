package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// DoHealthCheckPath computes a request path to the do action of health_check.
func DoHealthCheckPath() string {
	return fmt.Sprintf("/health-check")
}

// Health check
func (c *Client) DoHealthCheck(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDoHealthCheckRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDoHealthCheckRequest create the request corresponding to the do action endpoint of the health_check resource.
func (c *Client) NewDoHealthCheckRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
