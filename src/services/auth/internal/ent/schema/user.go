package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"github.com/varsotech/varsoapi/src/services/auth/internal/ent/mixins"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.SoftDeleteMixin{},
		mixins.BanMixin{},
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("uuid").Unique(),
		index.Fields("email").Unique(),
		index.Fields("username").Unique(),
		index.Fields("discord_user_id").Unique(),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.New()).Default(uuid.New).Unique(),
		field.String("email").MinLen(5).Optional(),
		field.String("username").Match(regexp.MustCompile("[a-zA-Z_]+$")).MinLen(1).Optional(),
		field.String("password").Optional(),
		field.String("salt").Optional(),
		field.String("discord_user_id").Optional(),
		field.String("name").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
