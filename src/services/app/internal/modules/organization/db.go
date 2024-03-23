package organization

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build/organization"
)

func GetOrganizations(ctx context.Context) ([]*build.Organization, error) {
	posts, err := ent.Database.Organization.Query().All(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed getting orgs from db")
	}

	return posts, nil
}

func UpsertOrganization(ctx context.Context, orgUUID uuid.UUID, name, websiteUrl, rssFeedUrl string) error {
	err := ent.Database.Organization.Create().
		SetID(orgUUID).
		SetName(name).
		SetWebsiteURL(websiteUrl).
		SetRssFeedURL(rssFeedUrl).
		OnConflictColumns(organization.FieldID).
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return errors.Wrapf(err, "failed upserting organization in db")
	}

	return nil
}
