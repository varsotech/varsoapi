package routes

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/varsotech/varsoapi/src/common/api"
	"github.com/varsotech/varsoapi/src/services/app/client/models"
	"github.com/varsotech/varsoapi/src/services/app/internal/modules/rssitem"
)

func BlurToggle(w *api.Writer, r *http.Request, p httprouter.Params, j *api.JWT, request models.BlurToggleRequest) (*any, *api.Error) {
	err := rssitem.ToggleBlur(r.Context(), uuid.MustParse(request.RssItemId))
	if err != nil {
		return nil, api.NewInternalError(err, "failed toggling blur")
	}

	return nil, nil
}
