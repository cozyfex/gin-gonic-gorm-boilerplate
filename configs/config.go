package configs

type Config struct {
	Mode   string
	Server ServerConfig
	DB     struct {
		Master   DBConfig
		Replicas []DBConfig
	}
	Logging LoggingConfig
}

type ServerConfig struct {
	Port int
}

type DBConfig struct {
	Type     string
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
	CharSet  string
	SSLMode  string
	Timezone string
}

type LoggingConfig struct {
	Level  string
	Format string
}
