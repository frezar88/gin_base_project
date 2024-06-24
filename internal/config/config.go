package config

// Пример конфигурационной структуры
type Config struct {
	DBHost     string
	DBPort     int
	APIKey     string
	OtherParam string
}

// LoadConfig загружает конфигурацию из файла или переменных окружения
func LoadConfig() *Config {
	// Ваша логика загрузки конфигурации здесь
	return &Config{
		DBHost:     "localhost",
		DBPort:     27017,
		APIKey:     "myapikey",
		OtherParam: "value",
	}
}
