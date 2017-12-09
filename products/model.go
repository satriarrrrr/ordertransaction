package products

// Product represents data from table products
type Product struct {
	ID       uint64 `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Quantity uint32 `json:"quantity,omitempty"`
	Price    uint32 `json:"price,omitempty"`
}
