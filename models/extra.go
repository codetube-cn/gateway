package models

type ExtraArg struct {
	Arg     string `json:"arg"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Repeat  bool   `json:"repeat"`
	Mapping struct {
		Key struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"key"`
		Value struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"value"`
	} `json:"mapping"`
}

type ExtraOption struct {
	Label string
	Value interface{}
}

type ExtraDefaultValue struct {
	Value interface{}
}

func (v ExtraDefaultValue) IsString() bool {
	_, ok := v.Value.(string)
	return ok
}
