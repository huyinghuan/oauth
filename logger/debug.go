package logger

import (
	"bytes"
	"fmt"
	"log"
)

var (
	//DEBUG
	debugBuffer bytes.Buffer
	debugLogger = log.New(&debugBuffer, "DEBUG:", log.Lshortfile)
	debug       = true
)

func flushDebug() {
	fmt.Print(&debugBuffer)
	debugBuffer.Reset()
}

//SetDebug 是否记录DEBUG， 默认输出
func SetDebug(flag bool) {
	debug = flag
}

//Debug 输出文件路径
func Debug(msg ...interface{}) {
	if !debug {
		return
	}
	defer flushDebug()
	debugLogger.Output(2, fmt.Sprint(msg...))

}

//Debugf 输出文件路径
func Debugf(msg string, args ...interface{}) {
	if !debug {
		return
	}
	defer flushDebug()
	debugLogger.Output(2, fmt.Sprintf(msg, args...))
}
