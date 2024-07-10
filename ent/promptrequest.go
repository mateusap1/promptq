// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/mateusap1/promptq/ent/promptrequest"
	"github.com/mateusap1/promptq/ent/promptresponse"
)

// PromptRequest is the model entity for the PromptRequest schema.
type PromptRequest struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Identifier holds the value of the "identifier" field.
	Identifier string `json:"identifier,omitempty"`
	// Prompt holds the value of the "prompt" field.
	Prompt string `json:"prompt,omitempty"`
	// IsQueued holds the value of the "is_queued" field.
	IsQueued bool `json:"is_queued,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PromptRequestQuery when eager-loading is set.
	Edges        PromptRequestEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PromptRequestEdges holds the relations/edges for other nodes in the graph.
type PromptRequestEdges struct {
	// PromptResponse holds the value of the prompt_response edge.
	PromptResponse *PromptResponse `json:"prompt_response,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PromptResponseOrErr returns the PromptResponse value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PromptRequestEdges) PromptResponseOrErr() (*PromptResponse, error) {
	if e.PromptResponse != nil {
		return e.PromptResponse, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: promptresponse.Label}
	}
	return nil, &NotLoadedError{edge: "prompt_response"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PromptRequest) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case promptrequest.FieldIsQueued:
			values[i] = new(sql.NullBool)
		case promptrequest.FieldID:
			values[i] = new(sql.NullInt64)
		case promptrequest.FieldIdentifier, promptrequest.FieldPrompt:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PromptRequest fields.
func (pr *PromptRequest) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case promptrequest.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case promptrequest.FieldIdentifier:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field identifier", values[i])
			} else if value.Valid {
				pr.Identifier = value.String
			}
		case promptrequest.FieldPrompt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field prompt", values[i])
			} else if value.Valid {
				pr.Prompt = value.String
			}
		case promptrequest.FieldIsQueued:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_queued", values[i])
			} else if value.Valid {
				pr.IsQueued = value.Bool
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PromptRequest.
// This includes values selected through modifiers, order, etc.
func (pr *PromptRequest) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryPromptResponse queries the "prompt_response" edge of the PromptRequest entity.
func (pr *PromptRequest) QueryPromptResponse() *PromptResponseQuery {
	return NewPromptRequestClient(pr.config).QueryPromptResponse(pr)
}

// Update returns a builder for updating this PromptRequest.
// Note that you need to call PromptRequest.Unwrap() before calling this method if this PromptRequest
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *PromptRequest) Update() *PromptRequestUpdateOne {
	return NewPromptRequestClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the PromptRequest entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *PromptRequest) Unwrap() *PromptRequest {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: PromptRequest is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *PromptRequest) String() string {
	var builder strings.Builder
	builder.WriteString("PromptRequest(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("identifier=")
	builder.WriteString(pr.Identifier)
	builder.WriteString(", ")
	builder.WriteString("prompt=")
	builder.WriteString(pr.Prompt)
	builder.WriteString(", ")
	builder.WriteString("is_queued=")
	builder.WriteString(fmt.Sprintf("%v", pr.IsQueued))
	builder.WriteByte(')')
	return builder.String()
}

// PromptRequests is a parsable slice of PromptRequest.
type PromptRequests []*PromptRequest
