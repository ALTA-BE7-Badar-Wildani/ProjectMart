package user

import entityDomain "go-ecommerce/entities/domain"

type UserRepositoryInterface interface {
	FindAll(limit int, offset int, filters []map[string]string, sorts []map[string]interface{}) ([]entityDomain.User, error)
	Find(id int) (entityDomain.User, error)
	FindBy(field string, value string) (entityDomain.User, error)
	CountAll() (int64, error)
	Store(user entityDomain.User) (entityDomain.User, error)
	Update(user entityDomain.User, id int) (entityDomain.User, error)
	Delete(id int) error
}