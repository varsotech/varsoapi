package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Comment struct {
	ent.Schema
}

func (Comment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("user_uuid"),
	}
}

func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Default(uuid.New).Unique().StorageKey("uuid"),
		field.UUID("user_uuid", uuid.New()),

		field.String("text"),
		field.Bool("was_edited").Default(false),

		field.Int64("total_votes"),
		field.Int64("upvotes"),
		field.Int64("downvotes"),
	}
}

func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("post", Post.Type).Ref("comments"),
	}
}
