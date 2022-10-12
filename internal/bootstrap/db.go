package bootstrap

import (
	"fmt"
	"github.com/danilashushkanov/student-service/internal/closer"
	"github.com/danilashushkanov/student-service/internal/config"
	_ "github.com/jackc/pgx/v4/stdlib" // необходим из-за драйвера pgx
	"github.com/jmoiron/sqlx"
)

func InitDB(cfg *config.Config) (*sqlx.DB, error) {
	conn, err := sqlx.Connect("pgx", "postgres://"+cfg.User+":"+cfg.Password+"@"+cfg.Host+":"+cfg.Port+"/"+cfg.DB)
	if err != nil {
		return nil, fmt.Errorf("can't connect to pg instance, %v", err)
	}

	closer.Add(conn.Close)

	return conn.Unsafe(), nil
}
