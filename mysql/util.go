package mysql

import "github.com/jinzhu/gorm"

func GetTestDB() (db *gorm.DB, err error) {
	cfg := Config{
		ConnectionStr: "root:mysqlpass@tcp(127.0.0.1:3306)/killoj?&parseTime=True",
	}
	return InitDB(cfg)
}
