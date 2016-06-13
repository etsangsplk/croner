package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// ShowJobPath computes a request path to the show action of job.
func ShowJobPath() string {
	return fmt.Sprintf("/job")
}

// Show information about cron job
func (c *Client) ShowJob(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowJobRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowJobRequest create the request corresponding to the show action endpoint of the job resource.
func (c *Client) NewShowJobRequest(ctx context.Context, path string) (*http.Request, error) {
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
