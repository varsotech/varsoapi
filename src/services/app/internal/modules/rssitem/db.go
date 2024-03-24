package rssitem

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build/newsitem"
)

func Upsert(ctx context.Context, feedUUID uuid.UUID, rssGuid, title, description, content, link string, links []string, publishTime, updateTime *time.Time, imageUrl, imageTitle string, categories []string, authors []*build.Person) error {
	creator := ent.Database.NewsItem.
		Create().
		SetFeedID(feedUUID).
		SetRssGUID(rssGuid).
		SetTitle(title).
		SetDescription(description).
		SetContent(content).
		SetLink(link).
		SetLinks(links).
		SetImageURL(imageUrl).
		SetImageTitle(imageTitle).
		SetCategories(categories).
		AddAuthors(authors...).
		OnConflictColumns(newsitem.FieldRssGUID)

	if publishTime != nil {
		creator.SetItemPublishTime(*publishTime)
	}

	if updateTime != nil {
		creator.SetItemUpdateTime(*updateTime)
	}

	return creator.Exec(ctx)
}
