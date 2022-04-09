package transaction_item

import entityDomain "go-ecommerce/entities/domain"

type TransactionItemRepositoryInterface interface {
	FindAll(limit int, offset int, filters []map[string]string, sorts []map[string]interface{}) ([]entityDomain.TransactionItem, error)
	Find(id int) (entityDomain.TransactionItem, error)
	FindBy(field string, value string) (entityDomain.TransactionItem, error)
	FindFirst(filters []map[string]string) (entityDomain.TransactionItem, error)
	CountAll() (int64, error)
	Store(transactionItem entityDomain.TransactionItem) (entityDomain.TransactionItem, error)
	Update(transactionItem entityDomain.TransactionItem, id int) (entityDomain.TransactionItem, error)
	Delete(id int) error
}