package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type UserVoteProblem struct {
	ID        int       `gorm:"column:id;primary_key" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserID    int       `gorm:"column:user_id" json:"user_id"`
	ProblemID int       `gorm:"column:problem_id" json:"problem_id"`
	IsOk      int       `gorm:"column:is_ok" json:"is_ok"`
}

// TableName sets the insert table name for this struct type
func (u *UserVoteProblem) TableName() string {
	return "user_vote_problem"
}
