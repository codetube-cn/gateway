package filters

import (
	"encoding/json"
	"net/http"
)

//CORSFilterDefaultValue CORSFilter 默认值
var CORSFilterDefaultValue = []CORSFilterValue{
	{Key: "Access-Control-Allow-Origin", Value: "*"},
	{Key: "Access-Control-Allow-Header", Value: "Origin, Content-Type, Cookie,X-CSRF-TOKEN, Accept,Authorization"},
	{Key: "Access-Control-Expose-Headers", Value: "Authorization,authenticated"},
	{Key: "Access-Control-Allow-Methods", Value: "GET, POST, PATCH, PUT, OPTIONS, DELETE"},
	{Key: "Access-Control-Allow-Credentials", Value: "true"},
}

//CORSFilterArgs 过滤器参数
var CORSFilterArgs = []ArgCombo{
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

type CORSFilter struct {
	FilterContract
	Value        []CORSFilterValue //值
	DefaultValue []CORSFilterValue //默认值
	Args         []ArgCombo
}

type CORSFilterValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewCORSFilter() FilterInterface {
	return &CORSFilter{}
}

//LoadValue 载入过滤器值，参数一般为 json
func (f *CORSFilter) LoadValue(v string) error {
	return json.Unmarshal([]byte(v), &f.Value)
}

//Apply 执行过滤器
func (f *CORSFilter) Apply(exchange *ServerWebExchange, response *http.Response) error {
	for _, v := range f.Value {
		response.Header.Add(v.Key, v.Value)
	}
	return nil
}
