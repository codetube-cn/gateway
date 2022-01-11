package models

import (
	"codetube.cn/core/errors"
	"codetube.cn/gateway/vars"
)

type ExtraArg struct {
	Arg          string     `json:"arg"`           //参数 key
	Name         string     `json:"name"`          //参数名称
	Type         string     `json:"type"`          //参数值类型，仅支持 string 和 int
	Repeat       bool       `json:"repeat"`        //是否允许重复
	DefaultValue *ExtraValue `json:"default_value";gorm:"embedded"` //默认值
	Value        *ExtraValue `json:"Value"`
}

type ExtraOption struct {
	Label string
	Value interface{}
}

func (v ExtraArg) GetValue() (interface{}, error) {
	return v.getValue(*v.Value)
}

func (v ExtraArg) GetDefaultValue() (interface{}, error) {
	return v.getValue(*v.DefaultValue)
}

func (v ExtraArg) getValue(extraValue ExtraValue) (interface{}, error) {
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

type ExtraValue struct {
	Value interface{}
}

func (v ExtraValue) IsString() bool {
	_, ok := v.Value.(string)
	return ok
}

func (v ExtraValue) IsInt() bool {
	_, ok := v.Value.(int)
	return ok
}

func (v ExtraValue) IsStringList() bool {
	_, ok := v.Value.([]string)
	return ok
}

func (v ExtraValue) IsIntList() bool {
	_, ok := v.Value.([]int)
	return ok
}

func (v ExtraValue) GetValue() interface{} {
	return v.Value
}
