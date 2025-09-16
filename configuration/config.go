package configuration

import (
	"context"
	"github.com/spf13/viper"
	"log/slog"
	"reflect"
)

var (
	configs   = make([]interface{}, 0)
	pathArray = []string{"./src/resources", "./src", "."}
)

func Register(config ...interface{}) {
	for _, c := range config {
		if reflect.TypeOf(c).Kind() != reflect.Ptr {
			panic("config must be pointer")
		}
	}
	configs = append(configs, config...)
}

func InitConfiguration() {

	viper.AutomaticEnv()

	for _, path := range pathArray {
		viper.SetConfigFile(path + "/application.yaml")
		viper.SetConfigType("yaml")
		_ = viper.MergeInConfig()
	}

	env, ok := viper.Get("env.name").(string)
	if ok {
		for _, path := range pathArray {
			viper.SetConfigFile(path + "/application-" + env + ".yaml")
			viper.SetConfigType("yaml")
			_ = viper.MergeInConfig()
		}
	}

	for _, config := range configs {
		err := viper.Unmarshal(&config)
		if err != nil {
			slog.Error("解析 config 錯誤: ", "config", config, "error", err)
		}
	}
}

func Get[T any](ctx context.Context) (config T, err error) {
	err = viper.Unmarshal(&config)
	if err != nil {
		slog.ErrorContext(ctx, "解析 config 錯誤", "error", err)
	}
	return
}
