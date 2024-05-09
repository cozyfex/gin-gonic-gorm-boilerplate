package parser

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/spf13/viper"

	"gin-gonic-gorm-boilerplate/configs"
	"gin-gonic-gorm-boilerplate/internal/util/logger"
)

func Replicas() *[]configs.DBConfig {
	var replicas []configs.DBConfig
	prefix := "DB_REPLICAS_"
	suffix := []string{"_TYPE", "_HOST", "_PORT", "_DBNAME", "_USER", "_PASSWORD", "_CHARSET", "_SSLMODE", "_TIMEZONE"}

	for _, env := range os.Environ() {
		if strings.HasPrefix(env, prefix) {
			logger.Info(env)
			pair := strings.SplitN(env, "=", 2)
			key := pair[0]

			if strings.HasSuffix(key, suffix[0]) {
				db := configs.DBConfig{}

				i := strings.TrimPrefix(key, prefix)
				i = strings.TrimSuffix(i, suffix[0])

				for _, s := range suffix {

					envKey := fmt.Sprintf("%s%s%s", prefix, i, s)
					envValue := os.Getenv(envKey)

					switch strings.Replace(s, "_", "", 1) {
					case "TYPE":
						db.Type = envValue
					case "HOST":
						db.Host = envValue
					case "PORT":
						port, err := strconv.Atoi(envValue)
						if err != nil {
							logger.Error(err)
							continue
						}
						db.Port = port
					case "DBNAME":
						db.DBName = envValue
					case "USER":
						db.User = envValue
					case "PASSWORD":
						db.Password = envValue
					case "CHARSET":
						db.CharSet = envValue
					case "SSLMODE":
						db.Host = envValue
					case "TIMEZONE":
						db.Timezone = envValue
					}
				}

				replicas = append(replicas, db)
			}
		}
	}

	return &replicas
}

func BindEnvs(envPrefix string, configs interface{}, configPrefix string) {
	t := reflect.TypeOf(configs)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Type.Kind() == reflect.Struct {
			prefix := fmt.Sprintf("%s%s.", configPrefix, field.Tag.Get("mapstructure"))
			BindEnvs(envPrefix, reflect.New(field.Type).Interface(), prefix)
		} else {
			prefix := func(p string) string {
				if p == "" {
					return ""
				} else {
					return p + "_"
				}
			}(envPrefix)
			key := configPrefix + field.Tag.Get("mapstructure")
			env := prefix + strings.Replace(strings.ToUpper(key), ".", "_", -1)
			err := viper.BindEnv(key, env)
			if err != nil {
				logger.Error("env mapping error")
			}
		}
	}
}
