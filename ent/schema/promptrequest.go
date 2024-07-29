package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PromptRequest holds the schema definition for the PromptRequest entity.
type PromptRequest struct {
	ent.Schema
}

// Fields of the PromptRequest.
func (PromptRequest) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("identifier", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("prompt"),
		field.Bool("is_queued").Default(false),
		field.Bool("is_answered").Default(false),
		field.Time("create_date").Default(time.Now),
	}
}

// Edges of the PromptRequest.
func (PromptRequest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("prompt_response", PromptResponse.Type).Unique(),
	}
}
