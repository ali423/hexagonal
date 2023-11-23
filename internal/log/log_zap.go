package log

import (
	"fmt"
	"github.com/ali423/hexagonal/cmd/shotener/config"
	sngelf "github.com/snovichkov/zap-gelf"
	"go.uber.org/zap"
	uberzap "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
	"io"
	"os"
	"strings"
)

const writersSeparator string = ","

var (
	Debug = 1
	Info  = 2
	Warn  = 3
	Error = 4
	Panic = 5
	Fatal = 6
)

// ZapLog is an implementation of Logger interface
type ZapLog struct {
	// Config contains zap config
	Config *config.Config
}

func NewZapLog(conf *config.Config) *ZapLog {
	return &ZapLog{
		Config: conf,
	}
}

func (z *ZapLog) Debug(msg string, fields Fields) {
	z.log(Debug, msg, fields)
}

func (z *ZapLog) Info(msg string, fields Fields) {
	z.log(Info, msg, fields)
}

func (z *ZapLog) Warn(msg string, fields Fields) {
	z.log(Warn, msg, fields)
}

func (z *ZapLog) Error(msg string, fields Fields) {
	z.log(Error, msg, fields)
}

func (z *ZapLog) Panic(msg string, fields Fields) {
	z.log(Panic, msg, fields)
}

func (z *ZapLog) Fatal(msg string, fields Fields) {
	z.log(Fatal, msg, fields)
}

func (z *ZapLog) log(lvl int, msg string, fields Fields) {
	fields["facility"] = fmt.Sprintf("%v", z.Config.Facility)

	writers, _ := z.getWriters()

	var writeSyncers []zapcore.WriteSyncer
	for _, w := range writers {
		if _, ok := w.(*gelf.UDPWriter); ok {
			z.writeToGrayLog(lvl, msg, fields)
			continue
		}
		writeSyncers = append(writeSyncers, zapcore.AddSync(w))
	}

	// config encoder
	encoderConfig := uberzap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	enc := zapcore.NewConsoleEncoder(encoderConfig)

	// create suger logger
	core := zapcore.NewCore(enc, zapcore.NewMultiWriteSyncer(writeSyncers...), uberzap.InfoLevel)
	sugar := uberzap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
	defer func() {
		if err := sugar.Sync(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	z.doLog(sugar, lvl, msg, fields)
}

func (z *ZapLog) doLog(sugar *zap.SugaredLogger, lvl int, msg string, fields Fields) {

	switch lvl {
	case Debug:
		sugar.Debugw(msg, FieldsToArray(fields)...)
	case Info:
		sugar.Infow(msg, FieldsToArray(fields)...)
	case Warn:
		sugar.Warnw(msg, FieldsToArray(fields)...)
	case Error:
		sugar.Errorw(msg, FieldsToArray(fields)...)
	case Panic:
		sugar.Panicw(msg, FieldsToArray(fields)...)
	case Fatal:
		sugar.Fatalw(msg, FieldsToArray(fields)...)
	}

}

func (z *ZapLog) writeToGrayLog(lvl int, msg string, fields Fields) {
	hostname, err := os.Hostname()
	if err != nil {
		return
	}
	core, err := sngelf.NewCore(
		sngelf.Addr(fmt.Sprintf("%s:%s", z.Config.Host, z.Config.Port)),
		sngelf.Host(hostname),
	)
	if err != nil {
		panic(err)
	}

	var logger = zap.New(
		core,
		zap.AddCaller(),
	)
	defer func(logger *uberzap.Logger) {
		err := logger.Sync()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(logger)

	sugar := logger.
		Sugar()

	z.doLog(sugar, lvl, msg, fields)
}

func (z *ZapLog) getWriters() (writers []io.Writer, err error) {
	cnf := z.Config
	ws := strings.Split(cnf.Writers, writersSeparator)

	writers = make([]io.Writer, 0, len(ws))
	for _, w := range ws {
		switch strings.TrimSpace(w) {
		case "stdout":
			writers = append(writers, os.Stdout)
		case "graylog":
			var graylogWriter *gelf.UDPWriter
			if graylogWriter, err = InitGraylogWriter(cnf.GraylogWriter); err != nil {
				break
			}
			graylogWriter.Facility = cnf.GraylogWriter.Facility

			writers = append(writers, graylogWriter)
		}
	}

	return writers, err
}

func NewRollingWriterConfig(cnf config.RollingLogWriter) RollingWriterConfig {
	return RollingWriterConfig{
		Filename:   cnf.Filename,
		MaxSize:    cnf.MaxSize,
		MaxAge:     cnf.MaxAge,
		MaxBackups: cnf.MaxBackups,
		LocalTime:  cnf.LocalTime,
		Compress:   cnf.Compress,
	}
}

func NewGraylogWriterConfig(cnf config.GraylogWriter) string {
	return fmt.Sprintf("%s:%s", cnf.Host, cnf.Port)
}

func InitGraylogWriter(cnf config.GraylogWriter) (*gelf.UDPWriter, error) {
	udpWriter, _ := gelf.NewUDPWriter(NewGraylogWriterConfig(cnf))
	return udpWriter, nil
}
