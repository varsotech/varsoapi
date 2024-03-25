package rssitem

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build"
	"github.com/varsotech/varsoapi/src/services/app/internal/ent/build/newsitem"
)

func Get(ctx context.Context, itemUUID uuid.UUID) (*build.NewsItem, error) {
	return ent.Database.NewsItem.Get(ctx, itemUUID)
}

func ToggleBlur(ctx context.Context, itemUUID uuid.UUID) error {
	newsItem, err := Get(ctx, itemUUID)
	if err != nil {
		return err
	}

	return ent.Database.NewsItem.Update().Where(newsitem.ID(itemUUID)).SetBlur(!newsItem.Blur).Exec(ctx)
}

func Upsert(ctx context.Context, feedUUID uuid.UUID, rssGuid, title, description, content, link string, links []string, publishTime, updateTime *time.Time, imageUrl, imageTitle string, categories []string, authors []*build.RSSAuthor) error {
	if rssGuid == "" {
		return fmt.Errorf("rss guid cannot be empty")
	}

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
		SetCategories(categories)

	if publishTime != nil {
		creator.SetItemPublishTime(*publishTime)
	}

	if updateTime != nil {
		creator.SetItemUpdateTime(*updateTime)
	}

	// Don't override - blur
	err := creator.OnConflictColumns(newsitem.FieldRssGUID).
		UpdateTitle().
		UpdateDescription().
		UpdateContent().
		UpdateLink().
		UpdateLinks().
		UpdateImageURL().
		UpdateImageTitle().
		UpdateCategories().
		Exec(ctx)
	if err != nil {
		return err
	}

	err = ent.Database.NewsItem.Update().ClearAuthors().AddAuthors(authors...).Where(newsitem.RssGUID(rssGuid)).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
