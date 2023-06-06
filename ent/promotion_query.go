// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"storage-app/ent/predicate"
	"storage-app/ent/promotion"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PromotionQuery is the builder for querying Promotion entities.
type PromotionQuery struct {
	config
	ctx        *QueryContext
	order      []promotion.OrderOption
	inters     []Interceptor
	predicates []predicate.Promotion
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PromotionQuery builder.
func (pq *PromotionQuery) Where(ps ...predicate.Promotion) *PromotionQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PromotionQuery) Limit(limit int) *PromotionQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PromotionQuery) Offset(offset int) *PromotionQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PromotionQuery) Unique(unique bool) *PromotionQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PromotionQuery) Order(o ...promotion.OrderOption) *PromotionQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// First returns the first Promotion entity from the query.
// Returns a *NotFoundError when no Promotion was found.
func (pq *PromotionQuery) First(ctx context.Context) (*Promotion, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{promotion.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PromotionQuery) FirstX(ctx context.Context) *Promotion {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Promotion ID from the query.
// Returns a *NotFoundError when no Promotion ID was found.
func (pq *PromotionQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{promotion.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PromotionQuery) FirstIDX(ctx context.Context) string {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Promotion entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Promotion entity is found.
// Returns a *NotFoundError when no Promotion entities are found.
func (pq *PromotionQuery) Only(ctx context.Context) (*Promotion, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{promotion.Label}
	default:
		return nil, &NotSingularError{promotion.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PromotionQuery) OnlyX(ctx context.Context) *Promotion {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Promotion ID in the query.
// Returns a *NotSingularError when more than one Promotion ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PromotionQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{promotion.Label}
	default:
		err = &NotSingularError{promotion.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PromotionQuery) OnlyIDX(ctx context.Context) string {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Promotions.
func (pq *PromotionQuery) All(ctx context.Context) ([]*Promotion, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Promotion, *PromotionQuery]()
	return withInterceptors[[]*Promotion](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PromotionQuery) AllX(ctx context.Context) []*Promotion {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Promotion IDs.
func (pq *PromotionQuery) IDs(ctx context.Context) (ids []string, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(promotion.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PromotionQuery) IDsX(ctx context.Context) []string {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PromotionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PromotionQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PromotionQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PromotionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PromotionQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PromotionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PromotionQuery) Clone() *PromotionQuery {
	if pq == nil {
		return nil
	}
	return &PromotionQuery{
		config:     pq.config,
		ctx:        pq.ctx.Clone(),
		order:      append([]promotion.OrderOption{}, pq.order...),
		inters:     append([]Interceptor{}, pq.inters...),
		predicates: append([]predicate.Promotion{}, pq.predicates...),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Price float64 `json:"price,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Promotion.Query().
//		GroupBy(promotion.FieldPrice).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PromotionQuery) GroupBy(field string, fields ...string) *PromotionGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PromotionGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = promotion.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Price float64 `json:"price,omitempty"`
//	}
//
//	client.Promotion.Query().
//		Select(promotion.FieldPrice).
//		Scan(ctx, &v)
func (pq *PromotionQuery) Select(fields ...string) *PromotionSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PromotionSelect{PromotionQuery: pq}
	sbuild.label = promotion.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PromotionSelect configured with the given aggregations.
func (pq *PromotionQuery) Aggregate(fns ...AggregateFunc) *PromotionSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PromotionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !promotion.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PromotionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Promotion, error) {
	var (
		nodes = []*Promotion{}
		_spec = pq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Promotion).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Promotion{config: pq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (pq *PromotionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PromotionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(promotion.Table, promotion.Columns, sqlgraph.NewFieldSpec(promotion.FieldID, field.TypeString))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, promotion.FieldID)
		for i := range fields {
			if fields[i] != promotion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PromotionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(promotion.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = promotion.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PromotionGroupBy is the group-by builder for Promotion entities.
type PromotionGroupBy struct {
	selector
	build *PromotionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PromotionGroupBy) Aggregate(fns ...AggregateFunc) *PromotionGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PromotionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PromotionQuery, *PromotionGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PromotionGroupBy) sqlScan(ctx context.Context, root *PromotionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PromotionSelect is the builder for selecting fields of Promotion entities.
type PromotionSelect struct {
	*PromotionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PromotionSelect) Aggregate(fns ...AggregateFunc) *PromotionSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PromotionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PromotionQuery, *PromotionSelect](ctx, ps.PromotionQuery, ps, ps.inters, v)
}

func (ps *PromotionSelect) sqlScan(ctx context.Context, root *PromotionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
