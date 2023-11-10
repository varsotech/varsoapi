package client

import (
	"fmt"

	"github.com/varsotech/varsoapi/src/common/api"
	"github.com/varsotech/varsoapi/src/services/auth/client/models"
)

func (c *Client) BanUser(uuid string) (*models.BanUserResponse, *api.Error) {
	return api.SendRequest[models.BanUserRequest, models.BanUserResponse](c.Client, c.AuthHeader, "POST", fmt.Sprintf("auth/user/ban/%s", uuid), &models.BanUserRequest{})
}
