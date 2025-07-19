package config

import (
	"os"
	"reflect"
	"strconv"
)

func Default() *Config {
	config := new(Config)
	return config
}

func Load() *Config {
	config := Default()
	loadFromEnv(config)

	return config
}

func loadFromEnv(v interface{}) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return
	}
	rv = rv.Elem()
	rt := rv.Type()

	for i := 0; i < rv.NumField(); i++ {
		fieldVal := rv.Field(i)
		fieldType := rt.Field(i)

		// env タグを取得
		envTag := fieldType.Tag.Get("env")
		defaultTag := fieldType.Tag.Get("default")
		var value string
		if envTag != "" {
			value = os.Getenv(envTag)
		}
		// 環境変数がなければ default を使用
		if value == "" {
			value = defaultTag
		}

		if value != "" && fieldVal.CanSet() {
			switch fieldVal.Kind() {
			case reflect.String:
				fieldVal.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if iv, err := strconv.Atoi(value); err == nil {
					fieldVal.SetInt(int64(iv))
				}
			case reflect.Bool:
				if bv, err := strconv.ParseBool(value); err == nil {
					fieldVal.SetBool(bv)
				}
			}
			continue
		}

		// ネストした構造体も再帰的に処理
		if fieldVal.Kind() == reflect.Struct {
			loadFromEnv(fieldVal.Addr().Interface())
		}
	}
}
