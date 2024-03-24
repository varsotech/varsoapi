package news

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/varsotech/varsoapi/src/services/app/internal/modules/organization"
	"github.com/varsotech/varsoapi/src/services/app/internal/modules/person"
	"github.com/varsotech/varsoapi/src/services/app/internal/modules/rssfeed"
	"github.com/varsotech/varsoapi/src/services/app/internal/modules/rssitem"
)

func SyncRSSFeeds(ctx context.Context) error {
	err := syncRSSFeeds(ctx)
	if err != nil {
		logrus.WithError(err).Errorf("failed initial syncing of rss feeds")
	}

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := syncRSSFeeds(ctx)
			if err != nil {
				logrus.WithError(err).Errorf("failed syncing rss feeds")
				continue
			}
		}

	}
}

func syncRSSFeeds(ctx context.Context) error {
	orgs, err := organization.GetOrganizations(ctx)
	if err != nil {
		return errors.Wrapf(err, "failed getting orgs for sync")
	}

	for _, org := range orgs {
		feeds, err := rssfeed.GetByOrganization(ctx, org.ID)
		if err != nil {
			logrus.WithError(err).WithField("orgUUID", org.ID.String()).Errorf("failed getting feeds by org for rss feed sync")
			continue
		}

		for _, feed := range feeds {
			goFeed, err := Parse(feed.RssFeedURL)
			if err != nil {
				logrus.WithError(err).WithField("rssFeedURL", feed.RssFeedURL).Errorf("failed parsing go feed")
				continue
			}

			// Upsert authors
			authorEmailToName := map[string]string{}
			for _, item := range goFeed.Items {
				for _, author := range item.Authors {
					authorEmailToName[author.Email] = author.Name
				}
			}

			authors, err := person.UpsertBulk(ctx, authorEmailToName)
			if err != nil {
				logrus.WithError(err).WithField("rssFeedURL", feed.RssFeedURL).Errorf("failed to bulk upsert authors")
				continue
			}

			err = rssfeed.Upsert(ctx, feed.RssFeedURL, org.ID)
			if err != nil {
				logrus.WithError(err).WithField("rssFeedURL", feed.RssFeedURL).Errorf("failed upserting rss feed for sync")
				continue
			}

			// TODO: Optimize by bulking
			for _, item := range goFeed.Items {
				imageUrl := ""
				imageTitle := ""
				if item.Image != nil {
					imageUrl = item.Image.URL
					imageTitle = item.Image.Title
				}
				err = rssitem.Upsert(ctx, feed.ID, item.GUID, item.Title, item.Description, item.Content, item.Link, item.Links, item.PublishedParsed, item.UpdatedParsed, imageUrl, imageTitle, item.Categories, authors)
				if err != nil {
					logrus.WithError(err).WithField("itemGUID", item.GUID).Errorf("failed upserting rss item")
				}
			}
		}
	}

	return nil
}
