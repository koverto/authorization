package authorization

import (
	"github.com/micro/go-micro/v2/client"
)

const Name = "com.koverto.svc.authorization"

type Client struct {
	AuthorizationService
}

func NewClient(client client.Client) *Client {
	return &Client{NewAuthorizationService(Name, client)}
}

func (c *Client) Name() string {
	return Name
}
