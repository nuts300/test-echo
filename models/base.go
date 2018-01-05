package models

import (
	"time"
)

type Base struct {
	ID        int        `json:"id" yaml:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createAt" yaml:"createAt"`
	UpdatedAt time.Time  `json:"updateAt" yaml:"updateAt"`
	DeletedAt *time.Time `json:"deleteAt" yaml:"deleteAt" sql:"index"`
}
