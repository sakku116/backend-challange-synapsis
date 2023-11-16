package service

import (
	"synapsis/domain/model"
	"synapsis/exception"
	"synapsis/repository"
	"synapsis/utils/helper"
)

type CartService struct {
	cartRepo    repository.ICartRepo
	productRepo repository.IProductRepo
}

type ICartService interface {
	GetCartItems(user_id string) ([]model.ProductOrder, error)
}

func NewCartService(cartRepo repository.ICartRepo) *CartService {
	return &CartService{
		cartRepo: cartRepo,
	}
}

func (slf *CartService) GetCartItems(user_id string) ([]model.ProductOrder, error) {
	// get latest unchecked-out cart or create if it doesn't exist
	cart, err := slf.cartRepo.GetLast(false, user_id)
	var newCart *model.Cart
	if err != nil {
		if err == exception.DbObjNotFound {
			newCart = &model.Cart{
				ID:     helper.GenerateUUID(),
				UserID: user_id,
			}
			err = slf.cartRepo.Create(newCart)
			if err != nil {
				return []model.ProductOrder{}, err
			}
			cart = newCart
		} else {
			return []model.ProductOrder{}, err
		}
	}

	// get all associated product orders
	orders, err := slf.cartRepo.GetAssociatedProductOrders(cart.ID)
	if err != nil {
		return []model.ProductOrder{}, err
	}

	return orders, nil
}

func (slf *CartService) DeleteItemFromCart(product_id string, user_id string) error {
	return nil
}

func (slf *CartService) CheckoutCart(user_id string) error {
	return nil
}
