// Code generated by goa v3.0.4, DO NOT EDIT.
//
// some-function client
//
// Command:
// $ goa gen github.com/seriwb/front-bff-msa/msa/design

package somefunction

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "some-function" service client.
type Client struct {
	AddEndpoint goa.Endpoint
}

// NewClient initializes a "some-function" service client given the endpoints.
func NewClient(add goa.Endpoint) *Client {
	return &Client{
		AddEndpoint: add,
	}
}

// Add calls the "add" endpoint of the "some-function" service.
func (c *Client) Add(ctx context.Context, p *AddPayload) (res int, err error) {
	var ires interface{}
	ires, err = c.AddEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}
