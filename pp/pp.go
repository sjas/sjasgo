package pp

import (
	"fmt"
	l "github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
	"golang.org/x/term"
	"github.com/gookit/color"
)

func pph(fillChar string,colorfg color.Color,colorbg color.Color,maxCharPosition int,showTime bool,input ...string){
	args:=make([]any,len(input))
    for i,v:=range input{args[i]=v}
	inputString:=strings.TrimSpace(fmt.Sprintln(args...))
	if maxCharPosition>100{inputString="\n"+inputString}
	if showTime{inputString=inputString+" * "+time.Now().Local().Format(time.RFC3339)}
	prettyPrint:=color.New(color.OpBold,colorfg,colorbg)
	inputStringLength:=len(inputString)
	terminalWidth:=110
	if term.IsTerminal(int(os.Stdout.Fd())){
		w,_,err:=term.GetSize(int(os.Stdout.Fd()))
		if err!=nil{l.Warn(err)
		}else{terminalWidth=w}
	}
	res:=inputString
	if len(fillChar)>0{res+=" "}
	if maxCharPosition<=terminalWidth&&maxCharPosition>(inputStringLength+1){
		res+=strings.Repeat(fillChar,maxCharPosition-len(inputString)-1)
	}else
	if maxCharPosition>terminalWidth&&maxCharPosition>(inputStringLength+1){
		res+=strings.Repeat(fillChar,terminalWidth-len(inputString)-1)
	}
	prettyPrint.Print(res)
}

func Long(input ...string){pph("*",color.FgLightWhite,color.BgLightRed,109,true,input...);fmt.Println()}
func LongGreen(input ...string){pph("*",color.FgLightWhite,color.BgGreen,109,true,input...);fmt.Println()}
func LongNoDateGreen(input ...string){pph("*",color.FgLightWhite,color.BgGreen,109,false,input...);fmt.Println()}
func LongNoDate(input ...string){pph("*",color.FgLightWhite,color.BgLightRed,109,false,input...);fmt.Println()}
func ShortRed(input ...string){pph("+",color.FgWhite,color.BgRed,55,false,input...);fmt.Println()}
func ShortYellow(input ...string){pph("+",color.FgWhite,color.BgYellow,55,false,input...);fmt.Println()}
func ShortGreen(input ...string){pph("+",color.FgWhite,color.BgGreen,55,false,input...);fmt.Println()}

func Red(input ...string){pph("",color.FgWhite,color.BgRed,55,false,input...)}
func Yellow(input ...string){pph("",color.FgWhite,color.BgYellow,55,false,input...)}
func Green(input ...string){pph("",color.FgWhite,color.BgGreen,55,false,input...)}
