package rssfeed

import (
	"context"

	"github.com/google/uuid"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build/organization"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build/rssfeed"
)

func Upsert(ctx context.Context, rssFeedUrl string, orgUuid uuid.UUID) error {
	return ent.Database.RSSFeed.Create().
		SetRssFeedURL(rssFeedUrl).
		SetOrganizationID(orgUuid).
		OnConflictColumns(rssfeed.FieldRssFeedURL).
		UpdateNewValues().
		Exec(ctx)
}

func GetByOrganization(ctx context.Context, orgUuid uuid.UUID) ([]*build.RSSFeed, error) {
	return ent.Database.RSSFeed.
		Query().
		Where(
			rssfeed.HasOrganizationWith(organization.ID(orgUuid)),
		).All(ctx)
}
