// Part of this code credit to https://github.com/gogs/gogs
// This code modify from original code https://github.com/gogs/gogs/blob/main/internal/conf/static.go
// Copyright 2020 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package configs

//DatabaseOpts struct
type DatabaseOpts struct {
	Type         string
	Host         string
	Name         string
	User         string
	Password     string
	SSLMode      string
	Path         string
	MaxOpenConns int
	MaxIdleConns int
}

// Database settings
var Database DatabaseOpts

// Indicates which database backend is currently being used.
var (
	UseSQLite3    bool
	UseMySQL      bool
	UsePostgreSQL bool
	UseMSSQL      bool
)
