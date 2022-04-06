package transaction_item

import (
	entityDomain "go-ecommerce/entities/domain"
	web "go-ecommerce/entities/web"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)
type TransactionItemRepository struct {
	db *gorm.DB
}

func NewTransactionItemRepository(db *gorm.DB) TransactionItemRepository {
	return TransactionItemRepository{
		db: db,
	}
}

func (repo TransactionItemRepository) FindAll(limit int, offset int, filters []map[string]string, sorts []map[string]interface{}) ([]entityDomain.TransactionItem, error) {
	transactionItems := []entityDomain.TransactionItem{}

	builder := repo.db.Preload("Transaction").Preload("Product").Limit(limit).Offset(offset)
	// Where filters
	for _, filter := range filters {
		builder.Where(filter["field"] + " " + filter["operator"] + " ?", filter["value"])
	}
	// OrderBy Filters
	for _, sort := range sorts {
		builder.Order(clause.OrderByColumn{Column: clause.Column{Name: sort["field"].(string)}, Desc: sort["desc"].(bool)})
	}
	tx := builder.Find(&transactionItems)
	if tx.Error != nil {
		return []entityDomain.TransactionItem{}, web.WebError{Code: 500, Message: tx.Error.Error()} 
	}
	return transactionItems, nil
}
func (repo TransactionItemRepository) CountAll() (int64, error) {
	var count int64
	tx := repo.db.Model(&entityDomain.TransactionItem{}).Count(&count)
	if tx.Error != nil {
		return -1, web.WebError{Code: 400, Message: tx.Error.Error()}
	}
	return count, nil
}

func (repo TransactionItemRepository) Find(id int) (entityDomain.TransactionItem, error) {
	transactionItem := entityDomain.TransactionItem{}
	tx := repo.db.Preload("Transaction").Preload("Product").Find(&transactionItem, id)
	if tx.Error != nil {
		return entityDomain.TransactionItem{}, web.WebError{Code: 500, Message: "server error"}
	} else if tx.RowsAffected <= 0 {
		return entityDomain.TransactionItem{}, web.WebError{Code: 400, Message: "cannot get transactionItem data with specified id"}
	}
	return transactionItem, nil
}

func (repo TransactionItemRepository) FindBy(field string, value string) (entityDomain.TransactionItem, error) {
	transactionItem := entityDomain.TransactionItem{}
	tx := repo.db.Preload("Transaction").Preload("Product").Where(field + " = ?", value).Find(&transactionItem)
	if tx.Error != nil {
		return entityDomain.TransactionItem{}, web.WebError{ Code: 500, Message: tx.Error.Error() }
	} else if tx.RowsAffected <= 0 {
		return entityDomain.TransactionItem{}, web.WebError{ Code: 400, Message: "The requested ID doesn't match with any record" }
	}
	return transactionItem, nil
}

func (repo TransactionItemRepository) Store(transactionItem entityDomain.TransactionItem) (entityDomain.TransactionItem, error) {
	
	tx := repo.db.Preload("Transaction").Preload("Product").Create(&transactionItem)
	if tx.Error != nil {
		return entityDomain.TransactionItem{}, web.WebError{Code: 500, Message: tx.Error.Error()}
	}
	return transactionItem, nil
}

func (repo TransactionItemRepository) Update(transactionItem entityDomain.TransactionItem, id int) (entityDomain.TransactionItem, error) {
	tx := repo.db.Save(&transactionItem)
	if tx.Error != nil {
		return entityDomain.TransactionItem{}, web.WebError{Code: 500, Message: tx.Error.Error()}
	}
	return transactionItem, nil
}

func (repo TransactionItemRepository) Delete(id int) error {
	tx := repo.db.Delete(&entityDomain.TransactionItem{}, id)
	if tx.Error != nil {
		return web.WebError{Code: 500, Message: tx.Error.Error()}
	}
	return nil
}
