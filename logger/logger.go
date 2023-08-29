package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

func getTime() string {
	return time.Now().Format("15:04:05")
}

func color(str, color string) {
	_, file, line, ok := runtime.Caller(2) // Adjusted caller depth

	if !ok {
		log.Fatalln("Unable to get caller")
	}

	cwd, err := os.Getwd()

	if err != nil {
		log.Fatalln("Unable to get cwd")
	}

	fmt.Printf(
		"\x1b[38;2;160;129;226mcatgir.ls >.< \x1b[0m| %s%s \x1b[0m| \x1b[38;2;160;129;226m%s:%d \x1b[0m| %s\n",
		color, getTime(), file[len(cwd)+1:], line, str,
	)
}

func Log(str string) {
	color(str, "\x1b[38;2;159;234;121m")
}

func Warn(str string) {
	color(str, "\x1b[38;2;242;223;104m")
}

func Error(str string, panic bool) {
	color(str, "\x1b[38;2;242;106;104m")

	if panic {
		os.Exit(1)
	}
}
