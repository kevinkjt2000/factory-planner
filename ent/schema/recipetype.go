package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RecipeType holds the schema definition for the RecipeType entity.
type RecipeType struct {
	ent.Schema
}

// Annotations of the RecipeType.
func (RecipeType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "recipe_type"},
	}
}

// Fields of the RecipeType.
func (RecipeType) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("category"),
		field.Bool("shapeless"),
		field.String("type"),
	}
}

// Edges of the RecipeType.
func (RecipeType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("recipe_type", Recipe.Type),
	}
}
