package schema

import (
	"entgo.io/ent"
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
		field.UUID("identifier", uuid.New()).Unique(),
		field.String("prompt"),
	}
}

// Edges of the PromptRequest.
func (PromptRequest) Edges() []ent.Edge {
	return nil
}
