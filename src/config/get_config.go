package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/goccy/go-json"
	"github.com/imdario/mergo"
	"os"
)

func GetConfig() Config {
	cfg := Config{}
	configFileBytes, readFileErr := os.ReadFile("./environment/config.json")
	if readFileErr != nil {
		panic(readFileErr)
	}

	var config Config
	unmarshallErr := json.Unmarshal(configFileBytes, &config)
	if unmarshallErr != nil {
		panic(unmarshallErr)
	}

	if mergeErr := mergo.Merge(&cfg, &config); mergeErr != nil {
		panic(mergeErr)
	}

	opts := env.Options{UseFieldNameByDefault: true}
	if parseEnvErr := env.ParseWithOptions(&cfg, opts); parseEnvErr != nil {
		panic(parseEnvErr)
	}

	ValidateConfig()
	return cfg
}
