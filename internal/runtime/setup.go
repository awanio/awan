package runtime

import (
	"database/sql"

	"github.com/awanio/awan/configs"
	"github.com/awanio/awan/internal/db"
	"github.com/awanio/awan/internal/env"
	"github.com/awanio/awan/pkg/helper"
	"gorm.io/gorm"
)

var (
	// DB var
	DB *gorm.DB
	// SQLDB var
	SQLDB *sql.DB
	// DBerror var
	DBerror error
)

// Setup method
func Setup() {

	envFileName := helper.FromBasepath("configs/.env.example")
	env.Load(envFileName)

	configs.Database.Type = env.DBType
	configs.Database.Path = env.DBPath
	configs.Database.MaxOpenConns = env.DBMaxOpenConns
	configs.Database.MaxIdleConns = env.DBMaxIdleConns

	db, sql, err := db.Factory(configs.Database)

	DB = db
	SQLDB = sql
	DBerror = err
}
