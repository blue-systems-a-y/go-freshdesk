package freshservice

import "fmt"

//Product ...
type Product struct {
	//Unique ID of the product.
	ID int `json:"id,omitempty"`
	//Name of the Product.
	Name string `json:"name,omitempty"`
	//ID of the asset type(must be a type in Freshservice)
	AssetTypeID int `json:"asset_type_id,omitempty"`
	//Manufacturer of the product
	Manufacturer string `json:"manufacturer,omitempty"`
	//Status of the product.
	Status string `json:"status,omitempty"`
	//Mode of procurement of the product.
	ModeOfProcurement string `json:"mode_of_procurement,omitempty"`
	//ID of the depreciation type.
	DepreciationTypeID int `json:"depreciation_type_id,omitempty"`
	//HTML Content of the product.
	Description string `json:"description,omitempty"`
	//Description of the product in plain text.
	DescriptionText string `json:"description_text,omitempty"`
	//Date and time when the product was created
	CreatedAt string `json:"created_at,omitempty"`
	//Date and time when the product was last updated
	UpdatedAt string `json:"updated_at,omitempty"`
}

//ProductItem holds one prouct
type ProductItem struct {
	Product Product `json:"product,omitempty"`
}

//ProuctList holds a list of products
type ProuctList struct {
	Products []Product `json:"products,omitempty"`
}

//ListProducts get products by page
func (c *Client) ListProducts(page int) ([]Product, error) {
	var list ProuctList
	err := c.ReadObject(fmt.Sprintf("/api/v2/products?page=%v", page), &list)
	return list.Products, err
}

//ListAllProducts read all products
func (c *Client) ListAllProducts() ([]Product, error) {
	all := make([]Product, 0)
	i := 1
	for true {
		page, err := c.ListProducts(i)
		if err != nil {
			return nil, err
		}
		if len(page) == 0 {
			break
		}
		all = append(all, page...)
		i++
	}
	return all, nil
}

//CreateProduct creates a procuct
func (c *Client) CreateProduct(prod *Product) (*Product, error) {
	var out ProductItem
	err := c.WriteObject("/api/v2/products", "POST", prod, &out)
	if err != nil {
		return nil, err
	}
	return &out.Product, nil
}
