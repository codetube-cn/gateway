package components

import (
	coreConfig "codetube.cn/core/config"
	"codetube.cn/core/errors"
	"codetube.cn/gateway/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// DB 数据库连接
var DB = newDatabases()

// GatewayDB 网关数据库连接
var GatewayDB *gorm.DB

// 数据库连接列表
type databases struct {
	Gateway *database // 网关使用的数据库连接
}

// 创建数据库连接列表
func newDatabases() *databases {
	return &databases{
		Gateway: &database{
			config: config.GatewayConfig.Mysql.TransToGormConfig(),
		},
	}
}

// 数据库连接
type database struct {
	config *coreConfig.MysqlConfigForGorm
	DB     *gorm.DB
}

// MysqlInit 初始化数据库连接
func (d *databases) MysqlInit() (err error) {
	err = d.Gateway.connect()
	if err != nil {
		return
	}
	GatewayDB = d.Gateway.DB
	//其他数据库...
	return
}

// 连接数据库
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
