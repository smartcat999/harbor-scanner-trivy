package opencve

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBWithOptions(options Options) (*gorm.DB, error) {
	con, err := gorm.Open(postgres.New(*options.PgConfig), options.GoOrmConfig)
	if err != nil {
		return nil, err
	}
	return con, err
}
