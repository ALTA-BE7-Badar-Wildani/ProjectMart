package user

import (
	entityDomain "go-ecommerce/entities/domain"
	web "go-ecommerce/entities/web"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (repo UserRepository) FindAll(limit int, offset int, filters []map[string]string, sorts []map[string]interface{}) ([]entityDomain.User, error) {
	users := []entityDomain.User{}


	builder := repo.db.Limit(limit).Offset(offset)
	// Where filters
	for _, filter := range filters {
		builder.Where(filter["field"] + " " + filter["operator"] + " ?", filter["value"])
	}
	// OrderBy Filters
	for _, sort := range sorts {
		builder.Order(clause.OrderByColumn{Column: clause.Column{Name: sort["field"].(string)}, Desc: sort["desc"].(bool)})
	}
	tx := builder.Find(&users)
	if tx.Error != nil {
		return []entityDomain.User{}, web.WebError{Code: 500, Message: tx.Error.Error()} 
	}
	return users, nil
}
func (repo UserRepository) CountAll() (int64, error) {
	var count int64
	tx := repo.db.Model(&entityDomain.User{}).Count(&count)
	if tx.Error != nil {
		return -1, web.WebError{Code: 400, Message: tx.Error.Error()}
	}
	return count, nil
}

func (repo UserRepository) Find(id int) (entityDomain.User, error) {
	panic("Implement me")
}

func (repo UserRepository) FindBy(field string, value string) (entityDomain.User, error) {
	panic("Implement me")
}

func (repo UserRepository) Store(user entityDomain.User) (entityDomain.User, error) {
	panic("Implement me")
}

func (repo UserRepository) Update(user entityDomain.User, id int) (entityDomain.User, error) {
	panic("Implement me")
}

func (repo UserRepository) Delete(id int) error {
	panic("Implement me")
}
