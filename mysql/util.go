package mysql

import (
	"github.com/si9ma/KillOJ-backend/kerror"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetTestDB() (db *gorm.DB, err error) {
	cfg := Config{
		ConnectionStr: "root:mysqlpass@tcp(127.0.0.1:3306)/killoj?&parseTime=True",
	}
	return InitDB(cfg)
}

func ApplyDBError(c *gin.Context, err error) (hasErr bool, isNotFoundErr bool) {
	if err == nil {
		return false, false
	}

	if gorm.IsRecordNotFoundError(err) {
		return true, true
	}

	_ = c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(kerror.ErrInternalServerErrorGeneral)
	return true, false
}
