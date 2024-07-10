package config

import (
	"io"
	"log"
	"os"
)

type Looger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func (l *Looger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Looger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Looger) Warning(v ...interface{}) {
	l.warning.Println(v...)
}

func (l *Looger) Error(v ...interface{}) {
	l.err.Println(v...)
}

func (l *Looger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *Looger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Looger) Warningf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}

func (l *Looger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}

func NewLogger(p string) *Looger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Looger{
		debug:   log.New(writer, "DEBUG", logger.Flags()),
		info:    log.New(writer, "INFO", logger.Flags()),
		warning: log.New(writer, "WARNING", logger.Flags()),
		err:     log.New(writer, "ERROR", logger.Flags()),
		writer:  writer,
	}
}
