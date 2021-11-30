// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/discountpool"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// DiscountPoolQuery is the builder for querying DiscountPool entities.
type DiscountPoolQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DiscountPool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DiscountPoolQuery builder.
func (dpq *DiscountPoolQuery) Where(ps ...predicate.DiscountPool) *DiscountPoolQuery {
	dpq.predicates = append(dpq.predicates, ps...)
	return dpq
}

// Limit adds a limit step to the query.
func (dpq *DiscountPoolQuery) Limit(limit int) *DiscountPoolQuery {
	dpq.limit = &limit
	return dpq
}

// Offset adds an offset step to the query.
func (dpq *DiscountPoolQuery) Offset(offset int) *DiscountPoolQuery {
	dpq.offset = &offset
	return dpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dpq *DiscountPoolQuery) Unique(unique bool) *DiscountPoolQuery {
	dpq.unique = &unique
	return dpq
}

// Order adds an order step to the query.
func (dpq *DiscountPoolQuery) Order(o ...OrderFunc) *DiscountPoolQuery {
	dpq.order = append(dpq.order, o...)
	return dpq
}

// First returns the first DiscountPool entity from the query.
// Returns a *NotFoundError when no DiscountPool was found.
func (dpq *DiscountPoolQuery) First(ctx context.Context) (*DiscountPool, error) {
	nodes, err := dpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{discountpool.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dpq *DiscountPoolQuery) FirstX(ctx context.Context) *DiscountPool {
	node, err := dpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DiscountPool ID from the query.
// Returns a *NotFoundError when no DiscountPool ID was found.
func (dpq *DiscountPoolQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{discountpool.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dpq *DiscountPoolQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := dpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DiscountPool entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one DiscountPool entity is not found.
// Returns a *NotFoundError when no DiscountPool entities are found.
func (dpq *DiscountPoolQuery) Only(ctx context.Context) (*DiscountPool, error) {
	nodes, err := dpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{discountpool.Label}
	default:
		return nil, &NotSingularError{discountpool.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dpq *DiscountPoolQuery) OnlyX(ctx context.Context) *DiscountPool {
	node, err := dpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DiscountPool ID in the query.
// Returns a *NotSingularError when exactly one DiscountPool ID is not found.
// Returns a *NotFoundError when no entities are found.
func (dpq *DiscountPoolQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{discountpool.Label}
	default:
		err = &NotSingularError{discountpool.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dpq *DiscountPoolQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := dpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DiscountPools.
func (dpq *DiscountPoolQuery) All(ctx context.Context) ([]*DiscountPool, error) {
	if err := dpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dpq *DiscountPoolQuery) AllX(ctx context.Context) []*DiscountPool {
	nodes, err := dpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DiscountPool IDs.
func (dpq *DiscountPoolQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := dpq.Select(discountpool.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dpq *DiscountPoolQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dpq *DiscountPoolQuery) Count(ctx context.Context) (int, error) {
	if err := dpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dpq *DiscountPoolQuery) CountX(ctx context.Context) int {
	count, err := dpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dpq *DiscountPoolQuery) Exist(ctx context.Context) (bool, error) {
	if err := dpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dpq *DiscountPoolQuery) ExistX(ctx context.Context) bool {
	exist, err := dpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DiscountPoolQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dpq *DiscountPoolQuery) Clone() *DiscountPoolQuery {
	if dpq == nil {
		return nil
	}
	return &DiscountPoolQuery{
		config:     dpq.config,
		limit:      dpq.limit,
		offset:     dpq.offset,
		order:      append([]OrderFunc{}, dpq.order...),
		predicates: append([]predicate.DiscountPool{}, dpq.predicates...),
		// clone intermediate query.
		sql:  dpq.sql.Clone(),
		path: dpq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AppID uuid.UUID `json:"app_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DiscountPool.Query().
//		GroupBy(discountpool.FieldAppID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (dpq *DiscountPoolQuery) GroupBy(field string, fields ...string) *DiscountPoolGroupBy {
	group := &DiscountPoolGroupBy{config: dpq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dpq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AppID uuid.UUID `json:"app_id,omitempty"`
//	}
//
//	client.DiscountPool.Query().
//		Select(discountpool.FieldAppID).
//		Scan(ctx, &v)
//
func (dpq *DiscountPoolQuery) Select(fields ...string) *DiscountPoolSelect {
	dpq.fields = append(dpq.fields, fields...)
	return &DiscountPoolSelect{DiscountPoolQuery: dpq}
}

func (dpq *DiscountPoolQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dpq.fields {
		if !discountpool.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dpq.path != nil {
		prev, err := dpq.path(ctx)
		if err != nil {
			return err
		}
		dpq.sql = prev
	}
	return nil
}

func (dpq *DiscountPoolQuery) sqlAll(ctx context.Context) ([]*DiscountPool, error) {
	var (
		nodes = []*DiscountPool{}
		_spec = dpq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &DiscountPool{config: dpq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, dpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (dpq *DiscountPoolQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dpq.querySpec()
	return sqlgraph.CountNodes(ctx, dpq.driver, _spec)
}

func (dpq *DiscountPoolQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dpq *DiscountPoolQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   discountpool.Table,
			Columns: discountpool.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: discountpool.FieldID,
			},
		},
		From:   dpq.sql,
		Unique: true,
	}
	if unique := dpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, discountpool.FieldID)
		for i := range fields {
			if fields[i] != discountpool.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dpq *DiscountPoolQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dpq.driver.Dialect())
	t1 := builder.Table(discountpool.Table)
	columns := dpq.fields
	if len(columns) == 0 {
		columns = discountpool.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dpq.sql != nil {
		selector = dpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range dpq.predicates {
		p(selector)
	}
	for _, p := range dpq.order {
		p(selector)
	}
	if offset := dpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DiscountPoolGroupBy is the group-by builder for DiscountPool entities.
type DiscountPoolGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dpgb *DiscountPoolGroupBy) Aggregate(fns ...AggregateFunc) *DiscountPoolGroupBy {
	dpgb.fns = append(dpgb.fns, fns...)
	return dpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dpgb *DiscountPoolGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dpgb.path(ctx)
	if err != nil {
		return err
	}
	dpgb.sql = query
	return dpgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dpgb *DiscountPoolGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := dpgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DiscountPoolGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(dpgb.fields) > 1 {
		return nil, errors.New("ent: DiscountPoolGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := dpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dpgb *DiscountPoolGroupBy) StringsX(ctx context.Context) []string {
	v, err := dpgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DiscountPoolGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dpgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discountpool.Label}
	default:
		err = fmt.Errorf("ent: DiscountPoolGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dpgb *DiscountPoolGroupBy) StringX(ctx context.Context) string {
	v, err := dpgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DiscountPoolGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(dpgb.fields) > 1 {
		return nil, errors.New("ent: DiscountPoolGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := dpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dpgb *DiscountPoolGroupBy) IntsX(ctx context.Context) []int {
	v, err := dpgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DiscountPoolGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dpgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discountpool.Label}
	default:
		err = fmt.Errorf("ent: DiscountPoolGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dpgb *DiscountPoolGroupBy) IntX(ctx context.Context) int {
	v, err := dpgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DiscountPoolGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(dpgb.fields) > 1 {
		return nil, errors.New("ent: DiscountPoolGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := dpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dpgb *DiscountPoolGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := dpgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DiscountPoolGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dpgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discountpool.Label}
	default:
		err = fmt.Errorf("ent: DiscountPoolGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dpgb *DiscountPoolGroupBy) Float64X(ctx context.Context) float64 {
	v, err := dpgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DiscountPoolGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(dpgb.fields) > 1 {
		return nil, errors.New("ent: DiscountPoolGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := dpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dpgb *DiscountPoolGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := dpgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dpgb *DiscountPoolGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dpgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discountpool.Label}
	default:
		err = fmt.Errorf("ent: DiscountPoolGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dpgb *DiscountPoolGroupBy) BoolX(ctx context.Context) bool {
	v, err := dpgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dpgb *DiscountPoolGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dpgb.fields {
		if !discountpool.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dpgb *DiscountPoolGroupBy) sqlQuery() *sql.Selector {
	selector := dpgb.sql.Select()
	aggregation := make([]string, 0, len(dpgb.fns))
	for _, fn := range dpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dpgb.fields)+len(dpgb.fns))
		for _, f := range dpgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dpgb.fields...)...)
}

// DiscountPoolSelect is the builder for selecting fields of DiscountPool entities.
type DiscountPoolSelect struct {
	*DiscountPoolQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (dps *DiscountPoolSelect) Scan(ctx context.Context, v interface{}) error {
	if err := dps.prepareQuery(ctx); err != nil {
		return err
	}
	dps.sql = dps.DiscountPoolQuery.sqlQuery(ctx)
	return dps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dps *DiscountPoolSelect) ScanX(ctx context.Context, v interface{}) {
	if err := dps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (dps *DiscountPoolSelect) Strings(ctx context.Context) ([]string, error) {
	if len(dps.fields) > 1 {
		return nil, errors.New("ent: DiscountPoolSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := dps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dps *DiscountPoolSelect) StringsX(ctx context.Context) []string {
	v, err := dps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (dps *DiscountPoolSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discountpool.Label}
	default:
		err = fmt.Errorf("ent: DiscountPoolSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dps *DiscountPoolSelect) StringX(ctx context.Context) string {
	v, err := dps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (dps *DiscountPoolSelect) Ints(ctx context.Context) ([]int, error) {
	if len(dps.fields) > 1 {
		return nil, errors.New("ent: DiscountPoolSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := dps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dps *DiscountPoolSelect) IntsX(ctx context.Context) []int {
	v, err := dps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (dps *DiscountPoolSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discountpool.Label}
	default:
		err = fmt.Errorf("ent: DiscountPoolSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dps *DiscountPoolSelect) IntX(ctx context.Context) int {
	v, err := dps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (dps *DiscountPoolSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(dps.fields) > 1 {
		return nil, errors.New("ent: DiscountPoolSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := dps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dps *DiscountPoolSelect) Float64sX(ctx context.Context) []float64 {
	v, err := dps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (dps *DiscountPoolSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discountpool.Label}
	default:
		err = fmt.Errorf("ent: DiscountPoolSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dps *DiscountPoolSelect) Float64X(ctx context.Context) float64 {
	v, err := dps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (dps *DiscountPoolSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(dps.fields) > 1 {
		return nil, errors.New("ent: DiscountPoolSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := dps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dps *DiscountPoolSelect) BoolsX(ctx context.Context) []bool {
	v, err := dps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (dps *DiscountPoolSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{discountpool.Label}
	default:
		err = fmt.Errorf("ent: DiscountPoolSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dps *DiscountPoolSelect) BoolX(ctx context.Context) bool {
	v, err := dps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dps *DiscountPoolSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dps.sql.Query()
	if err := dps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}