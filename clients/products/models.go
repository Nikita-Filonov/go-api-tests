package products

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
}

type GetProductsResponse struct {
	Skip     int       `json:"skip"`
	Limit    int       `json:"limit"`
	Total    int       `json:"total"`
	Products []Product `json:"products"`
}

type CreateProductRequest struct {
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
}
