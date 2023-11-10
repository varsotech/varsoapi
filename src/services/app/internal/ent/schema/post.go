package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Post struct {
	ent.Schema
}

func (Post) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("author_user_uuid"),
		index.Fields("total_votes"),
	}
}

func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Default(uuid.New).Unique().StorageKey("uuid"),
		field.UUID("author_user_uuid", uuid.New()),

		field.String("title"),
		field.String("cover_image"),

		field.Int64("total_votes"),
		field.Int64("upvotes"),
		field.Int64("downvotes"),
	}
}

func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comments", Comment.Type).Unique(),
	}
}
