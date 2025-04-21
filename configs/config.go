package configs

import (
	"fmt"
	"log"
	"reflect"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// AppConfig 应用配置
type AppConfig struct {
	AppName   string `mapstructure:"APP_NAME"`
	AppHost   string `mapstructure:"APP_HOST"`
	AppPort   string `mapstructure:"APP_PORT"`
	EmailType string `mapstructure:"EMAIL_TYPE"`
	FromEmail string `mapstructure:"FROM_EMAIL"`
	EmailSmtp string `mapstructure:"EMAIL_SMTP"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	DBDialect  string `mapstructure:"DB_DIALECT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PSW"`
	DBPath     string `mapstructure:"DB_PATH"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisDB       string `mapstructure:"REDIS_DB"`
	RedisPassword string `mapstructure:"REDIS_PSW"`
}

// LogConfig 日志配置
type LogConfig struct {
	LogFilePath     string `mapstructure:"LOG_FILE_PATH"`
	LogFileName     string `mapstructure:"LOG_FILE_NAME"`
	LogTimestampFmt string `mapstructure:"LOG_TIMESTAMP_FMT"`
	LogMaxAge       int64  `mapstructure:"LOG_MAX_AGE"`
	LogRotationTime int64  `mapstructure:"LOG_ROTATION_TIME"`
	LogLevel        string `mapstructure:"LOG_LEVEL"`
}

// SwaggerConfig Swagger配置
type SwaggerConfig struct {
	SwaggerHost    string `mapstructure:"SWAGGER_HOST"`
	SwaggerEnabled string `mapstructure:"SWAGGER_ENABLED"`
}

// Config 总配置结构
type Config struct {
	AppConfig     AppConfig      `mapstructure:"app"`
	DBConfig      DatabaseConfig `mapstructure:"database"`
	RedisConfig   RedisConfig    `mapstructure:"redis"`
	LogConfig     LogConfig      `mapstructure:"log"`
	SwaggerConfig SwaggerConfig  `mapstructure:"swagger"`
}

// DefaultConfigPath 默认配置文件路径
const DefaultConfigPath = "./configs/config.yml"

var (
	globalConfig  *Config      // 全局配置实例
	configLock    sync.RWMutex // 配置读写锁
	viperInstance *viper.Viper // viper实例
)

// Init 初始化配置
func Init(configPath string) error {
	viperInstance = viper.New()
	viperInstance.SetConfigFile(configPath)

	if err := viperInstance.ReadInConfig(); err != nil {
		return fmt.Errorf("配置文件读取失败: %w", err)
	}

	var config Config
	if err := viperInstance.Unmarshal(&config); err != nil {
		return fmt.Errorf("配置解析失败: %w", err)
	}

	globalConfig = &config
	go monitorConfigChanges()
	return nil
}

// LoadConfig 获取配置
func LoadConfig() (*Config, error) {
	configLock.RLock()
	defer configLock.RUnlock()

	if globalConfig == nil {
		return nil, fmt.Errorf("配置未初始化")
	}

	configCopy := *globalConfig
	return &configCopy, nil
}

// monitorConfigChanges 监听配置变更
func monitorConfigChanges() {
	viperInstance.WatchConfig()
	viperInstance.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("配置文件变更: %s", e.Name)

		var newConfig Config
		if err := viperInstance.Unmarshal(&newConfig); err != nil {
			log.Printf("新配置解析失败: %v", err)
			return
		}

		configLock.Lock()
		defer configLock.Unlock()

		oldConfig := *globalConfig
		changes := make(map[string][2]interface{})

		if !compareStructs(oldConfig, newConfig, "", changes) {
			log.Printf("配置类型不一致，变更被阻止")
			return
		}

		globalConfig = &newConfig

		for path, values := range changes {
			log.Printf("配置变更: %s 从 [%v] 变为 [%v]", path, values[0], values[1])
		}
	})
}

// compareStructs 比较结构体并收集变更
func compareStructs(oldObj, newObj interface{}, prefix string, changes map[string][2]interface{}) bool {
	oldVal := reflect.ValueOf(oldObj)
	newVal := reflect.ValueOf(newObj)

	if oldVal.Type() != newVal.Type() {
		return false
	}

	if oldVal.Kind() != reflect.Struct {
		return true
	}

	for i := 0; i < oldVal.NumField(); i++ {
		oldField := oldVal.Field(i)
		newField := newVal.Field(i)
		fieldName := oldVal.Type().Field(i).Name
		fullName := prefix + fieldName

		if oldField.Kind() == reflect.Struct {
			if !compareStructs(oldField.Interface(), newField.Interface(), fullName+".", changes) {
				return false
			}
			continue
		}

		if oldField.Kind() != newField.Kind() {
			return false
		}

		if !reflect.DeepEqual(oldField.Interface(), newField.Interface()) {
			changes[fullName] = [2]interface{}{oldField.Interface(), newField.Interface()}
		}
	}

	return true
}
