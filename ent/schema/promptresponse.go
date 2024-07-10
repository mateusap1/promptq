package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PromptResponse holds the schema definition for the PromptResponse entity.
type PromptResponse struct {
	ent.Schema
}

// Fields of the PromptResponse.
func (PromptResponse) Fields() []ent.Field {
	return []ent.Field{
		field.String("response"),
		field.Bool("is_answered").Default(false),
	}
}

// Edges of the PromptResponse.
func (PromptResponse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("prompt_request", PromptRequest.Type).
			Ref("prompt_response").
			Unique().
			Required(),
	}
}
