// Package adapters are the glue between components and external sources.
// # This manifest was generated by ymir. DO NOT EDIT.
package adapters

import (
	"fmt"
	"net"
	"strconv"
	"time"

	// mysql go lib.
	"github.com/go-sql-driver/mysql"

	"entgo.io/ent/dialect"
	sqlEnt "entgo.io/ent/dialect/sql"
	"github.com/rs/zerolog/log"
)

var YmirBlogMySQLOpen = sqlEnt.Open // YmirBlogMySQLOpen will invoke to test case.

// YmirBlogMySQL is data of instances.
type YmirBlogMySQL struct {
	NetworkDB
	driver *sqlEnt.Driver
}

// Open is open the connection of mysql.
func (ymbl *YmirBlogMySQL) Open() (*sqlEnt.Driver, error) {
	if ymbl.driver == nil {
		return nil, fmt.Errorf("driver was failed to connected")
	}
	return ymbl.driver, nil
}

// Connect is connected the connection of mysql.
func (ymbl *YmirBlogMySQL) Connect() (err error) {
	ymbl.driver, err = YmirBlogMySQLOpen(dialect.MySQL, ymbl.dsn())
	if err != nil {
		log.Error().Err(err).Msg("YmirBlogMySQLOpen is failed to open")
		return err
	}

	if ymbl.MaxIdleCons == 0 {
		ymbl.driver.DB().SetMaxIdleConns(0)
	} else {
		ymbl.driver.DB().SetMaxIdleConns(ymbl.MaxIdleCons)
	}
	return nil
}

// Disconnect is disconnect the connection of mysql.
func (ymbl *YmirBlogMySQL) Disconnect() error {
	return ymbl.driver.Close()
}

func (ymbl *YmirBlogMySQL) dsn() string {
	cfg := mysql.Config{
		User:                 ymbl.User,
		Passwd:               ymbl.Password,
		DBName:               ymbl.Database,
		Timeout:              time.Second * time.Duration(ymbl.ConnectionTimeout),
		ParseTime:            true,
		AllowNativePasswords: true,
		Params:               make(map[string]string),
	}
	if ymbl.Host != "" {
		if ymbl.Host[0] != '/' {
			cfg.Net = "tcp"
			cfg.Addr = ymbl.Host

			if ymbl.Port != 0 {
				cfg.Addr = net.JoinHostPort(ymbl.Host, strconv.Itoa(int(ymbl.Port)))
			}
		} else {
			cfg.Net = "unix"
			cfg.Addr = ymbl.Host
		}
	}
	return cfg.FormatDSN()
}

// WithYmirBlogMySQL option function to assign on adapters.
func WithYmirBlogMySQL(driver Driver[*sqlEnt.Driver]) Option {
	return func(a *Adapter) {
		if err := driver.Connect(); err != nil {
			panic(err)
		}
		open, err := driver.Open()
		if err != nil {
			panic(err)
		}
		a.YmirBlogMySQL = open
	}
}
