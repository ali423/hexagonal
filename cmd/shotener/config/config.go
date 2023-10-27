package config

import (
	"errors"
)

var (
	ErrMissingParam = errors.New("default values missing. Pass a map[string]interface{} to this function")
	ErrInvalidParam = errors.New("invalid param. map[string]interface{} expected")
)

type DBType uint32

const (
	DBPostgres = iota
	DBMySql
	DBSqlite
	DBRedis
)

// Config holds application configuration
type Config struct {
	// AppAddress is the address that server listens to
	// Format: 127.0.0.1
	AppAddress string

	// AppPort is the port that server binds
	AppPort int

	// AppCode indicates the service code
	AppCode int

	// DBType represents type of database
	DBType DBType

	// DBUsername database username
	DBUsername string

	// DBPassword database password
	DBPassword string

	// DBHostname database hostname.
	// Format: localhost | 127.0.0.1
	DBHostname string

	// DBName database name
	DBName string

	// DBPort database port
	DBPort int

	Writers string

	RollingLogWriter
	GraylogWriter

	SecurityKey string
}

type RollingLogWriter struct {
	Filename   string `long:"filename" description:"Filename is the file to write logs to" env:"LOGGING_FILENAME"`
	MaxSize    int    `long:"maxsize" description:"MaxSize is the maximum size in megabytes of the logger file before it gets rotated" env:"LOGGING_MAX_SIZE"`
	MaxAge     int    `long:"maxage" description:"MaxAge is the maximum number of days to retain old logger files based on the timestamp encoded in their filename" env:"LOGGING_MAX_AGE"`
	MaxBackups int    `long:"maxbackups" description:"MaxBackups is the maximum number of old logger files to retain" env:"LOGGING_MAX_BACKUPS"`
	LocalTime  bool   `long:"localtime" description:"LocalTime determines if the time used for formatting the timestamps in backup files is the computer's local time" env:"LOGGING_LOCAL_TIME"`
	Compress   bool   `long:"compress" description:"Compress determines if the rotated logger files should be compressed using gzip" env:"LOGGING_COMPRESS"`
}

type GraylogWriter struct {
	Host     string `long:"host" description:"graylog host" env:"LOGGING_GRAYLOG_HOST"`
	Port     string `long:"port" description:"graylog port" env:"LOGGING_GRAYLOG_PORT"`
	Facility string `long:"facility" description:"graylog port" env:"LOGGING_GRAYLOG_FACILITY"`
}
