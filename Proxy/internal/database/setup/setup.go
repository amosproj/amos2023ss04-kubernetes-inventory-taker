package setup

import (
	"database/sql"
	"os"
	"strings"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"k8s.io/klog"
)

const (
	timeoutSeconds     = 5
	dialTimeoutSeconds = 5
	DBRetryInterval    = 2 * time.Second  // Retry interval for DB connection
	DBTimeout          = 30 * time.Second // Timeout for DB connection
)

// SetupDBConnection setup database connection.
//
//nolint:revive
func SetupDBConnection() *bun.DB {
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

	pgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr("localhost:5432"),
		pgdriver.WithUser(dbUser),
		pgdriver.WithPassword(dbPassword),
		pgdriver.WithDatabase("postgres"),
		pgdriver.WithInsecure(true),
		pgdriver.WithTimeout(timeoutSeconds*time.Second),
		pgdriver.WithDialTimeout(dialTimeoutSeconds*time.Second),
	)

	sqldb := sql.OpenDB(pgconn)

	db := bun.NewDB(sqldb, pgdialect.New())

	testDBConnection(db)

	return db
}

// TestDBConnection attempts to establish a connection with the database and retries for 30 seconds if unsuccessful.
func testDBConnection(db *bun.DB) {
	start := time.Now()

	for {
		err := db.Ping()
		if err != nil {
			if time.Since(start) < DBTimeout {
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

				time.Sleep(DBRetryInterval)
			} else {
				klog.Fatalf("Could not connect to database after 30 seconds: %v", err)
			}
		} else {
			klog.Info("Successfully connected to database.")
		}
	}
}
