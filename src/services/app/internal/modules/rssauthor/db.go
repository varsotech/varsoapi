package rssauthor

import (
	"context"
	"fmt"

	"github.com/varsotech/varsoapi/src/services/app/internal/ent"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build/organization"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build/rssauthor"
)

func Upsert(ctx context.Context, org *build.Organization, name, email string) (*build.RSSAuthor, error) {
	if name == "" {
		return nil, fmt.Errorf("rss author name cannot be empty")
	}

	creator := ent.Database.RSSAuthor.Create().
		SetName(name).
		SetOrganization(org)

	if email != "" {
		creator.SetEmail(email)
	}

	err := creator.
		OnConflictColumns(rssauthor.FieldName, rssauthor.OrganizationColumn).
		UpdateEmail().
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return ent.Database.RSSAuthor.Query().Where(rssauthor.HasOrganizationWith(organization.ID(org.ID)), rssauthor.Name(name)).Only(ctx)
}
