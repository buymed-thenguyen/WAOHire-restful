package db

import "gorm.io/gorm"

var DB *gorm.DB

func InjectDB(db *gorm.DB) {
	DB = db
}

func BeginTx() *gorm.DB {
	return DB.Begin()
}
