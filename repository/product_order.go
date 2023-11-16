package repository

import (
	"synapsis/domain/model"
	"synapsis/exception"

	"gorm.io/gorm"
)

type ProductOrderRepo struct {
	db *gorm.DB
}

type IProductOrderRepo interface {
	Create(productOrder *model.ProductOrder) error
	Update(productOrder *model.ProductOrder) error
	GetByID(id string) (*model.ProductOrder, error)
}

func NewProductOrderRepo(db *gorm.DB) IProductOrderRepo {
	return &ProductOrderRepo{
		db: db,
	}
}

func (slf *ProductOrderRepo) Create(productOrder *model.ProductOrder) error {
	err := slf.db.Create(productOrder).Error
	if err != nil {
		return err
	}
	return nil
}

func (slf *ProductOrderRepo) Update(productOrder *model.ProductOrder) error {
	err := slf.db.Save(productOrder).Error
	if err != nil {
		return err
	}
	return nil
}

func (slf *ProductOrderRepo) GetByID(id string) (*model.ProductOrder, error) {
	var productOrder model.ProductOrder
	err := slf.db.Where("id = ?", id).First(&productOrder).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.DbObjNotFound
		} else {
			return nil, err
		}
	}
	return &productOrder, nil
}
