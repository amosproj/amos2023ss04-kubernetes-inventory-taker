package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"k8s.io/klog/v2"
)

// SetupDBConnection setup database connection.
func SetupDBConnection() *Queries {
	dbUser, exists := os.LookupEnv("DB_USER")
	if !exists {
		log.Println("DB_USER environment variable is not set. Trying dbUser = postgres")
		dbUser = "postgres"
	}

	dbPassword, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		log.Println("DB_PASSWORD environment variable is not set. Trying dbPassword = example")
		dbPassword = "example"
	}

	dbConnString := fmt.Sprintf("user=%s password=%s host=localhost port=5432 dbname=postgres pool_max_conns=10", dbUser, dbPassword)
	configDB, err := pgxpool.ParseConfig(dbConnString)
	if err != nil {
		klog.ErrorS(err, "Error while parsing dbConfig", "config String", dbConnString)
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), configDB)
	if err != nil {
		klog.ErrorS(err, "Error while building postgres pool from config")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	return New(pool)
}
