package user

import (
	"github.com/varsotech/varsoapi/src/services/auth/client/models"
	"github.com/varsotech/varsoapi/src/services/auth/internal/ent/build"
)

func TranslateUser(user *build.User) *models.User {
	return &models.User{
		Uuid:   user.ID.String(),
		Banned: !user.BanTime.IsZero(),
	}
}
