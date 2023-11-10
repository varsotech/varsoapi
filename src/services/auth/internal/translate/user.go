package translate

import (
	"github.com/varsotech/varsoapi/src/services/auth/client/models"
	"github.com/varsotech/varsoapi/src/services/auth/internal/ent/build"
)

func TranslateUser(user *build.User) *models.User {
	return &models.User{
		Uuid:          user.UUID.String(),
		DiscordUserId: user.DiscordUserID,
		Banned:        !user.BanTime.IsZero(),
	}
}
