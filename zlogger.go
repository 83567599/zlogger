package zlogger

import (
	"errors"
	"github.com/LK4D4/trylock"
	"log/syslog"
	"os"
	"path/filepath"
)

const (
	srvAddr = "127.0.0.1:514"
)

var writer *syslog.Writer
var cl trylock.Mutex
var logTag string

func getTag() string {
	if logTag != "" {
		return logTag
	}

	logTag = filepath.Base(os.Args[0])
	return logTag
}

func GetWriter() (*syslog.Writer, error) {
	if writer != nil {
		return writer, nil
	}

	if !cl.TryLock() {
		return nil, errors.New("try later")
	}
	defer cl.Unlock()

	var err error

	writer, err = syslog.Dial("udp", srvAddr, syslog.LOG_LOCAL5, getTag())

	return writer, err
}

func rewriteMsg(msg string) string {
	return msg
}

func Emerg(msg string) {
	wr, err := GetWriter()
	if err != nil {
		return
	}

	_ = wr.Emerg(rewriteMsg(msg))
}

func Crit(msg string) {
	wr, err := GetWriter()
	if err != nil {
		return
	}

	_ = wr.Crit(rewriteMsg(msg))
}

func Err(msg string) {
	wr, err := GetWriter()
	if err != nil {
		return
	}

	_ = wr.Err(rewriteMsg(msg))
}

func Warning(msg string) {
	wr, err := GetWriter()
	if err != nil {
		return
	}

	_ = wr.Warning(rewriteMsg(msg))
}

func Notice(msg string) {
	wr, err := GetWriter()
	if err != nil {
		return
	}

	_ = wr.Notice(rewriteMsg(msg))
}

func Info(msg string) {
	wr, err := GetWriter()
	if err != nil {
		return
	}

	_ = wr.Info(rewriteMsg(msg))
}

func Debug(msg string) {
	wr, err := GetWriter()
	if err != nil {
		return
	}

	_ = wr.Debug(rewriteMsg(msg))
}
