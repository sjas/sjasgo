package bash

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/gookit/color"
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

func Cmd(shellcommand string)string{
        cmd:=exec.Command(SHELLEXECUTABLE,SHELLFLAGS,"{ source ~/.bashrc;"+shellcommand+";}",)
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
