package logPkg

import (
	"strings"
	"time"
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
