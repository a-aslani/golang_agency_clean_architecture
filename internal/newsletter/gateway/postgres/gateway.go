package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/logger"

	"github.com/a-aslani/golang_agency_clean_architecture/configs"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/newsletter/model/entity"
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

func (r *gateway) SaveSubscriber(ctx context.Context, obj *entity.Subscriber) error {
	r.log.Info(ctx, "called")

	const query = `INSERT INTO subscribers (id, email, created) VALUES ($1, $2, $3)`

	_, err := r.db.ExecContext(ctx, query, obj.ID, obj.Email, obj.Created)
	if err != nil {
		return err
	}

	return nil
}

func (r *gateway) FindOneSubscriberByEmail(ctx context.Context, email string) (*entity.Subscriber, error) {
	r.log.Info(ctx, "called")

	const query = `SELECT id, email, created FROM subscribers WHERE email=$1 LIMIT 1`

	var obj entity.Subscriber

	err := r.db.QueryRowContext(ctx, query, email).Scan(&obj.ID, &obj.Email, &obj.Created)
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

func (r *gateway) FindOneSubscriberByID(ctx context.Context, subscriberID string) (*entity.Subscriber, error) {
	//TODO implement me
	panic("implement me")
}

func (r *gateway) ClearSubscribers(ctx context.Context) error {
	r.log.Info(ctx, "called")

	_, err := r.db.ExecContext(context.Background(), `DELETE FROM subscribers`)

	return err
}
