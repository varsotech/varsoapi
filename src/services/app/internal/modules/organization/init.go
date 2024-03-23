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

	models.Organization{
		Uuid:        "278ae4c8-b190-4543-93ec-4e28d558633e",
		UniqueName:  "btselem",
		Name:        "Btselem",
		Description: "Independent commentary and news from Israel & Palestine",
		WebsiteUrl:  "https://www.btselem.org",
		RssFeedUrl:  "https://www.btselem.org/rss",
	},

	models.Organization{
		Uuid:        "ad470590-8199-4906-b189-782acb7e672d",
		UniqueName:  "mondoweiss",
		Name:        "Mondoweiss",
		Description: "News & Opinion about Palestine, Israel & the United States",
		WebsiteUrl:  "https://mondoweiss.net",
		RssFeedUrl:  "https://mondoweiss.net/rss",
	},
}

func Initialize(ctx context.Context) error {
	for _, org := range orgs {
		err := UpsertOrganization(ctx, uuid.MustParse(org.Uuid), org.Name, org.Description, org.WebsiteUrl, org.RssFeedUrl)
		if err != nil {
			return errors.Wrapf(err, "failed upserting org with uuid '%s'", org.Uuid)
		}
	}

	return nil
}
