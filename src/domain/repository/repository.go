package repository

import (
	"context"
	"go-web-api/src/domain/filter"
	"go-web-api/src/domain/models"
)

type BaseRepository[TEntity any] interface {
	Create(ctx context.Context, entity TEntity) (TEntity, error)
	Update(ctx context.Context, id int, entity map[string]interface{}) (TEntity, error)
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (TEntity, error)
	GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (int64, *[]TEntity, error)
}

type ProductRepository interface {
	BaseRepository[models.Product]
}

type CategoryRepository interface {
	BaseRepository[models.Category]
}

type UserRepository interface {
	ExistsMobileNumber(ctx context.Context, mobileNumber string) (bool, error)
	ExistsUsername(ctx context.Context, username string) (bool, error)
	ExistsEmail(ctx context.Context, email string) (bool, error)
	FetchUserInfo(ctx context.Context, username string, password string) (models.User, error)
	GetDefaultRole(ctx context.Context) (roleId int, err error)
	CreateUser(ctx context.Context, u models.User) (models.User, error)
}
