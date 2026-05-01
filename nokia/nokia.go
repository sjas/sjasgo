package nokia

import (
	l "github.com/sirupsen/logrus"
	"github.com/sjas/sjasgo/bash"
	"github.com/scrapli/scrapligo/driver/options"
    "github.com/scrapli/scrapligo/platform"

)

func runCommandWrapper(host string,cmd string,mdcliEnabled bool)string{
    user:=bash.CmdToStringWithoutFullEnvironment("getuser .aduser")
	pass:=bash.CmdToStringWithoutFullEnvironment("getpass .adpass")
	var osType string
	if mdcliEnabled{
		osType="nokia_sros"
	}else{
		osType="nokia_sros_classic"
	}
    p,err:=platform.NewPlatform(
        osType,
        host,
        options.WithAuthNoStrictKey(),
        options.WithAuthUsername(user),
        options.WithAuthPassword(pass),
        options.WithTransportType("system"),
    );if err!=nil{l.Fatal(err)}
    d,err:=p.GetNetworkDriver();if err!=nil{l.Fatal(err)}
    err=d.Open();if err!=nil{l.Error(err)}
    defer d.Close()
    prompt,err:=d.Channel.GetPrompt();if err!=nil{l.Fatal(err)}
    l.Info("found prompt: ",string(prompt))

    res_bytes,err:=d.Channel.SendInput(cmd)
    if err!=nil{l.Fatal(err)}
    res:=string(res_bytes)
	l.Debug("ran commmand "+cmd+" on host "+host)
    l.Debug(res)
	return res
}

func Classic(host string,cmd string)string{
	mdcliEnabled:=false
	return runCommandWrapper(host,cmd,mdcliEnabled)
}

func Mdcli(host string,cmd string)string{
	mdcliEnabled:=true
	return runCommandWrapper(host,cmd,mdcliEnabled)
}
