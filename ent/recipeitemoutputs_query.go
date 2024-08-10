// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kevinkjt2000/factory-planner/ent/predicate"
	"github.com/kevinkjt2000/factory-planner/ent/recipeitemoutputs"
)

// RecipeItemOutputsQuery is the builder for querying RecipeItemOutputs entities.
type RecipeItemOutputsQuery struct {
	config
	ctx        *QueryContext
	order      []recipeitemoutputs.OrderOption
	inters     []Interceptor
	predicates []predicate.RecipeItemOutputs
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RecipeItemOutputsQuery builder.
func (rioq *RecipeItemOutputsQuery) Where(ps ...predicate.RecipeItemOutputs) *RecipeItemOutputsQuery {
	rioq.predicates = append(rioq.predicates, ps...)
	return rioq
}

// Limit the number of records to be returned by this query.
func (rioq *RecipeItemOutputsQuery) Limit(limit int) *RecipeItemOutputsQuery {
	rioq.ctx.Limit = &limit
	return rioq
}

// Offset to start from.
func (rioq *RecipeItemOutputsQuery) Offset(offset int) *RecipeItemOutputsQuery {
	rioq.ctx.Offset = &offset
	return rioq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rioq *RecipeItemOutputsQuery) Unique(unique bool) *RecipeItemOutputsQuery {
	rioq.ctx.Unique = &unique
	return rioq
}

// Order specifies how the records should be ordered.
func (rioq *RecipeItemOutputsQuery) Order(o ...recipeitemoutputs.OrderOption) *RecipeItemOutputsQuery {
	rioq.order = append(rioq.order, o...)
	return rioq
}

// First returns the first RecipeItemOutputs entity from the query.
// Returns a *NotFoundError when no RecipeItemOutputs was found.
func (rioq *RecipeItemOutputsQuery) First(ctx context.Context) (*RecipeItemOutputs, error) {
	nodes, err := rioq.Limit(1).All(setContextOp(ctx, rioq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{recipeitemoutputs.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rioq *RecipeItemOutputsQuery) FirstX(ctx context.Context) *RecipeItemOutputs {
	node, err := rioq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RecipeItemOutputs ID from the query.
// Returns a *NotFoundError when no RecipeItemOutputs ID was found.
func (rioq *RecipeItemOutputsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rioq.Limit(1).IDs(setContextOp(ctx, rioq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{recipeitemoutputs.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rioq *RecipeItemOutputsQuery) FirstIDX(ctx context.Context) int {
	id, err := rioq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RecipeItemOutputs entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RecipeItemOutputs entity is found.
// Returns a *NotFoundError when no RecipeItemOutputs entities are found.
func (rioq *RecipeItemOutputsQuery) Only(ctx context.Context) (*RecipeItemOutputs, error) {
	nodes, err := rioq.Limit(2).All(setContextOp(ctx, rioq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{recipeitemoutputs.Label}
	default:
		return nil, &NotSingularError{recipeitemoutputs.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rioq *RecipeItemOutputsQuery) OnlyX(ctx context.Context) *RecipeItemOutputs {
	node, err := rioq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RecipeItemOutputs ID in the query.
// Returns a *NotSingularError when more than one RecipeItemOutputs ID is found.
// Returns a *NotFoundError when no entities are found.
func (rioq *RecipeItemOutputsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rioq.Limit(2).IDs(setContextOp(ctx, rioq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{recipeitemoutputs.Label}
	default:
		err = &NotSingularError{recipeitemoutputs.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rioq *RecipeItemOutputsQuery) OnlyIDX(ctx context.Context) int {
	id, err := rioq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RecipeItemOutputsSlice.
func (rioq *RecipeItemOutputsQuery) All(ctx context.Context) ([]*RecipeItemOutputs, error) {
	ctx = setContextOp(ctx, rioq.ctx, ent.OpQueryAll)
	if err := rioq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RecipeItemOutputs, *RecipeItemOutputsQuery]()
	return withInterceptors[[]*RecipeItemOutputs](ctx, rioq, qr, rioq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rioq *RecipeItemOutputsQuery) AllX(ctx context.Context) []*RecipeItemOutputs {
	nodes, err := rioq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RecipeItemOutputs IDs.
func (rioq *RecipeItemOutputsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if rioq.ctx.Unique == nil && rioq.path != nil {
		rioq.Unique(true)
	}
	ctx = setContextOp(ctx, rioq.ctx, ent.OpQueryIDs)
	if err = rioq.Select(recipeitemoutputs.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rioq *RecipeItemOutputsQuery) IDsX(ctx context.Context) []int {
	ids, err := rioq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rioq *RecipeItemOutputsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rioq.ctx, ent.OpQueryCount)
	if err := rioq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rioq, querierCount[*RecipeItemOutputsQuery](), rioq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rioq *RecipeItemOutputsQuery) CountX(ctx context.Context) int {
	count, err := rioq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rioq *RecipeItemOutputsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rioq.ctx, ent.OpQueryExist)
	switch _, err := rioq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rioq *RecipeItemOutputsQuery) ExistX(ctx context.Context) bool {
	exist, err := rioq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RecipeItemOutputsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rioq *RecipeItemOutputsQuery) Clone() *RecipeItemOutputsQuery {
	if rioq == nil {
		return nil
	}
	return &RecipeItemOutputsQuery{
		config:     rioq.config,
		ctx:        rioq.ctx.Clone(),
		order:      append([]recipeitemoutputs.OrderOption{}, rioq.order...),
		inters:     append([]Interceptor{}, rioq.inters...),
		predicates: append([]predicate.RecipeItemOutputs{}, rioq.predicates...),
		// clone intermediate query.
		sql:  rioq.sql.Clone(),
		path: rioq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (rioq *RecipeItemOutputsQuery) GroupBy(field string, fields ...string) *RecipeItemOutputsGroupBy {
	rioq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RecipeItemOutputsGroupBy{build: rioq}
	grbuild.flds = &rioq.ctx.Fields
	grbuild.label = recipeitemoutputs.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (rioq *RecipeItemOutputsQuery) Select(fields ...string) *RecipeItemOutputsSelect {
	rioq.ctx.Fields = append(rioq.ctx.Fields, fields...)
	sbuild := &RecipeItemOutputsSelect{RecipeItemOutputsQuery: rioq}
	sbuild.label = recipeitemoutputs.Label
	sbuild.flds, sbuild.scan = &rioq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RecipeItemOutputsSelect configured with the given aggregations.
func (rioq *RecipeItemOutputsQuery) Aggregate(fns ...AggregateFunc) *RecipeItemOutputsSelect {
	return rioq.Select().Aggregate(fns...)
}

func (rioq *RecipeItemOutputsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rioq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rioq); err != nil {
				return err
			}
		}
	}
	for _, f := range rioq.ctx.Fields {
		if !recipeitemoutputs.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rioq.path != nil {
		prev, err := rioq.path(ctx)
		if err != nil {
			return err
		}
		rioq.sql = prev
	}
	return nil
}

func (rioq *RecipeItemOutputsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RecipeItemOutputs, error) {
	var (
		nodes = []*RecipeItemOutputs{}
		_spec = rioq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RecipeItemOutputs).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RecipeItemOutputs{config: rioq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rioq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (rioq *RecipeItemOutputsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rioq.querySpec()
	_spec.Node.Columns = rioq.ctx.Fields
	if len(rioq.ctx.Fields) > 0 {
		_spec.Unique = rioq.ctx.Unique != nil && *rioq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rioq.driver, _spec)
}

func (rioq *RecipeItemOutputsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(recipeitemoutputs.Table, recipeitemoutputs.Columns, sqlgraph.NewFieldSpec(recipeitemoutputs.FieldID, field.TypeInt))
	_spec.From = rioq.sql
	if unique := rioq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rioq.path != nil {
		_spec.Unique = true
	}
	if fields := rioq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, recipeitemoutputs.FieldID)
		for i := range fields {
			if fields[i] != recipeitemoutputs.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rioq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rioq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rioq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rioq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rioq *RecipeItemOutputsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rioq.driver.Dialect())
	t1 := builder.Table(recipeitemoutputs.Table)
	columns := rioq.ctx.Fields
	if len(columns) == 0 {
		columns = recipeitemoutputs.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rioq.sql != nil {
		selector = rioq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rioq.ctx.Unique != nil && *rioq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rioq.predicates {
		p(selector)
	}
	for _, p := range rioq.order {
		p(selector)
	}
	if offset := rioq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rioq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RecipeItemOutputsGroupBy is the group-by builder for RecipeItemOutputs entities.
type RecipeItemOutputsGroupBy struct {
	selector
	build *RecipeItemOutputsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (riogb *RecipeItemOutputsGroupBy) Aggregate(fns ...AggregateFunc) *RecipeItemOutputsGroupBy {
	riogb.fns = append(riogb.fns, fns...)
	return riogb
}

// Scan applies the selector query and scans the result into the given value.
func (riogb *RecipeItemOutputsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, riogb.build.ctx, ent.OpQueryGroupBy)
	if err := riogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RecipeItemOutputsQuery, *RecipeItemOutputsGroupBy](ctx, riogb.build, riogb, riogb.build.inters, v)
}

func (riogb *RecipeItemOutputsGroupBy) sqlScan(ctx context.Context, root *RecipeItemOutputsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(riogb.fns))
	for _, fn := range riogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*riogb.flds)+len(riogb.fns))
		for _, f := range *riogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*riogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := riogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RecipeItemOutputsSelect is the builder for selecting fields of RecipeItemOutputs entities.
type RecipeItemOutputsSelect struct {
	*RecipeItemOutputsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rios *RecipeItemOutputsSelect) Aggregate(fns ...AggregateFunc) *RecipeItemOutputsSelect {
	rios.fns = append(rios.fns, fns...)
	return rios
}

// Scan applies the selector query and scans the result into the given value.
func (rios *RecipeItemOutputsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rios.ctx, ent.OpQuerySelect)
	if err := rios.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RecipeItemOutputsQuery, *RecipeItemOutputsSelect](ctx, rios.RecipeItemOutputsQuery, rios, rios.inters, v)
}

func (rios *RecipeItemOutputsSelect) sqlScan(ctx context.Context, root *RecipeItemOutputsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rios.fns))
	for _, fn := range rios.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rios.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rios.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
