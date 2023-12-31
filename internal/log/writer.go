package log

import "gopkg.in/natefinch/lumberjack.v2"

type RollingWriterConfig struct {
	// Filename is the file to write logs to.
	// Backup log files will be retained in the same directory.
	// os.TempDir() if empty.
	Filename string

	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int

	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int

	// MaxBackups is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.)
	MaxBackups int

	// LocalTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time.  The default is to use UTC
	// time.
	LocalTime bool

	// Compress determines if the rotated log files should be compressed
	// using gzip. The default is not to perform compression.
	Compress bool
}

// NewRollingWriter creates and returns new instance of io.writer, which
// can be used as rolling-writer
func NewRollingWriter(config RollingWriterConfig) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   config.Filename,   // log file path
		MaxSize:    config.MaxSize,    // size of each log file unit: M
		MaxAge:     config.MaxAge,     // how many days can the file be saved at most
		MaxBackups: config.MaxBackups, // how many backups can log files save at most
		LocalTime:  config.LocalTime,  // localTime must be used for formatting timestamps
		Compress:   config.Compress,   // compress
	}
}
