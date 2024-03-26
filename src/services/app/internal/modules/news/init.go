package news

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/varsotech/varsoapi/src/services/app/client/models"
	"github.com/varsotech/varsoapi/src/services/app/internal/modules/organization"
	"github.com/varsotech/varsoapi/src/services/app/internal/modules/rssfeed"
)

var orgs = []struct {
	*models.Organization
	RssFeedUrl string
}{
	{
		Organization: &models.Organization{
			Uuid:        "f97d83e7-2f08-4f8c-9ec8-98279851608c",
			UniqueName:  "972mag",
			Name:        "+972 Magazine",
			Description: "Independent commentary and news from Israel &#38; Palestine",
			WebsiteUrl:  "https://www.972mag.com",
		},
		RssFeedUrl: "https://www.972mag.com/feed",
	},

	{
		Organization: &models.Organization{
			Uuid:        "278ae4c8-b190-4543-93ec-4e28d558633e",
			UniqueName:  "btselem",
			Name:        "Btselem",
			Description: "Independent commentary and news from Israel & Palestine",
			WebsiteUrl:  "https://www.btselem.org",
		},
		RssFeedUrl: "https://www.btselem.org/rss",
	},

	{
		Organization: &models.Organization{
			Uuid:        "ad470590-8199-4906-b189-782acb7e672d",
			UniqueName:  "mondoweiss",
			Name:        "Mondoweiss",
			Description: "News & Opinion about Palestine, Israel & the United States",
			WebsiteUrl:  "https://mondoweiss.net",
		},
		RssFeedUrl: "https://mondoweiss.net/rss",
	},

	{
		Organization: &models.Organization{
			Uuid:        "3b613eb0-2d13-4a74-bd94-4acb57c8fef9",
			UniqueName:  "bds",
			Name:        "BDS",
			Description: "Boycott, Divestment, Sanctions (BDS) is a Palestinian-led movement for freedom, justice and equality.",
			WebsiteUrl:  "https://bdsmovement.net",
		},
		RssFeedUrl: "https://bdsmovement.net/rss-feed.xml",
	},
}

func Initialize(ctx context.Context) error {
	for _, org := range orgs {
		orgUuid := uuid.MustParse(org.Uuid)

		err := organization.UpsertOrganization(ctx, orgUuid, org.UniqueName, org.Name, org.Description, org.WebsiteUrl)
		if err != nil {
			return errors.Wrapf(err, "failed upserting org with uuid '%s'", org.Uuid)
		}

		if org.RssFeedUrl != "" {
			err = rssfeed.Upsert(ctx, org.RssFeedUrl, orgUuid)
			if err != nil {
				return errors.Wrapf(err, "failed upserting RSS feed")
			}
		}
	}

	return nil
}
