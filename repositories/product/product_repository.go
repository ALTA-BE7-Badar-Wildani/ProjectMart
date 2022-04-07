package product

import (
	entityDomain "go-ecommerce/entities/domain"
	web "go-ecommerce/entities/web"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)
type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{
		db: db,
	}
}

func (repo ProductRepository) FindAll(limit int, offset int, filters []map[string]string, sorts []map[string]interface{}) ([]entityDomain.Product, error) {
	products := []entityDomain.Product{}

	builder := repo.db.Preload("User").Preload("Category").Limit(limit).Offset(offset)
	// Where filters
	for _, filter := range filters {
		builder.Where(filter["field"] + " " + filter["operator"] + " ?", filter["value"])
	}
	// OrderBy Filters
	for _, sort := range sorts {
		builder.Order(clause.OrderByColumn{Column: clause.Column{Name: sort["field"].(string)}, Desc: sort["desc"].(bool)})
	}
	tx := builder.Find(&products)
	if tx.Error != nil {
		return []entityDomain.Product{}, web.WebError{Code: 500, Message: tx.Error.Error()} 
	}
	return products, nil
}
func (repo ProductRepository) CountAll(filters []map[string]string) (int64, error) {
	var count int64
	builder := repo.db.Model(&entityDomain.Product{})
	// Where filters
	for _, filter := range filters {
		builder.Where(filter["field"] + " " + filter["operator"] + " ?", filter["value"])
	}
	tx := builder.Count(&count)
	if tx.Error != nil {
		return -1, web.WebError{Code: 400, Message: tx.Error.Error()}
	}
	return count, nil
}

func (repo ProductRepository) Find(id int) (entityDomain.Product, error) {
	product := entityDomain.Product{}
	tx := repo.db.Preload("User").Preload("Category").Find(&product, id)
	if tx.Error != nil {
		return entityDomain.Product{}, web.WebError{Code: 500, Message: "server error"}
	} else if tx.RowsAffected <= 0 {
		return entityDomain.Product{}, web.WebError{Code: 400, Message: "cannot get product data with specified id"}
	}
	return product, nil
}

func (repo ProductRepository) FindBy(field string, value string) (entityDomain.Product, error) {
	product := entityDomain.Product{}
	tx := repo.db.Preload("User").Preload("Category").Where(field + " = ?", value).Find(&product)
	if tx.Error != nil {
		return entityDomain.Product{}, web.WebError{ Code: 500, Message: tx.Error.Error() }
	} else if tx.RowsAffected <= 0 {
		return entityDomain.Product{}, web.WebError{ Code: 400, Message: "The requested ID doesn't match with any record" }
	}
	return product, nil
}

func (repo ProductRepository) Store(product entityDomain.Product) (entityDomain.Product, error) {
	
	tx := repo.db.Preload("User").Preload("Category").Create(&product)
	if tx.Error != nil {
		return entityDomain.Product{}, web.WebError{Code: 500, Message: tx.Error.Error()}
	}
	return product, nil
}

func (repo ProductRepository) Update(product entityDomain.Product, id int) (entityDomain.Product, error) {
	tx := repo.db.Save(&product)
	if tx.Error != nil {
		return entityDomain.Product{}, web.WebError{Code: 500, Message: tx.Error.Error()}
	}
	return product, nil
}

func (repo ProductRepository) Delete(id int) error {
	tx := repo.db.Delete(&entityDomain.Product{}, id)
	if tx.Error != nil {
		return web.WebError{Code: 500, Message: tx.Error.Error()}
	}
	return nil
}
