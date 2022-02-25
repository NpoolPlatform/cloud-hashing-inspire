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
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/appinvitationsetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppInvitationSettingQuery is the builder for querying AppInvitationSetting entities.
type AppInvitationSettingQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppInvitationSetting
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppInvitationSettingQuery builder.
func (aisq *AppInvitationSettingQuery) Where(ps ...predicate.AppInvitationSetting) *AppInvitationSettingQuery {
	aisq.predicates = append(aisq.predicates, ps...)
	return aisq
}

// Limit adds a limit step to the query.
func (aisq *AppInvitationSettingQuery) Limit(limit int) *AppInvitationSettingQuery {
	aisq.limit = &limit
	return aisq
}

// Offset adds an offset step to the query.
func (aisq *AppInvitationSettingQuery) Offset(offset int) *AppInvitationSettingQuery {
	aisq.offset = &offset
	return aisq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aisq *AppInvitationSettingQuery) Unique(unique bool) *AppInvitationSettingQuery {
	aisq.unique = &unique
	return aisq
}

// Order adds an order step to the query.
func (aisq *AppInvitationSettingQuery) Order(o ...OrderFunc) *AppInvitationSettingQuery {
	aisq.order = append(aisq.order, o...)
	return aisq
}

// First returns the first AppInvitationSetting entity from the query.
// Returns a *NotFoundError when no AppInvitationSetting was found.
func (aisq *AppInvitationSettingQuery) First(ctx context.Context) (*AppInvitationSetting, error) {
	nodes, err := aisq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appinvitationsetting.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aisq *AppInvitationSettingQuery) FirstX(ctx context.Context) *AppInvitationSetting {
	node, err := aisq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppInvitationSetting ID from the query.
// Returns a *NotFoundError when no AppInvitationSetting ID was found.
func (aisq *AppInvitationSettingQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aisq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appinvitationsetting.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aisq *AppInvitationSettingQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := aisq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppInvitationSetting entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one AppInvitationSetting entity is not found.
// Returns a *NotFoundError when no AppInvitationSetting entities are found.
func (aisq *AppInvitationSettingQuery) Only(ctx context.Context) (*AppInvitationSetting, error) {
	nodes, err := aisq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appinvitationsetting.Label}
	default:
		return nil, &NotSingularError{appinvitationsetting.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aisq *AppInvitationSettingQuery) OnlyX(ctx context.Context) *AppInvitationSetting {
	node, err := aisq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppInvitationSetting ID in the query.
// Returns a *NotSingularError when exactly one AppInvitationSetting ID is not found.
// Returns a *NotFoundError when no entities are found.
func (aisq *AppInvitationSettingQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aisq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appinvitationsetting.Label}
	default:
		err = &NotSingularError{appinvitationsetting.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aisq *AppInvitationSettingQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := aisq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppInvitationSettings.
func (aisq *AppInvitationSettingQuery) All(ctx context.Context) ([]*AppInvitationSetting, error) {
	if err := aisq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aisq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aisq *AppInvitationSettingQuery) AllX(ctx context.Context) []*AppInvitationSetting {
	nodes, err := aisq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppInvitationSetting IDs.
func (aisq *AppInvitationSettingQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := aisq.Select(appinvitationsetting.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aisq *AppInvitationSettingQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := aisq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aisq *AppInvitationSettingQuery) Count(ctx context.Context) (int, error) {
	if err := aisq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aisq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aisq *AppInvitationSettingQuery) CountX(ctx context.Context) int {
	count, err := aisq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aisq *AppInvitationSettingQuery) Exist(ctx context.Context) (bool, error) {
	if err := aisq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aisq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aisq *AppInvitationSettingQuery) ExistX(ctx context.Context) bool {
	exist, err := aisq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppInvitationSettingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aisq *AppInvitationSettingQuery) Clone() *AppInvitationSettingQuery {
	if aisq == nil {
		return nil
	}
	return &AppInvitationSettingQuery{
		config:     aisq.config,
		limit:      aisq.limit,
		offset:     aisq.offset,
		order:      append([]OrderFunc{}, aisq.order...),
		predicates: append([]predicate.AppInvitationSetting{}, aisq.predicates...),
		// clone intermediate query.
		sql:  aisq.sql.Clone(),
		path: aisq.path,
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
//	client.AppInvitationSetting.Query().
//		GroupBy(appinvitationsetting.FieldAppID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (aisq *AppInvitationSettingQuery) GroupBy(field string, fields ...string) *AppInvitationSettingGroupBy {
	group := &AppInvitationSettingGroupBy{config: aisq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aisq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aisq.sqlQuery(ctx), nil
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
//	client.AppInvitationSetting.Query().
//		Select(appinvitationsetting.FieldAppID).
//		Scan(ctx, &v)
//
func (aisq *AppInvitationSettingQuery) Select(fields ...string) *AppInvitationSettingSelect {
	aisq.fields = append(aisq.fields, fields...)
	return &AppInvitationSettingSelect{AppInvitationSettingQuery: aisq}
}

func (aisq *AppInvitationSettingQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aisq.fields {
		if !appinvitationsetting.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aisq.path != nil {
		prev, err := aisq.path(ctx)
		if err != nil {
			return err
		}
		aisq.sql = prev
	}
	return nil
}

func (aisq *AppInvitationSettingQuery) sqlAll(ctx context.Context) ([]*AppInvitationSetting, error) {
	var (
		nodes = []*AppInvitationSetting{}
		_spec = aisq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AppInvitationSetting{config: aisq.config}
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
	if err := sqlgraph.QueryNodes(ctx, aisq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (aisq *AppInvitationSettingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aisq.querySpec()
	_spec.Node.Columns = aisq.fields
	if len(aisq.fields) > 0 {
		_spec.Unique = aisq.unique != nil && *aisq.unique
	}
	return sqlgraph.CountNodes(ctx, aisq.driver, _spec)
}

func (aisq *AppInvitationSettingQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aisq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (aisq *AppInvitationSettingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appinvitationsetting.Table,
			Columns: appinvitationsetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appinvitationsetting.FieldID,
			},
		},
		From:   aisq.sql,
		Unique: true,
	}
	if unique := aisq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aisq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appinvitationsetting.FieldID)
		for i := range fields {
			if fields[i] != appinvitationsetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aisq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aisq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aisq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aisq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aisq *AppInvitationSettingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aisq.driver.Dialect())
	t1 := builder.Table(appinvitationsetting.Table)
	columns := aisq.fields
	if len(columns) == 0 {
		columns = appinvitationsetting.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aisq.sql != nil {
		selector = aisq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aisq.unique != nil && *aisq.unique {
		selector.Distinct()
	}
	for _, p := range aisq.predicates {
		p(selector)
	}
	for _, p := range aisq.order {
		p(selector)
	}
	if offset := aisq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aisq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AppInvitationSettingGroupBy is the group-by builder for AppInvitationSetting entities.
type AppInvitationSettingGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aisgb *AppInvitationSettingGroupBy) Aggregate(fns ...AggregateFunc) *AppInvitationSettingGroupBy {
	aisgb.fns = append(aisgb.fns, fns...)
	return aisgb
}

// Scan applies the group-by query and scans the result into the given value.
func (aisgb *AppInvitationSettingGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := aisgb.path(ctx)
	if err != nil {
		return err
	}
	aisgb.sql = query
	return aisgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (aisgb *AppInvitationSettingGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := aisgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (aisgb *AppInvitationSettingGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(aisgb.fields) > 1 {
		return nil, errors.New("ent: AppInvitationSettingGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := aisgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (aisgb *AppInvitationSettingGroupBy) StringsX(ctx context.Context) []string {
	v, err := aisgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aisgb *AppInvitationSettingGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = aisgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appinvitationsetting.Label}
	default:
		err = fmt.Errorf("ent: AppInvitationSettingGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (aisgb *AppInvitationSettingGroupBy) StringX(ctx context.Context) string {
	v, err := aisgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (aisgb *AppInvitationSettingGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(aisgb.fields) > 1 {
		return nil, errors.New("ent: AppInvitationSettingGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := aisgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (aisgb *AppInvitationSettingGroupBy) IntsX(ctx context.Context) []int {
	v, err := aisgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aisgb *AppInvitationSettingGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = aisgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appinvitationsetting.Label}
	default:
		err = fmt.Errorf("ent: AppInvitationSettingGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (aisgb *AppInvitationSettingGroupBy) IntX(ctx context.Context) int {
	v, err := aisgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (aisgb *AppInvitationSettingGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(aisgb.fields) > 1 {
		return nil, errors.New("ent: AppInvitationSettingGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := aisgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (aisgb *AppInvitationSettingGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := aisgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aisgb *AppInvitationSettingGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = aisgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appinvitationsetting.Label}
	default:
		err = fmt.Errorf("ent: AppInvitationSettingGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (aisgb *AppInvitationSettingGroupBy) Float64X(ctx context.Context) float64 {
	v, err := aisgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (aisgb *AppInvitationSettingGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(aisgb.fields) > 1 {
		return nil, errors.New("ent: AppInvitationSettingGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := aisgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (aisgb *AppInvitationSettingGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := aisgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aisgb *AppInvitationSettingGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = aisgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appinvitationsetting.Label}
	default:
		err = fmt.Errorf("ent: AppInvitationSettingGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (aisgb *AppInvitationSettingGroupBy) BoolX(ctx context.Context) bool {
	v, err := aisgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (aisgb *AppInvitationSettingGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range aisgb.fields {
		if !appinvitationsetting.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := aisgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aisgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (aisgb *AppInvitationSettingGroupBy) sqlQuery() *sql.Selector {
	selector := aisgb.sql.Select()
	aggregation := make([]string, 0, len(aisgb.fns))
	for _, fn := range aisgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(aisgb.fields)+len(aisgb.fns))
		for _, f := range aisgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(aisgb.fields...)...)
}

// AppInvitationSettingSelect is the builder for selecting fields of AppInvitationSetting entities.
type AppInvitationSettingSelect struct {
	*AppInvitationSettingQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (aiss *AppInvitationSettingSelect) Scan(ctx context.Context, v interface{}) error {
	if err := aiss.prepareQuery(ctx); err != nil {
		return err
	}
	aiss.sql = aiss.AppInvitationSettingQuery.sqlQuery(ctx)
	return aiss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (aiss *AppInvitationSettingSelect) ScanX(ctx context.Context, v interface{}) {
	if err := aiss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (aiss *AppInvitationSettingSelect) Strings(ctx context.Context) ([]string, error) {
	if len(aiss.fields) > 1 {
		return nil, errors.New("ent: AppInvitationSettingSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := aiss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (aiss *AppInvitationSettingSelect) StringsX(ctx context.Context) []string {
	v, err := aiss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (aiss *AppInvitationSettingSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = aiss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appinvitationsetting.Label}
	default:
		err = fmt.Errorf("ent: AppInvitationSettingSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (aiss *AppInvitationSettingSelect) StringX(ctx context.Context) string {
	v, err := aiss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (aiss *AppInvitationSettingSelect) Ints(ctx context.Context) ([]int, error) {
	if len(aiss.fields) > 1 {
		return nil, errors.New("ent: AppInvitationSettingSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := aiss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (aiss *AppInvitationSettingSelect) IntsX(ctx context.Context) []int {
	v, err := aiss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (aiss *AppInvitationSettingSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = aiss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appinvitationsetting.Label}
	default:
		err = fmt.Errorf("ent: AppInvitationSettingSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (aiss *AppInvitationSettingSelect) IntX(ctx context.Context) int {
	v, err := aiss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (aiss *AppInvitationSettingSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(aiss.fields) > 1 {
		return nil, errors.New("ent: AppInvitationSettingSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := aiss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (aiss *AppInvitationSettingSelect) Float64sX(ctx context.Context) []float64 {
	v, err := aiss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (aiss *AppInvitationSettingSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = aiss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appinvitationsetting.Label}
	default:
		err = fmt.Errorf("ent: AppInvitationSettingSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (aiss *AppInvitationSettingSelect) Float64X(ctx context.Context) float64 {
	v, err := aiss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (aiss *AppInvitationSettingSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(aiss.fields) > 1 {
		return nil, errors.New("ent: AppInvitationSettingSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := aiss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (aiss *AppInvitationSettingSelect) BoolsX(ctx context.Context) []bool {
	v, err := aiss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (aiss *AppInvitationSettingSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = aiss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appinvitationsetting.Label}
	default:
		err = fmt.Errorf("ent: AppInvitationSettingSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (aiss *AppInvitationSettingSelect) BoolX(ctx context.Context) bool {
	v, err := aiss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (aiss *AppInvitationSettingSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := aiss.sql.Query()
	if err := aiss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
