package config

import (
	"time"

	"github.com/spf13/viper"
)

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	viper.SetDefault("REPOSITORIES_PATH", defaultReposPath)
	viper.SetDefault("FILE_NAME", defaultFileName)
	viper.SetDefault("REPOSITORY_PREFIX", defaultRepoPrefix)
	viper.SetDefault("TARGET_REPOSITORY", defaultRepoPrefix+defaultTargetRepo)
	viper.SetDefault("START_DATE", time.Now().UTC().Format(time.RFC3339))

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
