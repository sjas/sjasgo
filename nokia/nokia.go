package nokia

import (
	"strings"
	"github.com/sjas/sjasgo/bash"
	"github.com/scrapli/scrapligo/driver/options"
    "github.com/scrapli/scrapligo/platform"
)

func runCommandWrapper(host string,mdcliEnabled bool,cmd ...string)string{
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

	var fullResBytes []byte
	for _,i:=range cmd{
		resBytes,err:=d.Channel.SendInput(i);if err!=nil{panic(err)}
		fullResBytes=append(fullResBytes,resBytes...)
		fullResBytes=append(fullResBytes,'\n')
		fullResBytes=append(fullResBytes,'\n')
	}
    res:=string(fullResBytes)
	return res
}

func Classic(host string,cmd ...string)string{
	mdcliEnabled:=false
	return runCommandWrapper(host,mdcliEnabled,cmd...)
}

func Mdcli(host string,cmd ...string)string{
	mdcliEnabled:=true
	return runCommandWrapper(host,mdcliEnabled,cmd...)
}
