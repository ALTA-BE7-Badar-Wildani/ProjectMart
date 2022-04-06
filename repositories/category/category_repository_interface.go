package category

import entityDomain "go-ecommerce/entities/domain"

type CategoryRepositoryInterface interface {
	FindAll(limit int, offset int, filters []map[string]string, sorts []map[string]interface{}) ([]entityDomain.Category, error)
	Find(id int) (entityDomain.Category, error)
	FindBy(field string, value string) (entityDomain.Category, error)
	CountAll() (int64, error)
	Store(category entityDomain.Category) (entityDomain.Category, error)
	Update(category entityDomain.Category, id int) (entityDomain.Category, error)
	Delete(id int) error
}