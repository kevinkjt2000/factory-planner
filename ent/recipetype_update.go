// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kevinkjt2000/factory-planner/ent/predicate"
	"github.com/kevinkjt2000/factory-planner/ent/recipe"
	"github.com/kevinkjt2000/factory-planner/ent/recipetype"
)

// RecipeTypeUpdate is the builder for updating RecipeType entities.
type RecipeTypeUpdate struct {
	config
	hooks    []Hook
	mutation *RecipeTypeMutation
}

// Where appends a list predicates to the RecipeTypeUpdate builder.
func (rtu *RecipeTypeUpdate) Where(ps ...predicate.RecipeType) *RecipeTypeUpdate {
	rtu.mutation.Where(ps...)
	return rtu
}

// SetCategory sets the "category" field.
func (rtu *RecipeTypeUpdate) SetCategory(s string) *RecipeTypeUpdate {
	rtu.mutation.SetCategory(s)
	return rtu
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (rtu *RecipeTypeUpdate) SetNillableCategory(s *string) *RecipeTypeUpdate {
	if s != nil {
		rtu.SetCategory(*s)
	}
	return rtu
}

// SetShapeless sets the "shapeless" field.
func (rtu *RecipeTypeUpdate) SetShapeless(b bool) *RecipeTypeUpdate {
	rtu.mutation.SetShapeless(b)
	return rtu
}

// SetNillableShapeless sets the "shapeless" field if the given value is not nil.
func (rtu *RecipeTypeUpdate) SetNillableShapeless(b *bool) *RecipeTypeUpdate {
	if b != nil {
		rtu.SetShapeless(*b)
	}
	return rtu
}

// SetType sets the "type" field.
func (rtu *RecipeTypeUpdate) SetType(s string) *RecipeTypeUpdate {
	rtu.mutation.SetType(s)
	return rtu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (rtu *RecipeTypeUpdate) SetNillableType(s *string) *RecipeTypeUpdate {
	if s != nil {
		rtu.SetType(*s)
	}
	return rtu
}

// AddRecipeTypeIDs adds the "recipe_type" edge to the Recipe entity by IDs.
func (rtu *RecipeTypeUpdate) AddRecipeTypeIDs(ids ...string) *RecipeTypeUpdate {
	rtu.mutation.AddRecipeTypeIDs(ids...)
	return rtu
}

// AddRecipeType adds the "recipe_type" edges to the Recipe entity.
func (rtu *RecipeTypeUpdate) AddRecipeType(r ...*Recipe) *RecipeTypeUpdate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rtu.AddRecipeTypeIDs(ids...)
}

// Mutation returns the RecipeTypeMutation object of the builder.
func (rtu *RecipeTypeUpdate) Mutation() *RecipeTypeMutation {
	return rtu.mutation
}

// ClearRecipeType clears all "recipe_type" edges to the Recipe entity.
func (rtu *RecipeTypeUpdate) ClearRecipeType() *RecipeTypeUpdate {
	rtu.mutation.ClearRecipeType()
	return rtu
}

// RemoveRecipeTypeIDs removes the "recipe_type" edge to Recipe entities by IDs.
func (rtu *RecipeTypeUpdate) RemoveRecipeTypeIDs(ids ...string) *RecipeTypeUpdate {
	rtu.mutation.RemoveRecipeTypeIDs(ids...)
	return rtu
}

// RemoveRecipeType removes "recipe_type" edges to Recipe entities.
func (rtu *RecipeTypeUpdate) RemoveRecipeType(r ...*Recipe) *RecipeTypeUpdate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rtu.RemoveRecipeTypeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rtu *RecipeTypeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, rtu.sqlSave, rtu.mutation, rtu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rtu *RecipeTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := rtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rtu *RecipeTypeUpdate) Exec(ctx context.Context) error {
	_, err := rtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtu *RecipeTypeUpdate) ExecX(ctx context.Context) {
	if err := rtu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rtu *RecipeTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(recipetype.Table, recipetype.Columns, sqlgraph.NewFieldSpec(recipetype.FieldID, field.TypeString))
	if ps := rtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rtu.mutation.Category(); ok {
		_spec.SetField(recipetype.FieldCategory, field.TypeString, value)
	}
	if value, ok := rtu.mutation.Shapeless(); ok {
		_spec.SetField(recipetype.FieldShapeless, field.TypeBool, value)
	}
	if value, ok := rtu.mutation.GetType(); ok {
		_spec.SetField(recipetype.FieldType, field.TypeString, value)
	}
	if rtu.mutation.RecipeTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipetype.RecipeTypeTable,
			Columns: []string{recipetype.RecipeTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(recipe.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtu.mutation.RemovedRecipeTypeIDs(); len(nodes) > 0 && !rtu.mutation.RecipeTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipetype.RecipeTypeTable,
			Columns: []string{recipetype.RecipeTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(recipe.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtu.mutation.RecipeTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipetype.RecipeTypeTable,
			Columns: []string{recipetype.RecipeTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(recipe.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recipetype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rtu.mutation.done = true
	return n, nil
}

// RecipeTypeUpdateOne is the builder for updating a single RecipeType entity.
type RecipeTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RecipeTypeMutation
}

// SetCategory sets the "category" field.
func (rtuo *RecipeTypeUpdateOne) SetCategory(s string) *RecipeTypeUpdateOne {
	rtuo.mutation.SetCategory(s)
	return rtuo
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (rtuo *RecipeTypeUpdateOne) SetNillableCategory(s *string) *RecipeTypeUpdateOne {
	if s != nil {
		rtuo.SetCategory(*s)
	}
	return rtuo
}

// SetShapeless sets the "shapeless" field.
func (rtuo *RecipeTypeUpdateOne) SetShapeless(b bool) *RecipeTypeUpdateOne {
	rtuo.mutation.SetShapeless(b)
	return rtuo
}

// SetNillableShapeless sets the "shapeless" field if the given value is not nil.
func (rtuo *RecipeTypeUpdateOne) SetNillableShapeless(b *bool) *RecipeTypeUpdateOne {
	if b != nil {
		rtuo.SetShapeless(*b)
	}
	return rtuo
}

// SetType sets the "type" field.
func (rtuo *RecipeTypeUpdateOne) SetType(s string) *RecipeTypeUpdateOne {
	rtuo.mutation.SetType(s)
	return rtuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (rtuo *RecipeTypeUpdateOne) SetNillableType(s *string) *RecipeTypeUpdateOne {
	if s != nil {
		rtuo.SetType(*s)
	}
	return rtuo
}

// AddRecipeTypeIDs adds the "recipe_type" edge to the Recipe entity by IDs.
func (rtuo *RecipeTypeUpdateOne) AddRecipeTypeIDs(ids ...string) *RecipeTypeUpdateOne {
	rtuo.mutation.AddRecipeTypeIDs(ids...)
	return rtuo
}

// AddRecipeType adds the "recipe_type" edges to the Recipe entity.
func (rtuo *RecipeTypeUpdateOne) AddRecipeType(r ...*Recipe) *RecipeTypeUpdateOne {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rtuo.AddRecipeTypeIDs(ids...)
}

// Mutation returns the RecipeTypeMutation object of the builder.
func (rtuo *RecipeTypeUpdateOne) Mutation() *RecipeTypeMutation {
	return rtuo.mutation
}

// ClearRecipeType clears all "recipe_type" edges to the Recipe entity.
func (rtuo *RecipeTypeUpdateOne) ClearRecipeType() *RecipeTypeUpdateOne {
	rtuo.mutation.ClearRecipeType()
	return rtuo
}

// RemoveRecipeTypeIDs removes the "recipe_type" edge to Recipe entities by IDs.
func (rtuo *RecipeTypeUpdateOne) RemoveRecipeTypeIDs(ids ...string) *RecipeTypeUpdateOne {
	rtuo.mutation.RemoveRecipeTypeIDs(ids...)
	return rtuo
}

// RemoveRecipeType removes "recipe_type" edges to Recipe entities.
func (rtuo *RecipeTypeUpdateOne) RemoveRecipeType(r ...*Recipe) *RecipeTypeUpdateOne {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rtuo.RemoveRecipeTypeIDs(ids...)
}

// Where appends a list predicates to the RecipeTypeUpdate builder.
func (rtuo *RecipeTypeUpdateOne) Where(ps ...predicate.RecipeType) *RecipeTypeUpdateOne {
	rtuo.mutation.Where(ps...)
	return rtuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rtuo *RecipeTypeUpdateOne) Select(field string, fields ...string) *RecipeTypeUpdateOne {
	rtuo.fields = append([]string{field}, fields...)
	return rtuo
}

// Save executes the query and returns the updated RecipeType entity.
func (rtuo *RecipeTypeUpdateOne) Save(ctx context.Context) (*RecipeType, error) {
	return withHooks(ctx, rtuo.sqlSave, rtuo.mutation, rtuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rtuo *RecipeTypeUpdateOne) SaveX(ctx context.Context) *RecipeType {
	node, err := rtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rtuo *RecipeTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := rtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtuo *RecipeTypeUpdateOne) ExecX(ctx context.Context) {
	if err := rtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rtuo *RecipeTypeUpdateOne) sqlSave(ctx context.Context) (_node *RecipeType, err error) {
	_spec := sqlgraph.NewUpdateSpec(recipetype.Table, recipetype.Columns, sqlgraph.NewFieldSpec(recipetype.FieldID, field.TypeString))
	id, ok := rtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RecipeType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, recipetype.FieldID)
		for _, f := range fields {
			if !recipetype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != recipetype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rtuo.mutation.Category(); ok {
		_spec.SetField(recipetype.FieldCategory, field.TypeString, value)
	}
	if value, ok := rtuo.mutation.Shapeless(); ok {
		_spec.SetField(recipetype.FieldShapeless, field.TypeBool, value)
	}
	if value, ok := rtuo.mutation.GetType(); ok {
		_spec.SetField(recipetype.FieldType, field.TypeString, value)
	}
	if rtuo.mutation.RecipeTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipetype.RecipeTypeTable,
			Columns: []string{recipetype.RecipeTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(recipe.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtuo.mutation.RemovedRecipeTypeIDs(); len(nodes) > 0 && !rtuo.mutation.RecipeTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipetype.RecipeTypeTable,
			Columns: []string{recipetype.RecipeTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(recipe.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtuo.mutation.RecipeTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipetype.RecipeTypeTable,
			Columns: []string{recipetype.RecipeTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(recipe.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &RecipeType{config: rtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recipetype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rtuo.mutation.done = true
	return _node, nil
}
