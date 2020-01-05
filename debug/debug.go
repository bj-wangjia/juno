package debug

import "encoding/json"

type Debug struct {
	Name string   `json:"name"`
	Msg  []string `json:"msg"`
	Node []*Debug `json:"node"`
}

func NewDebug(name string) *Debug {
	return &Debug{
		Name: name,
		Msg:  []string{},
	//	Node: []*Debug{},
	}
}

func (d *Debug) AddDebug(debug ...*Debug) {
	for _, v := range debug {
		d.Node = append(d.Node, v)
	}
}

func (d *Debug) AddDebugMsg(msg ...string) {
	for _, v := range msg {
		d.Msg = append(d.Msg, v)
	}
}

func (d *Debug) String() string {
	if res, err := json.Marshal(d); err == nil {
		return string(res)
	} else {
		return err.Error()
	}
}
