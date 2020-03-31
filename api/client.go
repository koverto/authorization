package authorization

import (
	"github.com/micro/go-micro/v2/client"
)

// Name is the identifying name of the Authorization service.
const Name = "com.koverto.svc.authorization"

// Client defines a client for the Authorization service.
type Client struct {
	AuthorizationService
}

// NewClient creates a new client for the Authorization service.
func NewClient(client client.Client) *Client {
	return &Client{NewAuthorizationService(Name, client)}
}

// Name returns the name of the Authorization service.
func (c *Client) Name() string {
	return Name
}
