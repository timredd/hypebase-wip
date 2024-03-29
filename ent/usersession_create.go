// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hypebase/ent/user"
	"hypebase/ent/usersession"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserSessionCreate is the builder for creating a UserSession entity.
type UserSessionCreate struct {
	config
	mutation *UserSessionMutation
	hooks    []Hook
}

// SetSessionID sets the "session_id" field.
func (usc *UserSessionCreate) SetSessionID(s string) *UserSessionCreate {
	usc.mutation.SetSessionID(s)
	return usc
}

// SetUserID sets the "user_id" field.
func (usc *UserSessionCreate) SetUserID(i int) *UserSessionCreate {
	usc.mutation.SetUserID(i)
	return usc
}

// SetLoginTime sets the "login_time" field.
func (usc *UserSessionCreate) SetLoginTime(t time.Time) *UserSessionCreate {
	usc.mutation.SetLoginTime(t)
	return usc
}

// SetLastSeenTime sets the "last_seen_time" field.
func (usc *UserSessionCreate) SetLastSeenTime(t time.Time) *UserSessionCreate {
	usc.mutation.SetLastSeenTime(t)
	return usc
}

// SetUser sets the "user" edge to the User entity.
func (usc *UserSessionCreate) SetUser(u *User) *UserSessionCreate {
	return usc.SetUserID(u.ID)
}

// Mutation returns the UserSessionMutation object of the builder.
func (usc *UserSessionCreate) Mutation() *UserSessionMutation {
	return usc.mutation
}

// Save creates the UserSession in the database.
func (usc *UserSessionCreate) Save(ctx context.Context) (*UserSession, error) {
	var (
		err  error
		node *UserSession
	)
	if len(usc.hooks) == 0 {
		if err = usc.check(); err != nil {
			return nil, err
		}
		node, err = usc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserSessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = usc.check(); err != nil {
				return nil, err
			}
			usc.mutation = mutation
			node, err = usc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(usc.hooks) - 1; i >= 0; i-- {
			mut = usc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, usc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (usc *UserSessionCreate) SaveX(ctx context.Context) *UserSession {
	v, err := usc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (usc *UserSessionCreate) check() error {
	if _, ok := usc.mutation.SessionID(); !ok {
		return &ValidationError{Name: "session_id", err: errors.New("ent: missing required field \"session_id\"")}
	}
	if _, ok := usc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New("ent: missing required field \"user_id\"")}
	}
	if _, ok := usc.mutation.LoginTime(); !ok {
		return &ValidationError{Name: "login_time", err: errors.New("ent: missing required field \"login_time\"")}
	}
	if _, ok := usc.mutation.LastSeenTime(); !ok {
		return &ValidationError{Name: "last_seen_time", err: errors.New("ent: missing required field \"last_seen_time\"")}
	}
	if _, ok := usc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New("ent: missing required edge \"user\"")}
	}
	return nil
}

func (usc *UserSessionCreate) sqlSave(ctx context.Context) (*UserSession, error) {
	_node, _spec := usc.createSpec()
	if err := sqlgraph.CreateNode(ctx, usc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (usc *UserSessionCreate) createSpec() (*UserSession, *sqlgraph.CreateSpec) {
	var (
		_node = &UserSession{config: usc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: usersession.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usersession.FieldID,
			},
		}
	)
	if value, ok := usc.mutation.SessionID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usersession.FieldSessionID,
		})
		_node.SessionID = value
	}
	if value, ok := usc.mutation.LoginTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usersession.FieldLoginTime,
		})
		_node.LoginTime = value
	}
	if value, ok := usc.mutation.LastSeenTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usersession.FieldLastSeenTime,
		})
		_node.LastSeenTime = value
	}
	if nodes := usc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usersession.UserTable,
			Columns: []string{usersession.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserSessionCreateBulk is the builder for creating many UserSession entities in bulk.
type UserSessionCreateBulk struct {
	config
	builders []*UserSessionCreate
}

// Save creates the UserSession entities in the database.
func (uscb *UserSessionCreateBulk) Save(ctx context.Context) ([]*UserSession, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uscb.builders))
	nodes := make([]*UserSession, len(uscb.builders))
	mutators := make([]Mutator, len(uscb.builders))
	for i := range uscb.builders {
		func(i int, root context.Context) {
			builder := uscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserSessionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uscb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uscb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, uscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uscb *UserSessionCreateBulk) SaveX(ctx context.Context) []*UserSession {
	v, err := uscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
