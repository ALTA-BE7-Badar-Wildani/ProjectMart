package transaction

import entityDomain "go-ecommerce/entities/domain"

type TransactionRepositoryInterface interface {
	FindAll(limit int, offset int, filters []map[string]string, sorts []map[string]interface{}) ([]entityDomain.Transaction, error)
	CountAll(filters []map[string]string) (int64, error)
	FindFirst(filters []map[string]string) (entityDomain.Transaction, error)
	Find(id int) (entityDomain.Transaction, error)
	Store(transaction entityDomain.Transaction) (entityDomain.Transaction, error)
	Update(transaction entityDomain.Transaction, id int) (entityDomain.Transaction, error)
	Delete(id int) error
}