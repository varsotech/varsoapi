package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type RSSFeed struct {
	ent.Schema
}

func (RSSFeed) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
		mixin.UpdateTime{},
	}
}

func (RSSFeed) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("rss_feed_url").Unique(),
	}
}

func (RSSFeed) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Default(uuid.New).Unique().StorageKey("uuid"),
		field.String("rss_feed_url").Unique(),
	}
}

func (RSSFeed) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("items", NewsItem.Type).Ref("feed"),
		edge.From("organization", Organization.Type).Ref("feeds").Unique(),
	}
}
