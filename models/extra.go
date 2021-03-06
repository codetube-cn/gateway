package models

import (
	"codetube.cn/core/errors"
	"codetube.cn/gateway/vars"
)

// ExtraArg 扩展参数
type ExtraArg struct {
	Arg          string      `json:"arg"`                           //参数 key
	Name         string      `json:"name"`                          //参数名称
	Type         string      `json:"type"`                          //参数值类型，仅支持 string 和 int
	Repeat       bool        `json:"repeat"`                        //是否允许重复
	DefaultValue *ExtraValue `json:"default_value";gorm:"embedded"` //默认值
	Value        *ExtraValue `json:"Value"`
}

// ExtraOption 扩展选项
type ExtraOption struct {
	Label string
	Value interface{}
}

// GetValue 获取扩展参数值
func (v *ExtraArg) GetValue() (interface{}, error) {
	return v.getValue(*v.Value)
}

// GetDefaultValue 获取扩展参数默认值
func (v *ExtraArg) GetDefaultValue() (interface{}, error) {
	return v.getValue(*v.DefaultValue)
}

//检查并取得扩展参数实际值
func (v *ExtraArg) getValue(extraValue ExtraValue) (interface{}, error) {
	ok := false
	if v.Type == "string" {
		if v.Repeat {
			ok = extraValue.IsStringList()
		} else {
			ok = extraValue.IsString()
		}
	} else if v.Type == "int" {
		if v.Repeat {
			ok = extraValue.IsIntList()
		} else {
			ok = extraValue.IsInt()
		}
	}

	if !ok {
		return nil, errors.Wrap("type["+v.Type+"]", vars.ErrGatewayValueTypeMismatch)
	}

	return extraValue.Value, nil
}

// ExtraValue 扩展参数值
type ExtraValue struct {
	Value interface{}
}

// IsString 扩展参数值是否是字符串
func (v *ExtraValue) IsString() bool {
	_, ok := v.Value.(string)
	return ok
}

// IsInt 扩展参数值是否是整数
func (v *ExtraValue) IsInt() bool {
	_, ok := v.Value.(int)
	return ok
}

// IsStringList 扩展参数值是否是字符串列表
func (v *ExtraValue) IsStringList() bool {
	_, ok := v.Value.([]string)
	return ok
}

// IsIntList 扩展参数值是否是整数列表
func (v *ExtraValue) IsIntList() bool {
	_, ok := v.Value.([]int)
	return ok
}

// GetValue 获取扩展参数实际值
func (v *ExtraValue) GetValue() interface{} {
	return v.Value
}
