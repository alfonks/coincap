package database

import (
	"coincap/pkg/cfg"
	"log"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbItfOnce sync.Once
	dbItf     DBItf
)

func Init(cfg *cfg.ConfigSchema) DBItf {
	dbItfOnce.Do(func() {
		funcName := "db.Init"
		gormDB, err := gorm.Open(postgres.Open(cfg.DB.Credential), &gorm.Config{})
		if err != nil {
			log.Fatalf("[%v] fail init db", funcName)
		}

		dbConfig, err := gormDB.DB()
		if err != nil {
			log.Fatalf("[%v] fail get gorm db to set up config", funcName)
		}

		dbConfig.SetMaxOpenConns(cfg.DB.MaxConnLifeTime)
		dbConfig.SetConnMaxLifetime(time.Duration(cfg.DB.MaxConnLifeTime) * time.Second)

		dbItf = wrapGromDB(gormDB)
	})

	return dbItf
}
