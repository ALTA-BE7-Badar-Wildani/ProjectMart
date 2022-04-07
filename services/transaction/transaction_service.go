package transaction

import (
	"go-ecommerce/entities/web"
	trRepo "go-ecommerce/repositories/transaction"
	userRepo "go-ecommerce/repositories/user"
	"reflect"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/copier"
)

type TransactionService struct {
	trRepo trRepo.TransactionRepositoryInterface
	userRepo userRepo.UserRepositoryInterface
}

func NewTransactionService(trRepo trRepo.TransactionRepositoryInterface, userRepo userRepo.UserRepositoryInterface) *TransactionService {
	return &TransactionService{
		trRepo: trRepo,
		userRepo: userRepo,
	}
}

/*
 * --------------------------
 * Get List of transaction 
 * --------------------------
 */
func (service TransactionService) FindAll(limit, page int, filters []map[string]string, sorts []map[string]interface{}, tokenReq interface{}) ([]web.TransactionResponse, error) {

	// Translate token
	token := tokenReq.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userIDReflect := reflect.ValueOf(claims).MapIndex(reflect.ValueOf("userID"))
	if reflect.ValueOf(userIDReflect.Interface()).Kind().String() != "float64" {
		return []web.TransactionResponse{}, web.WebError{ Code: 400, Message: "Invalid token, no userdata present" }
	}
	
	// get user data
	user, err := service.userRepo.Find(int(claims["userID"].(float64)))
	if err != nil {
		return []web.TransactionResponse{}, web.WebError{ Code: 400, Message: "No user matched with this authenticated user"}
	}

	// Offset calculation
	offset := (page - 1) * limit

	// repository action
	filters = append(filters, map[string]string {
		"field": "status",
		"operator": "!=",
		"value": "CART",
	})
	filters = append(filters, map[string]string {
		"field": "user_id",
		"operator": "=",
		"value": strconv.Itoa(int(user.ID)),
	})
	transactions, err := service.trRepo.FindAll(limit, offset, filters, sorts)

	// process to response
	transactionsRes := []web.TransactionResponse{}
	copier.Copy(&transactionsRes, &transactions)
	return transactionsRes, err
}

/*
 * --------------------------
 * Load pagination data 
 * --------------------------
 */
func (service TransactionService) GetPagination(page, limit int, filters []map[string]string, tokenReq interface{}) (web.Pagination, error) {
	// Translate token
	token := tokenReq.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userIDReflect := reflect.ValueOf(claims).MapIndex(reflect.ValueOf("userID"))
	if reflect.ValueOf(userIDReflect.Interface()).Kind().String() != "float64" {
		return web.Pagination{}, web.WebError{ Code: 400, Message: "Invalid token, no userdata present" }
	}
	
	// get user data
	user, err := service.userRepo.Find(int(claims["userID"].(float64)))
	if err != nil {
		return web.Pagination{}, web.WebError{ Code: 400, Message: "No user matched with this authenticated user"}
	}
	// repository action
	filters = append(filters, map[string]string {
		"field": "status",
		"operator": "!=",
		"value": "CART",
	})
	filters = append(filters, map[string]string {
		"field": "user_id",
		"operator": "=",
		"value": strconv.Itoa(int(user.ID)),
	})
	totalRows, err := service.trRepo.CountAll(filters)
	if err != nil {
		return web.Pagination{}, err
	}
	totalPages :=  totalRows / int64(limit)
	if totalPages <= 0 {
		totalPages = 1
	}

	return web.Pagination{
		Page: page,
		Limit: limit,
		TotalPages: int(totalPages),
	}, nil
}
