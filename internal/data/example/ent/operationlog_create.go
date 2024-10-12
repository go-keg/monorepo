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
	"github.com/go-keg/example/internal/data/example/ent/operationlog"
	"github.com/go-keg/example/internal/data/example/ent/user"
)

// OperationLogCreate is the builder for creating a OperationLog entity.
type OperationLogCreate struct {
	config
	mutation *OperationLogMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (olc *OperationLogCreate) SetCreatedAt(t time.Time) *OperationLogCreate {
	olc.mutation.SetCreatedAt(t)
	return olc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (olc *OperationLogCreate) SetNillableCreatedAt(t *time.Time) *OperationLogCreate {
	if t != nil {
		olc.SetCreatedAt(*t)
	}
	return olc
}

// SetUpdatedAt sets the "updated_at" field.
func (olc *OperationLogCreate) SetUpdatedAt(t time.Time) *OperationLogCreate {
	olc.mutation.SetUpdatedAt(t)
	return olc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (olc *OperationLogCreate) SetNillableUpdatedAt(t *time.Time) *OperationLogCreate {
	if t != nil {
		olc.SetUpdatedAt(*t)
	}
	return olc
}

// SetUserID sets the "user_id" field.
func (olc *OperationLogCreate) SetUserID(i int) *OperationLogCreate {
	olc.mutation.SetUserID(i)
	return olc
}

// SetType sets the "type" field.
func (olc *OperationLogCreate) SetType(s string) *OperationLogCreate {
	olc.mutation.SetType(s)
	return olc
}

// SetContext sets the "context" field.
func (olc *OperationLogCreate) SetContext(m map[string]interface{}) *OperationLogCreate {
	olc.mutation.SetContext(m)
	return olc
}

// SetUser sets the "user" edge to the User entity.
func (olc *OperationLogCreate) SetUser(u *User) *OperationLogCreate {
	return olc.SetUserID(u.ID)
}

// Mutation returns the OperationLogMutation object of the builder.
func (olc *OperationLogCreate) Mutation() *OperationLogMutation {
	return olc.mutation
}

// Save creates the OperationLog in the database.
func (olc *OperationLogCreate) Save(ctx context.Context) (*OperationLog, error) {
	olc.defaults()
	return withHooks(ctx, olc.sqlSave, olc.mutation, olc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (olc *OperationLogCreate) SaveX(ctx context.Context) *OperationLog {
	v, err := olc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (olc *OperationLogCreate) Exec(ctx context.Context) error {
	_, err := olc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (olc *OperationLogCreate) ExecX(ctx context.Context) {
	if err := olc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (olc *OperationLogCreate) defaults() {
	if _, ok := olc.mutation.CreatedAt(); !ok {
		v := operationlog.DefaultCreatedAt()
		olc.mutation.SetCreatedAt(v)
	}
	if _, ok := olc.mutation.UpdatedAt(); !ok {
		v := operationlog.DefaultUpdatedAt()
		olc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (olc *OperationLogCreate) check() error {
	if _, ok := olc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "OperationLog.user_id"`)}
	}
	if _, ok := olc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "OperationLog.type"`)}
	}
	if _, ok := olc.mutation.Context(); !ok {
		return &ValidationError{Name: "context", err: errors.New(`ent: missing required field "OperationLog.context"`)}
	}
	if len(olc.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "OperationLog.user"`)}
	}
	return nil
}

func (olc *OperationLogCreate) sqlSave(ctx context.Context) (*OperationLog, error) {
	if err := olc.check(); err != nil {
		return nil, err
	}
	_node, _spec := olc.createSpec()
	if err := sqlgraph.CreateNode(ctx, olc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	olc.mutation.id = &_node.ID
	olc.mutation.done = true
	return _node, nil
}

func (olc *OperationLogCreate) createSpec() (*OperationLog, *sqlgraph.CreateSpec) {
	var (
		_node = &OperationLog{config: olc.config}
		_spec = sqlgraph.NewCreateSpec(operationlog.Table, sqlgraph.NewFieldSpec(operationlog.FieldID, field.TypeInt))
	)
	_spec.OnConflict = olc.conflict
	if value, ok := olc.mutation.CreatedAt(); ok {
		_spec.SetField(operationlog.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := olc.mutation.UpdatedAt(); ok {
		_spec.SetField(operationlog.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := olc.mutation.GetType(); ok {
		_spec.SetField(operationlog.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := olc.mutation.Context(); ok {
		_spec.SetField(operationlog.FieldContext, field.TypeJSON, value)
		_node.Context = value
	}
	if nodes := olc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   operationlog.UserTable,
			Columns: []string{operationlog.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OperationLog.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OperationLogUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (olc *OperationLogCreate) OnConflict(opts ...sql.ConflictOption) *OperationLogUpsertOne {
	olc.conflict = opts
	return &OperationLogUpsertOne{
		create: olc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OperationLog.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (olc *OperationLogCreate) OnConflictColumns(columns ...string) *OperationLogUpsertOne {
	olc.conflict = append(olc.conflict, sql.ConflictColumns(columns...))
	return &OperationLogUpsertOne{
		create: olc,
	}
}

type (
	// OperationLogUpsertOne is the builder for "upsert"-ing
	//  one OperationLog node.
	OperationLogUpsertOne struct {
		create *OperationLogCreate
	}

	// OperationLogUpsert is the "OnConflict" setter.
	OperationLogUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *OperationLogUpsert) SetUpdatedAt(v time.Time) *OperationLogUpsert {
	u.Set(operationlog.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OperationLogUpsert) UpdateUpdatedAt() *OperationLogUpsert {
	u.SetExcluded(operationlog.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *OperationLogUpsert) ClearUpdatedAt() *OperationLogUpsert {
	u.SetNull(operationlog.FieldUpdatedAt)
	return u
}

// SetUserID sets the "user_id" field.
func (u *OperationLogUpsert) SetUserID(v int) *OperationLogUpsert {
	u.Set(operationlog.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *OperationLogUpsert) UpdateUserID() *OperationLogUpsert {
	u.SetExcluded(operationlog.FieldUserID)
	return u
}

// SetType sets the "type" field.
func (u *OperationLogUpsert) SetType(v string) *OperationLogUpsert {
	u.Set(operationlog.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *OperationLogUpsert) UpdateType() *OperationLogUpsert {
	u.SetExcluded(operationlog.FieldType)
	return u
}

// SetContext sets the "context" field.
func (u *OperationLogUpsert) SetContext(v map[string]interface{}) *OperationLogUpsert {
	u.Set(operationlog.FieldContext, v)
	return u
}

// UpdateContext sets the "context" field to the value that was provided on create.
func (u *OperationLogUpsert) UpdateContext() *OperationLogUpsert {
	u.SetExcluded(operationlog.FieldContext)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.OperationLog.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *OperationLogUpsertOne) UpdateNewValues() *OperationLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(operationlog.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OperationLog.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *OperationLogUpsertOne) Ignore() *OperationLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OperationLogUpsertOne) DoNothing() *OperationLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OperationLogCreate.OnConflict
// documentation for more info.
func (u *OperationLogUpsertOne) Update(set func(*OperationLogUpsert)) *OperationLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OperationLogUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OperationLogUpsertOne) SetUpdatedAt(v time.Time) *OperationLogUpsertOne {
	return u.Update(func(s *OperationLogUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OperationLogUpsertOne) UpdateUpdatedAt() *OperationLogUpsertOne {
	return u.Update(func(s *OperationLogUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *OperationLogUpsertOne) ClearUpdatedAt() *OperationLogUpsertOne {
	return u.Update(func(s *OperationLogUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *OperationLogUpsertOne) SetUserID(v int) *OperationLogUpsertOne {
	return u.Update(func(s *OperationLogUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *OperationLogUpsertOne) UpdateUserID() *OperationLogUpsertOne {
	return u.Update(func(s *OperationLogUpsert) {
		s.UpdateUserID()
	})
}

// SetType sets the "type" field.
func (u *OperationLogUpsertOne) SetType(v string) *OperationLogUpsertOne {
	return u.Update(func(s *OperationLogUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *OperationLogUpsertOne) UpdateType() *OperationLogUpsertOne {
	return u.Update(func(s *OperationLogUpsert) {
		s.UpdateType()
	})
}

// SetContext sets the "context" field.
func (u *OperationLogUpsertOne) SetContext(v map[string]interface{}) *OperationLogUpsertOne {
	return u.Update(func(s *OperationLogUpsert) {
		s.SetContext(v)
	})
}

// UpdateContext sets the "context" field to the value that was provided on create.
func (u *OperationLogUpsertOne) UpdateContext() *OperationLogUpsertOne {
	return u.Update(func(s *OperationLogUpsert) {
		s.UpdateContext()
	})
}

// Exec executes the query.
func (u *OperationLogUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OperationLogCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OperationLogUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *OperationLogUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *OperationLogUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// OperationLogCreateBulk is the builder for creating many OperationLog entities in bulk.
type OperationLogCreateBulk struct {
	config
	err      error
	builders []*OperationLogCreate
	conflict []sql.ConflictOption
}

// Save creates the OperationLog entities in the database.
func (olcb *OperationLogCreateBulk) Save(ctx context.Context) ([]*OperationLog, error) {
	if olcb.err != nil {
		return nil, olcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(olcb.builders))
	nodes := make([]*OperationLog, len(olcb.builders))
	mutators := make([]Mutator, len(olcb.builders))
	for i := range olcb.builders {
		func(i int, root context.Context) {
			builder := olcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OperationLogMutation)
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
					_, err = mutators[i+1].Mutate(root, olcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = olcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, olcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, olcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (olcb *OperationLogCreateBulk) SaveX(ctx context.Context) []*OperationLog {
	v, err := olcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (olcb *OperationLogCreateBulk) Exec(ctx context.Context) error {
	_, err := olcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (olcb *OperationLogCreateBulk) ExecX(ctx context.Context) {
	if err := olcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OperationLog.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OperationLogUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (olcb *OperationLogCreateBulk) OnConflict(opts ...sql.ConflictOption) *OperationLogUpsertBulk {
	olcb.conflict = opts
	return &OperationLogUpsertBulk{
		create: olcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OperationLog.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (olcb *OperationLogCreateBulk) OnConflictColumns(columns ...string) *OperationLogUpsertBulk {
	olcb.conflict = append(olcb.conflict, sql.ConflictColumns(columns...))
	return &OperationLogUpsertBulk{
		create: olcb,
	}
}

// OperationLogUpsertBulk is the builder for "upsert"-ing
// a bulk of OperationLog nodes.
type OperationLogUpsertBulk struct {
	create *OperationLogCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.OperationLog.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *OperationLogUpsertBulk) UpdateNewValues() *OperationLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(operationlog.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OperationLog.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *OperationLogUpsertBulk) Ignore() *OperationLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OperationLogUpsertBulk) DoNothing() *OperationLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OperationLogCreateBulk.OnConflict
// documentation for more info.
func (u *OperationLogUpsertBulk) Update(set func(*OperationLogUpsert)) *OperationLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OperationLogUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OperationLogUpsertBulk) SetUpdatedAt(v time.Time) *OperationLogUpsertBulk {
	return u.Update(func(s *OperationLogUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OperationLogUpsertBulk) UpdateUpdatedAt() *OperationLogUpsertBulk {
	return u.Update(func(s *OperationLogUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *OperationLogUpsertBulk) ClearUpdatedAt() *OperationLogUpsertBulk {
	return u.Update(func(s *OperationLogUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *OperationLogUpsertBulk) SetUserID(v int) *OperationLogUpsertBulk {
	return u.Update(func(s *OperationLogUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *OperationLogUpsertBulk) UpdateUserID() *OperationLogUpsertBulk {
	return u.Update(func(s *OperationLogUpsert) {
		s.UpdateUserID()
	})
}

// SetType sets the "type" field.
func (u *OperationLogUpsertBulk) SetType(v string) *OperationLogUpsertBulk {
	return u.Update(func(s *OperationLogUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *OperationLogUpsertBulk) UpdateType() *OperationLogUpsertBulk {
	return u.Update(func(s *OperationLogUpsert) {
		s.UpdateType()
	})
}

// SetContext sets the "context" field.
func (u *OperationLogUpsertBulk) SetContext(v map[string]interface{}) *OperationLogUpsertBulk {
	return u.Update(func(s *OperationLogUpsert) {
		s.SetContext(v)
	})
}

// UpdateContext sets the "context" field to the value that was provided on create.
func (u *OperationLogUpsertBulk) UpdateContext() *OperationLogUpsertBulk {
	return u.Update(func(s *OperationLogUpsert) {
		s.UpdateContext()
	})
}

// Exec executes the query.
func (u *OperationLogUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the OperationLogCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OperationLogCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OperationLogUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
