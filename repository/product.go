package repository

import (
	"synapsis/domain/model"
	"synapsis/exception"

	"gorm.io/gorm"
)

type ProductRepo struct {
	DB *gorm.DB
}

type IProductService interface {
	Create(product *model.Product) error
	GetList(category string, search string, page int, limit int, sort_by string, sort_order string) ([]model.Product, error)
	BulkCreate(products []model.Product) error
	GetByNameAndPrice(name string, price float64) (*model.Product, error)
}

func NewProductRepo(db *gorm.DB) IProductService {
	return &ProductRepo{
		DB: db,
	}
}

func (slf *ProductRepo) Create(product *model.Product) error {
	err := slf.DB.Create(product).Error
	if err == gorm.ErrRecordNotFound {
		return exception.DbObjNotFound
	} else if err != nil {
		return err
	}
	return nil
}

// GetList retrieves a list of products based on the given category, search keyword, page, limit, and sort order.
//
// Parameters:
// - category: a string representing the category of products to filter by. If empty, no category filter is applied.
// - search: a string representing the keyword to search for in the product names. If empty, no search filter is applied.
// - page: an integer representing the page number of the results to retrieve.
// - limit: an integer representing the maximum number of products per page.
// - sort_order: a string representing the sort order of the results. Can be "asc" or "desc".
//
// Returns:
// - products: a slice of model.Product containing the retrieved products.
// - error: an error if any occurred during the retrieval process.
func (slf *ProductRepo) GetList(category string, search string, page int, limit int, sort_by string, sort_order string) ([]model.Product, error) {
	var products []model.Product

	query := slf.DB.Model(&model.Product{})
	if category != "" {
		query = query.Joins("JOIN categories ON products.category_id = categories.id").
			Where("categories.name = ?", category)
	}
	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}
	err := query.Order("created_at " + sort_order).
		Limit(limit).
		Offset((page - 1) * limit).
		Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (slf *ProductRepo) BulkCreate(products []model.Product) error {
	err := slf.DB.Create(&products).Error
	if err != nil {
		return err
	}
	return nil
}

func (slf *ProductRepo) GetByNameAndPrice(name string, price float64) (*model.Product, error) {
	var product model.Product
	err := slf.DB.Where("name = ? AND price = ?", name, price).First(&product).Error
	if err == gorm.ErrRecordNotFound {
		return nil, exception.DbObjNotFound
	} else if err != nil {
		return nil, err
	}
	return &product, nil
}
