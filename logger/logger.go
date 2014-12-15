package logger

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"sync/atomic"
	"time"
)

var (
	colors map[string]string
	logNo  uint64
)

const (
	Black = (iota + 30)
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

type Worker struct {
	Minion  *log.Logger
	Color   int
	LogFile *os.File
}

type Info struct {
	Id      uint64
	Time    string
	Module  string
	Level   string
	Message string
	format  string
}

type Logger struct {
	Module string
	Worker *Worker
}

func (i *Info) Output() string {
	msg := fmt.Sprintf(i.format, i.Id, i.Time, i.Level, i.Message)
	return msg
}

func NewWorker(prefix string, flag int, color int, out io.Writer) *Worker {
	return &Worker{Minion: log.New(out, prefix, flag), Color: color, LogFile: nil}
}

func NewConsoleWorker(prefix string, flag int, color int) *Worker {
	return NewWorker(prefix, flag, color, os.Stdout)
}

func NewFileWorker(prefix string, flag int, color int, logFile *os.File) *Worker {
	return &Worker{Minion: log.New(logFile, prefix, flag), Color: color, LogFile: logFile}
}

func (w *Worker) Log(level string, calldepth int, info *Info) error {
	if w.Color != 0 {
		buf := &bytes.Buffer{}
		buf.Write([]byte(colors[level]))
		buf.Write([]byte(info.Output()))
		buf.Write([]byte("\033[0m"))
		return w.Minion.Output(calldepth+1, buf.String())
	} else {
		return w.Minion.Output(calldepth+1, info.Output())
	}
}

func colorString(color int) string {
	return fmt.Sprintf("\033[%dm", int(color))
}

func initColors() {
	colors = map[string]string{
		"CRITICAL": colorString(Magenta),
		"ERROR":    colorString(Red),
		"WARNING":  colorString(Yellow),
		"NOTICE":   colorString(Green),
		"DEBUG":    colorString(Cyan),
		"INFO":     colorString(White),
	}
}

func New(module string, color int) (*Logger, error) {
	initColors()
	newWorker := NewConsoleWorker("", 0, color)
	return &Logger{Module: module, Worker: newWorker}, nil
}

func NewDailyLogger(module string, color int, logPath string) (*Logger, error) {

	var logFile string
	const layout = "2006-01-02"
	now := time.Now()
	fileName := now.Format(layout)
	if len(logPath) == 0 {
		logFile = "./" + fileName + ".log"
	} else {
		logFile = logPath + "/" + fileName + ".log"
	}
	fileHandler, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	} else {
		initColors()
		newWorker := NewFileWorker("", 0, color, fileHandler)
		return &Logger{Module: module, Worker: newWorker}, nil
	}
}

func (l *Logger) Log(lvl string, message string) {
	var formatString string = "#%d %s ▶ %.3s %s"
	info := &Info{
		Id:      atomic.AddUint64(&logNo, 1),
		Time:    time.Now().Format("2006-01-02 15:04:05"),
		Module:  l.Module,
		Level:   lvl,
		Message: message,
		format:  formatString,
	}
	l.Worker.Log(lvl, 2, info)
}

func (l *Logger) Fatal(v ...interface{}) {
	message := fmt.Sprint(v)
	l.Log("CRITICAL", message)
	os.Exit(1)
}

func (l *Logger) Panic(v ...interface{}) {
	message := fmt.Sprint(v)
	l.Log("CRITICAL", message)
	panic(message)
}

func (l *Logger) Critical(v ...interface{}) {
	message := fmt.Sprint(v)
	l.Log("CRITICAL", message)
}

func (l *Logger) Error(v ...interface{}) {
	message := fmt.Sprint(v)
	l.Log("ERROR", message)
}

func (l *Logger) Warning(v ...interface{}) {
	message := fmt.Sprint(v)
	l.Log("WARNING", message)
}

func (l *Logger) Notice(v ...interface{}) {
	message := fmt.Sprint(v)
	l.Log("NOTICE", message)
}

func (l *Logger) Info(v ...interface{}) {
	message := fmt.Sprint(v)
	l.Log("INFO", message)
}

func (l *Logger) Debug(v ...interface{}) {
	message := fmt.Sprint(v)
	l.Log("DEBUG", message)
}

func (l *Logger) Strack(v ...interface{}) {
	message := fmt.Sprint(v)
	message += "\n"
	buf := make([]byte, 1024*1024)
	n := runtime.Stack(buf, true)
	message += string(buf[:n])
	message += "\n"
	l.Log("STRACK", message)
}

func (l *Logger) Close() {
	if l.Worker.LogFile != nil {
		l.Worker.LogFile.Close()
	}
}
