package catalog

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/varsotech/varsoapi/src/common/api"
	"github.com/varsotech/varsoapi/src/services/app/client/models"
)

func Search(_ *api.Writer, r *http.Request, params httprouter.Params, _ *api.JWT, request *models.CatalogSearchRequest) (*models.CatalogSearchResponse, *api.Error) {
	var returnedAlbums []*models.Album

	return &models.CatalogSearchResponse{
		Albums: returnedAlbums,
	}, nil
}
