package models

import "gorm.io/gorm"

// Gateway 网关模型
type Gateway struct {
	gorm.Model
	Code string //标识
	Name string //名称
	Host string //监听主机
	Port int    //监听端口
}