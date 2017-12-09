package util

import "os"

func Config(url string) string {
	if url != "" {
		return url
	}
	return "http://127.0.0.1:9200"
}

func EnvOrElse(env string, defaultParam string) string {
	if os.Getenv(env) != "" {
		return os.Getenv(env)
	}
	return defaultParam
}
