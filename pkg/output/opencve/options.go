package opencve

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Options struct {
	User        string
	Password    string
	Host        string
	Port        int64
	DBName      string
	Ssl         bool
	TimeZone    string
	PgConfig    *postgres.Config
	GoOrmConfig *gorm.Config
}

func (o *Options) Load() {
	dsn := ""
	if o.Host != "" {
		dsn = fmt.Sprintf("host=%s", o.Host)
	}
	if o.User != "" {
		dsn = fmt.Sprintf("%s user=%s", dsn, o.User)
	}
	if o.Password != "" {
		dsn = fmt.Sprintf("%s password=%s", dsn, o.Password)
	}
	if o.DBName != "" {
		dsn = fmt.Sprintf("%s dbname=%s", dsn, o.DBName)
	}
	if o.Port > 0 {
		dsn = fmt.Sprintf("%s port=%d", dsn, o.Port)
	}
	if o.Ssl == false {
		dsn = fmt.Sprintf("%s sslmode=%s", dsn, "disable")
	}
	if o.TimeZone != "" {
		dsn = fmt.Sprintf("%s TimeZone=%s", dsn, o.TimeZone)
	}
	if o.PgConfig == nil {
		o.PgConfig = &postgres.Config{
			DSN: dsn,
		}
	} else {
		o.PgConfig.DSN = dsn
	}
	if o.GoOrmConfig == nil {
		o.GoOrmConfig = &gorm.Config{}
	}
}
