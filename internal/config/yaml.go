// Package config пакет для работы с конфигурацией
package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/SKharchenko87/foodix/pkg/config"
	"gopkg.in/yaml.v3"
)

var errInvalidPath = errors.New("invalid config path")
var allowedExtensions = map[string]struct{}{".yaml": {}, ".yml": {}}

// YAMLConfig основная структура конфигурации
type YAMLConfig struct {
	Server config.Server `yaml:"server"`
	Repo   config.Repo   `yaml:"repo"`
	Logger config.Logger `yaml:"logger"`
}

// NewYAMLConfig возвращает новый объект YAMLConfig
func NewYAMLConfig() *YAMLConfig {
	res := &YAMLConfig{}
	return res
}

// Load функция загрузки конфигурации из указанного пути
func (yc *YAMLConfig) Load(path string) error {

	if path == "" {
		return fmt.Errorf("%w: path cannot be empty", errInvalidPath)
	}

	// Проверяем, что путь абсолютный или относительный, но не содержит ../
	cleanPath := filepath.Clean(path)
	if strings.Contains(cleanPath, "..") {
		return fmt.Errorf("%w: path traversal not allowed", errInvalidPath)
	}

	// Проверяем расширение файла
	if _, ok := allowedExtensions[filepath.Ext(cleanPath)]; !ok {
		return fmt.Errorf("%w: path extension not allowed", errInvalidPath)
	}

	file, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	err = yaml.Unmarshal(file, yc)
	if err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}

	return nil
}

// GetServer возвращает конфигурацию сервера
func (yc *YAMLConfig) GetServer() config.Server {
	return yc.Server
}

// GetRepo возвращает конфигурацию репозитория
func (yc *YAMLConfig) GetRepo() config.Repo {
	return yc.Repo
}

// GetLogger возвращает конфигурацию logger
func (yc *YAMLConfig) GetLogger() config.Logger {
	return yc.Logger
}
