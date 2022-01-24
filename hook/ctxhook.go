package hook

import (
	"github.com/carefreex-io/logger"
	log "github.com/sirupsen/logrus"
)

type CtxHook struct {
	Fields []string
}

func (hook *CtxHook) AddField(field string) {
	hook.Fields = append(hook.Fields, field)
}

func (hook *CtxHook) AddFields(fields []string) {
	hook.Fields = append(hook.Fields, fields...)
}

func (hook CtxHook) Levels() []log.Level {
	return log.AllLevels
}

func (hook CtxHook) Fire(entry *log.Entry) error {
	if entry.Context == nil {
		return nil
	}
	for _, field := range hook.Fields {
		entry.Data[logger.DefaultCustomOptions.FieldPrefix+field] = entry.Context.Value(field)
	}
	return nil
}
