package routes

import (
	"net/http"
	"slices"

	"github.com/julienschmidt/httprouter"
	"github.com/varsotech/varsoapi/src/common/api"
	"github.com/varsotech/varsoapi/src/services/app/client/models"
	"github.com/varsotech/varsoapi/src/services/app/internal/modules/news"
	"github.com/varsotech/varsoapi/src/services/app/internal/modules/organization"
)

func GetNews(w *api.Writer, r *http.Request, p httprouter.Params, j *api.JWT) (*models.GetNewsResponse, *api.Error) {
	orgs, err := organization.GetOrganizationsWithFeedItems(r.Context())
	if err != nil {
		return nil, api.NewInternalError(err, "failed getting organizations")
	}

	response := models.GetNewsResponse{
		Organizations: map[string]*models.Organization{},
		Featured:      nil,
		Latest:        &models.RSSFeed{},
	}

	// Populate organizations
	for _, org := range orgs {
		response.Organizations[org.ID.String()] = organization.TranslateOrganization(org)
	}

	// Populate latest articles
	for _, org := range orgs {
		for _, feed := range org.Edges.Feeds {
			response.Latest.Items = append(response.Latest.Items, news.TranslateRSSItems(feed.Edges.Items, org.ID)...)
		}
	}

	slices.SortFunc(response.Latest.Items, func(item1, item2 *models.RSSItem) int {
		if item1.PublishDate.AsTime().After(item2.PublishDate.AsTime()) {
			return -1
		}
		return 1
	})

	return &response, nil
}
