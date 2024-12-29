package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppName         string `mapstructure:"APP_NAME"`
	AppPort         string `mapstructure:"APP_PORT"`
	DBName          string `mapstructure:"DB_NAME"`
	DBHost          string `mapstructure:"DB_HOST"`
	DBPort          string `mapstructure:"DB_PORT"`
	DBUser          string `mapstructure:"DB_USER"`
	DBPassword      string `mapstructure:"DB_PSW"`
	RedisHost       string `mapstructure:"REDIS_HOST"`
	RedisDB         string `mapstructure:"REDIS_DB"`
	RedisPassword   string `mapstructure:"REDIS_PSW"`
	LogFilePath     string `mapstructure:"LOG_FILE_PATH"`
	LogFileName     string `mapstructure:"LOG_FILE_NAME"`
	LogTimestampFmt string `mapstructure:"LOG_TIMESTAMP_FMT"`
	LogMaxAge       int    `mapstructure:"LOG_MAX_AGE"`
	LogRotationTime int    `mapstructure:"LOG_ROTATION_TIME"`
	LogLevel        string `mapstructure:"LOG_LEVEL"`
	SwaggerHost     string `mapstructure:"SWAGGER_HOST"`
	QqSmtp          string `mapstructure:"QQ_SMTP"`
	FromEmail       string `mapstructure:"FROM_EMAIL"`
	Salt            string `mapstructure:"SALT"`
}

// LoadConfig 加载配置文件
func LoadConfig() (*Config, error) {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")

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
