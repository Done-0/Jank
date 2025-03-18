package configs

import (
	"log"

	"github.com/spf13/viper"
)

// AppConfig 存储应用相关配置
type AppConfig struct {
	AppName   string `mapstructure:"APP_NAME"`
	AppHost   string `mapstructure:"APP_HOST"`
	AppPort   string `mapstructure:"APP_PORT"`
	EmailType string `mapstructure:"EMAIL_TYPE"`
	FromEmail string `mapstructure:"FROM_EMAIL"`
	EmailSmtp string `mapstructure:"EMAIL_SMTP"`
}

// DatabaseConfig 存储数据库相关配置
type DatabaseConfig struct {
	DBDialect  string `mapstructure:"DB_DIALECT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PSW"`
	DBPath     string `mapstructure:"DB_PATH"`
}

// RedisConfig 存储Redis相关配置
type RedisConfig struct {
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisDB       string `mapstructure:"REDIS_DB"`
	RedisPassword string `mapstructure:"REDIS_PSW"`
}

// LogConfig 存储日志相关配置
type LogConfig struct {
	LogFilePath     string `mapstructure:"LOG_FILE_PATH"`
	LogFileName     string `mapstructure:"LOG_FILE_NAME"`
	LogTimestampFmt string `mapstructure:"LOG_TIMESTAMP_FMT"`
	LogMaxAge       int    `mapstructure:"LOG_MAX_AGE"`
	LogRotationTime int    `mapstructure:"LOG_ROTATION_TIME"`
	LogLevel        string `mapstructure:"LOG_LEVEL"`
}

// SwaggerConfig 存储Swagger相关配置
type SwaggerConfig struct {
	SwaggerHost string `mapstructure:"SWAGGER_HOST"`
}

// Config 存储所有配置项
type Config struct {
	AppConfig     AppConfig      `mapstructure:"app"`
	DBConfig      DatabaseConfig `mapstructure:"database"`
	RedisConfig   RedisConfig    `mapstructure:"redis"`
	LogConfig     LogConfig      `mapstructure:"log"`
	SwaggerConfig SwaggerConfig  `mapstructure:"swagger"`
}

// LoadConfig 加载配置文件
func LoadConfig() (*Config, error) {
	viper.SetConfigFile("./configs/config.yml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("配置文件加载失败：%v", err)
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("配置解析失败：%v", err)
		return nil, err
	}

	return &config, nil
}
