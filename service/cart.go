package service

import (
	"synapsis/domain/model"
	"synapsis/exception"
	"synapsis/repository"
	error_utils "synapsis/utils/error"
	"synapsis/utils/helper"
)

type CartService struct {
	cartRepo         repository.ICartRepo
	productRepo      repository.IProductRepo
	productOrderRepo repository.IProductOrderRepo
}

type ICartService interface {
	GetCartItems(user_id string) ([]model.ProductOrder, error)
	RemoveItemFromCart(order_id string, user_id string) error
}

func NewCartService(cartRepo repository.ICartRepo, productRepo repository.IProductRepo, productOrderRepo repository.IProductOrderRepo) *CartService {
	return &CartService{
		cartRepo:         cartRepo,
		productRepo:      productRepo,
		productOrderRepo: productOrderRepo,
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

func (slf *CartService) RemoveItemFromCart(order_id string, user_id string) error {
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
				return err
			}
			cart = newCart
		} else {
			return err
		}
	}

	// check if product is in cart
	orders, err := slf.cartRepo.GetAssociatedProductOrders(cart.ID)
	if err != nil {
		return err
	}
	var orderTarget model.ProductOrder
	for _, order := range orders {
		if order.ID == order_id {
			orderTarget = order
			break
		}
	}
	if orderTarget.ID == "" {
		return &error_utils.CustomErr{
			Code:    404,
			Message: "product is not in cart",
		}
	}

	// remove order from cart association
	err = slf.cartRepo.RemoveOrderAssociations(cart.ID, []model.ProductOrder{orderTarget})
	if err != nil {
		return err
	}

	// delete order
	err = slf.productOrderRepo.Delete(orderTarget.ID)
	if err != nil {
		return err
	}

	return nil
}

func (slf *CartService) CheckoutCart(user_id string) error {
	return nil
}
