package models

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

//validation interface
type Validator interface {
	Validate() (bool, map[string]interface{})
}
