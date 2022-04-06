package transaction

import entityDomain "go-ecommerce/entities/domain"

type TransactionRepositoryInterface interface {
	FindFirst(filters []map[string]string) (entityDomain.Transaction, error)
	Find(id int) (entityDomain.Transaction, error)
	Store(transaction entityDomain.Transaction) (entityDomain.Transaction, error)
	Update(transaction entityDomain.Transaction, id int) (entityDomain.Transaction, error)
	Delete(id int) error
}