package user

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/varsotech/varsoapi/src/common/api"
	"github.com/varsotech/varsoapi/src/services/auth/client"
	"github.com/varsotech/varsoapi/src/services/auth/client/models"
	"github.com/varsotech/varsoapi/src/services/auth/internal/ent"
	"github.com/varsotech/varsoapi/src/services/auth/internal/ent/build"
	"github.com/varsotech/varsoapi/src/services/auth/internal/ent/build/user"
	"github.com/varsotech/varsoapi/src/services/auth/internal/ent/mixins"
	"github.com/varsotech/varsoapi/src/services/auth/internal/translate"
)

func AdminInspectUser(_ *api.Writer, r *http.Request, params httprouter.Params, _ *api.JWT) (*models.GetUserByDiscordIdResponse, *api.Error) {
	discordUserId := params.ByName("discordid")

	user, err := ent.Database.User.Query().Where(user.DiscordUserIDEQ(discordUserId)).Only(mixins.SkipBanFilter(r.Context()))
	if err != nil {
		if _, ok := err.(*build.NotFoundError); ok {
			return nil, api.NewBadRequestError(err, "user not found").WithCode(client.UserNotFound)
		}
		return nil, api.NewInternalError(err, "failed querying db for getting user by discord id")
	}

	return &models.GetUserByDiscordIdResponse{User: translate.TranslateUser(user)}, nil
}
