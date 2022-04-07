package transaction

import (
	"errors"
	entityDomain "go-ecommerce/entities/domain"
	"go-ecommerce/entities/web"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)


type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (repo TransactionRepository) FindAll(limit int, offset int, filters []map[string]string, sorts []map[string]interface{}) ([]entityDomain.Transaction, error) {
	transactions := []entityDomain.Transaction{}

	builder := repo.db.Preload("Items").Preload("Items.Product").Preload("Items.Product.Category").Limit(limit).Offset(offset)
	// Where filters
	for _, filter := range filters {
		builder.Where(filter["field"] + " " + filter["operator"] + " ?", filter["value"])
	}
	// OrderBy Filters
	for _, sort := range sorts {
		builder.Order(clause.OrderByColumn{Column: clause.Column{Name: sort["field"].(string)}, Desc: sort["desc"].(bool)})
	}
	tx := builder.Find(&transactions)
	if tx.Error != nil {
		return []entityDomain.Transaction{}, web.WebError{Code: 500, Message: tx.Error.Error()} 
	}
	return transactions, nil
}

func (repo TransactionRepository) CountAll(filters []map[string]string) (int64, error) {
	var count int64

	builder := repo.db.Model(&entityDomain.Transaction{})
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


func (repo TransactionRepository) FindFirst(filters []map[string]string) (entityDomain.Transaction, error) {
	transactions := entityDomain.Transaction{}

	builder := repo.db.Preload("User")

	// Where filters
	for _, filter := range filters {
		builder.Where(filter["field"] + " " + filter["operator"] + " ?", filter["value"])
	}
	tx := builder.First(&transactions)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return entityDomain.Transaction{}, web.WebError{Code: 400, Message: "No transaction matched with the provided id"}
	} else if tx.Error != nil {
		return entityDomain.Transaction{}, web.WebError{Code: 500, Message: tx.Error.Error()} 
	}
	return transactions, nil
}

func (repo TransactionRepository) Find(id int) (entityDomain.Transaction, error) {
	transaction := entityDomain.Transaction{}
	tx := repo.db.Preload("User").Find(&transaction, id)
	if tx.Error != nil {
		return entityDomain.Transaction{}, web.WebError{Code: 500, Message: "server error"}
	} else if tx.RowsAffected <= 0 {
		return entityDomain.Transaction{}, web.WebError{Code: 400, Message: "cannot get transaction data with specified id"}
	}
	return transaction, nil
}

func (repo TransactionRepository) Store(transaction entityDomain.Transaction) (entityDomain.Transaction, error) {
	
	tx := repo.db.Preload("User").Create(&transaction)
	if tx.Error != nil {
		return entityDomain.Transaction{}, web.WebError{Code: 500, Message: tx.Error.Error()}
	}
	return transaction, nil
}

func (repo TransactionRepository) Update(transaction entityDomain.Transaction, id int) (entityDomain.Transaction, error) {
	tx := repo.db.Save(&transaction)
	if tx.Error != nil {
		return entityDomain.Transaction{}, web.WebError{Code: 500, Message: tx.Error.Error()}
	}
	return transaction, nil
}

func (repo TransactionRepository) Delete(id int) error {
	tx := repo.db.Delete(&entityDomain.Transaction{}, id)
	if tx.Error != nil {
		return web.WebError{Code: 500, Message: tx.Error.Error()}
	}
	return nil
}
