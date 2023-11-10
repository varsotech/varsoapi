package register

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/varsotech/varsoapi/src/common/api"
	"github.com/varsotech/varsoapi/src/services/auth/client/models"
	"github.com/varsotech/varsoapi/src/services/auth/internal/config"
	"github.com/varsotech/varsoapi/src/services/auth/internal/ent"
	"github.com/varsotech/varsoapi/src/services/auth/internal/ent/build"
	"golang.org/x/crypto/pbkdf2"
)

func Register(_ *api.Writer, r *http.Request, _ httprouter.Params, _ *api.JWT, request models.RegisterRequest) (*models.RegisterResponse, *api.Error) {
	// Validate password as it cannot be validated with Ent because it is not saved plaintext
	if len(request.Password) < config.PasswordMinLength {
		return nil, api.NewBadRequestError(nil, "password too short")
	}

	// Generate a random salt
	salt := make([]byte, config.SaltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, api.NewInternalError(err, "failed generating salt")
	}

	// Hash the password using PBKDF2
	hashedPassword := pbkdf2.Key([]byte(request.Password), salt, config.Iterations, config.PasswordHashLength, sha256.New)

	// Encode the salt and hashed password as base64 strings
	hashedPasswordString := base64.StdEncoding.EncodeToString(hashedPassword)
	saltString := base64.StdEncoding.EncodeToString(salt)

	user, err := ent.Database.User.Create().
		SetEmail(request.Email).
		SetUsername(request.Username).
		SetPassword(hashedPasswordString).
		SetSalt(saltString).
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

	return &models.RegisterResponse{Token: token}, nil
}
