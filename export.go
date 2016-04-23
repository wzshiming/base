package base

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	color "github.com/mitchellh/colorstring"
	"github.com/wzshiming/ffmt"
)

var (
	DEBUG_ = true
	err    = color.Color("[red]ERROR ")
	warn   = color.Color("[magenta]WARN  ")
	info   = color.Color("[blue]INFO  ")
	notice = color.Color("[cyan]NOTICE")
	pani   = color.Color("[dark_gray]PANIC ")
	debu   = color.Color("[yellow]DEBUG ")
	head   = color.Color(fmt.Sprintf("[light_green]\n%s ", filepath.Base(os.Args[0])))
)

func init() {
	log.SetFlags(log.Ltime)
	log.SetPrefix(head)
}

func ERR(v ...interface{}) {
	printf(err, v...)
}

func WARN(v ...interface{}) {
	printf(warn, v...)
}

func INFO(v ...interface{}) {
	printf(info, ffmt.Sputs(v...))
}

func NOTICE(v ...interface{}) {
	printf(notice, v...)
}

func PANIC(v ...interface{}) {
	panic(fmt.Errorf("%s\n%s", pani, fmt.Sprintln(v...)))
}

func printf(h string, v ...interface{}) {
	log.Printf("%s\n%s", h, fmt.Sprintln(v...))
}

func DEBUG(v ...interface{}) {
	if DEBUG_ {
		printf(debu, v...)
	}
}
