package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListRecipeTypes(t *testing.T) {
	db := NewDatabase()
	recipeTypes := db.ListRecipeTypes()
	require.NotEmpty(t, recipeTypes)
	require.Len(t, recipeTypes, 921)
}

func TestRecipeTypeSearch(t *testing.T) {
	db := NewDatabase()
	recipes := db.SearchRecipes(RecipeSearchCriteria{
		RecipeTypeIds: []string{"rt~minecraft~furnace"},
	})
	require.NotEmpty(t, recipes)
	for _, r := range recipes {
		require.Equal(t, r.RecipeTypeId, "rt~minecraft~furnace")
	}
}
