package nokia

import (
	"maps"
	l "github.com/sirupsen/logrus"
	"strings"
	"sync"

	"github.com/scrapli/scrapligo/driver/options"
	"github.com/scrapli/scrapligo/platform"
	"github.com/sjas/sjasgo/bash"
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
	l.Debug("classic command on host ",host)
	return runCommandWrapper(host,mdcliEnabled,cmd...)
}

func runCommandWorker(host string,mdcliEnabled bool,wg *sync.WaitGroup,c chan map[string]string,cmd ...string){
	defer wg.Done()
	res:=make(map[string]string)
	l.Debug("started worker")
	if mdcliEnabled{res[host]=Mdcli(host,cmd...)
	}else{res[host]=Classic(host,cmd...)}
	l.Debug("finished worker")
	c<-res
}

func Mdcli(host string,cmd ...string)string{
	mdcliEnabled:=true
	l.Debug("mdcli command on host ",host)
	return runCommandWrapper(host,mdcliEnabled,cmd...)
}

func RunCommandOnHostList(hostlist []string,mdcliEnabled bool,cmd ...string)map[string]string{
	res:=make(map[string]string)
	var wg sync.WaitGroup
	c:=make(chan map[string]string)
	for _,i:=range hostlist{
		wg.Add(1)
		l.Debug("trigger command worker ",i)
		runCommandWorker(i,mdcliEnabled,&wg,c,cmd...)
	}
	l.Debug("all workers started")
	go func(){wg.Wait();close(c)}()
	l.Debug("waitgroup finished and channel closed, accumulating result maps")
	for i:=range c{maps.Copy(res,i)}
	l.Debug("result maps accumulated")
	return res
}
