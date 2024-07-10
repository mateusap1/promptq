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

func newUUID() string {
	result := uuid.New()
	return result.String()
}

// Fields of the PromptRequest.
func (PromptRequest) Fields() []ent.Field {
	return []ent.Field{
		field.String("identifier").DefaultFunc(newUUID).Unique(),
		field.String("prompt"),
		field.Bool("queued").Default(false),
	}
}

// Edges of the PromptRequest.
func (PromptRequest) Edges() []ent.Edge {
	return nil
}
