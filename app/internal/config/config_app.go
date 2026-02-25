package config

type app struct {
	HandlerType string
	Port        string
}

func createApp() *app {
	return &app{
		HandlerType: getEnvWithDefault("APP_HANDLER_TYPE", "chi"),
		Port:        getEnvWithDefault("APP_PORT", "80"),
	}
}
