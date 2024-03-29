// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hypebase/ent/predicate"
	"hypebase/ent/servicetwitch"
	"hypebase/ent/user"
	"hypebase/ent/usersession"
	"sync"
	"time"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeServiceTwitch = "ServiceTwitch"
	TypeUser          = "User"
	TypeUserSession   = "UserSession"
)

// ServiceTwitchMutation represents an operation that mutates the ServiceTwitch nodes in the graph.
type ServiceTwitchMutation struct {
	config
	op            Op
	typ           string
	id            *int
	access_token  *string
	refresh_token *string
	scopes        *[]string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*ServiceTwitch, error)
	predicates    []predicate.ServiceTwitch
}

var _ ent.Mutation = (*ServiceTwitchMutation)(nil)

// servicetwitchOption allows management of the mutation configuration using functional options.
type servicetwitchOption func(*ServiceTwitchMutation)

// newServiceTwitchMutation creates new mutation for the ServiceTwitch entity.
func newServiceTwitchMutation(c config, op Op, opts ...servicetwitchOption) *ServiceTwitchMutation {
	m := &ServiceTwitchMutation{
		config:        c,
		op:            op,
		typ:           TypeServiceTwitch,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withServiceTwitchID sets the ID field of the mutation.
func withServiceTwitchID(id int) servicetwitchOption {
	return func(m *ServiceTwitchMutation) {
		var (
			err   error
			once  sync.Once
			value *ServiceTwitch
		)
		m.oldValue = func(ctx context.Context) (*ServiceTwitch, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().ServiceTwitch.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withServiceTwitch sets the old ServiceTwitch of the mutation.
func withServiceTwitch(node *ServiceTwitch) servicetwitchOption {
	return func(m *ServiceTwitchMutation) {
		m.oldValue = func(context.Context) (*ServiceTwitch, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ServiceTwitchMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ServiceTwitchMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *ServiceTwitchMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetAccessToken sets the "access_token" field.
func (m *ServiceTwitchMutation) SetAccessToken(s string) {
	m.access_token = &s
}

// AccessToken returns the value of the "access_token" field in the mutation.
func (m *ServiceTwitchMutation) AccessToken() (r string, exists bool) {
	v := m.access_token
	if v == nil {
		return
	}
	return *v, true
}

// OldAccessToken returns the old "access_token" field's value of the ServiceTwitch entity.
// If the ServiceTwitch object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ServiceTwitchMutation) OldAccessToken(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldAccessToken is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldAccessToken requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAccessToken: %w", err)
	}
	return oldValue.AccessToken, nil
}

// ResetAccessToken resets all changes to the "access_token" field.
func (m *ServiceTwitchMutation) ResetAccessToken() {
	m.access_token = nil
}

// SetRefreshToken sets the "refresh_token" field.
func (m *ServiceTwitchMutation) SetRefreshToken(s string) {
	m.refresh_token = &s
}

// RefreshToken returns the value of the "refresh_token" field in the mutation.
func (m *ServiceTwitchMutation) RefreshToken() (r string, exists bool) {
	v := m.refresh_token
	if v == nil {
		return
	}
	return *v, true
}

// OldRefreshToken returns the old "refresh_token" field's value of the ServiceTwitch entity.
// If the ServiceTwitch object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ServiceTwitchMutation) OldRefreshToken(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldRefreshToken is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldRefreshToken requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRefreshToken: %w", err)
	}
	return oldValue.RefreshToken, nil
}

// ResetRefreshToken resets all changes to the "refresh_token" field.
func (m *ServiceTwitchMutation) ResetRefreshToken() {
	m.refresh_token = nil
}

// SetScopes sets the "scopes" field.
func (m *ServiceTwitchMutation) SetScopes(s []string) {
	m.scopes = &s
}

// Scopes returns the value of the "scopes" field in the mutation.
func (m *ServiceTwitchMutation) Scopes() (r []string, exists bool) {
	v := m.scopes
	if v == nil {
		return
	}
	return *v, true
}

// OldScopes returns the old "scopes" field's value of the ServiceTwitch entity.
// If the ServiceTwitch object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ServiceTwitchMutation) OldScopes(ctx context.Context) (v []string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldScopes is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldScopes requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldScopes: %w", err)
	}
	return oldValue.Scopes, nil
}

// ResetScopes resets all changes to the "scopes" field.
func (m *ServiceTwitchMutation) ResetScopes() {
	m.scopes = nil
}

// Op returns the operation name.
func (m *ServiceTwitchMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (ServiceTwitch).
func (m *ServiceTwitchMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ServiceTwitchMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.access_token != nil {
		fields = append(fields, servicetwitch.FieldAccessToken)
	}
	if m.refresh_token != nil {
		fields = append(fields, servicetwitch.FieldRefreshToken)
	}
	if m.scopes != nil {
		fields = append(fields, servicetwitch.FieldScopes)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ServiceTwitchMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case servicetwitch.FieldAccessToken:
		return m.AccessToken()
	case servicetwitch.FieldRefreshToken:
		return m.RefreshToken()
	case servicetwitch.FieldScopes:
		return m.Scopes()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ServiceTwitchMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case servicetwitch.FieldAccessToken:
		return m.OldAccessToken(ctx)
	case servicetwitch.FieldRefreshToken:
		return m.OldRefreshToken(ctx)
	case servicetwitch.FieldScopes:
		return m.OldScopes(ctx)
	}
	return nil, fmt.Errorf("unknown ServiceTwitch field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ServiceTwitchMutation) SetField(name string, value ent.Value) error {
	switch name {
	case servicetwitch.FieldAccessToken:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAccessToken(v)
		return nil
	case servicetwitch.FieldRefreshToken:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRefreshToken(v)
		return nil
	case servicetwitch.FieldScopes:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetScopes(v)
		return nil
	}
	return fmt.Errorf("unknown ServiceTwitch field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ServiceTwitchMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ServiceTwitchMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ServiceTwitchMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown ServiceTwitch numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ServiceTwitchMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ServiceTwitchMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ServiceTwitchMutation) ClearField(name string) error {
	return fmt.Errorf("unknown ServiceTwitch nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ServiceTwitchMutation) ResetField(name string) error {
	switch name {
	case servicetwitch.FieldAccessToken:
		m.ResetAccessToken()
		return nil
	case servicetwitch.FieldRefreshToken:
		m.ResetRefreshToken()
		return nil
	case servicetwitch.FieldScopes:
		m.ResetScopes()
		return nil
	}
	return fmt.Errorf("unknown ServiceTwitch field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ServiceTwitchMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ServiceTwitchMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ServiceTwitchMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ServiceTwitchMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ServiceTwitchMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ServiceTwitchMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ServiceTwitchMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown ServiceTwitch unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ServiceTwitchMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown ServiceTwitch edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	email         *string
	password_hash *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetEmail sets the "email" field.
func (m *UserMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *UserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *UserMutation) ResetEmail() {
	m.email = nil
}

// SetPasswordHash sets the "password_hash" field.
func (m *UserMutation) SetPasswordHash(s string) {
	m.password_hash = &s
}

// PasswordHash returns the value of the "password_hash" field in the mutation.
func (m *UserMutation) PasswordHash() (r string, exists bool) {
	v := m.password_hash
	if v == nil {
		return
	}
	return *v, true
}

// OldPasswordHash returns the old "password_hash" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPasswordHash(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPasswordHash is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPasswordHash requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPasswordHash: %w", err)
	}
	return oldValue.PasswordHash, nil
}

// ResetPasswordHash resets all changes to the "password_hash" field.
func (m *UserMutation) ResetPasswordHash() {
	m.password_hash = nil
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m.password_hash != nil {
		fields = append(fields, user.FieldPasswordHash)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldEmail:
		return m.Email()
	case user.FieldPasswordHash:
		return m.PasswordHash()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldEmail:
		return m.OldEmail(ctx)
	case user.FieldPasswordHash:
		return m.OldPasswordHash(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldPasswordHash:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPasswordHash(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldPasswordHash:
		m.ResetPasswordHash()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}

// UserSessionMutation represents an operation that mutates the UserSession nodes in the graph.
type UserSessionMutation struct {
	config
	op             Op
	typ            string
	id             *int
	session_id     *string
	login_time     *time.Time
	last_seen_time *time.Time
	clearedFields  map[string]struct{}
	user           *int
	cleareduser    bool
	done           bool
	oldValue       func(context.Context) (*UserSession, error)
	predicates     []predicate.UserSession
}

var _ ent.Mutation = (*UserSessionMutation)(nil)

// usersessionOption allows management of the mutation configuration using functional options.
type usersessionOption func(*UserSessionMutation)

// newUserSessionMutation creates new mutation for the UserSession entity.
func newUserSessionMutation(c config, op Op, opts ...usersessionOption) *UserSessionMutation {
	m := &UserSessionMutation{
		config:        c,
		op:            op,
		typ:           TypeUserSession,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserSessionID sets the ID field of the mutation.
func withUserSessionID(id int) usersessionOption {
	return func(m *UserSessionMutation) {
		var (
			err   error
			once  sync.Once
			value *UserSession
		)
		m.oldValue = func(ctx context.Context) (*UserSession, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().UserSession.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUserSession sets the old UserSession of the mutation.
func withUserSession(node *UserSession) usersessionOption {
	return func(m *UserSessionMutation) {
		m.oldValue = func(context.Context) (*UserSession, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserSessionMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserSessionMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *UserSessionMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetSessionID sets the "session_id" field.
func (m *UserSessionMutation) SetSessionID(s string) {
	m.session_id = &s
}

// SessionID returns the value of the "session_id" field in the mutation.
func (m *UserSessionMutation) SessionID() (r string, exists bool) {
	v := m.session_id
	if v == nil {
		return
	}
	return *v, true
}

// OldSessionID returns the old "session_id" field's value of the UserSession entity.
// If the UserSession object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserSessionMutation) OldSessionID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldSessionID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldSessionID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSessionID: %w", err)
	}
	return oldValue.SessionID, nil
}

// ResetSessionID resets all changes to the "session_id" field.
func (m *UserSessionMutation) ResetSessionID() {
	m.session_id = nil
}

// SetUserID sets the "user_id" field.
func (m *UserSessionMutation) SetUserID(i int) {
	m.user = &i
}

// UserID returns the value of the "user_id" field in the mutation.
func (m *UserSessionMutation) UserID() (r int, exists bool) {
	v := m.user
	if v == nil {
		return
	}
	return *v, true
}

// OldUserID returns the old "user_id" field's value of the UserSession entity.
// If the UserSession object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserSessionMutation) OldUserID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUserID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUserID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserID: %w", err)
	}
	return oldValue.UserID, nil
}

// ResetUserID resets all changes to the "user_id" field.
func (m *UserSessionMutation) ResetUserID() {
	m.user = nil
}

// SetLoginTime sets the "login_time" field.
func (m *UserSessionMutation) SetLoginTime(t time.Time) {
	m.login_time = &t
}

// LoginTime returns the value of the "login_time" field in the mutation.
func (m *UserSessionMutation) LoginTime() (r time.Time, exists bool) {
	v := m.login_time
	if v == nil {
		return
	}
	return *v, true
}

// OldLoginTime returns the old "login_time" field's value of the UserSession entity.
// If the UserSession object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserSessionMutation) OldLoginTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldLoginTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldLoginTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLoginTime: %w", err)
	}
	return oldValue.LoginTime, nil
}

// ResetLoginTime resets all changes to the "login_time" field.
func (m *UserSessionMutation) ResetLoginTime() {
	m.login_time = nil
}

// SetLastSeenTime sets the "last_seen_time" field.
func (m *UserSessionMutation) SetLastSeenTime(t time.Time) {
	m.last_seen_time = &t
}

// LastSeenTime returns the value of the "last_seen_time" field in the mutation.
func (m *UserSessionMutation) LastSeenTime() (r time.Time, exists bool) {
	v := m.last_seen_time
	if v == nil {
		return
	}
	return *v, true
}

// OldLastSeenTime returns the old "last_seen_time" field's value of the UserSession entity.
// If the UserSession object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserSessionMutation) OldLastSeenTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldLastSeenTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldLastSeenTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastSeenTime: %w", err)
	}
	return oldValue.LastSeenTime, nil
}

// ResetLastSeenTime resets all changes to the "last_seen_time" field.
func (m *UserSessionMutation) ResetLastSeenTime() {
	m.last_seen_time = nil
}

// ClearUser clears the "user" edge to the User entity.
func (m *UserSessionMutation) ClearUser() {
	m.cleareduser = true
}

// UserCleared reports if the "user" edge to the User entity was cleared.
func (m *UserSessionMutation) UserCleared() bool {
	return m.cleareduser
}

// UserIDs returns the "user" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// UserID instead. It exists only for internal usage by the builders.
func (m *UserSessionMutation) UserIDs() (ids []int) {
	if id := m.user; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetUser resets all changes to the "user" edge.
func (m *UserSessionMutation) ResetUser() {
	m.user = nil
	m.cleareduser = false
}

// Op returns the operation name.
func (m *UserSessionMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (UserSession).
func (m *UserSessionMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserSessionMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.session_id != nil {
		fields = append(fields, usersession.FieldSessionID)
	}
	if m.user != nil {
		fields = append(fields, usersession.FieldUserID)
	}
	if m.login_time != nil {
		fields = append(fields, usersession.FieldLoginTime)
	}
	if m.last_seen_time != nil {
		fields = append(fields, usersession.FieldLastSeenTime)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserSessionMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case usersession.FieldSessionID:
		return m.SessionID()
	case usersession.FieldUserID:
		return m.UserID()
	case usersession.FieldLoginTime:
		return m.LoginTime()
	case usersession.FieldLastSeenTime:
		return m.LastSeenTime()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserSessionMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case usersession.FieldSessionID:
		return m.OldSessionID(ctx)
	case usersession.FieldUserID:
		return m.OldUserID(ctx)
	case usersession.FieldLoginTime:
		return m.OldLoginTime(ctx)
	case usersession.FieldLastSeenTime:
		return m.OldLastSeenTime(ctx)
	}
	return nil, fmt.Errorf("unknown UserSession field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserSessionMutation) SetField(name string, value ent.Value) error {
	switch name {
	case usersession.FieldSessionID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSessionID(v)
		return nil
	case usersession.FieldUserID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserID(v)
		return nil
	case usersession.FieldLoginTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLoginTime(v)
		return nil
	case usersession.FieldLastSeenTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastSeenTime(v)
		return nil
	}
	return fmt.Errorf("unknown UserSession field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserSessionMutation) AddedFields() []string {
	var fields []string
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserSessionMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserSessionMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown UserSession numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserSessionMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserSessionMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserSessionMutation) ClearField(name string) error {
	return fmt.Errorf("unknown UserSession nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserSessionMutation) ResetField(name string) error {
	switch name {
	case usersession.FieldSessionID:
		m.ResetSessionID()
		return nil
	case usersession.FieldUserID:
		m.ResetUserID()
		return nil
	case usersession.FieldLoginTime:
		m.ResetLoginTime()
		return nil
	case usersession.FieldLastSeenTime:
		m.ResetLastSeenTime()
		return nil
	}
	return fmt.Errorf("unknown UserSession field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserSessionMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.user != nil {
		edges = append(edges, usersession.EdgeUser)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserSessionMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case usersession.EdgeUser:
		if id := m.user; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserSessionMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserSessionMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserSessionMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cleareduser {
		edges = append(edges, usersession.EdgeUser)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserSessionMutation) EdgeCleared(name string) bool {
	switch name {
	case usersession.EdgeUser:
		return m.cleareduser
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserSessionMutation) ClearEdge(name string) error {
	switch name {
	case usersession.EdgeUser:
		m.ClearUser()
		return nil
	}
	return fmt.Errorf("unknown UserSession unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserSessionMutation) ResetEdge(name string) error {
	switch name {
	case usersession.EdgeUser:
		m.ResetUser()
		return nil
	}
	return fmt.Errorf("unknown UserSession edge %s", name)
}
