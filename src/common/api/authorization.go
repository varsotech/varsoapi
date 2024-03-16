package api

import (
	"fmt"
	"net/http"
)

type AccessLevel = int32

const (
	publicAccessLevel AccessLevel = iota
	internalAccessLevel
	userAccessLevel
)

type ProtectFunc = func(*http.Request, *JWT) (*JWT, *Error)

type Authorization struct {
	AccessLevel  AccessLevel
	protectFuncs []ProtectFunc
}

func newAuthorization(accessLevel AccessLevel, protectFunc ProtectFunc) Authorization {
	return Authorization{
		AccessLevel:  accessLevel,
		protectFuncs: []ProtectFunc{protectFunc},
	}
}

func (a *Authorization) Protect(r *http.Request) (jwt *JWT, aErr *Error) {
	for _, f := range a.protectFuncs {
		jwt, aErr = f(r, jwt)
		if aErr != nil {
			return nil, aErr
		}
	}

	return
}

func (a *Authorization) AddMiddleware(f ProtectFunc) {
	a.protectFuncs = append(a.protectFuncs, f)
}

func getJWT(r *http.Request) (*JWT, *Error) {
	authHeader := r.Header.Get("Authorization")

	token, err := UnmarshalJWTFromHeader(authHeader)
	if err != nil {
		return nil, NewUnauthorizedError(err, "failed unmarshaling jwt from header")
	}

	return token, nil
}

var Public = newAuthorization(publicAccessLevel, func(r *http.Request, _ *JWT) (*JWT, *Error) {
	return nil, nil
})

var User = newAuthorization(userAccessLevel, func(r *http.Request, _ *JWT) (*JWT, *Error) {
	jwt, discoErr := getJWT(r)
	if discoErr != nil {
		return nil, discoErr
	}

	if jwt.AccessLevel != userAccessLevel && jwt.AccessLevel != internalAccessLevel {
		return nil, NewUnauthorizedError(nil, fmt.Sprintf("insufficient access level %d", jwt.AccessLevel))
	}

	return jwt, nil
})

var Internal = newAuthorization(internalAccessLevel, func(r *http.Request, _ *JWT) (*JWT, *Error) {
	jwt, discoErr := getJWT(r)
	if discoErr != nil {
		return nil, discoErr
	}

	if jwt.AccessLevel != internalAccessLevel {
		return nil, NewUnauthorizedError(nil, fmt.Sprintf("insufficient access level %d", jwt.AccessLevel))
	}

	return jwt, nil
})
