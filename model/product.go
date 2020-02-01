package model

type Product struct {
	ID                 int64  `json:"productId"`
	ProductName        string `json:"productName"`
	ProductDescription string `json:"productDescription"`
	VendorName         string `json:"vendorName"`
}

type ProductService interface {
	Add(productName, productDescription, vendorName string) (*Product, error)
	Get(productID int64) (*Product, error)
	GetAll() ([]Product, error)
	Remove(productID int64) error
}
