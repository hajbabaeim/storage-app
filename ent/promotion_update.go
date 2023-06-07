// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"storage-app/ent/predicate"
	"storage-app/ent/promotion"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PromotionUpdate is the builder for updating Promotion entities.
type PromotionUpdate struct {
	config
	hooks    []Hook
	mutation *PromotionMutation
}

// Where appends a list predicates to the PromotionUpdate builder.
func (pu *PromotionUpdate) Where(ps ...predicate.Promotion) *PromotionUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetPid sets the "pid" field.
func (pu *PromotionUpdate) SetPid(s string) *PromotionUpdate {
	pu.mutation.SetPid(s)
	return pu
}

// SetPrice sets the "price" field.
func (pu *PromotionUpdate) SetPrice(f float64) *PromotionUpdate {
	pu.mutation.ResetPrice()
	pu.mutation.SetPrice(f)
	return pu
}

// AddPrice adds f to the "price" field.
func (pu *PromotionUpdate) AddPrice(f float64) *PromotionUpdate {
	pu.mutation.AddPrice(f)
	return pu
}

// SetExpirationDate sets the "expiration_date" field.
func (pu *PromotionUpdate) SetExpirationDate(t time.Time) *PromotionUpdate {
	pu.mutation.SetExpirationDate(t)
	return pu
}

// Mutation returns the PromotionMutation object of the builder.
func (pu *PromotionUpdate) Mutation() *PromotionMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PromotionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PromotionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PromotionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PromotionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PromotionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(promotion.Table, promotion.Columns, sqlgraph.NewFieldSpec(promotion.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Pid(); ok {
		_spec.SetField(promotion.FieldPid, field.TypeString, value)
	}
	if value, ok := pu.mutation.Price(); ok {
		_spec.SetField(promotion.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.AddedPrice(); ok {
		_spec.AddField(promotion.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.ExpirationDate(); ok {
		_spec.SetField(promotion.FieldExpirationDate, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{promotion.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PromotionUpdateOne is the builder for updating a single Promotion entity.
type PromotionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PromotionMutation
}

// SetPid sets the "pid" field.
func (puo *PromotionUpdateOne) SetPid(s string) *PromotionUpdateOne {
	puo.mutation.SetPid(s)
	return puo
}

// SetPrice sets the "price" field.
func (puo *PromotionUpdateOne) SetPrice(f float64) *PromotionUpdateOne {
	puo.mutation.ResetPrice()
	puo.mutation.SetPrice(f)
	return puo
}

// AddPrice adds f to the "price" field.
func (puo *PromotionUpdateOne) AddPrice(f float64) *PromotionUpdateOne {
	puo.mutation.AddPrice(f)
	return puo
}

// SetExpirationDate sets the "expiration_date" field.
func (puo *PromotionUpdateOne) SetExpirationDate(t time.Time) *PromotionUpdateOne {
	puo.mutation.SetExpirationDate(t)
	return puo
}

// Mutation returns the PromotionMutation object of the builder.
func (puo *PromotionUpdateOne) Mutation() *PromotionMutation {
	return puo.mutation
}

// Where appends a list predicates to the PromotionUpdate builder.
func (puo *PromotionUpdateOne) Where(ps ...predicate.Promotion) *PromotionUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PromotionUpdateOne) Select(field string, fields ...string) *PromotionUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Promotion entity.
func (puo *PromotionUpdateOne) Save(ctx context.Context) (*Promotion, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PromotionUpdateOne) SaveX(ctx context.Context) *Promotion {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PromotionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PromotionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PromotionUpdateOne) sqlSave(ctx context.Context) (_node *Promotion, err error) {
	_spec := sqlgraph.NewUpdateSpec(promotion.Table, promotion.Columns, sqlgraph.NewFieldSpec(promotion.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Promotion.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, promotion.FieldID)
		for _, f := range fields {
			if !promotion.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != promotion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Pid(); ok {
		_spec.SetField(promotion.FieldPid, field.TypeString, value)
	}
	if value, ok := puo.mutation.Price(); ok {
		_spec.SetField(promotion.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.AddedPrice(); ok {
		_spec.AddField(promotion.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.ExpirationDate(); ok {
		_spec.SetField(promotion.FieldExpirationDate, field.TypeTime, value)
	}
	_node = &Promotion{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{promotion.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
