package post

import (
	"github.com/varsotech/varsoapi/src/services/analytics/client/models"
	"github.com/varsotech/varsoapi/src/services/analytics/internal/ent/build"
)

func TranslatePosts(posts []*build.Post) []*models.Post {
	var translated []*models.Post

	for _, post := range posts {
		translated = append(translated, TranslatePost(post))
	}

	return translated
}

func TranslatePost(post *build.Post) *models.Post {
	return &models.Post{
		Uuid:  post.ID.String(),
		Title: post.Title,
	}
}
