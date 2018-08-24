package debugger

import (
	"bytes"
	js "encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

var (
	l *log.Logger
)

type Debugger struct {
	bf   *bytes.Buffer
	last time.Time
}

func Init(logPath string) {
	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err.Error())
	}
	l = log.New(f, "", log.LstdFlags|log.Lshortfile)
}

func Get() *Debugger {
	return &Debugger{
		bf:   bytes.NewBufferString(""),
		last: time.Now(),
	}
}

type Cond func() bool

func (d *Debugger) IfPrint(cond Cond) {
	if cond() {
		l.Println(d.bf.String())
	}
}

func (d *Debugger) M() {
	d.bf.WriteString(d.getCurrentDebuggerInfo(getCallInfo(1)))
}

func (d *Debugger) getCurrentDebuggerInfo(file string, line int) string {
	n := time.Now()
	rt := fmt.Sprintf("%s:%d. Cost time since last mark : %s . \n", file, line, n.Sub(d.last).String())
	d.last = n
	return rt
}

func getCallInfo(calldepth int) (file string, line int) {
	_, longfile, line, _ := runtime.Caller(calldepth + 1)
	for i := len(longfile) - 1; i > 0; i-- {
		if longfile[i] == '/' {
			file = longfile[i+1:]
			return
		}
	}
	file = longfile
	return
}

func Json(v interface{}) {
	bts, err := js.Marshal(v)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(bts[:]))
}
