package nokia

import (
	"strings"
	"github.com/sjas/sjasgo/bash"
	"github.com/scrapli/scrapligo/driver/options"
    "github.com/scrapli/scrapligo/platform"
)

func runCommandWrapper(host string,cmd string,mdcliEnabled bool)string{
    user:=strings.TrimSpace(bash.CmdToStringWithoutFullEnvironment("get .aduser"))
	pass:=strings.TrimSpace(bash.CmdToStringWithoutFullEnvironment("get .adpass"))
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
    );if err!=nil{panic(err)}
    d,err:=p.GetNetworkDriver();if err!=nil{panic(err)}
    err=d.Open();if err!=nil{panic(err)}
    defer d.Close()
    _,err=d.Channel.GetPrompt();if err!=nil{panic(err)}

    res_bytes,err:=d.Channel.SendInput(cmd)
    if err!=nil{panic(err)}
    res:=string(res_bytes)
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
