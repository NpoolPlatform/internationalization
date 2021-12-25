// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/internationalization/pkg/db/ent/message"
	"github.com/google/uuid"
)

// MessageCreate is the builder for creating a Message entity.
type MessageCreate struct {
	config
	mutation *MessageMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAppID sets the "app_id" field.
func (mc *MessageCreate) SetAppID(u uuid.UUID) *MessageCreate {
	mc.mutation.SetAppID(u)
	return mc
}

// SetMessageID sets the "message_id" field.
func (mc *MessageCreate) SetMessageID(u uuid.UUID) *MessageCreate {
	mc.mutation.SetMessageID(u)
	return mc
}

// SetLang sets the "lang" field.
func (mc *MessageCreate) SetLang(m message.Lang) *MessageCreate {
	mc.mutation.SetLang(m)
	return mc
}

// SetMessage sets the "message" field.
func (mc *MessageCreate) SetMessage(s string) *MessageCreate {
	mc.mutation.SetMessage(s)
	return mc
}

// SetCreateAt sets the "create_at" field.
func (mc *MessageCreate) SetCreateAt(u uint32) *MessageCreate {
	mc.mutation.SetCreateAt(u)
	return mc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (mc *MessageCreate) SetNillableCreateAt(u *uint32) *MessageCreate {
	if u != nil {
		mc.SetCreateAt(*u)
	}
	return mc
}

// SetUpdateAt sets the "update_at" field.
func (mc *MessageCreate) SetUpdateAt(u uint32) *MessageCreate {
	mc.mutation.SetUpdateAt(u)
	return mc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (mc *MessageCreate) SetNillableUpdateAt(u *uint32) *MessageCreate {
	if u != nil {
		mc.SetUpdateAt(*u)
	}
	return mc
}

// SetDeleteAt sets the "delete_at" field.
func (mc *MessageCreate) SetDeleteAt(u uint32) *MessageCreate {
	mc.mutation.SetDeleteAt(u)
	return mc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (mc *MessageCreate) SetNillableDeleteAt(u *uint32) *MessageCreate {
	if u != nil {
		mc.SetDeleteAt(*u)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MessageCreate) SetID(u uuid.UUID) *MessageCreate {
	mc.mutation.SetID(u)
	return mc
}

// Mutation returns the MessageMutation object of the builder.
func (mc *MessageCreate) Mutation() *MessageMutation {
	return mc.mutation
}

// Save creates the Message in the database.
func (mc *MessageCreate) Save(ctx context.Context) (*Message, error) {
	var (
		err  error
		node *Message
	)
	mc.defaults()
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MessageCreate) SaveX(ctx context.Context) *Message {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MessageCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MessageCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MessageCreate) defaults() {
	if _, ok := mc.mutation.CreateAt(); !ok {
		v := message.DefaultCreateAt()
		mc.mutation.SetCreateAt(v)
	}
	if _, ok := mc.mutation.UpdateAt(); !ok {
		v := message.DefaultUpdateAt()
		mc.mutation.SetUpdateAt(v)
	}
	if _, ok := mc.mutation.DeleteAt(); !ok {
		v := message.DefaultDeleteAt()
		mc.mutation.SetDeleteAt(v)
	}
	if _, ok := mc.mutation.ID(); !ok {
		v := message.DefaultID()
		mc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MessageCreate) check() error {
	if _, ok := mc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "app_id"`)}
	}
	if _, ok := mc.mutation.MessageID(); !ok {
		return &ValidationError{Name: "message_id", err: errors.New(`ent: missing required field "message_id"`)}
	}
	if _, ok := mc.mutation.Lang(); !ok {
		return &ValidationError{Name: "lang", err: errors.New(`ent: missing required field "lang"`)}
	}
	if v, ok := mc.mutation.Lang(); ok {
		if err := message.LangValidator(v); err != nil {
			return &ValidationError{Name: "lang", err: fmt.Errorf(`ent: validator failed for field "lang": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Message(); !ok {
		return &ValidationError{Name: "message", err: errors.New(`ent: missing required field "message"`)}
	}
	if _, ok := mc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := mc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := mc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (mc *MessageCreate) sqlSave(ctx context.Context) (*Message, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (mc *MessageCreate) createSpec() (*Message, *sqlgraph.CreateSpec) {
	var (
		_node = &Message{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: message.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: message.FieldID,
			},
		}
	)
	_spec.OnConflict = mc.conflict
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: message.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := mc.mutation.MessageID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: message.FieldMessageID,
		})
		_node.MessageID = value
	}
	if value, ok := mc.mutation.Lang(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: message.FieldLang,
		})
		_node.Lang = value
	}
	if value, ok := mc.mutation.Message(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: message.FieldMessage,
		})
		_node.Message = value
	}
	if value, ok := mc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: message.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := mc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: message.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := mc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: message.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Message.Create().
//		SetAppID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MessageUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (mc *MessageCreate) OnConflict(opts ...sql.ConflictOption) *MessageUpsertOne {
	mc.conflict = opts
	return &MessageUpsertOne{
		create: mc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (mc *MessageCreate) OnConflictColumns(columns ...string) *MessageUpsertOne {
	mc.conflict = append(mc.conflict, sql.ConflictColumns(columns...))
	return &MessageUpsertOne{
		create: mc,
	}
}

type (
	// MessageUpsertOne is the builder for "upsert"-ing
	//  one Message node.
	MessageUpsertOne struct {
		create *MessageCreate
	}

	// MessageUpsert is the "OnConflict" setter.
	MessageUpsert struct {
		*sql.UpdateSet
	}
)

// SetAppID sets the "app_id" field.
func (u *MessageUpsert) SetAppID(v uuid.UUID) *MessageUpsert {
	u.Set(message.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *MessageUpsert) UpdateAppID() *MessageUpsert {
	u.SetExcluded(message.FieldAppID)
	return u
}

// SetMessageID sets the "message_id" field.
func (u *MessageUpsert) SetMessageID(v uuid.UUID) *MessageUpsert {
	u.Set(message.FieldMessageID, v)
	return u
}

// UpdateMessageID sets the "message_id" field to the value that was provided on create.
func (u *MessageUpsert) UpdateMessageID() *MessageUpsert {
	u.SetExcluded(message.FieldMessageID)
	return u
}

// SetLang sets the "lang" field.
func (u *MessageUpsert) SetLang(v message.Lang) *MessageUpsert {
	u.Set(message.FieldLang, v)
	return u
}

// UpdateLang sets the "lang" field to the value that was provided on create.
func (u *MessageUpsert) UpdateLang() *MessageUpsert {
	u.SetExcluded(message.FieldLang)
	return u
}

// SetMessage sets the "message" field.
func (u *MessageUpsert) SetMessage(v string) *MessageUpsert {
	u.Set(message.FieldMessage, v)
	return u
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *MessageUpsert) UpdateMessage() *MessageUpsert {
	u.SetExcluded(message.FieldMessage)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *MessageUpsert) SetCreateAt(v uint32) *MessageUpsert {
	u.Set(message.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *MessageUpsert) UpdateCreateAt() *MessageUpsert {
	u.SetExcluded(message.FieldCreateAt)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *MessageUpsert) SetUpdateAt(v uint32) *MessageUpsert {
	u.Set(message.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *MessageUpsert) UpdateUpdateAt() *MessageUpsert {
	u.SetExcluded(message.FieldUpdateAt)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *MessageUpsert) SetDeleteAt(v uint32) *MessageUpsert {
	u.Set(message.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *MessageUpsert) UpdateDeleteAt() *MessageUpsert {
	u.SetExcluded(message.FieldDeleteAt)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(message.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *MessageUpsertOne) UpdateNewValues() *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(message.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Message.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *MessageUpsertOne) Ignore() *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MessageUpsertOne) DoNothing() *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MessageCreate.OnConflict
// documentation for more info.
func (u *MessageUpsertOne) Update(set func(*MessageUpsert)) *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MessageUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *MessageUpsertOne) SetAppID(v uuid.UUID) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateAppID() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateAppID()
	})
}

// SetMessageID sets the "message_id" field.
func (u *MessageUpsertOne) SetMessageID(v uuid.UUID) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetMessageID(v)
	})
}

// UpdateMessageID sets the "message_id" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateMessageID() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateMessageID()
	})
}

// SetLang sets the "lang" field.
func (u *MessageUpsertOne) SetLang(v message.Lang) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetLang(v)
	})
}

// UpdateLang sets the "lang" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateLang() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateLang()
	})
}

// SetMessage sets the "message" field.
func (u *MessageUpsertOne) SetMessage(v string) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateMessage() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateMessage()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *MessageUpsertOne) SetCreateAt(v uint32) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateCreateAt() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *MessageUpsertOne) SetUpdateAt(v uint32) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateUpdateAt() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *MessageUpsertOne) SetDeleteAt(v uint32) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateDeleteAt() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *MessageUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MessageCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MessageUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MessageUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: MessageUpsertOne.ID is not supported by MySQL driver. Use MessageUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MessageUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MessageCreateBulk is the builder for creating many Message entities in bulk.
type MessageCreateBulk struct {
	config
	builders []*MessageCreate
	conflict []sql.ConflictOption
}

// Save creates the Message entities in the database.
func (mcb *MessageCreateBulk) Save(ctx context.Context) ([]*Message, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Message, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessageMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MessageCreateBulk) SaveX(ctx context.Context) []*Message {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MessageCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MessageCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Message.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MessageUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (mcb *MessageCreateBulk) OnConflict(opts ...sql.ConflictOption) *MessageUpsertBulk {
	mcb.conflict = opts
	return &MessageUpsertBulk{
		create: mcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (mcb *MessageCreateBulk) OnConflictColumns(columns ...string) *MessageUpsertBulk {
	mcb.conflict = append(mcb.conflict, sql.ConflictColumns(columns...))
	return &MessageUpsertBulk{
		create: mcb,
	}
}

// MessageUpsertBulk is the builder for "upsert"-ing
// a bulk of Message nodes.
type MessageUpsertBulk struct {
	create *MessageCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(message.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *MessageUpsertBulk) UpdateNewValues() *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(message.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *MessageUpsertBulk) Ignore() *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MessageUpsertBulk) DoNothing() *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MessageCreateBulk.OnConflict
// documentation for more info.
func (u *MessageUpsertBulk) Update(set func(*MessageUpsert)) *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MessageUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *MessageUpsertBulk) SetAppID(v uuid.UUID) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateAppID() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateAppID()
	})
}

// SetMessageID sets the "message_id" field.
func (u *MessageUpsertBulk) SetMessageID(v uuid.UUID) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetMessageID(v)
	})
}

// UpdateMessageID sets the "message_id" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateMessageID() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateMessageID()
	})
}

// SetLang sets the "lang" field.
func (u *MessageUpsertBulk) SetLang(v message.Lang) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetLang(v)
	})
}

// UpdateLang sets the "lang" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateLang() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateLang()
	})
}

// SetMessage sets the "message" field.
func (u *MessageUpsertBulk) SetMessage(v string) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateMessage() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateMessage()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *MessageUpsertBulk) SetCreateAt(v uint32) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateCreateAt() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *MessageUpsertBulk) SetUpdateAt(v uint32) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateUpdateAt() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *MessageUpsertBulk) SetDeleteAt(v uint32) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateDeleteAt() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *MessageUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MessageCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MessageCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MessageUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}