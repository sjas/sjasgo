package main

import (
	"fmt"
	"github.com/sjas/sjasgo/bash"
	"github.com/sjas/sjasgo/pp"
)

func main(){
	pp.Header.Println("some test")
	fmt.Println(bash.Cmd("hostname -f"))
	fmt.Println(bash.Cmd("cat /etc/hosts"))
}
