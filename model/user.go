package model

import "time"

type User struct {
	ID           int       `gorm:"column:id;primary_key" json:"id"`
	GithubUserID string    `gorm:"column:github_user_id" json:"github_user_id"`
	GithubName   string    `gorm:"column:github_name" json:"github_name"`
	Email        string    `gorm:"column:email" json:"email" binding:"required,email,max=50"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"-"`
	Role         int       `gorm:"column:role" json:"-"`

	// password from user input, password should't response to user,
	// so set omitempty && set this field to nil before return
	Password         string `gorm:"-" json:"password,omitempty" binding:"max=30"`
	EncryptedPasswd  string `gorm:"column:passwd" json:"-"` // encrypted password in db
	NickName         string `gorm:"column:nick_name" json:"nick_name" binding:"required,max=50"`
	Signature        string `gorm:"column:signature" json:"signature" binding:"max=100"`
	NoInOrganization string `gorm:"column:no_in_organization" json:"no_in_organization" binding:"max=30"`
	Organization     string `gorm:"column:organization" json:"organization" binding:"max=50"`
	Site             string `gorm:"column:site" json:"site" binding:"max=50"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "user"
}
