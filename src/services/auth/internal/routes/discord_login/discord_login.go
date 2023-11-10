package discord_login

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/varsotech/varsoapi/src/common/api"
	"github.com/varsotech/varsoapi/src/common/config"
	"github.com/varsotech/varsoapi/src/services/auth/client"
	"github.com/varsotech/varsoapi/src/services/auth/client/models"
	"github.com/varsotech/varsoapi/src/services/auth/internal/ent"
	"github.com/varsotech/varsoapi/src/services/auth/internal/ent/build"
	"github.com/varsotech/varsoapi/src/services/auth/internal/ent/build/user"
	"github.com/varsotech/varsoapi/src/services/auth/internal/ent/mixins"
)

// DiscordUserLogin requires the internal secret instead of credentials, but only provides User access level
func DiscordUserLogin(_ *api.Writer, r *http.Request, _ httprouter.Params, _ *api.JWT, request models.DiscordLoginRequest) (*models.DiscordLoginResponse, *api.Error) {
	if request.InternalLoginAuthSecret != config.InternalLoginAuthSecret {
		return nil, api.NewUnauthorizedError(nil, "bad internal auth secret")
	}

	user, err := ent.Database.User.Query().Where(user.DiscordUserID(request.DiscordUserId)).Only(mixins.SkipSoftDeleteFilter(mixins.SkipBanFilter(r.Context())))
	if err != nil {
		if !build.IsNotFound(err) {
			return nil, api.NewInternalError(err, "failed querying db for discord login")
		}

		user, err = ent.Database.User.Create().
			SetDiscordUserID(request.DiscordUserId).
			Save(context.Background())
		if err != nil {
			return nil, api.NewInternalError(err, "failed creating user with discord login")
		}
	}

	if !user.BanTime.IsZero() {
		return nil, api.NewUnauthorizedError(nil, "user is banned").WithCode(client.YouAreBanned)
	}

	token, err := api.MarshalJWT(user.UUID.String(), api.Internal.AccessLevel)
	if err != nil {
		return nil, api.NewInternalError(err, "failed creating JWT token")
	}

	return &models.DiscordLoginResponse{Token: token}, nil
}
