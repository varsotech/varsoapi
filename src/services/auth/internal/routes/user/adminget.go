package user

import (
	"net/http"

	"github.com/google/uuid"
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

func AdminGetUser(_ *api.Writer, r *http.Request, params httprouter.Params, _ *api.JWT) (*models.GetUserByDiscordIdResponse, *api.Error) {
	uuidStr := params.ByName("uuid")
	uuid, err := uuid.Parse(uuidStr)
	if err != nil {
		return nil, api.NewInternalError(err, "failed parsing user uuid")
	}

	user, err := ent.Database.User.Query().Where(user.UUID(uuid)).Only(mixins.SkipBanFilter(r.Context()))
	if err != nil {
		if _, ok := err.(*build.NotFoundError); ok {
			return nil, api.NewBadRequestError(err, "user not found").WithCode(client.UserNotFound)
		}
		return nil, api.NewInternalError(err, "failed querying db for getting user by discord id")
	}

	return &models.GetUserByDiscordIdResponse{User: translate.TranslateUser(user)}, nil
}
