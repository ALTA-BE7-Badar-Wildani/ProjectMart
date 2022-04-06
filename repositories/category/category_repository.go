package category

import (
	entityDomain "go-ecommerce/entities/domain"
	web "go-ecommerce/entities/web"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)
type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return CategoryRepository{
		db: db,
	}
}

func (repo CategoryRepository) FindAll(limit int, offset int, filters []map[string]string, sorts []map[string]interface{}) ([]entityDomain.Category, error) {
	categories := []entityDomain.Category{}

	builder := repo.db.Limit(limit).Offset(offset)
	// Where filters
	for _, filter := range filters {
		builder.Where(filter["field"] + " " + filter["operator"] + " ?", filter["value"])
	}
	// OrderBy Filters
	for _, sort := range sorts {
		builder.Order(clause.OrderByColumn{Column: clause.Column{Name: sort["field"].(string)}, Desc: sort["desc"].(bool)})
	}
	tx := builder.Find(&categories)
	if tx.Error != nil {
		return []entityDomain.Category{}, web.WebError{Code: 500, Message: tx.Error.Error()} 
	}
	return categories, nil
}
func (repo CategoryRepository) CountAll() (int64, error) {
	var count int64
	tx := repo.db.Model(&entityDomain.Category{}).Count(&count)
	if tx.Error != nil {
		return -1, web.WebError{Code: 400, Message: tx.Error.Error()}
	}
	return count, nil
}

func (repo CategoryRepository) Find(id int) (entityDomain.Category, error) {
	category := entityDomain.Category{}
	tx := repo.db.Find(&category, id)
	if tx.Error != nil {
		return entityDomain.Category{}, web.WebError{Code: 500, Message: "server error"}
	} else if tx.RowsAffected <= 0 {
		return entityDomain.Category{}, web.WebError{Code: 400, Message: "cannot get category data with specified id"}
	}
	return category, nil
}

func (repo CategoryRepository) FindBy(field string, value string) (entityDomain.Category, error) {
	category := entityDomain.Category{}
	tx := repo.db.Where(field + " = ?", value).Find(&category)
	if tx.Error != nil {
		return entityDomain.Category{}, web.WebError{ Code: 500, Message: tx.Error.Error() }
	} else if tx.RowsAffected <= 0 {
		return entityDomain.Category{}, web.WebError{ Code: 400, Message: "The requested ID doesn't match with any record" }
	}
	return category, nil
}

func (repo CategoryRepository) Store(category entityDomain.Category) (entityDomain.Category, error) {
	
	tx := repo.db.Create(&category)
	if tx.Error != nil {
		return entityDomain.Category{}, web.WebError{Code: 500, Message: tx.Error.Error()}
	}
	return category, nil
}

func (repo CategoryRepository) Update(category entityDomain.Category, id int) (entityDomain.Category, error) {
	tx := repo.db.Save(&category)
	if tx.Error != nil {
		return entityDomain.Category{}, web.WebError{Code: 500, Message: tx.Error.Error()}
	}
	return category, nil
}

func (repo CategoryRepository) Delete(id int) error {
	tx := repo.db.Delete(&entityDomain.Category{}, id)
	if tx.Error != nil {
		return web.WebError{Code: 500, Message: tx.Error.Error()}
	}
	return nil
}
