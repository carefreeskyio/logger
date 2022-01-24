package hook

import (
	log "github.com/sirupsen/logrus"
)

type CtxHook struct {
	Prefix string
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
		entry.Data[hook.Prefix+field] = entry.Context.Value(field)
	}
	return nil
}
