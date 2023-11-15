package rest

import (
	"synapsis/domain/model"
	"time"

	"gorm.io/gorm"
)

type GetProductListResp struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

func (slf *GetProductListResp) ParseFromEntityList(products *[]model.Product) []*GetProductListResp {
	var result []*GetProductListResp
	for _, product := range *products {
		result = append(result, &GetProductListResp{
			ID:        product.ID,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
			DeletedAt: product.DeletedAt,
			Name:      product.Name,
			Price:     product.Price,
			Category:  product.Category,
		})
	}

	return result
}
