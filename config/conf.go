package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

type config struct {
	defaultdb   string
	connections connections
}

type connections struct {
	mysql mysql
}

// Config ...
var Config config

// Dir path for build
var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

// InitConfig - Load .env, init var, custom validation tag
func InitConfig() {
	envLoader()
}

func envLoader() {
	// load .env file
	err := godotenv.Load(filepath.Dir(basepath) + "/.env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
}

// Env .. single env loader
func env(key string, defaultValue string) string {
	value := os.Getenv(key)

	if len(value) > 0 {
		return value
	}
	return defaultValue

}
