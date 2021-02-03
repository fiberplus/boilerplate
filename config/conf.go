package config

import (
	"fmt" 
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

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
