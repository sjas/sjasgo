package textfsm

import (
	"fmt"
	l "github.com/sirupsen/logrus"
	"github.com/sirikothe/gotextfsm"
)

func Parse(input string,template string)[]map[string]any{
	fsm:=gotextfsm.TextFSM{}
	if err:=fsm.ParseString(template);err!=nil{l.Fatal(err)}
	parser:=gotextfsm.ParserOutput{}
	if err:=parser.ParseTextString(input,fsm,true);err!=nil{l.Fatal(err)}
	if l.IsLevelEnabled(l.TraceLevel){for _,i:=range parser.Dict{fmt.Println(i)}}
	return parser.Dict
}
