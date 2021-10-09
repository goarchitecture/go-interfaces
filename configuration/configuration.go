package configuration

import (
	"encoding/json"
	"fmt"
)

// Source интерфейс декларирует метод получения содержимого для конфига из любого источника
type Source interface {
	Contents() ([]byte, error)
}

type Configuration struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func Load(source Source) (*Configuration, error) {
	contents, err := source.Contents()
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения: %w", err)
	}

	var configuration Configuration
	if err := json.Unmarshal(contents, &configuration); err != nil {
		return nil, fmt.Errorf("ошибка парсинга: %w", err)
	}

	return &configuration, nil
}

// String отображает конфигурацию строкой (через JSON формат)
func (c *Configuration) String() string {
	jsonified, _ := json.Marshal(c)
	return string(jsonified)
}
