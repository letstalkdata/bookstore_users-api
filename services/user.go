package services

import (
	"net/http"

	"github.com/letstalkdata/bookstore_users-api/domain/users"
	"github.com/letstalkdata/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil

	return nil, &errors.RestErr{
		Status: http.StatusInternalServerError,
	}
}
