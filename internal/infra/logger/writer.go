package logger

import "gopkg.in/natefinch/lumberjack.v2"

type RollingWriterConfig struct {
	// Filename is the file to write logs to.
	// Backup logger files will be retained in the same directory.
	// os.TempDir() if empty.
	Filename string

	// MaxSize is the maximum size in megabytes of the logger file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int

	// MaxAge is the maximum number of days to retain old logger files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old logger files
	// based on age.
	MaxAge int

	// MaxBackups is the maximum number of old logger files to retain.  The default
	// is to retain all old logger files (though MaxAge may still cause them to get
	// deleted.)
	MaxBackups int

	// LocalTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time.  The default is to use UTC
	// time.
	LocalTime bool

	// Compress determines if the rotated logger files should be compressed
	// using gzip. The default is not to perform compression.
	Compress bool
}

// NewRollingWriter creates and returns new instance of io.writer, which
// can be used as rolling-writer
func NewRollingWriter(config RollingWriterConfig) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   config.Filename,   // logger file path
		MaxSize:    config.MaxSize,    // size of each logger file unit: M
		MaxAge:     config.MaxAge,     // how many days can the file be saved at most
		MaxBackups: config.MaxBackups, // how many backups can logger files save at most
		LocalTime:  config.LocalTime,  // localTime must be used for formatting timestamps
		Compress:   config.Compress,   // compress
	}
}
