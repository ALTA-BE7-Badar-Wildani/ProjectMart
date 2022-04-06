package cart

import (
	"go-ecommerce/entities/domain"
	"go-ecommerce/entities/web"
	productRepository "go-ecommerce/repositories/product"
	trRepository "go-ecommerce/repositories/transaction"
	trItemRepository "go-ecommerce/repositories/transaction_item"
	userRepository "go-ecommerce/repositories/user"
	"reflect"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/copier"
)

type CartService struct {
	trItemRepo trItemRepository.TransactionItemRepositoryInterface
	trRepo trRepository.TransactionRepositoryInterface
	userRepo userRepository.UserRepositoryInterface
	productRepo productRepository.ProductRepositoryInterface
}

func NewCartService(
	trItemRepo trItemRepository.TransactionItemRepositoryInterface, 
	trRepo trRepository.TransactionRepositoryInterface,
	userRepo userRepository.UserRepositoryInterface,
	productRepo productRepository.ProductRepositoryInterface,
) *CartService {
	return &CartService{
		trItemRepo: trItemRepo,
		trRepo: trRepo,
		userRepo: userRepo,
		productRepo: productRepo,
	}
}

/*
 * --------------------------
 * Get List of product 
 * --------------------------
 */
func (service CartService) FindAll(tokenReq interface{}) ([]web.CartResponse, error) {

	// Translate token
	token := tokenReq.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userIDReflect := reflect.ValueOf(claims).MapIndex(reflect.ValueOf("userID"))
	if reflect.ValueOf(userIDReflect.Interface()).Kind().String() != "float64" {
		return []web.CartResponse{}, web.WebError{ Code: 400, Message: "Invalid token, no userdata present" }
	}

	// get user data
	user, err := service.userRepo.Find(int(claims["userID"].(float64)))
	if err != nil {
		return []web.CartResponse{}, web.WebError{ Code: 400, Message: "No user matched with this authenticated user"}
	}

	

	// Get last User transaction (cart) 
	filters := []map[string]string {
		{
			"field": "user_id",
			"operator": "=",
			"value": strconv.Itoa(int(user.ID)),
		},
	}
	tr, err := service.trRepo.FindFirst(filters)
	if err != nil {
		webErr := err.(web.WebError)
		if webErr.Code == 400 {
			return []web.CartResponse{}, nil
		}
		return []web.CartResponse{}, err
	}


	// Get transaction items based on found transaction
	filters = []map[string]string {
		{
			"field": "transaction_id",
			"operator": "=",
			"value": strconv.Itoa(int(tr.ID)),
		},
	}
	trItems, err := service.trItemRepo.FindAll(0, 0, filters, []map[string]interface{}{})
	if err != nil {
		return []web.CartResponse{}, err
	}

	// Process to cart response
	cartsRes := []web.CartResponse{}
	copier.Copy(&cartsRes, &trItems)
	return cartsRes, err
}


/*
 * --------------------------
 *  Add to cart
 * --------------------------
 */
func (service CartService) Create(cartReq web.CartRequest, tokenReq interface{}) (web.CartResponse, error) {
	// Translate token
	token := tokenReq.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userIDReflect := reflect.ValueOf(claims).MapIndex(reflect.ValueOf("userID"))
	if reflect.ValueOf(userIDReflect.Interface()).Kind().String() != "float64" {
		return web.CartResponse{}, web.WebError{ Code: 400, Message: "Invalid token, no userdata present" }
	}

	// get user data
	user, err := service.userRepo.Find(int(claims["userID"].(float64)))
	if err != nil {
		return web.CartResponse{}, web.WebError{ Code: 400, Message: "No user matched with this authenticated user"}
	}

	// Get last User transaction (cart) 
	filters := []map[string]string {
		{
			"field": "user_id",
			"operator": "=",
			"value": strconv.Itoa(int(user.ID)),
		},
	}
	tr, err := service.trRepo.FindFirst(filters)
	if err != nil {
		webErr := err.(web.WebError)
		if webErr.Code != 400 {
			return web.CartResponse{}, err
		}
		// Make an empty transaction when there was none before
		tr, err = service.trRepo.Store(domain.Transaction{
			UserID: user.ID,
			Status: "CART",
		})
		if err != nil {
			return web.CartResponse{}, web.WebError{ Code: 500, Message: "Server error, cannot create transaction" }
		}
	}

	// get product data
	product, err := service.productRepo.Find(int(cartReq.ProductID))
	if err != nil {
		return web.CartResponse{}, err
	}

	// Convert request to domain entity
	trItem := domain.TransactionItem{}
	copier.Copy(&trItem, &cartReq)
	trItem.TransactionID = tr.ID
	trItem.ProductID = product.ID
	trItem.ProductPrice = int64(product.Price)
	trItem.SubTotal = int64(trItem.Qty) * int64(product.Price)

	// repository transaction item action
	trItem, err = service.trItemRepo.Store(trItem)
	if err != nil {
		return web.CartResponse{}, err
	}

	// get newly transaction item for getting preload data
	trItem, err = service.trItemRepo.Find(int(trItem.ID))
	if err != nil {
		return web.CartResponse{}, err
	}

	// Update transaction
	tr.TotalQty = tr.TotalQty + trItem.Qty
	tr.TotalPrice = tr.TotalPrice + trItem.SubTotal
	tr, err = service.trRepo.Update(tr, int(tr.ID))
	if err != nil {
		return web.CartResponse{}, err
	}

	// convert tr item to cart response
	cartRes := web.CartResponse{}
	copier.Copy(&cartRes, &trItem)

	return cartRes, nil
}