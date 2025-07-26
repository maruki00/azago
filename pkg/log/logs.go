package logPkg

import (
	"fmt"
	"strings"
	"time"

	cliPkg "github.com/maruki00/azago/pkg/cli"
)

func Log(args ...string) {
	var logLine = strings.Builder{}
	logLine.WriteString(fmt.Sprintf("%s[LOG  ] %s", cliPkg.Gray, time.Now().Format("2006/07/01 ")))
	for _, arg := range args {
		logLine.WriteString(arg)
		logLine.WriteString(" ")
	}
	println(logLine.String())
}

func Error(args ...string) {
	var logLine = strings.Builder{}
	logLine.WriteString(fmt.Sprintf("%s[ERROR] %s", cliPkg.Red, time.Now().Format("2006/07/01 ")))
	for _, arg := range args {
		logLine.WriteString(arg)
		logLine.WriteString(" ")
	}
	logLine.WriteString(cliPkg.Reset)
	println(logLine.String())
}


func Info(args ...string) {
	var logLine = strings.Builder{}
	logLine.WriteString(fmt.Sprintf("%s[INFO ] %s", cliPkg.Blue, time.Now().Format("2006/07/01 ")))
	for _, arg := range args {
		logLine.WriteString(arg)
		logLine.WriteString(" ")
	}
	logLine.WriteString(cliPkg.Reset)
	println(logLine.String())
}

