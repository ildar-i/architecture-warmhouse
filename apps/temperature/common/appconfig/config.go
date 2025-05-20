package appconfig

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	defaultPort = "8080"
)

// PostgresConfig конфигурация БД
type PostgresConfig struct {
	Host     string `mapstructure:"warmhouse_pg_host"`
	DBName   string `mapstructure:"warmhouse_pg_db"`
	UserName string `mapstructure:"warmhouse_pg_username"`
	Password string `mapstructure:"warmhouse_pg_password"`
	Port     string `mapstructure:"warmhouse_pg_port"`
}

// Param храним все необходимые данные для конфигурации сервиса
type Param struct {
	Pg PostgresConfig
}

type config struct {
	name    string
	version string
	port    string
	data    Param
}

type Config interface {
	// GetName получить имя приложения
	GetName() string
	// GetVersion получить версию приложения
	GetVersion() string
	// GetPort получить порт приложения
	GetPort() string
	// LoadConfig загрузка конфигурации из файла
	LoadConfig() error
	// GetPostgresConfig получить конфигурацию для ПГ
	GetPostgresConfig() PostgresConfig
}

// NewConfig получить новую конфигурацию
func NewConfig(name, version, port string) Config {
	return &config{
		name:    name,
		version: version,
		port:    port,
		data:    Param{},
	}
}

// LoadConfig загрузка env и файлов конфигураций сервиса
func (c *config) LoadConfig() error {
	//viper.AutomaticEnv()
	viper.SetConfigFile("./.env")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	var pg = c.data.Pg
	err := viper.Unmarshal(&pg)
	if err != nil {
		fmt.Printf("Unable to decode into map, %v", err)
	}

	fmt.Println("pg:")
	fmt.Println(pg)

	c.data.Pg = pg

	return nil
}

// GetPostgresConfig получить конфигурацию для ПГ
func (c *config) GetPostgresConfig() PostgresConfig {
	return c.data.Pg
}

// GetName получить имя приложения
func (c *config) GetName() string {
	return c.name
}

// GetVersion получить версию приложения
func (c *config) GetVersion() string {
	return c.version
}

// GetPort получить порт приложения
func (c *config) GetPort() string {
	if c.port == "" {
		c.port = defaultPort
	}

	return c.port
}
