package xtesting

import (
	"database/sql"
	"embed"
	"fmt"
	"testing"

	"github.com/Masterminds/sprig/v3"
	"github.com/PRPO-skupina-02/common/config"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/golang-migrate/migrate/v4"
	migratepostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/orgs/PRPO-skupina-02/Spored/database"
	"github.com/stretchr/testify/require"
	gormpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//go:embed all:fixtures/*.yml
var fixtureFS embed.FS

func GetTestDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.GetEnv("POSTGRES_IP"),
		config.GetEnv("POSTGRES_USERNAME"),
		config.GetEnv("POSTGRES_PASSWORD"),
		config.GetEnv("POSTGRES_TEST_DATABASE_NAME"),
		config.GetEnv("POSTGRES_PORT"))
}

func RecreateTestDatabase(t *testing.T) {
	prodDsn := database.GetProdDSN()

	prodDb, err := sql.Open("postgres", prodDsn)
	require.NoError(t, err)

	_, err = prodDb.Exec("DROP DATABASE IF EXISTS " + config.GetEnv("POSTGRES_TEST_DATABASE_NAME") + " WITH (FORCE)")
	require.NoError(t, err)

	_, err = prodDb.Exec("CREATE DATABASE " + config.GetEnv("POSTGRES_TEST_DATABASE_NAME"))
	require.NoError(t, err)

	err = prodDb.Close()
	require.NoError(t, err)
}

func PrepareTestDatabase(t *testing.T) (*gorm.DB, *testfixtures.Loader) {
	RecreateTestDatabase(t)

	dsn := GetTestDSN()
	db, err := gorm.Open(gormpostgres.Open(dsn), &gorm.Config{TranslateError: true})
	require.NoError(t, err)

	dbInstance, err := db.DB()
	require.NoError(t, err)

	driver, err := migratepostgres.WithInstance(dbInstance, &migratepostgres.Config{})
	require.NoError(t, err)

	migrationsFS := database.GetMigrationsFs()
	migrationsFSDriver, err := iofs.New(migrationsFS, "migrations")
	require.NoError(t, err)
	migrations, err := migrate.NewWithInstance("migrations", migrationsFSDriver, "postgres", driver)
	require.NoError(t, err)

	err = migrations.Up()
	require.NoError(t, err)

	fixtures, err := testfixtures.New(
		testfixtures.Template(),
		testfixtures.TemplateFuncs(sprig.FuncMap()),
		testfixtures.Database(dbInstance),
		testfixtures.Dialect("postgres"),
		testfixtures.FS(fixtureFS),
		testfixtures.Directory("fixtures"),
	)
	require.NoError(t, err)

	return db, fixtures
}
