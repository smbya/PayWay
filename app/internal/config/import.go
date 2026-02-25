package config

import "os"

func getEnvWithDefault(key string, def string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		return def
	} else {
		return val
	}
}
