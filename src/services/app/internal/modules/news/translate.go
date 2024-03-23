package news

import (
	"github.com/mmcdole/gofeed"
	"github.com/varsotech/varsoapi/src/services/app/client/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TranslateRSSFeed(feed *gofeed.Feed) *models.RSSFeed {
	if feed == nil {
		return &models.RSSFeed{}
	}

	return &models.RSSFeed{
		Title: feed.Title,
		Items: TranslateRSSItems(feed.Items),
	}
}

func TranslateRSSItems(items []*gofeed.Item) []*models.RSSItem {
	var translated []*models.RSSItem
	for _, item := range items {
		translated = append(translated, TranslateRSSItem(item))
	}
	return translated
}

func TranslateRSSItem(item *gofeed.Item) *models.RSSItem {
	if item == nil {
		return &models.RSSItem{}
	}

	var updateDate *timestamppb.Timestamp
	if item.UpdatedParsed != nil {
		updateDate = timestamppb.New(*item.UpdatedParsed)
	}

	var publishDate *timestamppb.Timestamp
	if item.UpdatedParsed != nil {
		publishDate = timestamppb.New(*item.PublishedParsed)
	}

	return &models.RSSItem{
		Title:       item.Title,
		Description: item.Description,
		Content:     item.Content,
		Link:        item.Link,
		Links:       item.Links,
		UpdateDate:  updateDate,
		PublishDate: publishDate,
		Authors:     TranslateRSSAuthors(item.Authors),
		Uuid:        item.GUID,
		Image:       TranslateRSSImage(item.Image),
		Categories:  item.Categories,
	}
}

func TranslateRSSAuthors(authors []*gofeed.Person) []*models.RSSAuthor {
	var translated []*models.RSSAuthor
	for _, author := range authors {
		translated = append(translated, TranslateRSSAuthor(author))
	}
	return translated
}

func TranslateRSSAuthor(author *gofeed.Person) *models.RSSAuthor {
	if author == nil {
		return &models.RSSAuthor{}
	}

	return &models.RSSAuthor{
		Email: author.Email,
		Name:  author.Name,
	}
}

func TranslateRSSImage(image *gofeed.Image) *models.RSSImage {
	if image == nil {
		return &models.RSSImage{}
	}

	return &models.RSSImage{
		Url:   image.URL,
		Title: image.Title,
	}
}
