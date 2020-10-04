package migration

import (
	"database/sql"

	migrate "github.com/rubenv/sql-migrate"
)

type (
	Config struct {
		DBConnection   *sql.DB
		DBDriver       string
		MigrationDir   string
		MigrationTable string
		RunMigrations  bool
	}

	Migration struct {
		config          *Config
		migrationSource *migrate.FileMigrationSource
	}
)

func NewMigration(config *Config) (migrationInstance *Migration) {
	migrationInstance = &Migration{
		config: config,
		migrationSource: &migrate.FileMigrationSource{
			Dir: config.MigrationDir,
		},
	}
	migrate.SetTable(config.MigrationTable)
	return
}

func (d *Migration) Sync() (err error) {
	if !d.config.RunMigrations {
		return
	}

	_, err = migrate.Exec(
		d.config.DBConnection,
		d.config.DBDriver,
		d.migrationSource,
		migrate.Up,
	)
	return
}

func (d *Migration) Clear() (err error) {
	_, err = migrate.Exec(
		d.config.DBConnection,
		d.config.DBDriver,
		d.migrationSource,
		migrate.Down,
	)
	return
}
