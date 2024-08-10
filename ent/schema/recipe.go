package schema

import (
	"entgo.io/ent"

	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Recipe holds the schema definition for the Recipe entity.
type Recipe struct {
	ent.Schema
}

// Annotations of the Recipe.
func (Recipe) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "recipe"},
	}
}

// Fields of the Recipe.
func (Recipe) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("recipe_type_id").Optional(),
	}
}

// Edges of the Recipe.
func (Recipe) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("recipe_type", RecipeType.Type).
			Ref("recipe_type").
			Field("recipe_type_id").
			Unique(),
	}
}
