package role

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/varsotech/varsoapi/src/services/auth/client"
	"github.com/varsotech/varsoapi/src/services/auth/client/models"
)

const (
	KeyAdmin = "admin"
)

var roles = []*models.Role{
	{
		Uuid: "3ec70e30-afa7-4e85-b161-c7a71e8c4e0a",
		Key:  KeyAdmin,
		Permissions: []string{
			client.KeyNewsItemImageBlur,
		},
	},
}

func Initialize(ctx context.Context) error {
	for _, role := range roles {
		err := Upsert(ctx, uuid.MustParse(role.Uuid), role.Key, role.Permissions)
		if err != nil {
			return errors.Wrapf(err, "failed upserting role")
		}
	}

	return nil
}
