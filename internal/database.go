package internal

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// useful psql tips
// \dt              lists table names
// \d+ <tablename>  shows column type information about tables

// TODO automate migrations db.AutoMigrate(&RecipeType{})

type RecipeType struct {
	Id       string `gorm:"primaryKey"`
	Category string
	Type     string `gorm:"column:TYPE"`
}

type RecipeItemOutputs struct {
	ItemOutputsKey              int `gorm:"primaryKey"`
	ItemOutputsValueItemId      string
	ItemOutputsValueProbability float64
	ItemOutputsValueStackSize   int
	RecipeId                    string `gorm:"primaryKey"`
}

type Database struct {
	db *gorm.DB
}

func (db *Database) Close() {
	wrappedDb, _ := db.db.DB()
	wrappedDb.Close()
}

func NewDatabase() Database {
	db, err := gorm.Open(postgres.Open("postgres://postgres:example@localhost:5432/gtnh_recipes"), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		fmt.Printf("gorm.Open: %v\n", err)
		panic(1)
	}

	return Database{
		db: db,
	}
}

func (db *Database) SearchRecipeItemOutputs(query string) []string {
	var matches []string
	if err := db.db.Table("recipe_item_outputs").Select("recipe_id").Where("item_outputs_value_item_id LIKE ?", "%"+query+"%").Limit(10).Find(&matches).Error; err != nil {
		fmt.Printf("SearchRecipeItemOutputs: %v\n", err)
	}
	return matches
}

func (db *Database) GetRecipeTypes() []string {
	var types []string
	if err := db.db.Model(&RecipeType{}).Pluck(`"TYPE"`, &types).Error; err != nil {
		fmt.Printf("GetRecipeTypes: %v\n", err)
		return []string{"error"}
	}

	return types
}
