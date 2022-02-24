package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Config for the environment
type Config struct {
	// Debug           bool   `envconfig:"DEBUG"`
	// Addr            string `envconfig:"ADDR" default:"8081"`
	MyAuthUser string `envconfig:"MY_AUTH_USERNAME"`
	MyAuthPass string `envconfig:"MY_AUTH_PASSWORD"`
	// MySQLDatasource string `envconfig:"MYSQLDATASOURCE"`
	// DbSecrets       string `envconfig:"DB_SECRET"`
}

// type DBSecrets struct {
// 	Password string `json:"password,omitempty"`
// 	DBName   string `json:"dbname,omitempty"`
// 	Engine   string `json:"engine,omitempty"`
// 	Port     int    `json:"port,omitempty"`
// 	Host     string `json:"host,omitempty"`
// 	Username string `json:"username,omitempty"`
// }

// this extracts the JSON secret value from the environment and builds the
// datasource used to connect to mysql
// func (cfg *Config) parseDbSecrets() error {
// 	// if cfg.DbSecrets == "" {
// 	// 	log.Printf("no DB secrets value provided")
// 	// 	return nil
// 	// }

// 	// dbsecrets := &DBSecrets{}

// 	// err := json.Unmarshal([]byte(cfg.DbSecrets), dbsecrets)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// cfg.MySQLDatasource = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
// 	// 	dbsecrets.Username,
// 	// 	dbsecrets.Password,
// 	// 	dbsecrets.Host,
// 	// 	dbsecrets.DBName,
// 	// )
// 	log.Printf("Configured MYSQL datasource")
// 	return nil
// }

func New() (*Config, error) {
	cfg := new(Config)
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse environment config")
	}
	// err = cfg.parseDbSecrets()
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed parse DB secret")
	// }

	return cfg, nil
}

func SetupDbConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load ")
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	log.Printf("Connecting to DB ...")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Can't connect to DB")
	}
	log.Printf("DB connection Succeded !")
	return db
}
