// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/bugs"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/predicate"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/repository"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/teams"
)

// TeamsQuery is the builder for querying Teams entities.
type TeamsQuery struct {
	config
	ctx              *QueryContext
	order            []OrderFunc
	inters           []Interceptor
	predicates       []predicate.Teams
	withRepositories *RepositoryQuery
	withBugs         *BugsQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TeamsQuery builder.
func (tq *TeamsQuery) Where(ps ...predicate.Teams) *TeamsQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TeamsQuery) Limit(limit int) *TeamsQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TeamsQuery) Offset(offset int) *TeamsQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TeamsQuery) Unique(unique bool) *TeamsQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TeamsQuery) Order(o ...OrderFunc) *TeamsQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryRepositories chains the current query on the "repositories" edge.
func (tq *TeamsQuery) QueryRepositories() *RepositoryQuery {
	query := (&RepositoryClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(teams.Table, teams.FieldID, selector),
			sqlgraph.To(repository.Table, repository.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, teams.RepositoriesTable, teams.RepositoriesColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBugs chains the current query on the "bugs" edge.
func (tq *TeamsQuery) QueryBugs() *BugsQuery {
	query := (&BugsClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(teams.Table, teams.FieldID, selector),
			sqlgraph.To(bugs.Table, bugs.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, teams.BugsTable, teams.BugsColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Teams entity from the query.
// Returns a *NotFoundError when no Teams was found.
func (tq *TeamsQuery) First(ctx context.Context) (*Teams, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{teams.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TeamsQuery) FirstX(ctx context.Context) *Teams {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Teams ID from the query.
// Returns a *NotFoundError when no Teams ID was found.
func (tq *TeamsQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{teams.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TeamsQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Teams entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Teams entity is found.
// Returns a *NotFoundError when no Teams entities are found.
func (tq *TeamsQuery) Only(ctx context.Context) (*Teams, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{teams.Label}
	default:
		return nil, &NotSingularError{teams.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TeamsQuery) OnlyX(ctx context.Context) *Teams {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Teams ID in the query.
// Returns a *NotSingularError when more than one Teams ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TeamsQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{teams.Label}
	default:
		err = &NotSingularError{teams.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TeamsQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TeamsSlice.
func (tq *TeamsQuery) All(ctx context.Context) ([]*Teams, error) {
	ctx = setContextOp(ctx, tq.ctx, "All")
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Teams, *TeamsQuery]()
	return withInterceptors[[]*Teams](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TeamsQuery) AllX(ctx context.Context) []*Teams {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Teams IDs.
func (tq *TeamsQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = setContextOp(ctx, tq.ctx, "IDs")
	if err := tq.Select(teams.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TeamsQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TeamsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, "Count")
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TeamsQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TeamsQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TeamsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, "Exist")
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TeamsQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TeamsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TeamsQuery) Clone() *TeamsQuery {
	if tq == nil {
		return nil
	}
	return &TeamsQuery{
		config:           tq.config,
		ctx:              tq.ctx.Clone(),
		order:            append([]OrderFunc{}, tq.order...),
		inters:           append([]Interceptor{}, tq.inters...),
		predicates:       append([]predicate.Teams{}, tq.predicates...),
		withRepositories: tq.withRepositories.Clone(),
		withBugs:         tq.withBugs.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithRepositories tells the query-builder to eager-load the nodes that are connected to
// the "repositories" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TeamsQuery) WithRepositories(opts ...func(*RepositoryQuery)) *TeamsQuery {
	query := (&RepositoryClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withRepositories = query
	return tq
}

// WithBugs tells the query-builder to eager-load the nodes that are connected to
// the "bugs" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TeamsQuery) WithBugs(opts ...func(*BugsQuery)) *TeamsQuery {
	query := (&BugsClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withBugs = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TeamName string `json:"team_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Teams.Query().
//		GroupBy(teams.FieldTeamName).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (tq *TeamsQuery) GroupBy(field string, fields ...string) *TeamsGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TeamsGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = teams.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TeamName string `json:"team_name,omitempty"`
//	}
//
//	client.Teams.Query().
//		Select(teams.FieldTeamName).
//		Scan(ctx, &v)
func (tq *TeamsQuery) Select(fields ...string) *TeamsSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TeamsSelect{TeamsQuery: tq}
	sbuild.label = teams.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TeamsSelect configured with the given aggregations.
func (tq *TeamsQuery) Aggregate(fns ...AggregateFunc) *TeamsSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TeamsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !teams.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TeamsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Teams, error) {
	var (
		nodes       = []*Teams{}
		_spec       = tq.querySpec()
		loadedTypes = [2]bool{
			tq.withRepositories != nil,
			tq.withBugs != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Teams).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Teams{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withRepositories; query != nil {
		if err := tq.loadRepositories(ctx, query, nodes,
			func(n *Teams) { n.Edges.Repositories = []*Repository{} },
			func(n *Teams, e *Repository) { n.Edges.Repositories = append(n.Edges.Repositories, e) }); err != nil {
			return nil, err
		}
	}
	if query := tq.withBugs; query != nil {
		if err := tq.loadBugs(ctx, query, nodes,
			func(n *Teams) { n.Edges.Bugs = []*Bugs{} },
			func(n *Teams, e *Bugs) { n.Edges.Bugs = append(n.Edges.Bugs, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TeamsQuery) loadRepositories(ctx context.Context, query *RepositoryQuery, nodes []*Teams, init func(*Teams), assign func(*Teams, *Repository)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Teams)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Repository(func(s *sql.Selector) {
		s.Where(sql.InValues(teams.RepositoriesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.teams_repositories
		if fk == nil {
			return fmt.Errorf(`foreign-key "teams_repositories" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "teams_repositories" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (tq *TeamsQuery) loadBugs(ctx context.Context, query *BugsQuery, nodes []*Teams, init func(*Teams), assign func(*Teams, *Bugs)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Teams)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Bugs(func(s *sql.Selector) {
		s.Where(sql.InValues(teams.BugsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.teams_bugs
		if fk == nil {
			return fmt.Errorf(`foreign-key "teams_bugs" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "teams_bugs" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (tq *TeamsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TeamsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   teams.Table,
			Columns: teams.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: teams.FieldID,
			},
		},
		From:   tq.sql,
		Unique: true,
	}
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, teams.FieldID)
		for i := range fields {
			if fields[i] != teams.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TeamsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(teams.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = teams.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TeamsGroupBy is the group-by builder for Teams entities.
type TeamsGroupBy struct {
	selector
	build *TeamsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TeamsGroupBy) Aggregate(fns ...AggregateFunc) *TeamsGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TeamsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, "GroupBy")
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TeamsQuery, *TeamsGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TeamsGroupBy) sqlScan(ctx context.Context, root *TeamsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TeamsSelect is the builder for selecting fields of Teams entities.
type TeamsSelect struct {
	*TeamsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TeamsSelect) Aggregate(fns ...AggregateFunc) *TeamsSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TeamsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, "Select")
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TeamsQuery, *TeamsSelect](ctx, ts.TeamsQuery, ts, ts.inters, v)
}

func (ts *TeamsSelect) sqlScan(ctx context.Context, root *TeamsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
