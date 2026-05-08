package nokia

import (
	"errors"
	l "github.com/sirupsen/logrus"
	"os"
)

func getFileNameForHost(host string)string{return "/home/sjas/c/"+host}

func CheckIfConfigsExistForAllHosts(list []string)bool{
	for _,i:=range list{
		fileName:=getFileNameForHost(i)
		if _,err:=os.Stat(fileName);errors.Is(err, os.ErrNotExist){l.Fatal("missing config: "+fileName)}
    }
    l.Info("checked if all configs exist")
	return true
}

func CheckIfConfigExists(host string)bool{
	fileName:=getFileNameForHost(host)
	if _,err:=os.Stat(fileName);errors.Is(err, os.ErrNotExist){l.Fatal("missing config: "+fileName)}
    l.Info("checked if config exists")
	return true
}
