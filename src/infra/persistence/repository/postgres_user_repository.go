package repository

import (
	"context"
	"go-web-api/src/config"
	"go-web-api/src/constant"
	"go-web-api/src/domain/models"
	"go-web-api/src/infra/persistence/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const userFilterExp string = "username = ?"
const countFilterExp string = "count(*) > 0"

type PostgresUserRepository struct {
	*BaseRepository[models.User]
}

func NewUserRepository(cfg *config.Config) *PostgresUserRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{} // Kontrol edilecek
	return &PostgresUserRepository{BaseRepository: (*BaseRepository[models.User])(NewBaseRepository[models.User](cfg, preloads))}
}

func (r *PostgresUserRepository) CreateUser(ctx context.Context, u models.User) (models.User, error) {

	roleId, err := r.GetDefaultRole(ctx)
	if err != nil {
		return u, err
	}
	tx := r.database.WithContext(ctx).Begin()
	err = tx.Create(&u).Error
	if err != nil {
		tx.Rollback()
		return u, err
	}
	err = tx.Create(&models.UserRole{RoleId: roleId, UserId: u.Id}).Error
	if err != nil {
		tx.Rollback()
		return u, err
	}
	tx.Commit()
	return u, nil
}

func (r *PostgresUserRepository) FetchUserInfo(ctx context.Context, username string, password string) (models.User, error) {
	var user models.User
	err := r.database.WithContext(ctx).
		Model(&models.User{}).
		Where(userFilterExp, username).
		Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Role")
		}).
		Find(&user).Error

	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *PostgresUserRepository) ExistsEmail(ctx context.Context, email string) (bool, error) {
	var exists bool
	if err := r.database.WithContext(ctx).Model(&models.User{}).
		Select(countFilterExp).
		Where("email = ?", email).
		Find(&exists).
		Error; err != nil {
		return false, err
	}
	return exists, nil
}

func (r *PostgresUserRepository) ExistsUsername(ctx context.Context, username string) (bool, error) {
	var exists bool
	if err := r.database.WithContext(ctx).Model(&models.User{}).
		Select(countFilterExp).
		Where(userFilterExp, username).
		Find(&exists).
		Error; err != nil {
		return false, err
	}
	return exists, nil
}

func (r *PostgresUserRepository) ExistsMobileNumber(ctx context.Context, mobileNumber string) (bool, error) {
	var exists bool
	if err := r.database.WithContext(ctx).Model(&models.User{}).
		Select(countFilterExp).
		Where("mobile_number = ?", mobileNumber).
		Find(&exists).
		Error; err != nil {
		return false, err
	}
	return exists, nil
}

func (r *PostgresUserRepository) GetDefaultRole(ctx context.Context) (roleId int, err error) {

	if err = r.database.WithContext(ctx).Model(&models.Role{}).
		Select("id").
		Where("name = ?", constant.DefaultRoleName).
		First(&roleId).Error; err != nil {
		return 0, err
	}
	return roleId, nil
}
