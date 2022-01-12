package filter

import (
	"encoding/json"
	"net/http"
	"strings"
)

var StripPrefixDefaultValue = ""

//StripPrefixFilterArgs 过滤器参数
var StripPrefixFilterArgs = &Arg{
	Name:         "去除前缀",
	Key:          "prefix",
	Repeat:       false,
	InputType:    InputTypeSelect,
	DefaultValue: "no",
	Options: []Option{
		{Label: "是", Value: "yes"},
		{Label: "否", Value: "no"},
	},
}

// StripPrefixFilter strip prefix 过滤器
type StripPrefixFilter struct {
	FilterContract
	Value        string //值
	DefaultValue string //默认值
}

// NewStripPrefixFilter 创建过滤器
func NewStripPrefixFilter() FilterInterface {
	return &StripPrefixFilter{}
}

//LoadValue 载入过滤器值
func (f *StripPrefixFilter) LoadValue(v string) error {
	return json.Unmarshal([]byte(v), &f.Value)
}

//Apply 执行过滤器
func (f *StripPrefixFilter) Apply(exchange *ServerWebExchange, response *http.Response) error {
	if f.Value == "yes" {
		path := exchange.Request.URL.Path
		pathList := strings.Split(path, "/")
		//暂时只去除最前面的一节
		exchange.Request.URL.Path = strings.Join(pathList[2:], "/")
	}
	return nil
}
