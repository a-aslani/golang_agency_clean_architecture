package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/a-aslani/golang_agency_clean_architecture/configs"
	"github.com/a-aslani/golang_agency_clean_architecture/internal/support/model/entity"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/lib/pq"
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

func (r *gateway) SaveContactForm(ctx context.Context, obj *entity.ContactForm) error {
	r.log.Info(ctx, "called")

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	const query = `INSERT INTO contact_us (id, name, email, message, created_at) VALUES ($1, $2, $3, $4, $5)`

	_, err = tx.ExecContext(ctx, query, obj.ID.String(), obj.Name.String(), obj.Email.String(), obj.Message.String(), obj.CreatedAt)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	for _, file := range obj.Files {

		const query = `INSERT INTO contact_us_file (contact_us_id, file_id) VALUES ($1, $2)`

		_, err = tx.ExecContext(ctx, query, obj.ID.String(), file.ID.String())
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *gateway) FindFilesByIDs(ctx context.Context, ids []string) ([]*entity.File, error) {
	r.log.Info(ctx, "called")

	const query = `SELECT id, name, path, created_at FROM files WHERE id = ANY($1::text[])`

	rows, err := r.db.QueryContext(ctx, query, pq.Array(ids))
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err = rows.Close()
	}(rows)

	files := make([]*entity.File, 0)

	for rows.Next() {

		file := new(entity.File)
		err = rows.Scan(&file.ID, &file.Name, &file.Path, &file.Created)
		if err != nil {
			return nil, err
		}

		files = append(files, file)

	}

	return files, nil
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

func (r *gateway) FindRolesByCodes(ctx context.Context, codes []string) ([]*entity.Role, error) {
	r.log.Info(ctx, "called")

	const query = `SELECT id, code, name FROM roles WHERE code = ANY($1::text[])`

	rows, err := r.db.QueryContext(ctx, query, pq.Array(codes))
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err = rows.Close()
	}(rows)

	roles := make([]*entity.Role, 0)

	for rows.Next() {

		role := new(entity.Role)
		err = rows.Scan(&role.ID, &role.Code, &role.Name)
		if err != nil {
			return nil, err
		}

		roles = append(roles, role)

	}

	return roles, nil
}

func (r *gateway) FindChatIdsByRoles(ctx context.Context, roles []*entity.Role) ([]int64, error) {
	r.log.Info(ctx, "called")

	roleIds := make([]string, 0)

	for _, role := range roles {
		roleIds = append(roleIds, role.ID.String())
	}

	const query = `SELECT chat_id FROM telegram_chat_ids WHERE role_id = ANY($1::text[])`

	rows, err := r.db.QueryContext(ctx, query, pq.Array(roleIds))
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err = rows.Close()
	}(rows)

	chatIds := make([]int64, 0)

	for rows.Next() {

		var chatId int64

		err = rows.Scan(&chatId)
		if err != nil {
			return nil, err
		}

		chatIds = append(chatIds, chatId)

	}

	return chatIds, nil
}

func (r *gateway) SiteVerify(ctx context.Context, secret, token string) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) SendMessage(ctx context.Context, chatId int64, text string, parseMode string) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) CommandHandling(ctx context.Context, cmd func(tgbotapi.Update) string) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) SaveRole(ctx context.Context, obj *entity.Role) error {
	r.log.Info(ctx, "called")

	const query = `
		INSERT INTO roles(id, name, code) VALUES ($1, $2, $3)
	`

	_, err := r.db.ExecContext(ctx, query, obj.ID, obj.Name, obj.Code)
	if err != nil {
		return err
	}

	return nil
}

func (r *gateway) SaveTelegramChatID(ctx context.Context, obj *entity.TelegramChatID) error {
	r.log.Info(ctx, "called")

	const query = `
		INSERT INTO telegram_chat_ids(id, chat_id, role_id) VALUES ($1, $2, $3)
	`

	_, err := r.db.ExecContext(ctx, query, obj.ID, obj.ChatID, obj.RoleID)
	if err != nil {
		return err
	}

	return nil
}

func (r *gateway) ClearContactUsFileTable(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM contact_us_file`)
	return err
}

func (r *gateway) ClearContactUsTable(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM contact_us`)
	return err
}

func (r *gateway) ClearTelegramChatIdsTable(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM telegram_chat_ids`)
	return err
}

func (r *gateway) ClearRolesTable(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM roles`)
	return err
}
