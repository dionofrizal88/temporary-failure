package strategy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnectToRedisDB(t *testing.T) {
	t.Run("if can connect into redis", func(t *testing.T) {
		db := RdbConf{
			Host:     "localhost",
			Port:     "6379",
			Username: "",
			Password: "",
		}

		rdb, err := db.ConnectToRedisDB()

		assert.NoError(t, err)
		assert.NotNil(t, rdb)
	})

	t.Run("if can't connect into redis", func(t *testing.T) {
		db := RdbConf{
			Host:     "localhost",
			Port:     "1234",
			Username: "user",
			Password: "",
		}

		rdb, err := db.ConnectToRedisDB()

		assert.Error(t, err)
		assert.Nil(t, rdb)
	})
}
