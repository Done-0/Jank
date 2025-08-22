// Package configs 提供应用程序配置加载和更新功能
// 创建者：Done-0
// 创建时间：2025-08-05
package configs

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// AppConfig 应用配置
type AppConfig struct {
	AppName    string      `mapstructure:"APP_NAME"` // 应用名称
	AppHost    string      `mapstructure:"APP_HOST"` // 应用主机
	AppPort    string      `mapstructure:"APP_PORT"` // 应用端口
	CORSConfig CORSConfig  `mapstructure:"CORS"`     // CORS 跨域配置
	Email      EmailConfig `mapstructure:"EMAIL"`    // 邮箱配置
	JWT        JWTConfig   `mapstructure:"JWT"`      // JWT 认证配置
	User       UserConfig  `mapstructure:"USER"`     // 用户相关配置
}

// EmailConfig 邮箱配置
type EmailConfig struct {
	EmailType string `mapstructure:"EMAIL_TYPE"` // 邮箱类型
	FromEmail string `mapstructure:"FROM_EMAIL"` // 发件人邮箱
	EmailSmtp string `mapstructure:"EMAIL_SMTP"` // 邮件SMTP服务器
}

// JWTConfig JWT 认证配置
type JWTConfig struct {
	Secret        string `mapstructure:"SECRET"`         // JWT 签名密钥
	ExpireTime    int64  `mapstructure:"EXPIRE_TIME"`    // Token 有效期（小时）
	RefreshExpire int64  `mapstructure:"REFRESH_EXPIRE"` // 刷新 Token 有效期（小时）
}

// UserConfig 用户相关配置 - 主流Casbin RBAC配置
type UserConfig struct {
	// 管理员账户配置
	AdminEmail    string `mapstructure:"ADMIN_EMAIL"`    // 管理员邮箱
	AdminPassword string `mapstructure:"ADMIN_PASSWORD"` // 管理员密码
	AdminNickname string `mapstructure:"ADMIN_NICKNAME"` // 管理员昵称
	AdminRole     string `mapstructure:"ADMIN_ROLE"`     // 管理员角色
	DefaultRole   string `mapstructure:"DEFAULT_ROLE"`   // 默认用户角色
	// 注册控制
	AllowRegister bool `mapstructure:"ALLOW_REGISTER"` // 是否允许用户注册
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	DBDialect  string `mapstructure:"DB_DIALECT"` // 数据库类型
	DBName     string `mapstructure:"DB_NAME"`    // 数据库名称
	DBHost     string `mapstructure:"DB_HOST"`    // 数据库主机
	DBPort     string `mapstructure:"DB_PORT"`    // 数据库端口
	DBUser     string `mapstructure:"DB_USER"`    // 数据库用户
	DBPassword string `mapstructure:"DB_PSW"`     // 数据库密码
	DBPath     string `mapstructure:"DB_PATH"`    // 数据库路径
}

// LogConfig 日志配置
type LogConfig struct {
	LogFilePath     string `mapstructure:"LOG_FILE_PATH"`     // 日志文件路径
	LogFileName     string `mapstructure:"LOG_FILE_NAME"`     // 日志文件名
	LogTimestampFmt string `mapstructure:"LOG_TIMESTAMP_FMT"` // 日志时间戳格式
	LogMaxAge       int64  `mapstructure:"LOG_MAX_AGE"`       // 日志保留天数
	LogRotationTime int64  `mapstructure:"LOG_ROTATION_TIME"` // 日志轮转时间（小时）
	LogLevel        string `mapstructure:"LOG_LEVEL"`         // 日志级别
}

// RedisConfig Redis 配置
type RedisConfig struct {
	RedisHost     string `mapstructure:"REDIS_HOST"` // Redis 服务器地址
	RedisPort     string `mapstructure:"REDIS_PORT"` // Redis 服务器端口
	RedisPassword string `mapstructure:"REDIS_PSW"`  // Redis 密码
	RedisDB       string `mapstructure:"REDIS_DB"`   // Redis 数据库索引
}

// CasbinConfig Casbin 权限配置
type CasbinConfig struct {
	ModelPath  string `mapstructure:"MODEL_PATH"`  // 模型文件路径
	PolicyPath string `mapstructure:"POLICY_PATH"` // 策略文件路径
	DBAdapter  bool   `mapstructure:"DB_ADAPTER"`  // 是否使用数据库适配器
}

// CORSConfig CORS 跨域配置
type CORSConfig struct {
	AllowOrigins     []string `mapstructure:"ALLOW_ORIGINS"`     // 允许的源
	AllowCredentials bool     `mapstructure:"ALLOW_CREDENTIALS"` // 是否允许携带凭证
	MaxAge           int64    `mapstructure:"MAX_AGE"`           // 预检请求缓存时间（小时）
}

// PluginConfig 插件配置
type PluginConfig struct {
	// 插件目录和文件
	PluginDir        string `mapstructure:"PLUGIN_DIR"`         // 插件目录
	PluginConfigFile string `mapstructure:"PLUGIN_CONFIG_FILE"` // 插件配置文件
	PluginBinDir     string `mapstructure:"PLUGIN_BIN_DIR"`     // 插件二进制文件目录
	PluginMainFile   string `mapstructure:"PLUGIN_MAIN_FILE"`   // 插件主文件名

	// 构建相关
	BuildScriptDir      string `mapstructure:"BUILD_SCRIPT_DIR"`      // 构建脚本目录
	BuildScriptFile     string `mapstructure:"BUILD_SCRIPT_FILE"`     // 构建脚本文件名
	BuildTimeoutMinutes int    `mapstructure:"BUILD_TIMEOUT_MINUTES"` // 构建超时时间（分钟）
}

// ThemeConfig 主题配置
type ThemeConfig struct {
	// 主题目录和文件
	ThemeDir        string `mapstructure:"THEME_DIR"`         // 主题目录
	ThemeConfigFile string `mapstructure:"THEME_CONFIG_FILE"` // 主题配置文件名

	// Frontend 主题配置（用户前端）
	FrontendDefaultTheme    string `mapstructure:"FRONTEND_DEFAULT_THEME"`     // Frontend 默认主题 ID
	FrontendLastActiveTheme string `mapstructure:"FRONTEND_LAST_ACTIVE_THEME"` // Frontend 上次激活的主题 ID

	// Console 主题配置（管理后台）
	ConsoleDefaultTheme    string `mapstructure:"CONSOLE_DEFAULT_THEME"`     // Console 默认主题 ID
	ConsoleLastActiveTheme string `mapstructure:"CONSOLE_LAST_ACTIVE_THEME"` // Console 上次激活的主题 ID

	// 构建相关
	BuildScriptDir      string `mapstructure:"BUILD_SCRIPT_DIR"`      // 构建脚本目录
	BuildScriptFile     string `mapstructure:"BUILD_SCRIPT_FILE"`     // 构建脚本文件名
	BuildTimeoutMinutes int    `mapstructure:"BUILD_TIMEOUT_MINUTES"` // 构建超时时间（分钟）
}

// Config 总配置结构
type Config struct {
	AppConfig    AppConfig      `mapstructure:"APP"`      // 应用配置
	DBConfig     DatabaseConfig `mapstructure:"DATABASE"` // 数据库配置
	LogConfig    LogConfig      `mapstructure:"LOG"`      // 日志配置
	RedisConfig  RedisConfig    `mapstructure:"REDIS"`    // Redis 配置
	CasbinConfig CasbinConfig   `mapstructure:"CASBIN"`   // Casbin 权限配置
	PluginConfig PluginConfig   `mapstructure:"PLUGIN"`   // 插件配置
	ThemeConfig  ThemeConfig    `mapstructure:"THEME"`    // 主题配置
}

// DefaultConfigPath 默认配置文件路径
const DefaultConfigPath = "./configs/configs.yaml"

var (
	configInstance  *Config      // 全局配置实例
	configMutex     sync.RWMutex // 配置读写锁
	viperController *viper.Viper // viper 实例
)

// New 初始化配置
// 参数：
//
//	configPath: 配置文件路径
//
// 返回值：
//
//	error: 初始化过程中的错误
func New(configPath string) error {
	viperController = viper.New()
	viperController.SetConfigFile(configPath)

	if err := viperController.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := viperController.Unmarshal(&config); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	configInstance = &config
	go monitorConfigChanges()
	return nil
}

// GetConfig 获取配置
// 返回值：
//
//	*Config: 配置副本
//	error: 获取过程中的错误
func GetConfig() (*Config, error) {
	configMutex.RLock()
	defer configMutex.RUnlock()

	if configInstance == nil {
		return nil, fmt.Errorf("config not initialized")
	}

	configCopy := *configInstance
	return &configCopy, nil
}

// monitorConfigChanges 监听配置变更
func monitorConfigChanges() {
	viperController.WatchConfig()
	viperController.OnConfigChange(func(e fsnotify.Event) {
		var newConfig Config
		if err := viperController.Unmarshal(&newConfig); err != nil {
			log.Printf("failed to unmarshal new config: %v", err)
			return
		}

		configMutex.Lock()
		defer configMutex.Unlock()

		oldConfig := *configInstance
		changes := make(map[string][2]any)

		if !compareStructs(oldConfig, newConfig, "", changes) {
			log.Printf("config type mismatch, changes blocked")
			return
		}

		configInstance = &newConfig

		for path, values := range changes {
			log.Printf("config item [%s] changed: %v -> %v", path, values[0], values[1])
		}
	})
}

// compareStructs 比较结构体并收集变更
// 参数：
//
//	oldObj: 旧结构体
//	newObj: 新结构体
//	prefix: 字段路径前缀
//	changes: 记录变更的映射
//
// 返回值：
//
//	bool: 结构体类型是否一致
func compareStructs(oldObj, newObj any, prefix string, changes map[string][2]any) bool {
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
			changes[fullName] = [2]any{oldField.Interface(), newField.Interface()}
		}
	}

	return true
}

// UpdateField 更新配置字段
// 参数：
//
//	updateFunc: 更新函数
//
// 返回值：
//
//	error: 更新过程中的错误
func UpdateField(updateFunc func(*Config)) error {
	configMutex.Lock()
	defer configMutex.Unlock()

	oldConfig := *configInstance
	updateFunc(configInstance)

	configFile := viperController.ConfigFileUsed()
	content, err := os.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	newContent := string(content)

	var updateContent func(reflect.Value, reflect.Value, reflect.Type)
	updateContent = func(oldVal, newVal reflect.Value, t reflect.Type) {
		for i := 0; i < oldVal.NumField(); i++ {
			oldField, newField := oldVal.Field(i), newVal.Field(i)
			if tag := t.Field(i).Tag.Get("mapstructure"); tag != "" {
				if oldField.Kind() == reflect.Struct {
					updateContent(oldField, newField, oldField.Type())
				} else if !reflect.DeepEqual(oldField.Interface(), newField.Interface()) {
					var old, new string
					if oldField.Kind() == reflect.Slice || oldField.Kind() == reflect.Array {
						// 数组类型
						var oldElems, newElems []string
						for i := 0; i < oldField.Len(); i++ {
							elem := oldField.Index(i)
							if elem.Kind() == reflect.String {
								oldElems = append(oldElems, fmt.Sprintf(`"%s"`, elem.String()))
							} else {
								oldElems = append(oldElems, fmt.Sprintf("%v", elem.Interface()))
							}
						}
						for i := 0; i < newField.Len(); i++ {
							elem := newField.Index(i)
							if elem.Kind() == reflect.String {
								newElems = append(newElems, fmt.Sprintf(`"%s"`, elem.String()))
							} else {
								newElems = append(newElems, fmt.Sprintf("%v", elem.Interface()))
							}
						}
						old, new = fmt.Sprintf("[%s]", strings.Join(oldElems, ", ")), fmt.Sprintf("[%s]", strings.Join(newElems, ", "))

						for _, pattern := range []string{fmt.Sprintf(`%s: %s`, tag, old), fmt.Sprintf(`%s: []`, tag)} {
							if strings.Contains(newContent, pattern) {
								newContent = strings.ReplaceAll(newContent, pattern, fmt.Sprintf(`%s: %s`, tag, new))
								break
							}
						}
					} else {
						// 非数组类型
						old, new = fmt.Sprintf("%v", oldField.Interface()), fmt.Sprintf("%v", newField.Interface())
						for _, pattern := range []string{
							fmt.Sprintf(`%s: "%s"`, tag, old),
							fmt.Sprintf(`%s: %s`, tag, old),
							fmt.Sprintf(`%s: ""`, tag),
						} {
							if strings.Contains(newContent, pattern) {
								newContent = strings.ReplaceAll(newContent, pattern, fmt.Sprintf(`%s: "%s"`, tag, new))
								break
							}
						}
					}
				}
			}
		}
	}

	updateContent(reflect.ValueOf(oldConfig), reflect.ValueOf(*configInstance), reflect.TypeOf(oldConfig))

	if newContent != string(content) {
		return os.WriteFile(configFile, []byte(newContent), 0644)
	}

	return nil
}
