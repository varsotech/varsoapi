package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mmcdole/gofeed"
	"github.com/sirupsen/logrus"
	"github.com/varsotech/varsoapi/src/common/api"
	"github.com/varsotech/varsoapi/src/services/app/client/models"
	"github.com/varsotech/varsoapi/src/services/app/internal/modules/news"
	"github.com/varsotech/varsoapi/src/services/app/internal/modules/organization"
)

func GetNews(w *api.Writer, r *http.Request, p httprouter.Params, j *api.JWT) (*models.GetNewsResponse, *api.Error) {
	orgs, err := organization.GetOrganizations(r.Context())
	if err != nil {
		return nil, api.NewInternalError(err, "failed getting organizations")
	}

	fp := gofeed.NewParser()
	fp.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36"

	response := models.GetNewsResponse{
		Feeds: []*models.RSSFeed{},
	}

	for _, org := range orgs {
		feed, err := fp.ParseURL(org.RssFeedURL)
		if err != nil {
			logrus.WithField("uniqueName", org.UniqueName).WithError(err).Errorf("failed to parse rss feed")
			continue
		}

		response.Feeds = append(response.Feeds, news.TranslateRSSFeed(feed))
	}

	return &response, nil
}
