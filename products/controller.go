package products

import (
	"net/http"
	"strconv"

	"github.com/satriarrrrr/store/helpers"
	"goji.io/pat"
)

// IController interface controller
type IController interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
}

// Controller implements IController
type Controller struct {
	Service IService
}

// NewProductsController create a new controller
func NewProductsController(service IService) *Controller {
	return &Controller{Service: service}
}

// GetProducts return list of products
func (c *Controller) GetProducts(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.ParseUint(r.FormValue("limit"), 10, 64)
	if limit == 0 {
		limit = 10
	}
	req := GetProductsRequest{Limit: uint16(limit)}
	resp, err := c.Service.GetProducts(req)
	if err != nil {
		helpers.ResponseJSON(w, "Failed to load products", 9999, 200)
	} else {
		helpers.ResponseJSON(w, resp, 1000, 200)
	}
	return
}

// GetProductByID get product by id
func (c *Controller) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(pat.Param(r, "id"), 10, 64)
	if err != nil {
		helpers.ResponseJSON(w, "Invalid input", 9999, 200)
	}
	req := GetProductByIDRequest{ID: id}
	resp, err := c.Service.GetProductByID(req)
	if err != nil {
		helpers.ResponseJSON(w, "Failed to load product", 9999, 200)
	} else {
		helpers.ResponseJSON(w, resp, 1000, 200)
	}
	return
}
