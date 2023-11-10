package client

import (
	"fmt"

	"github.com/varsotech/varsoapi/src/common/api"
	"github.com/varsotech/varsoapi/src/services/auth/client/models"
)

func (c *Client) GetUserByDiscordId(discordId string) (*models.GetUserByDiscordIdResponse, *api.Error) {
	return api.SendRequest[any, models.GetUserByDiscordIdResponse](c.Client, c.AuthHeader, "GET", fmt.Sprintf("auth/user/admininspect/%s", discordId), nil)
}
