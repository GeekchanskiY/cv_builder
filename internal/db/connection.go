package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	_ "github.com/lib/pq"

	"github.com/GeekchanskiY/cv_builder/internal/config"
)

var (
	connection *sql.DB
	once       sync.Once
)

func Connect() (*sql.DB, error) {
	var err error = nil
	once.Do(func() {
		config.LoadConfig()
		var (
			host     = os.Getenv("db_host")
			port     = os.Getenv("db_port")
			user     = os.Getenv("db_user")
			password = os.Getenv("db_password")
			dbname   = os.Getenv("db_name")
		)
		log.Println("Connecting to database as user: " + os.Getenv("db_user"))
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
		log.Print("Connecting to database")
		connection, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Printf("Failed to open database connection: %v", err)
			return
		}

		err = connection.Ping()
		if err != nil {
			log.Printf("Failed to ping database: %v", err)
			return
		}

		// Run migrations
		_, filename, _, _ := runtime.Caller(0)
		configDir := filepath.Dir(filename)

		migrationsPath := filepath.Join(configDir, "/migrations")
		log.Println("Reading migrations from: " + migrationsPath)
		dir, err := os.Open(migrationsPath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		files, err := dir.Readdir(0)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		for i, j := 0, len(files)-1; i < j; i, j = i+1, j-1 {
			files[i], files[j] = files[j], files[i]
		}
		for _, file := range files {
			tmp := strings.Split(file.Name(), ".")
			if tmp[len(tmp)-1] == "sql" {
				sqlData, err := os.ReadFile(filepath.Join(migrationsPath, file.Name()))
				if err != nil {
					panic(err)
				}
				log.Println("Running migration: " + file.Name())

				_, err = connection.Exec(string(sqlData))
				if err != nil {
					log.Panicln(err)
				}
			}
		}
		log.Println("Connected to database")
		err = dir.Close()
		if err != nil {
			log.Panicln(err)
		}
	},
	)
	return connection, err
}

func GetDB() *sql.DB {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	return db
}
