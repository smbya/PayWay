package config

type Config struct {
	Db  *db
	App *app
}

func CreateConfig() *Config {
	return &Config{
		Db:  createDb(),
		App: createApp(),
	}
}
