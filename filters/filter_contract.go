package filters

import (
	"net/http"
)

var (
	InputTypeText        = "text"     //输入方式：文本输入
	InputTypeSelect      = "select"   //输入方式：选择
	InputTypeMultiSelect = "checkbox" //输入方式：多选
	InputTypeAdvance     = "advance"  //输入方式：高级模式
)

//FilterInterface 过滤器接口
type FilterInterface interface {
	//LoadValue 载入过滤器值，参数一般为 json
	LoadValue(v string) error
	//Apply 应用过滤器
	Apply(exchange *ServerWebExchange, response *http.Response) error
}

//FilterContract 过滤器通用协议
type FilterContract struct {
	Name         string   //名称
	Code         string   //标识
	InputType    string   //输入方式：1|text|文本输入，2|select|选项，3|checkbox|多选，4|advance|高级模式（填写设定的args）
	Args         []Arg    //参数，只有高级输入模式才需要，正常使用 []Arg 即可，但更高要求时可以使用 []ArgCombo
	Options      []Option //选项，有输入方式为选择类型时需要使用
	DefaultValue string   //默认值，不同过滤器需要单独指定，为 string 或 []string
}

//Arg 参数
type Arg struct {
	Name         string   //参数名称
	Key          string   //参数 Key
	Repeat       bool     //是否允许重复，例如文本输入时，可以新增多组输入框
	InputType    string   //输入方式：1|text|文本输入，2|select|选项，3|checkbox|多选
	DefaultValue string   //默认值，不同过滤器需要单独指定，为 string 或 []string
	Options      []Option //选项，有输入方式为选择类型时需要使用
}

//ArgCombo 参数组合，如果某个参数需要多个参数组合到一起时，作为整体进行操作，值也会保存为一组对象
type ArgCombo struct {
	Name   string //参数名称，如果不显示可以留空
	Repeat bool   //是否允许重复，例如文本输入时，可以新增多组输入框
	Args   []Arg  //子参数
}

//Option 选项
type Option struct {
	Label string //显示名称
	Value string //值
}

type ServerWebExchange struct {
	Request *http.Request
}

type ResponseFilter func(*http.Response)


type GatewayFilter func(exchange *ServerWebExchange) ResponseFilter

func BuildServerWebExchange(request *http.Request) *ServerWebExchange {
	return &ServerWebExchange{Request: request}
}
