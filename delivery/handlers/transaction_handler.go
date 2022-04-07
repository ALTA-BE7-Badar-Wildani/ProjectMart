package handlers

import (
	"go-ecommerce/config"
	"go-ecommerce/delivery/helpers"
	"go-ecommerce/entities/web"
	transactionService "go-ecommerce/services/transaction"
	"reflect"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionService *transactionService.TransactionService
}

func NewTransactionHandler(service *transactionService.TransactionService) TransactionHandler {
	return TransactionHandler{
		transactionService: service,
	}
}


/*
 * -------------------------------------------
 * Show All transactions based on available queries
 * -------------------------------------------
 */
func (handler TransactionHandler) Index(c echo.Context) error {

	// Translate query param to map of filters
	filters := []map[string]string{}
	dateStart := c.QueryParam("dateStart") 
	dateEnd := c.QueryParam("dateEnd") 
	if dateStart != "" {
		filters = append(filters, map[string]string{
			"field": "created_at",
			"operator": ">=",
			"value": dateStart,
		})
	}
	if dateEnd != "" {
		filters = append(filters, map[string]string{
			"field": "created_at",
			"operator": "<=",
			"value": dateEnd,
		})
	}
	status := c.QueryParam("status")
	if status != "" {
		filters = append(filters, map[string]string{
			"field": "status",
			"operator": "=",
			"value": status,
		})
	}

	// Sort parameter
	sorts := []map[string]interface{} {}
	sortDate := c.QueryParam("sortDate") 
	if sortDate != "" {
		switch sortDate {
		case "1":
			sorts = append(sorts, map[string]interface{} {
				"field": "created_at",
				"desc": true,
			})
		case "0":
			sorts = append(sorts, map[string]interface{} {
				"field": "created_at",
				"desc": false,
			})
		}
	}

	links := map[string]string{}
	links["self"] = config.Get().App.BaseUrl + "/api/transactions?limit=" + c.QueryParam("limit") + "&page=" + c.QueryParam("page")

	// pagination param
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(400, helpers.MakeErrorResponse("ERROR", 400, "Limit Parameter format is invalid", links))
	}
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		links := map[string]string {"self": config.Get().App.BaseUrl}
		return c.JSON(400, helpers.MakeErrorResponse("ERROR", 400, "page Parameter format is invalid", links))
	}


	// Get all transactions
	token := c.Get("user")
	transactionsRes, err := handler.transactionService.FindAll(limit, page, filters, sorts, token)
	if err != nil {
		if reflect.TypeOf(err).String() == "web.WebError" {
			webErr := err.(web.WebError)
			return c.JSON(webErr.Code, helpers.MakeErrorResponse("ERROR", webErr.Code, webErr.Error(), links))
		}
		return c.JSON(500, helpers.MakeErrorResponse("ERROR", 500, err.Error(), links))
	}
	
	// Get pagination data
	pagination, err := handler.transactionService.GetPagination(page, limit, filters, token)
	if err != nil {
		if reflect.TypeOf(err).String() == "web.WebError" {
			webErr := err.(web.WebError)
			return c.JSON(webErr.Code, helpers.MakeErrorResponse("ERROR", webErr.Code, webErr.Error(), links))
		}
		return c.JSON(500, helpers.MakeErrorResponse("ERROR", 500, err.Error(), links))
	}

	links["first"] = config.Get().App.BaseUrl + "/api/transactions?limit=" + c.QueryParam("limit") + "&page=1"
	links["last"] = config.Get().App.BaseUrl + "/api/transactions?limit=" + c.QueryParam("limit") + "&page=" + strconv.Itoa(pagination.TotalPages)
	if pagination.Page > 1 {
		links["prev"] = config.Get().App.BaseUrl + "/api/transactions?limit=" + c.QueryParam("limit") + "&page=" + strconv.Itoa(pagination.Page - 1)
	}
	if pagination.Page < pagination.TotalPages {
		links["next"] = config.Get().App.BaseUrl + "/api/transactions?limit=" + c.QueryParam("limit") + "&page=" + strconv.Itoa(pagination.Page + 1)
	}
	
	// success response
	return c.JSON(200, web.SuccessListResponse{
		Status: "OK",
		Code: 200,
		Error: nil,
		Links: links,
		Data: transactionsRes,
		Pagination: pagination,
	})
}
