package repository

import (
	"synapsis/domain/model"
	"synapsis/exception"

	"gorm.io/gorm"
)

type CartRepo struct {
	DB *gorm.DB
}

type ICartRepo interface {
	Create(cart *model.Cart) error
	GetLast(is_checkout bool, user_id string) (*model.Cart, error)
	GetAssociatedProductOrders(id string) ([]model.ProductOrder, error)
	AppendProductOrders(id string, orders []model.ProductOrder) error
}

func NewCartRepo(db *gorm.DB) ICartRepo {
	return &CartRepo{
		DB: db,
	}
}

func (slf *CartRepo) Create(cart *model.Cart) error {
	err := slf.DB.Create(cart).Error
	if err != nil {
		return err
	}
	return nil
}

func (slf *CartRepo) GetLast(is_checkout bool, user_id string) (*model.Cart, error) {
	var cart model.Cart
	err := slf.DB.Model(&model.Cart{}).
		Preload("ProductOrders").
		Order("updated_at desc").
		Where("user_id = ?", user_id).
		Where("is_checkout = ?", is_checkout).
		First(&cart).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.DbObjNotFound
		} else {
			return nil, err
		}
	}
	return &cart, nil
}

func (slf *CartRepo) GetAssociatedProductOrders(id string) ([]model.ProductOrder, error) {
	var productOrders []model.ProductOrder
	err := slf.DB.Model(&model.Cart{ID: id}).
		Preload("Product"). // actually its preloading from ProductOrders instead Cart (depending on result model / not from table where operation is performed (cart))
		Association("ProductOrders").
		Find(&productOrders)
	if err != nil {
		return []model.ProductOrder{}, err
	}
	return productOrders, nil
}

func (slf *CartRepo) AppendProductOrders(id string, orders []model.ProductOrder) error {
	err := slf.DB.Model(&model.Cart{ID: id}).Association("ProductOrders").Append(orders)
	if err != nil {
		return err
	}
	return nil
}
