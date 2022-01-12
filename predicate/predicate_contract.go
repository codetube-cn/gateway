package predicate

import "net/http"

var (
	InputTypeText        = "text"     //输入方式：文本输入
	InputTypeSelect      = "select"   //输入方式：选择
	InputTypeMultiSelect = "checkbox" //输入方式：多选
)

//PredicateInterface 断言接口
type PredicateInterface interface {
	//LoadValue 载入断言值，参数一般为 json
	LoadValue(v string) error
	//Match 匹配断言
	Match(request *http.Request) bool
}

//PredicateContract 断言通用协议
type PredicateContract struct {
	Name         string   //名称
	Code         string   //标识
	InputType    string   //输入方式：1|text|文本输入，2|select|选项，3|checkbox|多选，4|advance|高级模式（填写设定的args）
	Options      []Option //选项，有输入方式为选择类型时需要使用
	DefaultValue string   //默认值，不同过滤器需要单独指定，为 string 或 []string
}

//Option 选项
type Option struct {
	Label string //显示名称
	Value string //值
}
