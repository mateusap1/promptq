package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique(),
		field.Bytes("password_hash"),
		field.Bytes("salt"),
		field.String("full_name"),
		field.Bool("email_verified"),
		field.String("reset_token").Nillable(),
		field.Time("reset_token_expires").Nillable(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sessions", Session.Type),
		edge.To("prompt_requests", PromptRequest.Type),
	}
}
