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

func GetOrganizationsWithFeedItems(ctx context.Context) ([]*build.Organization, error) {
	posts, err := ent.Database.Organization.Query().
		WithFeeds(func(rq *build.RSSFeedQuery) {
			rq.WithItems(func(niq *build.NewsItemQuery) {
				niq.WithAuthors()
			})
		}).
		All(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed getting orgs from db")
	}

	return posts, nil
}

func UpsertOrganization(ctx context.Context, orgUUID uuid.UUID, uniqueName, name, description, websiteUrl string) error {
	err := ent.Database.Organization.Create().
		SetID(orgUUID).
		SetName(name).
		SetUniqueName(uniqueName).
		SetDescription(description).
		SetWebsiteURL(websiteUrl).
		OnConflictColumns(organization.FieldID).
		UpdateName().
		UpdateDescription().
		UpdateWebsiteURL().
		Exec(ctx)
	if err != nil {
		return errors.Wrapf(err, "failed upserting organization in db")
	}

	return nil
}
