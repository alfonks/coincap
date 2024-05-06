package cfg

import (
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func readJSONConfig(cfg *ConfigSchema, basePath, cfgLocation string, fileList []string) {
	funcName := "cfg.readJSONConfig"
	v := viper.New()
	v.AddConfigPath(fmt.Sprintf("%s/%s/", basePath, cfgLocation))
	for _, file := range fileList {
		v.SetConfigName(file)
		err := v.MergeInConfig()
		if err != nil {
			log.Fatalf("[%v] fail merge config for file %v, error: %v", funcName, file, err)
		}
	}

	err := v.Unmarshal(cfg, func(decoderConfig *mapstructure.DecoderConfig) {
		decoderConfig.TagName = "json"
	})
	if err != nil {
		log.Fatalf("[%v] fail unmarshal config, error: %v", funcName, err)
	}
}
