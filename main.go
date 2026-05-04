package main

import (
// 	"fmt"
// 	l "github.com/sirupsen/logrus"
//	"github.com/sjas/sjasgo/bash"
 	"github.com/sjas/sjasgo/pp"
// 	"github.com/sjas/sjasgo/nokia"
//	"strconv"
//	"time"
)

func main(){
// 	l.SetReportCaller(true)
// 	l.SetLevel(l.DebugLevel)

// 	pp.Long("date test")
 	pp.LongGreen("date test")
//	pp.LongNoDate("nodate test")
	pp.LongNoDateGreen("nodate test")
//	pp.ShortRed("red")
//	pp.ShortYellow("yellow")
//	pp.ShortGreen("green")

// 	start:=time.Now()
// 	fmt.Println(bash.CmdToString("hostname -f"))
// 	end:=time.Now()
// 	fmt.Println(end.Sub(start))
// 	start=time.Now()
// 	fmt.Println(bash.CmdToStringWithoutFullEnvironment("hostname -f"))
// 	end=time.Now()
// 	fmt.Println(end.Sub(start))

//	fmt.Println(bash.CmdToString("cat /etc/hosts"))
//	for _,i:=range(bash.CmdToStringSlice("ls -alh")){fmt.Println(i)}
//	for _,i:=range(bash.CmdToStringSliceWithCall("ls -alh")){fmt.Println(i)}
//	for _,i:=range(bash.CmdToStringSliceWithCall("ls -lisha")){fmt.Println(i)}
//	pp.ShortRed("delim")
//	for idx,i:=range(bash.CmdToStringSlice("ls -lisha")){fmt.Println(strconv.Itoa(idx)+": "+i)}
//	for idx,i:=range(bash.CmdToStringSliceWithCall("ls -lisha")){fmt.Println(strconv.Itoa(idx)+": "+i)}
// 	pp.Red("asdf")
// 	pp.Yellow("qwer")
// 	pp.Green("zxcv")
}
