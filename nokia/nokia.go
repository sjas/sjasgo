package nokia

import (
	"maps"
	l "github.com/sirupsen/logrus"
	"os"
	"strings"
	"sync"
	"time"

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
	baseDir:="/dev/shm/nokialogs"
	err:=os.MkdirAll(baseDir,0755);if err!=nil{l.Fatal(err)}
	t:=time.Now().Local()
	formattedTime:=t.Format("YYYYMMDDTHHMMSS")
	fileName:=host+formattedTime+".log"
	filePath:=baseDir+"/"+fileName
	f,err:=os.OpenFile(filePath,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644);if err!=nil{l.Fatal(err)}
    p,err:=platform.NewPlatform(
        osType,
        host,
        options.WithAuthNoStrictKey(),
        options.WithAuthUsername(user),
        options.WithAuthPassword(pass),
        options.WithTransportType("standard"),
		options.WithTimeoutOps(30*time.Second),
		options.WithChannelLog(f),
    );if err!=nil{l.Error(err)}
    d,err:=p.GetNetworkDriver();if err!=nil{l.Error(err)}
    err=d.Open();if err!=nil{l.Error(err)}
    defer d.Close()
    //_,err=d.Channel.GetPrompt();if err!=nil{l.Error(err)}
	if mdcliEnabled{_,err=d.Channel.SendInput("environment more false");if err!=nil{l.Error(err)}
	}else{_,err=d.Channel.SendInput("environment no more");if err!=nil{l.Error(err)}}

	var fullResBytes []byte
	for _,i:=range cmd{
		resBytes,err:=d.Channel.SendInput(i);if err!=nil{l.Error(err)}
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
	l.Debug("started worker ",host)
	if mdcliEnabled{res[host]=Mdcli(host,cmd...)
	}else{res[host]=Classic(host,cmd...)}
	l.Debug("finished worker ",host)
	c<-res
	l.Debug("pushed to channel ",host)
	l.Debug(res)
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
		go runCommandWorker(i,mdcliEnabled,&wg,c,cmd...)
	}
	l.Debug("all workers started")
	go func(){wg.Wait();close(c)}()
	l.Debug("waitgroup finished and channel closed, accumulating result maps")
	for i:=range c{maps.Copy(res,i)}
	l.Debug("result maps accumulated")
	return res
}
