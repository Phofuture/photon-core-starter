package configuration

import (
	"context"
	"fmt"
	"log/slog"
	"os"
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

	//把所有環境變數塞進 viper
	for _, e := range os.Environ() {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) == 2 {
			key := strings.ToLower(parts[0]) // 轉小寫，避免 key 太醜
			viper.Set(key, parts[1])
		}
	}

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

	for _, key := range viper.AllKeys() {
		fmt.Printf("%s = %v\n", key, viper.Get(key))
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
