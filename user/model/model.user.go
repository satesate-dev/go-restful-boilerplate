package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"

	uuid "github.com/satori/go.uuid"
)

type (
	UserModel struct {
		ID        uuid.UUID     `json:"id"`
		Name      string        `json:"name"`
		Username  string        `json:"username"`
		Password  string        `json:"password"`
		CreatedAt time.Time     `json:"created_at"`
		CreatedBy uuid.UUID     `json:"created_by"`
		UpdatedAt pq.NullTime   `json:"updated_at"`
		UpdatedBy uuid.NullUUID `json:"updated_by"`
	}

	UserResponse struct {
		ID        uuid.UUID `json:"id"`
		Name      string    `json:"name"`
		Username  string    `json:"username"`
		Password  string    `json:"-"`
		CreatedAt time.Time `json:"created_at"`
		CreatedBy uuid.UUID `json:"created_by"`
		UpdatedAt time.Time `json:"updated_at"`
		UpdatedBy uuid.UUID `json:"updated_by"`
	}
)

func (u UserModel) Response() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Username:  u.Username,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		CreatedBy: u.CreatedBy,
		UpdatedAt: u.UpdatedAt.Time,
		UpdatedBy: u.UpdatedBy.UUID,
	}
}

func GetAllUser(ctx context.Context, db *sql.DB) ([]UserModel, error) {
	query := fmt.Sprintf(`
		SELECT 
			id,
			name,
			username,
			password,
			created_at,
			created_by,
			updated_at,
			updated_by
		FROM
			"user"
	`)

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var users []UserModel
	for rows.Next() {
		var user UserModel
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Password,
			&user.CreatedAt,
			&user.CreatedBy,
			&user.UpdatedAt,
			&user.UpdatedBy,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
