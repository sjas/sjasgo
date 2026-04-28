package pp

import (
	"fmt"
	l "github.com/sirupsen/logrus"
	"os"
	"strings"
	"golang.org/x/term"
	"github.com/gookit/color"
)

/*
pph() {
        [[ $# == 0 ]]&&echo "usage: pph _char _maxcharpos _asciicolor YOUR_RANDOM_STRING_STUFF"&&return
        local _char="$1";shift
        local _maxcharpos=$1;shift
        local _asciicolor="$1";shift
        local _inputstring="${*}"
        #c="$(IFS=';' read -sdR -p $'\e[6n' ROW COL;echo "${ROW#*[}")"  ## cursorpos, unneeded
        local _inputstringsize=${#_inputstring}  ## we need stringlength instead
        local _terminalwidth=
        _terminalwidth="$(tput cols)"
        local _fillstringwidth=$(( _maxcharpos - 1 - _inputstringsize ))

        printf "\n\e[%sm%s" "${_asciicolor}" "${_inputstring}"
        if [[ ${_maxcharpos} -le ${_terminalwidth} ]]&&[[ ${_maxcharpos} -gt $(( _inputstringsize + 1 )) ]]
        then
                printf " "
                printf "${_char}%.0s" $( seq 1 ${_fillstringwidth} )
        elif [[ ${_maxcharpos} -gt ${_terminalwidth} ]]&&[[ ${_maxcharpos} -gt $(( _inputstringsize + 1 )) ]]
                then
                _fillstringwidth=$(( _terminalwidth - 1 - _inputstringsize ))
                printf " "
                printf "${_char}%.0s" $( seq 1 ${_fillstringwidth} )
        fi
        printf "\e[m\n"
}
*/

func pph(fillChar string,colorfg color.Color,colorbg color.Color,maxCharPosition int,input ...string){
	args:=make([]any,len(input))
    for i,v:=range input{args[i]=v}
	inputString:=strings.TrimSpace(fmt.Sprintln(args...))
	prettyPrint:=color.New(color.OpBold,colorfg,colorbg)
	inputStringLength:=len(inputString)
	terminalWidth,_,err:=term.GetSize(int(os.Stdout.Fd()));if err!=nil{l.Fatal(err)}
	res:=inputString+" "
	if maxCharPosition<=terminalWidth && maxCharPosition>(inputStringLength+1){
		res+=strings.Repeat(fillChar,maxCharPosition-len(inputString)-1)
	} else
	if maxCharPosition>terminalWidth&&maxCharPosition>(inputStringLength+1){
		res+=strings.Repeat(fillChar,terminalWidth-len(inputString)-1)
	}
	prettyPrint.Println(res)
}

/*
pp() { pph '*' 109 '41;1' "$* * $(TZ=UTC date +%FT%T%N)";}
ppnocolor() { pph '*' 109 '' "$* * $(TZ=UTC date +%FT%T%N)";}
ppp() { pph '*' 109 '41;1' "$*";}
ppnodate() { pph '*' 109 '41;1' "$*";}
pps() { pph '+' 80 '42' "$*";}
ppss() { pph '+' 50 '33' "$*";}
*/

func Long(input ...string){pph("*",color.FgLightWhite,color.BgLightRed,109,input...)}
func ShortRed(input ...string){pph("+",color.FgWhite,color.BgRed,55,input...)}
func ShortYellow(input ...string){pph("+",color.FgWhite,color.BgYellow,55,input...)}
func ShortGreen(input ...string){pph("+",color.FgWhite,color.BgGreen,55,input...)}
