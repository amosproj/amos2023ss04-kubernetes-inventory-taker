package setup

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"k8s.io/klog"
)

const (
	connectionTimeout = 5 * time.Second
	dialTimeout       = 5 * time.Second
	retryInterval     = 2 * time.Second  // retry interval for testDBConnection
	retryTimeout      = 30 * time.Second // final timeout for testDBConnection
)

// DBConnection setup database connection.
func DBConnection() *bun.DB {
	dbUser, exists := os.LookupEnv("DB_USER")
	if !exists {
		klog.Warning("DB_USER environment variable is not set. Trying dbUser = postgres")

		dbUser = "postgres"
	}

	dbPassword, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		klog.Warning("DB_PASSWORD environment variable is not set. Trying dbPassword = example")

		dbPassword = "example"
	}

	dbURI, exists := os.LookupEnv("DB_URI")
	if !exists {
		klog.Warning("DB_URI environment variable is not set. Trying dbURI = localhost")

		dbURI = "localhost"
	}

	pgAddr := fmt.Sprintf("%s:5432", dbURI)

	pgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(pgAddr),
		pgdriver.WithUser(dbUser),
		pgdriver.WithPassword(dbPassword),
		pgdriver.WithDatabase("postgres"),
		pgdriver.WithInsecure(true),
		pgdriver.WithTimeout(connectionTimeout),
		pgdriver.WithDialTimeout(dialTimeout),
	)

	sqldb := sql.OpenDB(pgconn)

	dbConn := bun.NewDB(sqldb, pgdialect.New())

	// This line can toggle verbose debugging output for every query that is send
	// db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	dbConn.AddQueryHook(bundebug.NewQueryHook())

	testDBConnection(dbConn)

	return dbConn
}

// TestDBConnection attempts to establish a connection with the database and retries for 30 seconds if unsuccessful.
func testDBConnection(db *bun.DB) {
	start := time.Now()

	for {
		err := db.Ping()
		if err != nil {
			if time.Since(start) < retryTimeout {
				// Check if the error is due to database not started
				switch {
				case strings.Contains(err.Error(), "connection refused"):
					// Check if the error is due to database not started
					klog.Warningf("Database is not up yet: %v. Retrying...", err)

				case strings.Contains(err.Error(), "authentication failed"):
					// Check if the error is due to wrong authentication data
					klog.Fatalf("Wrong database authentication data: %v", err)

				default:
					// Other database error
					klog.Warningf("Could not connect to database: %v. Retrying...", err)
				}

				time.Sleep(retryInterval)
			} else {
				klog.Fatalf("Could not connect to database after 30 seconds: %v", err)
			}
		} else {
			klog.Info("Successfully connected to database.")
			return
		}
	}
}
