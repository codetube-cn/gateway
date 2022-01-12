package filter

import (
	"encoding/json"
	"net/http"
)

//SetHeaderDefaultValue SetHeader 默认值
var SetHeaderDefaultValue []SetHeaderFilterValue

//SetHeaderFilterArgs 过滤器参数
var SetHeaderFilterArgs = []ArgCombo{
	{
		Name:   "",
		Repeat: true,
		Args: []Arg{
			{
				Name:      "Header",
				Key:       "key",
				InputType: InputTypeText,
			},
			{
				Name:      "Value",
				Key:       "value",
				InputType: InputTypeText,
			},
		},
	},
}

// SetHeaderFilter set header 过滤器
type SetHeaderFilter struct {
	FilterContract
	Value        []SetHeaderFilterValue //值
	DefaultValue []SetHeaderFilterValue //默认值
	Args         []ArgCombo
}

// SetHeaderFilterValue 过滤器值
type SetHeaderFilterValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// NewSetHeaderFilter 创建过滤器
func NewSetHeaderFilter() FilterInterface {
	return &SetHeaderFilter{}
}

//LoadValue 载入过滤器值，参数一般为 json
func (f *SetHeaderFilter) LoadValue(v string) error {
	return json.Unmarshal([]byte(v), &f.Value)
}

//Apply 执行过滤器
func (f *SetHeaderFilter) Apply(exchange *ServerWebExchange, response *http.Response) error {
	for _, v := range f.Value {
		response.Header.Add(v.Key, v.Value)
	}
	return nil
}
