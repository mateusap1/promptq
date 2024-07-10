// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mateusap1/promptq/ent/predicate"
	"github.com/mateusap1/promptq/ent/promptrequest"
)

// PromptRequestUpdate is the builder for updating PromptRequest entities.
type PromptRequestUpdate struct {
	config
	hooks    []Hook
	mutation *PromptRequestMutation
}

// Where appends a list predicates to the PromptRequestUpdate builder.
func (pru *PromptRequestUpdate) Where(ps ...predicate.PromptRequest) *PromptRequestUpdate {
	pru.mutation.Where(ps...)
	return pru
}

// SetIdentifier sets the "identifier" field.
func (pru *PromptRequestUpdate) SetIdentifier(s string) *PromptRequestUpdate {
	pru.mutation.SetIdentifier(s)
	return pru
}

// SetNillableIdentifier sets the "identifier" field if the given value is not nil.
func (pru *PromptRequestUpdate) SetNillableIdentifier(s *string) *PromptRequestUpdate {
	if s != nil {
		pru.SetIdentifier(*s)
	}
	return pru
}

// SetPrompt sets the "prompt" field.
func (pru *PromptRequestUpdate) SetPrompt(s string) *PromptRequestUpdate {
	pru.mutation.SetPrompt(s)
	return pru
}

// SetNillablePrompt sets the "prompt" field if the given value is not nil.
func (pru *PromptRequestUpdate) SetNillablePrompt(s *string) *PromptRequestUpdate {
	if s != nil {
		pru.SetPrompt(*s)
	}
	return pru
}

// SetState sets the "state" field.
func (pru *PromptRequestUpdate) SetState(s string) *PromptRequestUpdate {
	pru.mutation.SetState(s)
	return pru
}

// SetNillableState sets the "state" field if the given value is not nil.
func (pru *PromptRequestUpdate) SetNillableState(s *string) *PromptRequestUpdate {
	if s != nil {
		pru.SetState(*s)
	}
	return pru
}

// Mutation returns the PromptRequestMutation object of the builder.
func (pru *PromptRequestUpdate) Mutation() *PromptRequestMutation {
	return pru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pru *PromptRequestUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pru.sqlSave, pru.mutation, pru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pru *PromptRequestUpdate) SaveX(ctx context.Context) int {
	affected, err := pru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pru *PromptRequestUpdate) Exec(ctx context.Context) error {
	_, err := pru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pru *PromptRequestUpdate) ExecX(ctx context.Context) {
	if err := pru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pru *PromptRequestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(promptrequest.Table, promptrequest.Columns, sqlgraph.NewFieldSpec(promptrequest.FieldID, field.TypeInt))
	if ps := pru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pru.mutation.Identifier(); ok {
		_spec.SetField(promptrequest.FieldIdentifier, field.TypeString, value)
	}
	if value, ok := pru.mutation.Prompt(); ok {
		_spec.SetField(promptrequest.FieldPrompt, field.TypeString, value)
	}
	if value, ok := pru.mutation.State(); ok {
		_spec.SetField(promptrequest.FieldState, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{promptrequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pru.mutation.done = true
	return n, nil
}

// PromptRequestUpdateOne is the builder for updating a single PromptRequest entity.
type PromptRequestUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PromptRequestMutation
}

// SetIdentifier sets the "identifier" field.
func (pruo *PromptRequestUpdateOne) SetIdentifier(s string) *PromptRequestUpdateOne {
	pruo.mutation.SetIdentifier(s)
	return pruo
}

// SetNillableIdentifier sets the "identifier" field if the given value is not nil.
func (pruo *PromptRequestUpdateOne) SetNillableIdentifier(s *string) *PromptRequestUpdateOne {
	if s != nil {
		pruo.SetIdentifier(*s)
	}
	return pruo
}

// SetPrompt sets the "prompt" field.
func (pruo *PromptRequestUpdateOne) SetPrompt(s string) *PromptRequestUpdateOne {
	pruo.mutation.SetPrompt(s)
	return pruo
}

// SetNillablePrompt sets the "prompt" field if the given value is not nil.
func (pruo *PromptRequestUpdateOne) SetNillablePrompt(s *string) *PromptRequestUpdateOne {
	if s != nil {
		pruo.SetPrompt(*s)
	}
	return pruo
}

// SetState sets the "state" field.
func (pruo *PromptRequestUpdateOne) SetState(s string) *PromptRequestUpdateOne {
	pruo.mutation.SetState(s)
	return pruo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (pruo *PromptRequestUpdateOne) SetNillableState(s *string) *PromptRequestUpdateOne {
	if s != nil {
		pruo.SetState(*s)
	}
	return pruo
}

// Mutation returns the PromptRequestMutation object of the builder.
func (pruo *PromptRequestUpdateOne) Mutation() *PromptRequestMutation {
	return pruo.mutation
}

// Where appends a list predicates to the PromptRequestUpdate builder.
func (pruo *PromptRequestUpdateOne) Where(ps ...predicate.PromptRequest) *PromptRequestUpdateOne {
	pruo.mutation.Where(ps...)
	return pruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pruo *PromptRequestUpdateOne) Select(field string, fields ...string) *PromptRequestUpdateOne {
	pruo.fields = append([]string{field}, fields...)
	return pruo
}

// Save executes the query and returns the updated PromptRequest entity.
func (pruo *PromptRequestUpdateOne) Save(ctx context.Context) (*PromptRequest, error) {
	return withHooks(ctx, pruo.sqlSave, pruo.mutation, pruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pruo *PromptRequestUpdateOne) SaveX(ctx context.Context) *PromptRequest {
	node, err := pruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pruo *PromptRequestUpdateOne) Exec(ctx context.Context) error {
	_, err := pruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pruo *PromptRequestUpdateOne) ExecX(ctx context.Context) {
	if err := pruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pruo *PromptRequestUpdateOne) sqlSave(ctx context.Context) (_node *PromptRequest, err error) {
	_spec := sqlgraph.NewUpdateSpec(promptrequest.Table, promptrequest.Columns, sqlgraph.NewFieldSpec(promptrequest.FieldID, field.TypeInt))
	id, ok := pruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PromptRequest.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, promptrequest.FieldID)
		for _, f := range fields {
			if !promptrequest.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != promptrequest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pruo.mutation.Identifier(); ok {
		_spec.SetField(promptrequest.FieldIdentifier, field.TypeString, value)
	}
	if value, ok := pruo.mutation.Prompt(); ok {
		_spec.SetField(promptrequest.FieldPrompt, field.TypeString, value)
	}
	if value, ok := pruo.mutation.State(); ok {
		_spec.SetField(promptrequest.FieldState, field.TypeString, value)
	}
	_node = &PromptRequest{config: pruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{promptrequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pruo.mutation.done = true
	return _node, nil
}
