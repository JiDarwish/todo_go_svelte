package models

import "gorm.io/gorm"

type TodoList struct {
  gorm.Model          `json:"-"`
  Title               string
  Description         string
  TodoItems           []Todo `gorm:"foreignKey:ID"`
}
