package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/widyan/go-codebase/middleware/entity"
	"github.com/widyan/go-codebase/middleware/interfaces"

	"github.com/sirupsen/logrus"
)

type Repository struct {
	// CQRS design pattern best practice. memisahkan DB untuk write dan read
	DBWrite *sql.DB
	DBRead  *sql.DB
	Logger  *logrus.Logger
}

func CreateRepository(dbWrite, dbRead *sql.DB, logger *logrus.Logger) interfaces.RepositoryInterface {
	return &Repository{dbWrite, dbRead, logger}
}

func TimeStamapNow() time.Time {
	location, _ := time.LoadLocation("Asia/Jakarta")
	return time.Now().In(location)
}

func (r *Repository) AddUser(ctx context.Context, user entity.User) (err error) {
	query := `INSERT INTO auth."user" (id, "name", email, "role", is_active, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7)`
	_, err = r.DBWrite.ExecContext(ctx, query, user.ID, user.Name, user.Email, user.Role, user.IsActive, TimeStamapNow(), TimeStamapNow())
	return
}
