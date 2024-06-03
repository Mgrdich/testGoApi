package repository

import (
	"context"

	db2 "testGoApi/internal/db"
	db "testGoApi/internal/db/sqlc"
	"testGoApi/internal/models"
)

type UserRepositoryImpl struct {
	q *db.Queries
}

func NewUserRepositoryImpl(queries *db.Queries) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		q: queries,
	}
}

func mapDBUserToModelUser(user *db.User) *models.User {
	model := &models.User{
		ID:       user.ID.Bytes,
		Username: user.Username.String,
		Password: user.Password.String,
		Role:     models.BasicRole,
	}

	if user.Role == db.RoleAdmin {
		model.Role = models.AdminRole
	}

	return model
}

func (r *UserRepositoryImpl) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	dbUser, err := r.q.GetByUsername(ctx, db2.ToText(username))

	if err != nil {
		return nil, err
	}

	user := mapDBUserToModelUser(&dbUser)

	return user, nil
}

func (r *UserRepositoryImpl) Save(ctx context.Context, param models.CreateUser) (*models.User, error) {
	dbParam := db.CreateUserParams{
		Username: db2.ToText(param.Username),
		Password: db2.ToText(param.Password),
		Role:     db.RoleUser,
	}

	if param.Role == models.AdminRole {
		dbParam.Role = db.RoleAdmin
	}

	dbUser, err := r.q.CreateUser(ctx, dbParam)

	if err != nil {
		return nil, err
	}

	user := mapDBUserToModelUser(&dbUser)

	return user, nil
}
