package posts

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/varsotech/varsoapi/src/common/api"
	"github.com/varsotech/varsoapi/src/services/app/client/models"
	"github.com/varsotech/varsoapi/src/services/app/internal/modules/post"
)

func GetPosts(_ *api.Writer, r *http.Request, params httprouter.Params, _ *api.JWT) (*models.GetPostsResponse, *api.Error) {
	posts, err := post.GetPosts(r.Context())
	if err != nil {
		return nil, api.NewInternalError(err, "failed getting posts")
	}

	return &models.GetPostsResponse{
		Posts: post.TranslatePosts(posts),
	}, nil
}
