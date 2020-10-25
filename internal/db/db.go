// Part of this code credit to https://github.com/gogs/gogs
// This code modify from original code https://github.com/gogs/gogs/blob/main/internal/db/db.go
// Copyright 2020 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the https://github.com/gogs/gogs/blob/main/LICENSE link.

package db

import (
	"database/sql"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/awanio/awan/configs"
	"github.com/awanio/awan/internal/env"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// parsePostgreSQLHostPort parses given input in various forms defined in
// https://www.postgresql.org/docs/current/static/libpq-connect.html#LIBPQ-CONNSTRING
// and returns proper host and port number.
func parsePostgreSQLHostPort(info string) (host, port string) {
	host, port = "127.0.0.1", "5432"
	if strings.Contains(info, ":") && !strings.HasSuffix(info, "]") {
		idx := strings.LastIndex(info, ":")
		host = info[:idx]
		port = info[idx+1:]
	} else if len(info) > 0 {
		host = info
	}
	return host, port
}

func parseMSSQLHostPort(info string) (host, port string) {
	host, port = "127.0.0.1", "1433"
	if strings.Contains(info, ":") {
		host = strings.Split(info, ":")[0]
		port = strings.Split(info, ":")[1]
	} else if strings.Contains(info, ",") {
		host = strings.Split(info, ",")[0]
		port = strings.TrimSpace(strings.Split(info, ",")[1])
	} else if len(info) > 0 {
		host = info
	}
	return host, port
}

// parseDSN takes given database options and returns parsed DSN.
func parseDSN(opts configs.DatabaseOpts) (dsn string, err error) {
	// In case the database name contains "?" with some parameters
	concate := "?"
	if strings.Contains(opts.Name, concate) {
		concate = "&"
	}

	switch opts.Type {
	case "mysql":
		if opts.Host[0] == '/' { // Looks like a unix socket
			dsn = fmt.Sprintf("%s:%s@unix(%s)/%s%scharset=utf8mb4&parseTime=true",
				opts.User, opts.Password, opts.Host, opts.Name, concate)
		} else {
			dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s%scharset=utf8mb4&parseTime=true",
				opts.User, opts.Password, opts.Host, opts.Name, concate)
		}

	case "postgres":
		host, port := parsePostgreSQLHostPort(opts.Host)
		if host[0] == '/' { // looks like a unix socket
			dsn = fmt.Sprintf("postgres://%s:%s@:%s/%s%ssslmode=%s&host=%s",
				url.QueryEscape(opts.User), url.QueryEscape(opts.Password), port, opts.Name, concate, opts.SSLMode, host)
		} else {
			dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s%ssslmode=%s",
				url.QueryEscape(opts.User), url.QueryEscape(opts.Password), host, port, opts.Name, concate, opts.SSLMode)
		}

	case "mssql":
		host, port := parseMSSQLHostPort(opts.Host)
		dsn = fmt.Sprintf("server=%s; port=%s; database=%s; user id=%s; password=%s;",
			host, port, opts.Name, opts.User, opts.Password)

	case "sqlite3":
		dsn = "file:" + opts.Path + "?cache=shared&mode=rwc"

	default:
		return "", errors.Errorf("unrecognized dialect: %s", opts.Type)
	}

	return dsn, nil
}

func openDB(opts configs.DatabaseOpts, cfg *gorm.Config) (*gorm.DB, error) {
	dsn, err := parseDSN(opts)
	if err != nil {
		return nil, errors.Wrap(err, "parse DSN")
	}

	var dialector gorm.Dialector
	switch opts.Type {
	case "mysql":
		dialector = mysql.Open(dsn)
	case "postgres":
		dialector = postgres.Open(dsn)
	case "mssql":
		dialector = sqlserver.Open(dsn)
	case "sqlite3":
		dialector = sqlite.Open(dsn)
	default:
		panic("unreachable")
	}

	return gorm.Open(dialector, cfg)
}

// Tables is the list of struct-to-table mappings.
//
// NOTE: Lines are sorted in alphabetical order, each letter in its own line.
// var Tables = []interface{}{
// 	new(Access), new(AccessToken),
// 	new(LFSObject), new(LoginSource),
// }

// Run method
func Run() (*gorm.DB, *sql.DB, error) {

	// My config option
	var database configs.DatabaseOpts
	database.Type = env.DBType
	database.Path = env.DBPath
	database.MaxOpenConns = env.DBMaxOpenConns
	database.MaxIdleConns = env.DBMaxIdleConns

	db, err := openDB(database, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		NowFunc: func() time.Time {
			return time.Now().UTC().Truncate(time.Microsecond)
		},
	})
	if err != nil {
		return nil, nil, errors.Wrap(err, "open database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, errors.Wrap(err, "get underlying *sql.DB")
	}
	sqlDB.SetMaxOpenConns(database.MaxOpenConns)
	sqlDB.SetMaxIdleConns(database.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Minute)

	switch database.Type {
	case "postgres":
		configs.UsePostgreSQL = true
	case "mysql":
		configs.UseMySQL = true
		db = db.Set("gorm:table_options", "ENGINE=InnoDB").
			Session(&gorm.Session{
				WithConditions: true,
			})
	case "sqlite3":
		configs.UseSQLite3 = true
	case "mssql":
		configs.UseMSSQL = true
	default:
		panic("unreachable")
	}

	return db, sqlDB, nil
}
