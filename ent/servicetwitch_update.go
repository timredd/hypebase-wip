// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hypebase/ent/predicate"
	"hypebase/ent/servicetwitch"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ServiceTwitchUpdate is the builder for updating ServiceTwitch entities.
type ServiceTwitchUpdate struct {
	config
	hooks    []Hook
	mutation *ServiceTwitchMutation
}

// Where adds a new predicate for the ServiceTwitchUpdate builder.
func (stu *ServiceTwitchUpdate) Where(ps ...predicate.ServiceTwitch) *ServiceTwitchUpdate {
	stu.mutation.predicates = append(stu.mutation.predicates, ps...)
	return stu
}

// SetAccessToken sets the "access_token" field.
func (stu *ServiceTwitchUpdate) SetAccessToken(s string) *ServiceTwitchUpdate {
	stu.mutation.SetAccessToken(s)
	return stu
}

// SetRefreshToken sets the "refresh_token" field.
func (stu *ServiceTwitchUpdate) SetRefreshToken(s string) *ServiceTwitchUpdate {
	stu.mutation.SetRefreshToken(s)
	return stu
}

// SetScopes sets the "scopes" field.
func (stu *ServiceTwitchUpdate) SetScopes(s []string) *ServiceTwitchUpdate {
	stu.mutation.SetScopes(s)
	return stu
}

// Mutation returns the ServiceTwitchMutation object of the builder.
func (stu *ServiceTwitchUpdate) Mutation() *ServiceTwitchMutation {
	return stu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (stu *ServiceTwitchUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(stu.hooks) == 0 {
		affected, err = stu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServiceTwitchMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			stu.mutation = mutation
			affected, err = stu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(stu.hooks) - 1; i >= 0; i-- {
			mut = stu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, stu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (stu *ServiceTwitchUpdate) SaveX(ctx context.Context) int {
	affected, err := stu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (stu *ServiceTwitchUpdate) Exec(ctx context.Context) error {
	_, err := stu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stu *ServiceTwitchUpdate) ExecX(ctx context.Context) {
	if err := stu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (stu *ServiceTwitchUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   servicetwitch.Table,
			Columns: servicetwitch.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: servicetwitch.FieldID,
			},
		},
	}
	if ps := stu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := stu.mutation.AccessToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: servicetwitch.FieldAccessToken,
		})
	}
	if value, ok := stu.mutation.RefreshToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: servicetwitch.FieldRefreshToken,
		})
	}
	if value, ok := stu.mutation.Scopes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: servicetwitch.FieldScopes,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, stu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{servicetwitch.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ServiceTwitchUpdateOne is the builder for updating a single ServiceTwitch entity.
type ServiceTwitchUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ServiceTwitchMutation
}

// SetAccessToken sets the "access_token" field.
func (stuo *ServiceTwitchUpdateOne) SetAccessToken(s string) *ServiceTwitchUpdateOne {
	stuo.mutation.SetAccessToken(s)
	return stuo
}

// SetRefreshToken sets the "refresh_token" field.
func (stuo *ServiceTwitchUpdateOne) SetRefreshToken(s string) *ServiceTwitchUpdateOne {
	stuo.mutation.SetRefreshToken(s)
	return stuo
}

// SetScopes sets the "scopes" field.
func (stuo *ServiceTwitchUpdateOne) SetScopes(s []string) *ServiceTwitchUpdateOne {
	stuo.mutation.SetScopes(s)
	return stuo
}

// Mutation returns the ServiceTwitchMutation object of the builder.
func (stuo *ServiceTwitchUpdateOne) Mutation() *ServiceTwitchMutation {
	return stuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (stuo *ServiceTwitchUpdateOne) Select(field string, fields ...string) *ServiceTwitchUpdateOne {
	stuo.fields = append([]string{field}, fields...)
	return stuo
}

// Save executes the query and returns the updated ServiceTwitch entity.
func (stuo *ServiceTwitchUpdateOne) Save(ctx context.Context) (*ServiceTwitch, error) {
	var (
		err  error
		node *ServiceTwitch
	)
	if len(stuo.hooks) == 0 {
		node, err = stuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServiceTwitchMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			stuo.mutation = mutation
			node, err = stuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(stuo.hooks) - 1; i >= 0; i-- {
			mut = stuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, stuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (stuo *ServiceTwitchUpdateOne) SaveX(ctx context.Context) *ServiceTwitch {
	node, err := stuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (stuo *ServiceTwitchUpdateOne) Exec(ctx context.Context) error {
	_, err := stuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stuo *ServiceTwitchUpdateOne) ExecX(ctx context.Context) {
	if err := stuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (stuo *ServiceTwitchUpdateOne) sqlSave(ctx context.Context) (_node *ServiceTwitch, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   servicetwitch.Table,
			Columns: servicetwitch.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: servicetwitch.FieldID,
			},
		},
	}
	id, ok := stuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ServiceTwitch.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := stuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, servicetwitch.FieldID)
		for _, f := range fields {
			if !servicetwitch.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != servicetwitch.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := stuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := stuo.mutation.AccessToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: servicetwitch.FieldAccessToken,
		})
	}
	if value, ok := stuo.mutation.RefreshToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: servicetwitch.FieldRefreshToken,
		})
	}
	if value, ok := stuo.mutation.Scopes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: servicetwitch.FieldScopes,
		})
	}
	_node = &ServiceTwitch{config: stuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, stuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{servicetwitch.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
