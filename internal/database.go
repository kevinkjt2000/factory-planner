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
	Id       string `gorm:"column:id;primaryKey"`
	Category string `gorm:"column:category"`
	Type     string `gorm:"column:TYPE"`
}

type Database struct {
	db *gorm.DB
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

func (db *Database) GetRecipeTypes() []string {
	var types []string
	if err := db.db.Table("recipe_type").Select(`"TYPE"`).Find(&types).Error; err != nil {
		fmt.Printf("GetRecipeTypes: %v\n", err)
		return []string{"error"}
	}

	return types
}
