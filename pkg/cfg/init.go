package cfg

import (
	"log"
	"path/filepath"
	"sync"
)

var (
	configOnce sync.Once
	config     ConfigSchema
)

func Init() *ConfigSchema {
	configOnce.Do(func() {
		funcName := "cfg.Init"

		configListJSON := []string{
			"coincap-config",
			"coincap-config-secret",
		}
		path, err := filepath.Abs(".")
		if err != nil {
			log.Fatalf("[%v] cant find absolute path: %v", funcName, err)
		}
		configLocation := "config"
		readJSONConfig(&config, path, configLocation, configListJSON)
	})

	return &config
}
