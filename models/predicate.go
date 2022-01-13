package models

import (
	"codetube.cn/core/model"
	"gorm.io/gorm"
)

// PredicateExtra 断言扩展参数
type PredicateExtra struct {
	Options      []*ExtraOption //扩展选项
	DefaultValue interface{}    `json:"default_value"` //默认值
}

// Predicate 断言模型
type Predicate struct {
	gorm.Model
	GatewayId  uint       //网关ID
	Name       string     //名称
	Code       string     //标识
	InputType  int        //输入方式
	ValueType  string     //值类型
	Extra      model.JSON `gorm:"type:json"` //扩展参数
	SortNumber int        //排序序号
}
