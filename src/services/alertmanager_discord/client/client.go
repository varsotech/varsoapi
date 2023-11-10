package client

import (
	"github.com/varsotech/varsoapi/src/services/auth/client"
)

type Client struct {
	*client.Client
}

func NewClient(client *client.Client) *Client {
	return &Client{
		Client: client,
	}
}
