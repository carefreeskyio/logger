package hook

import (
	log "github.com/sirupsen/logrus"
	"path"
	"runtime"
	"strconv"
)

type CallerHook struct {
	Prefix string
}

func (hook CallerHook) Levels() []log.Level {
	return log.AllLevels
}
func (hook CallerHook) Fire(entry *log.Entry) error {
	if pc, file, line, ok := runtime.Caller(8); ok {
		funcName := runtime.FuncForPC(pc).Name()
		entry.Data[hook.Prefix+"source"] = file + ":" + strconv.Itoa(line)
		entry.Data[hook.Prefix+"func"] = path.Base(funcName)
	}
	return nil
}
