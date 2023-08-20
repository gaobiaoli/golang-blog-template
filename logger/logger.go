package logger

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

type LogFormatter struct {
}

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	Prefix := "[Blog] "
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", Prefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s\n", Prefix, timestamp, levelColor, entry.Level, entry.Message)

	}
	return b.Bytes(), nil
}

func init() {
	mlog := logrus.New()
	mlog.SetOutput(os.Stdout)
	mlog.SetLevel(logrus.DebugLevel)
	mlog.SetFormatter(&LogFormatter{})
	Log = mlog
}

func Logger(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Log.Debug(r.Host + r.URL.Path + " " + r.Method)
		h(w, r)
	}
}
