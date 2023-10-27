package config

import (
	"strings"

	"github.com/spf13/viper"
)

type ViperLoader struct{}

func (v *ViperLoader) Load(params ...interface{}) (*Config, error) {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	c := &Config{
		AppAddress: viper.GetString("app.address"),
		AppPort:    viper.GetInt("app.port"),
		AppCode:    viper.GetInt("app.code"),
		DBType:     convertStrToDB(viper.GetString("db.type")),
		DBUsername: viper.GetString("db.username"),
		DBPassword: viper.GetString("db.password"),
		DBHostname: viper.GetString("db.hostname"),
		DBName:     viper.GetString("db.name"),
		DBPort:     viper.GetInt("db.port"),
		Writers:    viper.GetString("logging.writers"),
		RollingLogWriter: RollingLogWriter{
			Filename:   viper.GetString("logging.filename"),
			MaxSize:    viper.GetInt("logging.max.size"),
			MaxAge:     viper.GetInt("logging.max.age"),
			MaxBackups: viper.GetInt("logging.max.backups"),
			LocalTime:  viper.GetBool("logging.local.time"),
			Compress:   viper.GetBool("logging.compress"),
		},
		GraylogWriter: GraylogWriter{
			Host:     viper.GetString("logging.graylog_host"),
			Port:     viper.GetString("logging.graylog_port"),
			Facility: viper.GetString("logging.graylog_facility"),
		},
		SecurityKey: viper.GetString("security.key"),
	}

	return c, nil
}

// convertStrToDB converts string to DBType. It returns postgres if there isn't any match
func convertStrToDB(s string) DBType {
	if strings.ToLower(s) == "mysql" {
		return DBMySql
	} else if strings.ToLower(s) == "sqlite" {
		return DBSqlite
	}
	return DBPostgres
}
