package products

import "log"

// IService service abstractions
type IService interface {
	GetProducts(req GetProductsRequest) (GetProductsResponse, error)
	GetProductByID(req GetProductByIDRequest) (GetProductByIDResponse, error)
}

// Service implements IService
type Service struct {
	Logger     *log.Logger
	Repository IRepository
}

// GetProductsRequest request
type GetProductsRequest struct {
	Limit uint16 `json:"limit,omitempty"`
}

// GetProductsResponse response
type GetProductsResponse struct {
	Products []Product `json:"products,omitempty"`
}

// GetProductByIDRequest request
type GetProductByIDRequest struct {
	ID uint64 `json:"id,omitempty"`
}

// GetProductByIDResponse response
type GetProductByIDResponse struct {
	Product Product `json:"product,omitempty"`
}

// NewProductsService initialize service for products
func NewProductsService(logger *log.Logger, repo IRepository) *Service {
	return &Service{
		Logger:     logger,
		Repository: repo,
	}
}

// GetProducts get products
func (s *Service) GetProducts(req GetProductsRequest) (GetProductsResponse, error) {
	products, err := s.Repository.GetProducts(req.Limit)
	if err != nil {
		s.Logger.Printf("[%s] %v", "ERROR", err)
	}
	resp := GetProductsResponse{Products: products}
	return resp, err
}

// GetProductByID get product by id
func (s *Service) GetProductByID(req GetProductByIDRequest) (GetProductByIDResponse, error) {
	product, err := s.Repository.GetProductByID(req.ID)
	if err != nil {
		s.Logger.Printf("[%s] %v", "ERROR", err)
	}
	resp := GetProductByIDResponse{Product: product}
	return resp, err
}
