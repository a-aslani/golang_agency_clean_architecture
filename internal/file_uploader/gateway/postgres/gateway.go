package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/a-aslani/golang_agency_clean_architecture/configs"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/logger"

	"github.com/a-aslani/golang_agency_clean_architecture/internal/file_uploader/model/entity"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type gateway struct {
	appData framework.ApplicationData
	config  *configs.Config
	log     logger.Logger
	db      *sql.DB
}

// NewGateway ...
func NewGateway(log logger.Logger, appData framework.ApplicationData, cfg *configs.Config) (*gateway, error) {

	cdb := cfg.Servers[appData.AppName].PostgresDB

	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s", cdb.Host, cdb.Port, cdb.Name, cdb.User, cdb.Password, cdb.SSLMode)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	return &gateway{
		log:     log,
		appData: appData,
		config:  cfg,
		db:      db,
	}, nil
}

func (r *gateway) SaveFilePath(ctx context.Context, obj *entity.File) error {
	r.log.Info(ctx, "called")

	const query = `
		INSERT INTO files(id, name, path, created_at) VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.ExecContext(ctx, query, obj.ID, obj.Name, obj.Path, obj.Created)
	if err != nil {
		return err
	}

	return nil
}

func (r *gateway) ClearFilesTable(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM files`)
	return err
}
