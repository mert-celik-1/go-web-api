package dependency

import (
	"go-web-api/src/config"
	"go-web-api/src/domain/models"
	contractRepository "go-web-api/src/domain/repository"
	database "go-web-api/src/infra/persistence/database"
	infraRepository "go-web-api/src/infra/persistence/repository"
)

func GetUserRepository(cfg *config.Config) contractRepository.UserRepository {
	return infraRepository.NewUserRepository(cfg)
}

func GetProductRepository(cfg *config.Config) contractRepository.ProductRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{}
	return infraRepository.NewBaseRepository[models.Product](cfg, preloads)
}

func GetCategoryRepository(cfg *config.Config) contractRepository.CategoryRepository {
	var preloads []database.PreloadEntity = []database.PreloadEntity{}
	return infraRepository.NewBaseRepository[models.Category](cfg, preloads)
}
