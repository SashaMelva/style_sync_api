package storage

import (
	"context"
	"database/sql"

	"github.com/SashaMelva/style_sync_api/internal/config"
	_ "github.com/jackc/pgx/stdlib"
	"go.uber.org/zap"
)

type Storage struct {
	Logger       *zap.SugaredLogger
	ConnectionDB *sql.DB
}

func New(confDB *config.ConfigDB, log *zap.SugaredLogger) *Storage {
	dsn := "postgres://" + confDB.User + ":" + confDB.Password + "@" + confDB.Host + ":" + confDB.Port + "/" + confDB.NameDB
	storage, err := sql.Open("pgx", dsn)

	if err != nil {
		log.Fatal("Cannot open pgx driver: %w", err)
	}

	log.Debug("DSN connection database " + dsn)

	return &Storage{
		Logger:       log,
		ConnectionDB: storage,
	}
}

func (s *Storage) Connect(ctx context.Context) error {
	err := s.ConnectionDB.PingContext(ctx)
	return err
}

func (s *Storage) Close(ctx context.Context) error {
	s.ConnectionDB.Close()
	return nil
}
