package migration

import (
	"fmt"
	"go-web-api/src/constant"
	"go-web-api/src/domain/models"
	"go-web-api/src/infra/persistence/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const countStarExp = "count(*)"

func Up1() {
	database := database.GetDb()

	createTables(database)
	createDefaultUserInformation(database)
	createCategory(database)
	createProduct(database)
}

func createTables(database *gorm.DB) {
	tables := []interface{}{}

	// Models
	tables = addNewTable(database, models.Category{}, tables)
	tables = addNewTable(database, models.Product{}, tables)

	// User
	tables = addNewTable(database, models.User{}, tables)
	tables = addNewTable(database, models.Role{}, tables)
	tables = addNewTable(database, models.UserRole{}, tables)

	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}

func createDefaultUserInformation(database *gorm.DB) {

	adminRole := models.Role{Name: constant.AdminRoleName}
	createRoleIfNotExists(database, &adminRole)

	defaultRole := models.Role{Name: constant.DefaultRoleName}
	createRoleIfNotExists(database, &defaultRole)

	u := models.User{Username: constant.DefaultUserName, FirstName: "Test", LastName: "Test",
		MobileNumber: "09111112222", Email: "admin@admin.com"}
	pass := "12345678"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)

	createAdminUserIfNotExists(database, &u, adminRole.Id)

}

func createRoleIfNotExists(database *gorm.DB, r *models.Role) {
	exists := 0
	database.
		Model(&models.Role{}).
		Select("1").
		Where("name = ?", r.Name).
		First(&exists)
	if exists == 0 {
		database.Create(r)
	}
}

func createAdminUserIfNotExists(database *gorm.DB, u *models.User, roleId int) {
	exists := 0
	database.
		Model(&models.User{}).
		Select("1").
		Where("username = ?", u.Username).
		First(&exists)
	if exists == 0 {
		database.Create(u)
		ur := models.UserRole{UserId: u.Id, RoleId: roleId}
		database.Create(&ur)
	}
}

func createProduct(database *gorm.DB) {
	count := 0
	database.
		Model(&models.Product{}).
		Select(countStarExp).
		Find(&count)
	if count == 0 {
		database.Create(&models.Product{Name: "P1", Price: 10, CategoryId: 1})
	}
}

func createCategory(database *gorm.DB) {
	count := 0
	database.
		Model(&models.Category{}).
		Select(countStarExp).
		Find(&count)
	if count == 0 {
		database.Create(&models.Category{Name: "C1", BaseModel: models.BaseModel{Id: 1}, Products: []models.Product{
			{Name: "P1"},
		}})
	}
}
