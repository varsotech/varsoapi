package login

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/varsotech/varsoapi/src/common/api"
	"github.com/varsotech/varsoapi/src/services/auth/client"
	"github.com/varsotech/varsoapi/src/services/auth/client/models"
	"github.com/varsotech/varsoapi/src/services/auth/internal/config"
	"github.com/varsotech/varsoapi/src/services/auth/internal/ent"
	"github.com/varsotech/varsoapi/src/services/auth/internal/ent/build"
	"github.com/varsotech/varsoapi/src/services/auth/internal/ent/build/user"
	"github.com/varsotech/varsoapi/src/services/auth/internal/ent/mixins"
	"golang.org/x/crypto/pbkdf2"
)

func Login(_ *api.Writer, r *http.Request, _ httprouter.Params, _ *api.JWT, request models.LoginRequest) (*models.LoginResponse, *api.Error) {
	// Determine if provided username or email for later querying
	logrus.WithField("request.UsernameOrEmail", request.UsernameOrEmail).Info("Login request")
	userPredicate := user.Username(request.UsernameOrEmail)
	if strings.Contains(request.UsernameOrEmail, "@") {
		userPredicate = user.Email(request.UsernameOrEmail)
	}

	user, err := ent.Database.User.Query().Where(userPredicate).Only(mixins.SkipBanFilter(r.Context()))
	if err != nil {
		if _, ok := err.(*build.NotFoundError); ok {
			return nil, api.NewUnauthorizedError(err, "user not found")
		}
		return nil, api.NewInternalError(err, "failed querying db for login")
	}

	salt, err := base64.StdEncoding.DecodeString(user.Salt)
	if err != nil {
		return nil, api.NewInternalError(err, "failed decoding salt")
	}

	// Hash the password using PBKDF2
	hashedPassword := pbkdf2.Key([]byte(request.Password), salt, config.Iterations, config.PasswordHashLength, sha256.New)
	hashedPasswordString := base64.StdEncoding.EncodeToString(hashedPassword)

	if user.Password != hashedPasswordString {
		return nil, api.NewUnauthorizedError(nil, "password is incorrect")
	}

	if !user.BanTime.IsZero() {
		return nil, api.NewUnauthorizedError(nil, "user is banned, can't login").WithCode(client.YouAreBanned)
	}

	token, err := api.MarshalJWT(user.UUID.String(), api.User.AccessLevel)
	if err != nil {
		return nil, api.NewInternalError(err, "failed creating JWT token")
	}

	// Return the signed JWT
	return &models.LoginResponse{Token: token}, nil
}
