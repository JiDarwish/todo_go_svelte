package models

import (
 "gorm.io/gorm"
 "gorm.io/datatypes"
)

type Todo struct {
  gorm.Model                `json:"-"`
  Title                     string
  Description               string
  Deadline                  datatypes.Date
}
