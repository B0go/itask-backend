package db

import (
	"github.com/B0go/itask-backend/config"
	"github.com/B0go/itask-backend/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db  *gorm.DB
	Cfg *config.Config
}

func (d *Database) MustConnect() {
	db, err := gorm.Open(postgres.Open(d.Cfg.DatabaseConnString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	d.Db = db
}

func (d *Database) MustMigrate() {
	d.Db.AutoMigrate(&model.Task{})
}
