// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mateusap1/promptq/ent/predicate"
	"github.com/mateusap1/promptq/ent/promptrequest"
	"github.com/mateusap1/promptq/ent/promptresponse"
)

// PromptResponseUpdate is the builder for updating PromptResponse entities.
type PromptResponseUpdate struct {
	config
	hooks    []Hook
	mutation *PromptResponseMutation
}

// Where appends a list predicates to the PromptResponseUpdate builder.
func (pru *PromptResponseUpdate) Where(ps ...predicate.PromptResponse) *PromptResponseUpdate {
	pru.mutation.Where(ps...)
	return pru
}

// SetResponse sets the "response" field.
func (pru *PromptResponseUpdate) SetResponse(s string) *PromptResponseUpdate {
	pru.mutation.SetResponse(s)
	return pru
}

// SetNillableResponse sets the "response" field if the given value is not nil.
func (pru *PromptResponseUpdate) SetNillableResponse(s *string) *PromptResponseUpdate {
	if s != nil {
		pru.SetResponse(*s)
	}
	return pru
}

// SetCreateDate sets the "create_date" field.
func (pru *PromptResponseUpdate) SetCreateDate(t time.Time) *PromptResponseUpdate {
	pru.mutation.SetCreateDate(t)
	return pru
}

// SetNillableCreateDate sets the "create_date" field if the given value is not nil.
func (pru *PromptResponseUpdate) SetNillableCreateDate(t *time.Time) *PromptResponseUpdate {
	if t != nil {
		pru.SetCreateDate(*t)
	}
	return pru
}

// SetPromptRequestID sets the "prompt_request" edge to the PromptRequest entity by ID.
func (pru *PromptResponseUpdate) SetPromptRequestID(id int) *PromptResponseUpdate {
	pru.mutation.SetPromptRequestID(id)
	return pru
}

// SetPromptRequest sets the "prompt_request" edge to the PromptRequest entity.
func (pru *PromptResponseUpdate) SetPromptRequest(p *PromptRequest) *PromptResponseUpdate {
	return pru.SetPromptRequestID(p.ID)
}

// Mutation returns the PromptResponseMutation object of the builder.
func (pru *PromptResponseUpdate) Mutation() *PromptResponseMutation {
	return pru.mutation
}

// ClearPromptRequest clears the "prompt_request" edge to the PromptRequest entity.
func (pru *PromptResponseUpdate) ClearPromptRequest() *PromptResponseUpdate {
	pru.mutation.ClearPromptRequest()
	return pru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pru *PromptResponseUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pru.sqlSave, pru.mutation, pru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pru *PromptResponseUpdate) SaveX(ctx context.Context) int {
	affected, err := pru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pru *PromptResponseUpdate) Exec(ctx context.Context) error {
	_, err := pru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pru *PromptResponseUpdate) ExecX(ctx context.Context) {
	if err := pru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pru *PromptResponseUpdate) check() error {
	if _, ok := pru.mutation.PromptRequestID(); pru.mutation.PromptRequestCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PromptResponse.prompt_request"`)
	}
	return nil
}

func (pru *PromptResponseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(promptresponse.Table, promptresponse.Columns, sqlgraph.NewFieldSpec(promptresponse.FieldID, field.TypeInt))
	if ps := pru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pru.mutation.Response(); ok {
		_spec.SetField(promptresponse.FieldResponse, field.TypeString, value)
	}
	if value, ok := pru.mutation.CreateDate(); ok {
		_spec.SetField(promptresponse.FieldCreateDate, field.TypeTime, value)
	}
	if pru.mutation.PromptRequestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   promptresponse.PromptRequestTable,
			Columns: []string{promptresponse.PromptRequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(promptrequest.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pru.mutation.PromptRequestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   promptresponse.PromptRequestTable,
			Columns: []string{promptresponse.PromptRequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(promptrequest.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{promptresponse.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pru.mutation.done = true
	return n, nil
}

// PromptResponseUpdateOne is the builder for updating a single PromptResponse entity.
type PromptResponseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PromptResponseMutation
}

// SetResponse sets the "response" field.
func (pruo *PromptResponseUpdateOne) SetResponse(s string) *PromptResponseUpdateOne {
	pruo.mutation.SetResponse(s)
	return pruo
}

// SetNillableResponse sets the "response" field if the given value is not nil.
func (pruo *PromptResponseUpdateOne) SetNillableResponse(s *string) *PromptResponseUpdateOne {
	if s != nil {
		pruo.SetResponse(*s)
	}
	return pruo
}

// SetCreateDate sets the "create_date" field.
func (pruo *PromptResponseUpdateOne) SetCreateDate(t time.Time) *PromptResponseUpdateOne {
	pruo.mutation.SetCreateDate(t)
	return pruo
}

// SetNillableCreateDate sets the "create_date" field if the given value is not nil.
func (pruo *PromptResponseUpdateOne) SetNillableCreateDate(t *time.Time) *PromptResponseUpdateOne {
	if t != nil {
		pruo.SetCreateDate(*t)
	}
	return pruo
}

// SetPromptRequestID sets the "prompt_request" edge to the PromptRequest entity by ID.
func (pruo *PromptResponseUpdateOne) SetPromptRequestID(id int) *PromptResponseUpdateOne {
	pruo.mutation.SetPromptRequestID(id)
	return pruo
}

// SetPromptRequest sets the "prompt_request" edge to the PromptRequest entity.
func (pruo *PromptResponseUpdateOne) SetPromptRequest(p *PromptRequest) *PromptResponseUpdateOne {
	return pruo.SetPromptRequestID(p.ID)
}

// Mutation returns the PromptResponseMutation object of the builder.
func (pruo *PromptResponseUpdateOne) Mutation() *PromptResponseMutation {
	return pruo.mutation
}

// ClearPromptRequest clears the "prompt_request" edge to the PromptRequest entity.
func (pruo *PromptResponseUpdateOne) ClearPromptRequest() *PromptResponseUpdateOne {
	pruo.mutation.ClearPromptRequest()
	return pruo
}

// Where appends a list predicates to the PromptResponseUpdate builder.
func (pruo *PromptResponseUpdateOne) Where(ps ...predicate.PromptResponse) *PromptResponseUpdateOne {
	pruo.mutation.Where(ps...)
	return pruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pruo *PromptResponseUpdateOne) Select(field string, fields ...string) *PromptResponseUpdateOne {
	pruo.fields = append([]string{field}, fields...)
	return pruo
}

// Save executes the query and returns the updated PromptResponse entity.
func (pruo *PromptResponseUpdateOne) Save(ctx context.Context) (*PromptResponse, error) {
	return withHooks(ctx, pruo.sqlSave, pruo.mutation, pruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pruo *PromptResponseUpdateOne) SaveX(ctx context.Context) *PromptResponse {
	node, err := pruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pruo *PromptResponseUpdateOne) Exec(ctx context.Context) error {
	_, err := pruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pruo *PromptResponseUpdateOne) ExecX(ctx context.Context) {
	if err := pruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pruo *PromptResponseUpdateOne) check() error {
	if _, ok := pruo.mutation.PromptRequestID(); pruo.mutation.PromptRequestCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PromptResponse.prompt_request"`)
	}
	return nil
}

func (pruo *PromptResponseUpdateOne) sqlSave(ctx context.Context) (_node *PromptResponse, err error) {
	if err := pruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(promptresponse.Table, promptresponse.Columns, sqlgraph.NewFieldSpec(promptresponse.FieldID, field.TypeInt))
	id, ok := pruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PromptResponse.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, promptresponse.FieldID)
		for _, f := range fields {
			if !promptresponse.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != promptresponse.FieldID {
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
	if value, ok := pruo.mutation.Response(); ok {
		_spec.SetField(promptresponse.FieldResponse, field.TypeString, value)
	}
	if value, ok := pruo.mutation.CreateDate(); ok {
		_spec.SetField(promptresponse.FieldCreateDate, field.TypeTime, value)
	}
	if pruo.mutation.PromptRequestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   promptresponse.PromptRequestTable,
			Columns: []string{promptresponse.PromptRequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(promptrequest.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pruo.mutation.PromptRequestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   promptresponse.PromptRequestTable,
			Columns: []string{promptresponse.PromptRequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(promptrequest.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PromptResponse{config: pruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{promptresponse.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pruo.mutation.done = true
	return _node, nil
}
