package logs

import (
	"github.com/astaxie/beego/logs"
)

var Logging *logs.BeeLogger

func init() {
	Logging = logs.NewLogger(100000)
	Logging.SetLogger("console", "")
	Logging.SetLogger("file", `{"filename":"out.log"}`)
	Logging.EnableFuncCallDepth(true)
	Logging.Async()
}

func Debugf(format string, args ...interface{}) {
	Logging.Debug(format, args...)
	Logging.Flush()
}

func Infof(format string, args ...interface{}) {
	Logging.Info(format, args...)
	//logging.Flush()
}

func Warningf(format string, args ...interface{}) {
	Logging.Warning(format, args...)
	Logging.Flush()
}

func Errorf(format string, args ...interface{}) {
	Logging.Error(format, args...)
	Logging.Flush()
}

/*import (
	"fmt"
	"log"
	"os"
)

const (
	LEVEL_DEBUG = iota
	LEVEL_INFO
	LEVEL_WARN
	LEVEL_ERROR
	LEVEL_FATAL
)

var (
	logger   *log.Logger
	logLevel int = LEVEL_DEBUG
)

func init() {
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func SetLevel(level int) {
	logLevel = level
}

func logByLevel(level int, prefix string, format string, v ...interface{}) {
	if level < logLevel {
		return
	}
	logger.SetPrefix(prefix)
	logger.Println(fmt.Sprintf(format, v...))
}

func Debug(format string, v ...interface{}) {
	logByLevel(LEVEL_DEBUG, "[D]", format, v...)
}

func Info(format string, v ...interface{}) {
	logByLevel(LEVEL_INFO, "[I]", format, v...)
}

func Warn(format string, v ...interface{}) {
	logByLevel(LEVEL_WARN, "[W]", format, v...)
}

func Error(format string, v ...interface{}) {
	logByLevel(LEVEL_ERROR, "[E]", format, v...)
}

func Fatal(format string, v ...interface{}) {
	logByLevel(LEVEL_FATAL, "[F]", format, v...)
}*/
