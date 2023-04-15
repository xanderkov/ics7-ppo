package db

import (
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"errors"
	"fmt"
	"go.uber.org/zap"
	err_c "hospital/src/internal/models/errors"
	"hospital/src/internal/modules/config"
	"hospital/src/internal/modules/db/ent"
	"hospital/src/internal/modules/db/trace_driver"
	"time"
)

func connectDB(cfg config.Config, logger *zap.Logger) (*ent.Client, error) {
	/* тут используем cgf и формируем строку подключения */
	connString := "host=localhost port=5432 user=postgres dbname=server_db password=postgres sslmode=disable"

	db, err := sql.Open(dialect.Postgres, connString)
	if err != nil {
		return nil, fmt.Errorf("ошибка при подключении к БД: %w", err)
	}

	logLevel := trace_driver.Warn
	if cfg.TraceSQLCommands {
		logLevel = trace_driver.Info
	}

	traceDriver := trace_driver.NewTraceDriver(db, NewLogger(
		logger,
		trace_driver.Config{
			SlowThreshold: time.Duration(cfg.SQLSlowThreshold) * time.Second,
			LogLevel:      logLevel,
		}))

	var opts []ent.Option
	opts = append(opts, ent.Driver(traceDriver))

	client := ent.NewClient(opts...)

	return client, nil
}

func WrapError(err error) error {
	if errors.Is(err, &ent.NotFoundError{}) {
		return err_c.ErrDatabaseRecordNotFound
	}

	return err
}
