// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/mateusap1/promptq/ent/promptrequest"
	"github.com/mateusap1/promptq/ent/promptresponse"
)

// PromptRequestCreate is the builder for creating a PromptRequest entity.
type PromptRequestCreate struct {
	config
	mutation *PromptRequestMutation
	hooks    []Hook
}

// SetIdentifier sets the "identifier" field.
func (prc *PromptRequestCreate) SetIdentifier(u uuid.UUID) *PromptRequestCreate {
	prc.mutation.SetIdentifier(u)
	return prc
}

// SetNillableIdentifier sets the "identifier" field if the given value is not nil.
func (prc *PromptRequestCreate) SetNillableIdentifier(u *uuid.UUID) *PromptRequestCreate {
	if u != nil {
		prc.SetIdentifier(*u)
	}
	return prc
}

// SetPrompt sets the "prompt" field.
func (prc *PromptRequestCreate) SetPrompt(s string) *PromptRequestCreate {
	prc.mutation.SetPrompt(s)
	return prc
}

// SetIsQueued sets the "is_queued" field.
func (prc *PromptRequestCreate) SetIsQueued(b bool) *PromptRequestCreate {
	prc.mutation.SetIsQueued(b)
	return prc
}

// SetNillableIsQueued sets the "is_queued" field if the given value is not nil.
func (prc *PromptRequestCreate) SetNillableIsQueued(b *bool) *PromptRequestCreate {
	if b != nil {
		prc.SetIsQueued(*b)
	}
	return prc
}

// SetIsAnswered sets the "is_answered" field.
func (prc *PromptRequestCreate) SetIsAnswered(b bool) *PromptRequestCreate {
	prc.mutation.SetIsAnswered(b)
	return prc
}

// SetNillableIsAnswered sets the "is_answered" field if the given value is not nil.
func (prc *PromptRequestCreate) SetNillableIsAnswered(b *bool) *PromptRequestCreate {
	if b != nil {
		prc.SetIsAnswered(*b)
	}
	return prc
}

// SetCreateDate sets the "create_date" field.
func (prc *PromptRequestCreate) SetCreateDate(t time.Time) *PromptRequestCreate {
	prc.mutation.SetCreateDate(t)
	return prc
}

// SetNillableCreateDate sets the "create_date" field if the given value is not nil.
func (prc *PromptRequestCreate) SetNillableCreateDate(t *time.Time) *PromptRequestCreate {
	if t != nil {
		prc.SetCreateDate(*t)
	}
	return prc
}

// SetPromptResponseID sets the "prompt_response" edge to the PromptResponse entity by ID.
func (prc *PromptRequestCreate) SetPromptResponseID(id int) *PromptRequestCreate {
	prc.mutation.SetPromptResponseID(id)
	return prc
}

// SetNillablePromptResponseID sets the "prompt_response" edge to the PromptResponse entity by ID if the given value is not nil.
func (prc *PromptRequestCreate) SetNillablePromptResponseID(id *int) *PromptRequestCreate {
	if id != nil {
		prc = prc.SetPromptResponseID(*id)
	}
	return prc
}

// SetPromptResponse sets the "prompt_response" edge to the PromptResponse entity.
func (prc *PromptRequestCreate) SetPromptResponse(p *PromptResponse) *PromptRequestCreate {
	return prc.SetPromptResponseID(p.ID)
}

// Mutation returns the PromptRequestMutation object of the builder.
func (prc *PromptRequestCreate) Mutation() *PromptRequestMutation {
	return prc.mutation
}

// Save creates the PromptRequest in the database.
func (prc *PromptRequestCreate) Save(ctx context.Context) (*PromptRequest, error) {
	prc.defaults()
	return withHooks(ctx, prc.sqlSave, prc.mutation, prc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (prc *PromptRequestCreate) SaveX(ctx context.Context) *PromptRequest {
	v, err := prc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (prc *PromptRequestCreate) Exec(ctx context.Context) error {
	_, err := prc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (prc *PromptRequestCreate) ExecX(ctx context.Context) {
	if err := prc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (prc *PromptRequestCreate) defaults() {
	if _, ok := prc.mutation.Identifier(); !ok {
		v := promptrequest.DefaultIdentifier()
		prc.mutation.SetIdentifier(v)
	}
	if _, ok := prc.mutation.IsQueued(); !ok {
		v := promptrequest.DefaultIsQueued
		prc.mutation.SetIsQueued(v)
	}
	if _, ok := prc.mutation.IsAnswered(); !ok {
		v := promptrequest.DefaultIsAnswered
		prc.mutation.SetIsAnswered(v)
	}
	if _, ok := prc.mutation.CreateDate(); !ok {
		v := promptrequest.DefaultCreateDate()
		prc.mutation.SetCreateDate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (prc *PromptRequestCreate) check() error {
	if _, ok := prc.mutation.Identifier(); !ok {
		return &ValidationError{Name: "identifier", err: errors.New(`ent: missing required field "PromptRequest.identifier"`)}
	}
	if _, ok := prc.mutation.Prompt(); !ok {
		return &ValidationError{Name: "prompt", err: errors.New(`ent: missing required field "PromptRequest.prompt"`)}
	}
	if _, ok := prc.mutation.IsQueued(); !ok {
		return &ValidationError{Name: "is_queued", err: errors.New(`ent: missing required field "PromptRequest.is_queued"`)}
	}
	if _, ok := prc.mutation.IsAnswered(); !ok {
		return &ValidationError{Name: "is_answered", err: errors.New(`ent: missing required field "PromptRequest.is_answered"`)}
	}
	if _, ok := prc.mutation.CreateDate(); !ok {
		return &ValidationError{Name: "create_date", err: errors.New(`ent: missing required field "PromptRequest.create_date"`)}
	}
	return nil
}

func (prc *PromptRequestCreate) sqlSave(ctx context.Context) (*PromptRequest, error) {
	if err := prc.check(); err != nil {
		return nil, err
	}
	_node, _spec := prc.createSpec()
	if err := sqlgraph.CreateNode(ctx, prc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	prc.mutation.id = &_node.ID
	prc.mutation.done = true
	return _node, nil
}

func (prc *PromptRequestCreate) createSpec() (*PromptRequest, *sqlgraph.CreateSpec) {
	var (
		_node = &PromptRequest{config: prc.config}
		_spec = sqlgraph.NewCreateSpec(promptrequest.Table, sqlgraph.NewFieldSpec(promptrequest.FieldID, field.TypeInt))
	)
	if value, ok := prc.mutation.Identifier(); ok {
		_spec.SetField(promptrequest.FieldIdentifier, field.TypeUUID, value)
		_node.Identifier = value
	}
	if value, ok := prc.mutation.Prompt(); ok {
		_spec.SetField(promptrequest.FieldPrompt, field.TypeString, value)
		_node.Prompt = value
	}
	if value, ok := prc.mutation.IsQueued(); ok {
		_spec.SetField(promptrequest.FieldIsQueued, field.TypeBool, value)
		_node.IsQueued = value
	}
	if value, ok := prc.mutation.IsAnswered(); ok {
		_spec.SetField(promptrequest.FieldIsAnswered, field.TypeBool, value)
		_node.IsAnswered = value
	}
	if value, ok := prc.mutation.CreateDate(); ok {
		_spec.SetField(promptrequest.FieldCreateDate, field.TypeTime, value)
		_node.CreateDate = value
	}
	if nodes := prc.mutation.PromptResponseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   promptrequest.PromptResponseTable,
			Columns: []string{promptrequest.PromptResponseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(promptresponse.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PromptRequestCreateBulk is the builder for creating many PromptRequest entities in bulk.
type PromptRequestCreateBulk struct {
	config
	err      error
	builders []*PromptRequestCreate
}

// Save creates the PromptRequest entities in the database.
func (prcb *PromptRequestCreateBulk) Save(ctx context.Context) ([]*PromptRequest, error) {
	if prcb.err != nil {
		return nil, prcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(prcb.builders))
	nodes := make([]*PromptRequest, len(prcb.builders))
	mutators := make([]Mutator, len(prcb.builders))
	for i := range prcb.builders {
		func(i int, root context.Context) {
			builder := prcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PromptRequestMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, prcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, prcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, prcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (prcb *PromptRequestCreateBulk) SaveX(ctx context.Context) []*PromptRequest {
	v, err := prcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (prcb *PromptRequestCreateBulk) Exec(ctx context.Context) error {
	_, err := prcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (prcb *PromptRequestCreateBulk) ExecX(ctx context.Context) {
	if err := prcb.Exec(ctx); err != nil {
		panic(err)
	}
}
