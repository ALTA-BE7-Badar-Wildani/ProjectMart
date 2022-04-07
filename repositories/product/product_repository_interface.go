package product

import entityDomain "go-ecommerce/entities/domain"

type ProductRepositoryInterface interface {
	FindAll(limit int, offset int, filters []map[string]string, sorts []map[string]interface{}) ([]entityDomain.Product, error)
	Find(id int) (entityDomain.Product, error)
	FindBy(field string, value string) (entityDomain.Product, error)
	CountAll(filters []map[string]string) (int64, error)
	Store(product entityDomain.Product) (entityDomain.Product, error)
	Update(product entityDomain.Product, id int) (entityDomain.Product, error)
	Delete(id int) error
}