package main

import (
	"fmt"
	"github.com/sjas/sjasgo/bash"
	"github.com/sjas/sjasgo/pp"
)

func main(){
	pp.Long("some test")
	pp.ShortRed("red")
	pp.ShortYellow("yellow")
	pp.ShortGreen("green")
	fmt.Println(bash.Cmd("hostname -f"))
	fmt.Println(bash.Cmd("cat /etc/hosts"))
}
