package repository

import (
	"context"

	"github.com/widyan/go-codebase/middleware/entity"
)

func (r *Repository) GetUserBasedOnEmail(ctx context.Context, email string) (users []entity.User, err error) {
	var user entity.User
	rows, err := r.DBRead.QueryContext(ctx,
		`SELECT id, "name", email, "role", is_active FROM auth."user" where email = $1`, email,
	)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.IsActive)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	return
}
