package model

import (
	"time"
)

type Problem struct {
	ID          int       `gorm:"column:id;primary_key" json:"id"`
	Name        string    `gorm:"column:name" json:"name" binding:"required,max=100"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"-"`
	OwnerID     int       `gorm:"column:owner_id" json:"owner_id"`
	Desc        string    `gorm:"column:desc" json:"desc"`
	Input       string    `gorm:"column:input" json:"input"`
	Output      string    `gorm:"column:output" json:"output"`
	Hint        string    `gorm:"column:hint" json:"hint"`
	Source      string    `gorm:"column:source" json:"source" binding:"max=100"`
	TimeLimit   int       `gorm:"column:time_limit" json:"time_limit" binding:"required,min=100,max=180000"`     // min = 100ms,max=3 minute
	MemoryLimit int       `gorm:"column:memory_limit" json:"memory_limit" binding:"required,min=100,max=204800"` // min = 100KB`,max = 200MB
	Difficulty  int       `gorm:"column:difficulty" json:"difficulty" binding:"required"`
	IsPublic    bool      `gorm:"column:is_public" json:"is_public"`
	CatalogID   int       `gorm:"column:catalog_id" json:"catalog_id"`
	Tags        []string  `gorm:"-" json:"tags"`
}

// TableName sets the insert table name for this struct type
func (p *Problem) TableName() string {
	return "problem"
}
