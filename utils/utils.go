package utils

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"os"
)

var (
	APP_Address string
	DB_Driver, DB_Connect string
)

func InitConf(confPath string) {
	if _, err := os.Stat(confPath); !os.IsNotExist(err) {
		if config, err := toml.LoadFile(confPath); err == nil {
			APP_Address = config.Get("app.address").(string)

			DB_Driver = config.Get("db.driver").(string)
			DB_Connect = config.Get("db.connect").(string)
		}
	}
	if driver := os.Getenv("DB_DRIVER"); driver != "" {
		DB_Driver = driver
	}
	if connect := os.Getenv("DB_CONNECT"); connect != "" {
		DB_Connect = connect
	}

	fmt.Println("config: ", map[string]interface{}{
		"address":      APP_Address,
		"db_driver":    DB_Driver,
		"db_connect":   DB_Connect,
	})
}
