package ulog

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init(){
	initLog()
}

func initLog() {
	log.SetFormatter(&log.JSONFormatter{})//设置日志的输出格式为json格式，还可以设置为text格式
	log.SetOutput(os.Stdout)//设置日志的输出为标准输出
	log.SetLevel(log.InfoLevel)//设置日志的显示级别，这一级别以及更高级别的日志信息将会输出
}

type Logger struct {
	StdPath string
}

var LoggerClient *Logger

func

func NewRR(){
	log.SetOutput()
	s1 := &Student{
		Age:  18,
		Name: "kaiying",
		Sex:  1,
	}
	log.Info(s1)
}
