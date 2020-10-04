package api

import (
	"context"

	"github.com/satesate-dev/go-restful-boilerplate/user"

	"github.com/satesate-dev/go-restful-boilerplate/user/model"
)

func GetAllUser(ctx context.Context) (interface{}, error) {
	users, err := model.GetAllUser(ctx, user.DBPool)
	if err != nil {
		return nil, err
	}
	var response []model.UserResponse
	for _, user := range users {
		response = append(response, user.Response())
	}

	return response, nil
}
