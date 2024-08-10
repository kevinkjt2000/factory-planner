package internal

import (
	"context"
	"fmt"

	"github.com/kevinkjt2000/factory-planner/components"
	"github.com/kevinkjt2000/factory-planner/ent"
	"github.com/kevinkjt2000/factory-planner/ent/recipe"
	"github.com/kevinkjt2000/factory-planner/ent/recipetype"

	_ "github.com/lib/pq"
)

// useful psql tips
// \dt              lists table names
// \d+ <tablename>  shows column type information about tables

type Database struct {
	client *ent.Client
}

func (db *Database) Close() {
	db.client.Close()
}

func NewDatabase() Database {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres password=example dbname=gtnh_2_6_1 sslmode=disable")
	if err != nil {
		fmt.Printf("gorm.Open: %v\n", err)
		panic(1)
	}

	return Database{
		client: client,
	}
}

type RecipeSearchCriteria struct {
	RecipeTypeIds []string
	InputFilter   string
}

func (db *Database) SearchRecipes(rsc RecipeSearchCriteria) []components.Recipe {
	query := db.client.Recipe.Query().Limit(20)
	if len(rsc.RecipeTypeIds) > 0 {
		query = query.Where(recipe.HasRecipeTypeWith(recipetype.IDIn(rsc.RecipeTypeIds...)))
	}
	if rsc.InputFilter != "" {
		query = query.Where(recipe.IDEQ(rsc.InputFilter))
	}
	entRecipes, err := query.All(context.Background())
	if err != nil {
		fmt.Printf("recipe query: %v\n", err)
		return []components.Recipe{}
	}
	recipes := make([]components.Recipe, len(entRecipes))
	for i, rt := range entRecipes {
		recipes[i] = components.Recipe{
			Id:           rt.ID,
			RecipeTypeId: rt.RecipeTypeID,
		}
	}
	return recipes
}

func (db *Database) ListRecipeTypes() []string {
	entRecipeTypes, err := db.client.RecipeType.Query().All(context.Background())
	if err != nil {
		fmt.Printf("list recipe types: %v\n", err)
		return []string{}
	}
	recipeTypes := make([]string, len(entRecipeTypes))
	for i, rt := range entRecipeTypes {
		recipeTypes[i] = rt.ID
	}
	return recipeTypes
}
