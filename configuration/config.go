package configuration

import (
	"context"
	"log/slog"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

var (
	configs   = make([]interface{}, 0)
	pathArray = []string{"./src/resources", "./src", "."}
	postfixes = []string{".yaml", ".yml"}
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

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	for _, path := range pathArray {
		for _, postfix := range postfixes {
			viper.SetConfigFile(path + "/config" + postfix)
			viper.SetConfigType("yaml")
			_ = viper.MergeInConfig()
		}
	}

	for _, path := range pathArray {
		for _, postfix := range postfixes {
			viper.SetConfigFile(path + "/application" + postfix)
			viper.SetConfigType("yaml")
			_ = viper.MergeInConfig()
		}
	}

	env, ok := viper.Get("env.name").(string)
	if ok {
		for _, path := range pathArray {
			for _, postfix := range postfixes {
				viper.SetConfigFile(path + "/config-" + env + postfix)
				viper.SetConfigType("yaml")
				_ = viper.MergeInConfig()
			}
		}
		for _, path := range pathArray {
			for _, postfix := range postfixes {
				viper.SetConfigFile(path + "/application-" + env + postfix)
				viper.SetConfigType("yaml")
				_ = viper.MergeInConfig()
			}
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
