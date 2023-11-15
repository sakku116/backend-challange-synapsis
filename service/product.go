package service

import (
	"synapsis/domain/model"
	"synapsis/repository"
)

type ProductService struct {
	productRepo repository.IProductService
}

type IProductService interface {
	GetList(category string, search string, page int, limit int, sort_by string, sort_order string) ([]model.Product, error)
	GetCategoryList() ([]string, error)
	AddItemToCart(id string) error
}

func NewProductService(productRepo repository.IProductService) IProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

func (slf *ProductService) GetList(category string, search string, page int, limit int, sort_by string, sort_order string) ([]model.Product, error) {
	products, err := slf.productRepo.GetList(category, search, page, limit, sort_by, sort_order)
	if err != nil {
		return []model.Product{}, err
	}
	return products, nil
}

func (slf *ProductService) GetCategoryList() ([]string, error) {
	return nil, nil
}

func (slf *ProductService) AddItemToCart(id string) error {
	return nil
}
