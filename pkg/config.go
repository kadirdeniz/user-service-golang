package pkg

var Configs Config

type Config struct {
	Application ApplicationConfig `yaml:"application"`
	Database    DatabaseConfig    `yaml:"database"`
	Redis       RedisConfig       `yaml:"redis"`
	JWT         JWTConfig         `yaml:"jwt"`
}

type ApplicationConfig struct {
	Port        string `yaml:"port"`
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Environment string `yaml:"environment"`
	LogLevel    string `yaml:"log_level"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type JWTConfig struct {
	Secret     string `yaml:"secret"`
	Issuer     string `yaml:"issuer"`
	Expiration int64  `yaml:"expiration"`
}
