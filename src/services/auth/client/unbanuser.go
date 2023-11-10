package client

import (
	"fmt"

	"github.com/varsotech/varsoapi/src/common/api"
	"github.com/varsotech/varsoapi/src/services/auth/client/models"
)

func (c *Client) UnbanUser(uuid string) (*models.UnbanUserResponse, *api.Error) {
	return api.SendRequest[models.UnbanUserRequest, models.UnbanUserResponse](c.Client, c.AuthHeader, "POST", fmt.Sprintf("auth/user/unban/%s", uuid), &models.UnbanUserRequest{})
}
