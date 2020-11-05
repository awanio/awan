// Part of this code credit to https://github.com/kataras
// This code modify from original code https://github.com/iris-contrib/examples/blob/master/database/mongodb/env/env.go

package env

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/awanio/awan/pkg/helper"
	"github.com/joho/godotenv"
)

var (
	// DBType the option is: mysql, postgres, mssql and sqlite3
	DBType string
	// DBHost var
	DBHost string
	// DBName var
	DBName string
	// DBUser var
	DBUser string
	// DBPassword var
	DBPassword string
	// DBSSLMode var
	DBSSLMode string
	// DBPath var
	DBPath string
	// DBMaxOpenConns var
	DBMaxOpenConns int
	// DBMaxIdleConns var
	DBMaxIdleConns int
	// JWTSecret var
	JWTSecret string
)

func parse() {
	DBType = getDefault("DB_TYPE", "sqlite3")
	DBPath = helper.FromBasepath(getDefault("DB_PATH", "awan.db"))
	// dbHost format should: host:port
	DBHost = getDefault("DB_HOST", "127.0.0.1:5432")
	DBName = getDefault("DB_NAME", "awan")
	DBUser = getDefault("DB_USER", "awan")
	DBPassword = getDefault("DB_PASSWORD", "")

	dbMaxOpenConns := getDefault("DB_MAX_OPEN_CONNS", "30")
	maxOpen, _ := strconv.Atoi(dbMaxOpenConns)
	DBMaxOpenConns = maxOpen

	dbMaxIdleConns := getDefault("DB_MAX_IDLE_CONNS", "30")
	maxIdle, _ := strconv.Atoi(dbMaxIdleConns)
	DBMaxIdleConns = maxIdle

	JWTSecret = getDefault("JWT_SECRET", "")

}

// Load loads environment variables that are being used across the whole app.
// Loading from file(s), i.e .env or dev.env
//
// Example of a 'dev.env':
// PORT=8080
// DSN=mongodb://localhost:27017
//
// After `Load` the callers can get an environment variable via `os.Getenv`.
func Load(envFileName string) {

	// If more than one filename passed with comma separated then load from all
	// of these, a env file can be a partial too.
	envFiles := strings.Split(envFileName, ",")
	for _, envFile := range envFiles {

		fmt.Println(envFile)

		if filepath.Ext(envFile) == "" {
			envFile += ".env"
		}

		if fileExists(envFile) {
			log.Printf("Loading environment variables from file: %s\n", envFile)

			if err := godotenv.Load(envFile); err != nil {
				panic(fmt.Sprintf("error loading environment variables from [%s]: %v", envFile, err))
			}
		} else {
			fmt.Println(fmt.Sprintf("file [%s] not exist", envFile))
		}
	}

	// envMap, _ := godotenv.Read(envFiles...)
	// for k, v := range envMap {
	// 	log.Printf("◽ %s=%s\n", k, v)
	// }

	parse()
}

func getDefault(key string, def string) string {
	value := os.Getenv(key)
	if value == "" {
		os.Setenv(key, def)
		value = def
	}

	return value
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
