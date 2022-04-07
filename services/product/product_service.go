package product

import (
	"go-ecommerce/entities/domain"
	web "go-ecommerce/entities/web"
	productRepository "go-ecommerce/repositories/product"
	userRepository "go-ecommerce/repositories/user"
	"reflect"

	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/copier"
)

type ProductService struct {
	productRepo productRepository.ProductRepositoryInterface
	userRepo userRepository.UserRepositoryInterface
}

func NewProductService(repository productRepository.ProductRepositoryInterface, userRepository userRepository.UserRepositoryInterface) *ProductService {
	return &ProductService{
		productRepo: repository,
		userRepo: userRepository,
	}
}

/*
 * --------------------------
 * Get List of product 
 * --------------------------
 */
func (service ProductService) FindAll(limit, page int, filters []map[string]string, sorts []map[string]interface{}) ([]web.ProductResponse, error) {

	offset := (page - 1) * limit

	productsRes := []web.ProductResponse{}
	products, err := service.productRepo.FindAll(limit, offset, filters, sorts)
	copier.Copy(&productsRes, &products)
	return productsRes, err
}

/*
 * --------------------------
 * Load pagination data 
 * --------------------------
 */
func (service ProductService) GetPagination(limit, page int, filters []map[string]string) (web.Pagination, error) {
	totalRows, err := service.productRepo.CountAll(filters)
	if err != nil {
		return web.Pagination{}, err
	}
	totalPages :=  totalRows / int64(limit)
	if totalPages % int64(limit) > 0 {
		totalPages++
	}

	return web.Pagination{
		Page: page,
		Limit: limit,
		TotalPages: int(totalPages),
	}, nil
}

/*
 * --------------------------
 * Get single product data based on ID
 * --------------------------
 */
func (service ProductService) Find(id int) (web.ProductResponse, error) {
	
	product, err := service.productRepo.Find(id)
	productRes  := web.ProductResponse{}
	copier.Copy(&productRes, &product)

	return productRes, err
}


/*
 * --------------------------
 * Create product resource
 * --------------------------
 */
func (service ProductService) Create(productRequest web.ProductRequest, tokenReq interface{}) (web.ProductResponse, error) {
	// convert product to domain entity
	product := domain.Product{}
	copier.Copy(&product, &productRequest)

	// Translate token
	token := tokenReq.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userIDReflect := reflect.ValueOf(claims).MapIndex(reflect.ValueOf("userID"))
	if reflect.ValueOf(userIDReflect.Interface()).Kind().String() != "float64" {
		return web.ProductResponse{}, web.WebError{ Code: 400, Message: "Invalid token, no userdata present" }
	}

	// get user data
	user, err := service.userRepo.Find(int(claims["userID"].(float64)))
	if err != nil {
		return web.ProductResponse{}, web.WebError{ Code: 400, Message: "No user matched with this authenticated user"}
	}
	product.UserID = user.ID

	// repository action
	product, err = service.productRepo.Store(product)
	if err != nil {
		return web.ProductResponse{}, err
	}

	// get product data
	productRes, err := service.Find(int(product.ID))
	if err != nil {
		return web.ProductResponse{}, web.WebError{ Code: 500, Message: "Cannot get newly created product" }
	}

	return productRes, nil
}


/*
 * --------------------------
 * Update product resource
 * --------------------------
 */
func (service ProductService) Update(productRequest web.ProductRequest, id int, tokenReq interface{}) (web.ProductResponse, error) {

	// Find product
	product, err := service.productRepo.Find(id)
	if err != nil {
		return web.ProductResponse{}, web.WebError{ Code: 400, Message: "The requested ID doesn't match with any record" }
	}

	// Translate token
	token := tokenReq.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userIDReflect := reflect.ValueOf(claims).MapIndex(reflect.ValueOf("userID"))
	if reflect.ValueOf(userIDReflect.Interface()).Kind().String() != "float64" {
		return web.ProductResponse{}, web.WebError{ Code: 400, Message: "Invalid token, no userdata present" }
	}

	// get user data
	user, err := service.userRepo.Find(int(claims["userID"].(float64)))
	if err != nil {
		return web.ProductResponse{}, web.WebError{ Code: 400, Message: "No user matched with this authenticated user"}
	}
	if product.UserID != user.ID {
		return web.ProductResponse{}, web.WebError{ Code: 401, Message: "Cannot update product that belongs to someone else" }
	}

	// Copy request to found product
	copier.CopyWithOption(&product, &productRequest, copier.Option{ IgnoreEmpty: true, DeepCopy: true })

	// repository action
	product, err = service.productRepo.Update(product, id)
	if err != nil {
		return web.ProductResponse{}, err
	}

	// get product data
	productRes, err := service.Find(int(product.ID))
	if err != nil {
		return web.ProductResponse{}, web.WebError{ Code: 500, Message: "Cannot get newly created product" }
	}

	return productRes, err
}

/*
 * --------------------------
 * Delete resource data 
 * --------------------------
 */
func (service ProductService) Delete(id int, tokenReq interface{}) error {
	// Find product
	product, err := service.productRepo.Find(id)
	if err != nil {
		return web.WebError{ Code: 400, Message: "The requested ID doesn't match with any record" }
	}

	// Translate token
	token := tokenReq.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userIDReflect := reflect.ValueOf(claims).MapIndex(reflect.ValueOf("userID"))
	if reflect.ValueOf(userIDReflect.Interface()).Kind().String() != "float64" {
		return web.WebError{ Code: 400, Message: "Invalid token, no userdata present" }
	}

	// get user data
	user, err := service.userRepo.Find(int(claims["userID"].(float64)))
	if err != nil {
		return web.WebError{ Code: 400, Message: "No user matched with this authenticated user"}
	}
	if product.UserID != user.ID {
		return web.WebError{ Code: 401, Message: "Cannot update product that belongs to someone else" }
	}
	
	// Repository action
	err = service.productRepo.Delete(id)
	return err
}