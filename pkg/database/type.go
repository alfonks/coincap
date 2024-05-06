package database

import "gorm.io/gorm"

type (
	DBItf interface {
		Exec(sql string, values ...interface{}) DBItf
		Raw(sql string, values ...interface{}) DBItf
		Begin() DBItf
		Commit() DBItf
		Scan(dest interface{}) DBItf
		Error() error
	}

	db struct {
		conn *gorm.DB
	}
)
