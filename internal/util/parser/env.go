package parser

import (
	"fmt"
	"gin-gonic-gorm-boilerplate/configs"
	"gin-gonic-gorm-boilerplate/internal/util/logger"
	"os"
	"strconv"
	"strings"
)

func ReplicaParser() *[]configs.DBConfig {
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
