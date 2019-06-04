package models

import (
  "github.com/jinzhu/gorm"
)

type Budget struct {
  gorm.Model
  Amount float64 `json:"amount"`
}

type Budgets []Budget