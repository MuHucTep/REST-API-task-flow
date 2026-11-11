package config

import "os"

type Config struct {
	HTTPAddr string

	DBHost string
	DBPort string
	DBUser string
	DBPswd string
	DBName string
}

func Load() Config {
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = ":8000"
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = ":5433"
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "taskflow"
	}

	dbPswd := os.Getenv("DB_PSWD")
	if dbPswd == "" {
		dbPswd = "taskflow"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "taskflow"
	}

	return Config{
		HTTPAddr: addr,
		DBHost:   dbHost,
		DBPort:   dbPort,
		DBUser:   dbUser,
		DBPswd:   dbPswd,
		DBName:   dbName,
	}
}
