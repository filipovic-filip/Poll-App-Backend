// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"filip.filipovic/polling-app/model/ent/polloption"
	"filip.filipovic/polling-app/model/ent/predicate"
)

// PollOptionDelete is the builder for deleting a PollOption entity.
type PollOptionDelete struct {
	config
	hooks    []Hook
	mutation *PollOptionMutation
}

// Where appends a list predicates to the PollOptionDelete builder.
func (pod *PollOptionDelete) Where(ps ...predicate.PollOption) *PollOptionDelete {
	pod.mutation.Where(ps...)
	return pod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pod *PollOptionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pod.sqlExec, pod.mutation, pod.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pod *PollOptionDelete) ExecX(ctx context.Context) int {
	n, err := pod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pod *PollOptionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(polloption.Table, sqlgraph.NewFieldSpec(polloption.FieldID, field.TypeInt))
	if ps := pod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pod.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pod.mutation.done = true
	return affected, err
}

// PollOptionDeleteOne is the builder for deleting a single PollOption entity.
type PollOptionDeleteOne struct {
	pod *PollOptionDelete
}

// Where appends a list predicates to the PollOptionDelete builder.
func (podo *PollOptionDeleteOne) Where(ps ...predicate.PollOption) *PollOptionDeleteOne {
	podo.pod.mutation.Where(ps...)
	return podo
}

// Exec executes the deletion query.
func (podo *PollOptionDeleteOne) Exec(ctx context.Context) error {
	n, err := podo.pod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{polloption.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (podo *PollOptionDeleteOne) ExecX(ctx context.Context) {
	if err := podo.Exec(ctx); err != nil {
		panic(err)
	}
}
