// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kevinkjt2000/factory-planner/ent/recipeitemoutputs"
)

// RecipeItemOutputsCreate is the builder for creating a RecipeItemOutputs entity.
type RecipeItemOutputsCreate struct {
	config
	mutation *RecipeItemOutputsMutation
	hooks    []Hook
}

// Mutation returns the RecipeItemOutputsMutation object of the builder.
func (rioc *RecipeItemOutputsCreate) Mutation() *RecipeItemOutputsMutation {
	return rioc.mutation
}

// Save creates the RecipeItemOutputs in the database.
func (rioc *RecipeItemOutputsCreate) Save(ctx context.Context) (*RecipeItemOutputs, error) {
	return withHooks(ctx, rioc.sqlSave, rioc.mutation, rioc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rioc *RecipeItemOutputsCreate) SaveX(ctx context.Context) *RecipeItemOutputs {
	v, err := rioc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rioc *RecipeItemOutputsCreate) Exec(ctx context.Context) error {
	_, err := rioc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rioc *RecipeItemOutputsCreate) ExecX(ctx context.Context) {
	if err := rioc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rioc *RecipeItemOutputsCreate) check() error {
	return nil
}

func (rioc *RecipeItemOutputsCreate) sqlSave(ctx context.Context) (*RecipeItemOutputs, error) {
	if err := rioc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rioc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rioc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rioc.mutation.id = &_node.ID
	rioc.mutation.done = true
	return _node, nil
}

func (rioc *RecipeItemOutputsCreate) createSpec() (*RecipeItemOutputs, *sqlgraph.CreateSpec) {
	var (
		_node = &RecipeItemOutputs{config: rioc.config}
		_spec = sqlgraph.NewCreateSpec(recipeitemoutputs.Table, sqlgraph.NewFieldSpec(recipeitemoutputs.FieldID, field.TypeInt))
	)
	return _node, _spec
}

// RecipeItemOutputsCreateBulk is the builder for creating many RecipeItemOutputs entities in bulk.
type RecipeItemOutputsCreateBulk struct {
	config
	err      error
	builders []*RecipeItemOutputsCreate
}

// Save creates the RecipeItemOutputs entities in the database.
func (riocb *RecipeItemOutputsCreateBulk) Save(ctx context.Context) ([]*RecipeItemOutputs, error) {
	if riocb.err != nil {
		return nil, riocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(riocb.builders))
	nodes := make([]*RecipeItemOutputs, len(riocb.builders))
	mutators := make([]Mutator, len(riocb.builders))
	for i := range riocb.builders {
		func(i int, root context.Context) {
			builder := riocb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RecipeItemOutputsMutation)
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
					_, err = mutators[i+1].Mutate(root, riocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, riocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, riocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (riocb *RecipeItemOutputsCreateBulk) SaveX(ctx context.Context) []*RecipeItemOutputs {
	v, err := riocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (riocb *RecipeItemOutputsCreateBulk) Exec(ctx context.Context) error {
	_, err := riocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (riocb *RecipeItemOutputsCreateBulk) ExecX(ctx context.Context) {
	if err := riocb.Exec(ctx); err != nil {
		panic(err)
	}
}
