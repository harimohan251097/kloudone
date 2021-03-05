package database

import (
	"fmt"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once
var dba *gorm.DB

// GetInstance - Returns a DB instance
func GetInstance() *gorm.DB {
	once.Do(func() {
		// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
		user := viper.GetString("database.user")
		password := viper.GetString("database.password")
		host := viper.GetString("database.host")
		port := viper.GetString("database.port")
		dbname := viper.GetString("database.dbname")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		dba = db
		if err != nil {
			log.Panic().Msgf("Error connecting to the database at %s:%s/%s", host, port, dbname)
		}
		sqlDB, err := dba.DB()
		if err != nil {
			log.Panic().Msgf("Error getting GORM DB definition")
		}
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)

		log.Info().Msgf("Successfully established connection to %s:%s/%s", host, port, dbname)
	})
	return dba
}
