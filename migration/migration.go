package migration

import (
	"database/sql"
	"embed"
	_ "github.com/jackc/pgx/v5/stdlib"
	migrate "github.com/rubenv/sql-migrate"
	"io/fs"
	"net/http"
	"sso/internal/config"
)

var (
	//go:embed sql/*.sql
	migrations embed.FS
)

func Apply() (applied int, err error) {
	subFS, _ := fs.Sub(migrations, "sql")
	staticFS := http.FS(subFS)
	var migrationSource = &migrate.HttpFileSystemMigrationSource{
		FileSystem: staticFS,
	}

	dsn := config.Get().Postgre.GetDSN()
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return -1, err
	}
	defer db.Close()

	return migrate.Exec(db, "postgres", migrationSource, migrate.Up)
}
