package bash

import (
	"bytes"
	"os"
	"os/exec"
	"strings"

	"github.com/gookit/color"
	"fmt"
)

// SCRIPTENV env var is used in local .bashrc to exclude the interactive session contents
// bashrc is loaded as all aliases and functions are sources from there, so i get the proper environment
// a plain .bashrc will not run with the aliases loaded as it exits upon $- not being marked interactive
const (
        SCRIPTENVFLAG=`SCRIPTENV="1"`
        SHELLEXECUTABLE=`/bin/bash`
        //SHELLFLAGS=`-xc`
        SHELLFLAGS=`-c`
)
func cmdToStringWrapper(shellcommand string, fullEnvironment bool)string{
		var cmd *exec.Cmd
		if fullEnvironment{
			cmd=exec.Command(SHELLEXECUTABLE,SHELLFLAGS,"{ source ~/.bashrc;"+shellcommand+";}")
		}else{
			cmd=exec.Command(SHELLEXECUTABLE,SHELLFLAGS,shellcommand)
		}
		fmt.Printf("%T\n",cmd)
        cmd.Env=append(os.Environ(),SCRIPTENVFLAG)
        //cmd.Env=append(os.Environ(),"BASH_ENV="+os.Getenv("HOME")+"/.bashrc")
        var stdout,stderr bytes.Buffer
        cmd.Stdout=&stdout
        cmd.Stderr=&stderr
        err:=cmd.Run()
        stderrcolor:=color.New(color.OpBold,color.FgLightRed,color.BgBlack)
        if stderr.String()!=""{stderrcolor.Println(stderr.String())}
        errcolor:=color.New(color.OpBold,color.FgYellow,color.BgBlack)
        if err!=nil{errcolor.Println(err)}
        return stdout.String()
}

func CmdToStringWithoutFullEnvironment(shellcommand string)string{
	return cmdToStringWrapper(shellcommand,false)
}

func CmdToString(shellcommand string)string{
	return cmdToStringWrapper(shellcommand,true)
}

func CmdToStringSlice(shellcommand string)[]string{
	res_string:=CmdToString(shellcommand)
	res:=strings.Split(res_string,"\n")
	//drop last element if empty. probably only needed because i use a newline via $PROMPT_COMMAND
	if len(res[len(res)-1])==0{res=res[:len(res)-1]}
	return res
}

func CmdToStringSliceWithCall(shellcommand string)[]string{
	prompt:=CmdToString(`echo $USER@$(hostname):$(pwd) $(if [[ $UID -eq 0 ]];then echo "#";else echo "$";fi)`)
	res_string:=prompt+shellcommand+"\n\n"
	res_string+=CmdToString(shellcommand)+"\n"+prompt
	temp:=strings.Split(res_string,"\n")
 	res:=temp[:len(temp)-1]
 	return res
}
