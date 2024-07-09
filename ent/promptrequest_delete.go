// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"promptq/ent/predicate"
	"promptq/ent/promptrequest"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PromptRequestDelete is the builder for deleting a PromptRequest entity.
type PromptRequestDelete struct {
	config
	hooks    []Hook
	mutation *PromptRequestMutation
}

// Where appends a list predicates to the PromptRequestDelete builder.
func (prd *PromptRequestDelete) Where(ps ...predicate.PromptRequest) *PromptRequestDelete {
	prd.mutation.Where(ps...)
	return prd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (prd *PromptRequestDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, prd.sqlExec, prd.mutation, prd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (prd *PromptRequestDelete) ExecX(ctx context.Context) int {
	n, err := prd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (prd *PromptRequestDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(promptrequest.Table, sqlgraph.NewFieldSpec(promptrequest.FieldID, field.TypeInt))
	if ps := prd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, prd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	prd.mutation.done = true
	return affected, err
}

// PromptRequestDeleteOne is the builder for deleting a single PromptRequest entity.
type PromptRequestDeleteOne struct {
	prd *PromptRequestDelete
}

// Where appends a list predicates to the PromptRequestDelete builder.
func (prdo *PromptRequestDeleteOne) Where(ps ...predicate.PromptRequest) *PromptRequestDeleteOne {
	prdo.prd.mutation.Where(ps...)
	return prdo
}

// Exec executes the deletion query.
func (prdo *PromptRequestDeleteOne) Exec(ctx context.Context) error {
	n, err := prdo.prd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{promptrequest.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (prdo *PromptRequestDeleteOne) ExecX(ctx context.Context) {
	if err := prdo.Exec(ctx); err != nil {
		panic(err)
	}
}
