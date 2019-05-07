package model

import (
	"testing"

	"github.com/si9ma/KillOJ-common/kjson"

	"github.com/jinzhu/gorm"
	"github.com/si9ma/KillOJ-common/mysql"
)

func TestBelongsTo(t *testing.T) {
	var (
		db  *gorm.DB
		err error
	)

	if db, err = mysql.GetTestDB(); err != nil {
		t.Fatal("init mysql fail", err)
	}

	var (
		problem Problem
		submit  Submit
	)

	db.First(&problem)
	db.Model(&problem).Related(&submit)
	t.Log(kjson.MarshalString(problem))
	t.Log(kjson.MarshalString(submit))

	var (
		problem1 Problem
		submit1  Submit
	)

	db.First(&submit1)
	db.Model(&submit1).Related(&problem1)
	t.Log(kjson.MarshalString(problem1))
	t.Log(kjson.MarshalString(submit1))
}

func TestHaveMany(t *testing.T) {
	var (
		db  *gorm.DB
		err error
	)

	if db, err = mysql.GetTestDB(); err != nil {
		t.Fatal("init mysql fail", err)
	}

	var (
		problem         Problem
		problemTestCase []ProblemTestCase
	)

	db.First(&problem)
	db.Model(&problem).Related(&problemTestCase)
	t.Log(kjson.MarshalString(problem))
	t.Log(kjson.MarshalString(problemTestCase))
}
