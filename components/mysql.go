package components

import (
	"codetube.cn/core/config"
	"codetube.cn/core/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB = newDatabases()

type databases struct {
	Gateway *database
}

func newDatabases() *databases {
	return &databases{
		Gateway: &database{
			config: &config.MysqlConfig{
				Dsn:     "codetube:djlwz123@tcp(www.daijulong.com:53306)/codetube_gateway?charset=utf8mb4&parseTime=True&loc=Local",
				Maxidle: 5,
				Maxopen: 10,
			},
		},
	}
}

type database struct {
	config *config.MysqlConfig
	DB     *gorm.DB
}

func (d *databases) MysqlInit() (err error) {
	err = d.Gateway.connect()
	if err != nil {
		return
	}
	//其他数据库...
	return
}

func (d *database) connect() error {
	if d.config == nil {
		return errors.Wrap("connect database error", errors.ErrConfigNotExist)
	}
	if d.config.Dsn == "" || d.config.Maxidle < 1 || d.config.Maxopen < 1 {
		return errors.Wrap("connect database error", errors.ErrConfigNotExist)
	}
	db, err := gorm.Open(mysql.Open(d.config.Dsn), &gorm.Config{})
	if err != nil {
		db = nil
		return errors.Wrap("connect database error", err)
	}
	d.DB = db
	//d.DB.SingularTable(true)
	sqlDb, err := d.DB.DB()
	sqlDb.SetMaxIdleConns(d.config.Maxidle)
	sqlDb.SetMaxOpenConns(d.config.Maxopen)
	sqlDb.SetConnMaxLifetime(time.Hour)
	log.Println("init database success")
	return nil
}