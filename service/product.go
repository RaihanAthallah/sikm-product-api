package service

import (
	"product-api/apperror"
	"product-api/model"
	"product-api/repository"
	"strings"
)

var (
	blackListedWords = []string{
		"termurah",
		"terbaik",
		"diskon",
		"promo",
	}
)

type ProductService interface {
	GetList() ([]model.Product, error)
	GetByID(id int) (*model.ProductDetail, error)
	Store(*model.ProductRequest) error
	Delete(*model.Product) error
	Update(*model.Product) error
}

type productServiceImpl struct {
	productRepo repository.ProductRepository
}

func NewProductService(prdRepo repository.ProductRepository) ProductService {
	return &productServiceImpl{
		productRepo: prdRepo,
	}
}

func (s *productServiceImpl) GetList() ([]model.Product, error) {
	return s.productRepo.FindAll()
}

func (s *productServiceImpl) GetByID(id int) (*model.ProductDetail, error) {
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &model.ProductDetail{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		User:     product.User,
		Category: product.Category,
	}, nil
}

func (s *productServiceImpl) isValidName(name string) bool {
	for _, word := range blackListedWords {
		if strings.Contains(strings.ToLower(name), word) {
			return false
		}
	}
	return true
}

func (s *productServiceImpl) Store(payload *model.ProductRequest) error {
	if !s.isValidName(payload.Name) {
		return apperror.ErrInvalidProductName
	}

	product := model.Product{
		Name:       payload.Name,
		Price:      payload.Price,
		CategoryID: payload.CategoryID,
		PostedBy:   uint(payload.PostedBy),
	}
	return s.productRepo.Insert(&product)
}

func (s *productServiceImpl) Update(*model.Product) error {
	// TODO: add code
	return nil
}

func (s *productServiceImpl) Delete(*model.Product) error {
	// TODO: add code
	return nil
}
