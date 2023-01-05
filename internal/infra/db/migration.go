package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // driver sql

	"github.com/golang-migrate/migrate"
	_mysql "github.com/golang-migrate/migrate/database/mysql" // driver sql (migrate)
)

type Migration struct {
	Migrate *migrate.Migrate
}

// Up deverá rodar as migrations de criação
func (m *Migration) Up() error {
	err := m.Migrate.Up()

	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

// Down deverá rodar as migrations de reversão
func (m *Migration) Down() error {
	return m.Migrate.Down()
}

// Deverá criar um *Migration configurado e pronto para rodar as migrações
func CreateMigration(dbConn *sql.DB, migrationsFolderLocation string) (*Migration, error) {
	driver, err := _mysql.WithInstance(dbConn, &_mysql.Config{})
	if err != nil {
		return nil, err
	}

	pathToMigrate := fmt.Sprintf("file://%s", migrationsFolderLocation)

	m, err := migrate.NewWithDatabaseInstance(pathToMigrate, "mysql", driver)
	if err != nil {
		return nil, err
	}

	return &Migration{Migrate: m}, nil
}
