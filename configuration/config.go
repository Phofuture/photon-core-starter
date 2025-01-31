package configuration

import (
	"context"
	"fmt"
	"log/slog"
	"reflect"
	"regexp"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
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
	readCommandLineArgs()

	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	for _, path := range pathArray {
		viper.AddConfigPath(path)
	}
	_ = viper.ReadInConfig()

	env, ok := viper.Get("env.name").(string)
	if ok {
		viper.SetConfigName("application-" + env)
		viper.SetConfigType("yaml")
		_ = viper.ReadInConfig()
	}
	for _, config := range configs {
		err := viper.Unmarshal(&config)
		if err != nil {
			slog.Error("解析 config 錯誤: ", "config", config, "error", err)
		}
	}
}

func readCommandLineArgs() {
	pflag.Parse()
	args := pflag.Args()

	regex := regexp.MustCompile(`([a-zA-Z][a-zA-Z0-9_-]*)=(.*)$`)
	for _, arg := range args {
		fmt.Println("arg", arg)
		if matches := regex.FindStringSubmatch(arg); matches != nil {
			viper.Set(matches[1], matches[2])
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
