package model

import (
	"time"
)

type ProblemInGroup struct {
	ID        int       `gorm:"column:id;primary_key" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	ProblemID int       `gorm:"column:problem_id" json:"problem_id"`
	GroupID   int       `gorm:"column:group_id" json:"group_id"`
}

// TableName sets the insert table name for this struct type
func (p *ProblemInGroup) TableName() string {
	return "problem_in_group"
}
