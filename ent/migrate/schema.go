// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// RecipeColumns holds the columns for the "recipe" table.
	RecipeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "recipe_type_id", Type: field.TypeString, Nullable: true},
	}
	// RecipeTable holds the schema information for the "recipe" table.
	RecipeTable = &schema.Table{
		Name:       "recipe",
		Columns:    RecipeColumns,
		PrimaryKey: []*schema.Column{RecipeColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "recipe_recipe_type_recipe_type",
				Columns:    []*schema.Column{RecipeColumns[1]},
				RefColumns: []*schema.Column{RecipeTypeColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RecipeItemOutputsColumns holds the columns for the "recipe_item_outputs" table.
	RecipeItemOutputsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// RecipeItemOutputsTable holds the schema information for the "recipe_item_outputs" table.
	RecipeItemOutputsTable = &schema.Table{
		Name:       "recipe_item_outputs",
		Columns:    RecipeItemOutputsColumns,
		PrimaryKey: []*schema.Column{RecipeItemOutputsColumns[0]},
	}
	// RecipeTypeColumns holds the columns for the "recipe_type" table.
	RecipeTypeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "category", Type: field.TypeString},
		{Name: "shapeless", Type: field.TypeBool},
		{Name: "type", Type: field.TypeString},
	}
	// RecipeTypeTable holds the schema information for the "recipe_type" table.
	RecipeTypeTable = &schema.Table{
		Name:       "recipe_type",
		Columns:    RecipeTypeColumns,
		PrimaryKey: []*schema.Column{RecipeTypeColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		RecipeTable,
		RecipeItemOutputsTable,
		RecipeTypeTable,
	}
)

func init() {
	RecipeTable.ForeignKeys[0].RefTable = RecipeTypeTable
	RecipeTable.Annotation = &entsql.Annotation{
		Table: "recipe",
	}
	RecipeTypeTable.Annotation = &entsql.Annotation{
		Table: "recipe_type",
	}
}
