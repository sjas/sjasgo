package nokia

import (
	"errors"
	l "github.com/sirupsen/logrus"
	"os"
)

func getFileNameForHost(host string)string{return "/home/sjas/c/"+host}

func getFileNameIfConfigExists(host string)string{
	fileName:=getFileNameForHost(host)
	if _,err:=os.Stat(fileName);errors.Is(err, os.ErrNotExist){l.Fatal("missing config: "+fileName)}
    l.Info("checked if config exists")
	return fileName
}

func GetConfig(host string)string{
	fileName:=getFileNameIfConfigExists(host)
	b,err:=os.ReadFile(fileName);if err!=nil{l.Fatal(err)}
	configFileString:=string(b)
	return configFileString
}
