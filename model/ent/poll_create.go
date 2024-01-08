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

// PollCreate is the builder for creating a Poll entity.
type PollCreate struct {
	config
	mutation *PollMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *PollCreate) SetName(s string) *PollCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *PollCreate) SetDescription(s string) *PollCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pc *PollCreate) SetNillableDescription(s *string) *PollCreate {
	if s != nil {
		pc.SetDescription(*s)
	}
	return pc
}

// AddPollOptionIDs adds the "poll_options" edge to the PollOption entity by IDs.
func (pc *PollCreate) AddPollOptionIDs(ids ...int) *PollCreate {
	pc.mutation.AddPollOptionIDs(ids...)
	return pc
}

// AddPollOptions adds the "poll_options" edges to the PollOption entity.
func (pc *PollCreate) AddPollOptions(p ...*PollOption) *PollCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPollOptionIDs(ids...)
}

// SetCreatedByID sets the "created_by" edge to the User entity by ID.
func (pc *PollCreate) SetCreatedByID(id int) *PollCreate {
	pc.mutation.SetCreatedByID(id)
	return pc
}

// SetNillableCreatedByID sets the "created_by" edge to the User entity by ID if the given value is not nil.
func (pc *PollCreate) SetNillableCreatedByID(id *int) *PollCreate {
	if id != nil {
		pc = pc.SetCreatedByID(*id)
	}
	return pc
}

// SetCreatedBy sets the "created_by" edge to the User entity.
func (pc *PollCreate) SetCreatedBy(u *User) *PollCreate {
	return pc.SetCreatedByID(u.ID)
}

// Mutation returns the PollMutation object of the builder.
func (pc *PollCreate) Mutation() *PollMutation {
	return pc.mutation
}

// Save creates the Poll in the database.
func (pc *PollCreate) Save(ctx context.Context) (*Poll, error) {
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PollCreate) SaveX(ctx context.Context) *Poll {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PollCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PollCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PollCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Poll.name"`)}
	}
	return nil
}

func (pc *PollCreate) sqlSave(ctx context.Context) (*Poll, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PollCreate) createSpec() (*Poll, *sqlgraph.CreateSpec) {
	var (
		_node = &Poll{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(poll.Table, sqlgraph.NewFieldSpec(poll.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(poll.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(poll.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := pc.mutation.PollOptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   poll.PollOptionsTable,
			Columns: []string{poll.PollOptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(polloption.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   poll.CreatedByTable,
			Columns: []string{poll.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_created_polls = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PollCreateBulk is the builder for creating many Poll entities in bulk.
type PollCreateBulk struct {
	config
	err      error
	builders []*PollCreate
}

// Save creates the Poll entities in the database.
func (pcb *PollCreateBulk) Save(ctx context.Context) ([]*Poll, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Poll, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PollMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PollCreateBulk) SaveX(ctx context.Context) []*Poll {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PollCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PollCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}