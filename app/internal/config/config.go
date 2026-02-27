package config

import "log/slog"

type Config struct {
	Db     *db
	App    *app
	Logger *slog.Logger
}

func CreateConfig() *Config {

	return &Config{
		Db:     createDb(),
		App:    createApp(),
		Logger: createLogger(),
	}
}
