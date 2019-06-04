package main

import (
	"github.com/bargenson/go-budget-api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type database struct {
	db *gorm.DB
}

func createDatabase() *database {
	db, err := gorm.Open("sqlite3", "test.db")
  if err != nil {
    panic(err)
  }
  defer db.Close()
  db.AutoMigrate(&models.Budget{})
  return &database{db}
}