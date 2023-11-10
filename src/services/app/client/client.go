package client

import (
	"github.com/varsotech/varsoapi/src/services/auth/client"
)

type Client struct {
	c *client.Client
}

func NewClient(client *client.Client) *Client {
	return &Client{
		c: client,
	}
}
