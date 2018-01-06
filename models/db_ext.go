package models

import (
	"time"
)

type DbExt struct {
	CreatedAt time.Time  `json:"createAt" yaml:"createAt"`
	UpdatedAt time.Time  `json:"updateAt" yaml:"updateAt"`
	DeletedAt *time.Time `json:"deleteAt" yaml:"deleteAt" sql:"index"`
}
