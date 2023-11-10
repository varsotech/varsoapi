package easyregister

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/varsotech/varsoapi/src/common/api"
	"github.com/varsotech/varsoapi/src/services/auth/client/models"
	"github.com/varsotech/varsoapi/src/services/auth/internal/config"
	"github.com/varsotech/varsoapi/src/services/auth/internal/ent"
	"github.com/varsotech/varsoapi/src/services/auth/internal/ent/build"
)

func EasyRegister(_ *api.Writer, r *http.Request, _ httprouter.Params, _ *api.JWT, request models.EasyRegisterRequest) (*models.EasyRegisterResponse, *api.Error) {
	if len(request.Name) < config.NameMinLength {
		return nil, api.NewBadRequestError(nil, "password too short")
	}

	user, err := ent.Database.User.Create().
		SetName(request.Name).
		Save(r.Context())
	if err != nil {
		if _, ok := err.(*build.ConstraintError); ok {
			return nil, api.NewBadRequestError(err, "user might already exist")
		}
		if _, ok := err.(*build.ValidationError); ok {
			return nil, api.NewBadRequestError(err, "invalid registration input")
		}
		return nil, api.NewInternalError(err, "failed creating user in database")
	}

	token, err := api.MarshalJWT(user.UUID.String(), api.User.AccessLevel)
	if err != nil {
		return nil, api.NewInternalError(err, "failed creating JWT token")
	}

	return &models.EasyRegisterResponse{Token: token}, nil
}
