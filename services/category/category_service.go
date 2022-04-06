package product

import (
	"go-ecommerce/entities/domain"
	web "go-ecommerce/entities/web"
	productRepository "go-ecommerce/repositories/product"

	"github.com/jinzhu/copier"
)

type ProductService struct {
	productRepo productRepository.ProductRepositoryInterface
}

func NewProductService(repository productRepository.ProductRepositoryInterface) *ProductService {
	return &ProductService{
		productRepo: repository,
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
func (service ProductService) GetPagination(page, limit int) (web.Pagination, error) {
	totalRows, err := service.productRepo.CountAll()
	if err != nil {
		return web.Pagination{}, err
	}
	totalPages :=  totalRows / int64(limit)

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
func (service ProductService) Create(productRequest web.ProductRequest) (web.ProductResponse, error) {
	
	product := domain.Product{}
	copier.Copy(&product, &productRequest)

	product, err := service.productRepo.Store(product)
	if err != nil {
		return web.ProductResponse{}, err
	}

	productRes := web.ProductResponse{}
	copier.Copy(&productRes, &product)

	return productRes, nil
}


/*
 * --------------------------
 * Update product resource
 * --------------------------
 */
func (service ProductService) Update(productRequest web.ProductRequest, id int) (web.ProductResponse, error) {

	// Find product
	product, err := service.productRepo.Find(id)
	if err != nil {
		return web.ProductResponse{}, web.WebError{ Code: 400, Message: "The requested ID doesn't match with any record" }
	}

	product, err = service.productRepo.Update(product, id)

	// Convert product domain to product response
	productRes := web.ProductResponse{}
	copier.Copy(&productRes, &product)

	return productRes, err
}

/*
 * --------------------------
 * Delete resource data 
 * --------------------------
 */
func (service ProductService) Delete(id int) error {
	// Find product
	_, err := service.productRepo.Find(id)
	if err != nil {
		return web.WebError{ Code: 400, Message: "The requested ID doesn't match with any record" }
	}
	
	// Copy request to found product
	err = service.productRepo.Delete(id)
	return err
}