// Package config пакет для настройки конфигурации
package config

// Server структура описывающая конфигурацию сервера
type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// Repo описание конфигурации репозитория источника данных
type Repo struct {
	Name string `yaml:"name"`
	DSN  string `yaml:"dsn"`
}

// Logger конфигурация логирования
type Logger struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

// Config интерфейс для работы с конфигурациями
type Config interface {
	Load(path string) error
	GetServer() Server
	GetRepo() Repo
	GetLogger() Logger
}
