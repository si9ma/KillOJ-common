package model

import (
	"time"
)

type Catalog struct {
	ID        int       `gorm:"column:id;primary_key" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
}

// TableName sets the insert table name for this struct type
func (c *Catalog) TableName() string {
	return "catalog"
}
