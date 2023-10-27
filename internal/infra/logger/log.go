package logger

import "fmt"

const (
	ValDebug   = "DEBUG"
	ValInfo    = "INFO"
	ValWarning = "WARNING"
	ValError   = "ERROR"
	ValPanic   = "PANIC"
	ValFatal   = "FATAL"

	ValCorrelationIdQueue = "QUEUE"
)

// Logger is an interface which will be used in different parts of this software
// example logger: "user logged in", {"username": "test", "time": "12:34:56"}
type Logger interface {
	Debug(string, Fields)
	Info(string, Fields)
	Warn(string, Fields)
	Error(string, Fields)
	Panic(string, Fields)
	Fatal(string, Fields)
}

type Fields map[string]interface{}

func FieldsToArray(f Fields) []interface{} {
	var ret []interface{}
	for k, v := range f {
		if fmt.Sprintf("%v", k) != "facility" {
			k = fmt.Sprintf("ctxt_%v", k)
		}
		ret = append(ret, k, fmt.Sprintf("%v", v))
	}
	return ret
}
