package opencve

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDB(t *testing.T) {
	opts := Options{
		User:        "opencve",
		Password:    "opencve",
		Host:        "172.31.19.22",
		Port:        5432,
		DBName:      "opencve",
		Ssl:         false,
		TimeZone:    "Asia/Shanghai",
		PgConfig:    nil,
		GoOrmConfig: nil,
	}
	_, err := NewDBWithOptions(opts)
	assert.Equal(t, err, nil)
}
