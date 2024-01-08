// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"filip.filipovic/polling-app/model/ent/poll"
	"filip.filipovic/polling-app/model/ent/polloption"
	"filip.filipovic/polling-app/model/ent/user"
)

// PollOptionCreate is the builder for creating a PollOption entity.
type PollOptionCreate struct {
	config
	mutation *PollOptionMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (poc *PollOptionCreate) SetName(s string) *PollOptionCreate {
	poc.mutation.SetName(s)
	return poc
}

// SetVoteCount sets the "vote_count" field.
func (poc *PollOptionCreate) SetVoteCount(i int) *PollOptionCreate {
	poc.mutation.SetVoteCount(i)
	return poc
}

// AddUsersVotedIDs adds the "users_voted" edge to the User entity by IDs.
func (poc *PollOptionCreate) AddUsersVotedIDs(ids ...int) *PollOptionCreate {
	poc.mutation.AddUsersVotedIDs(ids...)
	return poc
}

// AddUsersVoted adds the "users_voted" edges to the User entity.
func (poc *PollOptionCreate) AddUsersVoted(u ...*User) *PollOptionCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return poc.AddUsersVotedIDs(ids...)
}

// SetPollID sets the "poll" edge to the Poll entity by ID.
func (poc *PollOptionCreate) SetPollID(id int) *PollOptionCreate {
	poc.mutation.SetPollID(id)
	return poc
}

// SetNillablePollID sets the "poll" edge to the Poll entity by ID if the given value is not nil.
func (poc *PollOptionCreate) SetNillablePollID(id *int) *PollOptionCreate {
	if id != nil {
		poc = poc.SetPollID(*id)
	}
	return poc
}

// SetPoll sets the "poll" edge to the Poll entity.
func (poc *PollOptionCreate) SetPoll(p *Poll) *PollOptionCreate {
	return poc.SetPollID(p.ID)
}

// Mutation returns the PollOptionMutation object of the builder.
func (poc *PollOptionCreate) Mutation() *PollOptionMutation {
	return poc.mutation
}

// Save creates the PollOption in the database.
func (poc *PollOptionCreate) Save(ctx context.Context) (*PollOption, error) {
	return withHooks(ctx, poc.sqlSave, poc.mutation, poc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (poc *PollOptionCreate) SaveX(ctx context.Context) *PollOption {
	v, err := poc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (poc *PollOptionCreate) Exec(ctx context.Context) error {
	_, err := poc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (poc *PollOptionCreate) ExecX(ctx context.Context) {
	if err := poc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (poc *PollOptionCreate) check() error {
	if _, ok := poc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "PollOption.name"`)}
	}
	if _, ok := poc.mutation.VoteCount(); !ok {
		return &ValidationError{Name: "vote_count", err: errors.New(`ent: missing required field "PollOption.vote_count"`)}
	}
	return nil
}

func (poc *PollOptionCreate) sqlSave(ctx context.Context) (*PollOption, error) {
	if err := poc.check(); err != nil {
		return nil, err
	}
	_node, _spec := poc.createSpec()
	if err := sqlgraph.CreateNode(ctx, poc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	poc.mutation.id = &_node.ID
	poc.mutation.done = true
	return _node, nil
}

func (poc *PollOptionCreate) createSpec() (*PollOption, *sqlgraph.CreateSpec) {
	var (
		_node = &PollOption{config: poc.config}
		_spec = sqlgraph.NewCreateSpec(polloption.Table, sqlgraph.NewFieldSpec(polloption.FieldID, field.TypeInt))
	)
	if value, ok := poc.mutation.Name(); ok {
		_spec.SetField(polloption.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := poc.mutation.VoteCount(); ok {
		_spec.SetField(polloption.FieldVoteCount, field.TypeInt, value)
		_node.VoteCount = value
	}
	if nodes := poc.mutation.UsersVotedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   polloption.UsersVotedTable,
			Columns: polloption.UsersVotedPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := poc.mutation.PollIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   polloption.PollTable,
			Columns: []string{polloption.PollColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(poll.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.poll_poll_options = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PollOptionCreateBulk is the builder for creating many PollOption entities in bulk.
type PollOptionCreateBulk struct {
	config
	err      error
	builders []*PollOptionCreate
}

// Save creates the PollOption entities in the database.
func (pocb *PollOptionCreateBulk) Save(ctx context.Context) ([]*PollOption, error) {
	if pocb.err != nil {
		return nil, pocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pocb.builders))
	nodes := make([]*PollOption, len(pocb.builders))
	mutators := make([]Mutator, len(pocb.builders))
	for i := range pocb.builders {
		func(i int, root context.Context) {
			builder := pocb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PollOptionMutation)
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
					_, err = mutators[i+1].Mutate(root, pocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pocb *PollOptionCreateBulk) SaveX(ctx context.Context) []*PollOption {
	v, err := pocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pocb *PollOptionCreateBulk) Exec(ctx context.Context) error {
	_, err := pocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pocb *PollOptionCreateBulk) ExecX(ctx context.Context) {
	if err := pocb.Exec(ctx); err != nil {
		panic(err)
	}
}
