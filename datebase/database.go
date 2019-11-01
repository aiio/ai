package database

import (
	"fmt"
	"github.com/aiio/ai/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var (
	DB      *gorm.DB
	err     error
	dialect = config.V.Section("DB").Key("CONNECTION").String()
	args    string
	debug   bool
)

func init() {
	debug, _ := config.V.Section("").Key("APP_DEBUG").Bool()

	switch dialect {
	case "mysql":
		args = fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
			config.V.Section("DB").Key("USERNAME").String(),
			config.V.Section("DB").Key("PASSWORD").String(),
			"tcp",
			config.V.Section("DB").Key("HOST").String(),
			config.V.Section("DB").Key("PORT").String(),
			config.V.Section("DB").Key("DATABASE").String(),
		)
		break
	case "sqlite3":
		args = config.V.Section("DB").Key("HOST").String()
		break
	default:
		panic("not support this dialect:" + dialect)

	}
	DB, err = gorm.Open(dialect, args)
	if err != nil {
		log.Panic(err)
	}
	if debug {
		DB.LogMode(true)
	}
}
