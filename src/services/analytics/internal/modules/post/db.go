package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/varsotech/varsoapi/src/services/analytics/internal/ent"
	"github.com/varsotech/varsoapi/src/services/analytics/internal/ent/build"
)

/*
	Users can submit posts in either the main chart or the friends chart.
*/

func GetPosts(ctx context.Context) ([]*build.Post, error) {
	posts, err := ent.Database.Post.Query().All(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed getting posts from db")
	}

	return posts, nil
}

func CreatePost(ctx context.Context, userUUID uuid.UUID, title string, coverImage string) error {
	err := ent.Database.Post.Create().
		SetAuthorUserUUID(userUUID).
		SetTitle(title).
		SetCoverImage(coverImage).
		Exec(ctx)
	if err != nil {
		return errors.Wrapf(err, "failed creating post in db")
	}

	return nil
}
