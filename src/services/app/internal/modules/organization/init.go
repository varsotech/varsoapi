package organization

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/varsotech/varsoapi/src/services/app/client/models"
)

var orgs = []models.Organization{
	models.Organization{
		Uuid:        "f97d83e7-2f08-4f8c-9ec8-98279851608c",
		UniqueName:  "972mag",
		Name:        "+972 Magazine",
		Description: "Independent commentary and news from Israel &#38; Palestine",
		WebsiteUrl:  "https://www.972mag.com",
		RssFeedUrl:  "https://www.972mag.com/feed",
	},
}

func Initialize(ctx context.Context) error {
	for _, org := range orgs {
		err := UpsertOrganization(ctx, uuid.MustParse(org.Uuid), org.Name, org.WebsiteUrl, org.RssFeedUrl)
		if err != nil {
			return errors.Wrapf(err, "failed upserting org with uuid '%s'", org.Uuid)
		}
	}

	return nil
}
