package predicate

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
)

// PathPredicate path 断言
type PathPredicate struct {
	Contract
	Value        string //值
	DefaultValue string //默认值
}

// NewPathPredicate 创建 path 断言
func NewPathPredicate() Interface {
	return &PathPredicate{Value: "", DefaultValue: ""}
}

//LoadValue 载入断言值，参数一般为 json
func (p *PathPredicate) LoadValue(v string) error {
	return json.Unmarshal([]byte(v), &p.Value)
}

//Match 匹配断言
func (p *PathPredicate) Match(request *http.Request) bool {
	//如果完全相等，直接通过
	if p.Value == request.URL.Path {
		return true
	}
	//无特殊标识，匹配失败
	if !strings.ContainsAny(p.Value, ":*") {
		return false
	}
	//有特殊标识的，分别切割，按段匹配
	//规则分段
	patterns := strings.Split(strings.Trim(p.Value, "/"), "/")
	//path 分段
	paths := strings.Split(strings.Trim(request.URL.Path, "/"), "/")

	//规则分段 必不能少于 path 分段，path 更多，而规则没那么，肯定是匹配不上的
	pathLen := len(paths)
	if len(patterns) < pathLen {
		return false
	}

	//是否已经存在可选的分段
	//在已经在前面匹配过可选分段后，后续不允许再出现必选分段
	existOptionalPath := false

	//循环匹配，以可能较多的 values 为准
	for k, v := range patterns {
		firstChar := string(v[0])
		//非特殊匹配
		if firstChar != "?" && firstChar != ":" {
			//未匹配上，结束
			if v != paths[k] {
				return false
			}
			//匹配上了，直接下一段
			continue
		}
		//格式标识符：
		// \d: 数字
		// \c: 非数字，纯字母
		// \a: 不限，等价于 \c\d\s
		// \s: 特殊字符，仅支持：“.”、“_”、"-"
		// \D: 数字+特殊字符，等价于 \d + \s + “_” + "-"
		// \C: 非数字+特殊字符，等价于 \c + \s
		//如果是必选值，则对应的 paths 也必须要有
		isMatchRule := string(v[1]) == "\\" //第二个字符，如果是 \ 则需要进行规则匹配
		rule := "a"                         //匹配规则，默认不限
		if isMatchRule {
			rule = string(v[2]) //第三个字符，匹配的规则
		}
		//如果是必选匹配分段，则一定要匹配上
		if firstChar == ":" {
			//前面已经出现了可选分段，不再允许后面有必选分段
			if existOptionalPath {
				return false
			}
			//到某分段无法匹配，直接退出
			if pathLen <= k || paths[k] == "" || !p.matchRule(rule, paths[k]) {
				return false
			}
		} else {
			//标记已经出现过可选分段
			existOptionalPath = true

			//非必选匹配分段，可以没有，但有就必须要匹配
			if pathLen <= k {
				continue
			}
			if paths[k] != "" && !p.matchRule(rule, paths[k]) {
				return false
			}
		}
	}

	return true
}

//值匹配规则
func (p *PathPredicate) matchRule(rule string, value string) bool {
	if value == "" {
		return false
	}
	pattern := ""
	switch rule {
	case "d":
		pattern = "^[\\d]+$"
	case "c":
		pattern = "^[A-Za-z]+$"
	case "a":
		pattern = "^[\\w\\d\\.\\-_]+$"
	case "s":
		pattern = "^[-_\\.]+$"
	case "D":
		pattern = "^[\\d\\-_\\.]+$"
	case "C":
		pattern = "^[A-Za-z\\-_\\.]+$"
	}
	if pattern == "" {
		return false
	}

	matched, _ := regexp.Match(pattern, []byte(value))
	return matched
}
