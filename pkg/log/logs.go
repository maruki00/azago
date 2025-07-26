package logPkg

import (
	"fmt"
	"strings"
	"time"

	"github.com/maruki00/azago/internal/common"
)


func Log(args ...string){
	var logLine = strings.Builder{}
	logLine.WriteString("[azago] ")
	logLine.WriteString(time.Now().Format("2006/07/01 "))
	for _,arg := range args {
		logLine.WriteString(arg)
		logLine.WriteString(" ")
	}
	println(logLine.String())
}

func Error(args ...string){
	var logLine = strings.Builder{}
	logLine.WriteString(fmt.Sprintf("%s[mazago] %s" , common.Red,time.Now().Format("2006/07/01 ")))
	for _,arg := range args {
		logLine.WriteString(arg)
		logLine.WriteString(" ")
	}
	logLine.WriteString(common.)
	println(logLine.String())
}

func Info(args ...string){
	var logLine = strings.Builder{}
	logLine.WriteString("[azago] ")
	logLine.WriteString(time.Now().Format("2006/07/01 "))
	for _,arg := range args {
		logLine.WriteString(arg)
		logLine.WriteString(" ")
	}
	println(logLine.String())
}

