package models

import (
	"codetube.cn/core/model"
	"gorm.io/gorm"
)

// FilterExtra 过滤器扩展参数
type FilterExtra struct {
	Args struct { //扩展参数
		Args []*ExtraArg
	}
	Options struct { //扩展选项
		Options []*ExtraOption
	}
	DefaultValue struct { //默认值
		DefaultValue ExtraValue
	}
}

// Filter 过滤器模型
type Filter struct {
	gorm.Model
	GatewayId  uint       //网关ID
	Name       string     //过滤器名称
	Code       string     //过滤器标识
	InputType  int        //输入方式
	ValueType  string     //值类型
	Extra      model.JSON `json:"queryParam"` //扩展参数
	SortNumber int        //排序序号
}
