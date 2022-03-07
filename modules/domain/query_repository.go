package domain

import (
	"codebase/go-codebase/entity"
	"codebase/go-codebase/helper/null"
	"context"
	"fmt"
)

const (
	SampleQuery = ""
)

func (r Repository) GetOneUser(ctx context.Context) (user entity.Users, err error) {
	query := "select ss.id, ss.fullname, ss.no_hp, ss.is_attend, ss.created_at from users ss limit 1"
	rows, err := r.DBRead.QueryContext(ctx, query)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var fullname, noHp string
		var isAttend bool

		var createdAt null.NullString
		err = rows.Scan(&id, &fullname, &noHp, &isAttend, &createdAt)
		if err != nil {
			return
		}

		user.ID = id
		user.Fullname = fullname
		user.NoHP = noHp
		user.IsAttend = isAttend
		user.CreatedAt = createdAt.String
	}
	return
}

func (r Repository) GetAllUsers(ctx context.Context) (users []entity.Users, err error) {
	query := "select ss.id, ss.fullname, ss.no_hp, ss.is_attend, ss.created_at from users ss"
	rows, err := r.DBRead.QueryContext(ctx, query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var fullname, noHp string
		var isAttend bool

		var createdAt null.NullString
		err = rows.Scan(&id, &fullname, &noHp, &isAttend, &createdAt)
		if err != nil {
			return
		}

		users = append(users, entity.Users{id, fullname, noHp, isAttend, createdAt.String})
	}

	if len(users) == 0 {
		err = fmt.Errorf("Data kosong")
		return
	}
	return
}
