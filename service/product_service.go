package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gosimple/slug"
	"github.com/khaizbt/imkg-ecommerce/entity"
	"github.com/khaizbt/imkg-ecommerce/model"
	"github.com/khaizbt/imkg-ecommerce/repository"
)

type (
	productService struct {
		product_service repository.ElasticRepository
		user_service    repository.UserRepository
	}

	ProductService interface {
		InsertProduct(input entity.ProductInput) error
		GetProduct(inputSearchProduct entity.SearchInput) ([]model.Product, error)
		UpdateProduct(productId string, input entity.ProductInput) error
		DeleteProduct(input entity.ProductInput) error
	}
)

func NewProductService(productRepository repository.ElasticRepository, userService repository.UserRepository) *productService {
	return &productService{productRepository, userService}
}

func (s *productService) InsertProduct(input entity.ProductInput) error {
	product := model.Product{
		Title:       input.Title,
		Description: input.Description,
		SKU:         input.SKU,
		Stock:       input.Stock,
		Price:       input.Price,
		CategoryId:  input.CategoryID,
		InStock:     input.InStock,
		UserID:      input.UserID,
		CreatedAt:   time.Time{},
	}

	product.Slug = slug.Make(input.Title)

	err := s.product_service.InsertData(product)

	if err != nil {
		log.Printf(err.Error())
		return err
	}

	return nil
}

func (s *productService) GetProduct(inputSearchProduct entity.SearchInput) ([]model.Product, error) {
	fmt.Println(inputSearchProduct.Brand)

	var brand model.ProductBrand
	var category model.ProductCategory
	var err error
	if inputSearchProduct.Brand != "" {

		brand, err = s.user_service.GetBrandIdByName(inputSearchProduct.Brand)

		if err != nil {
			log.Printf(err.Error())
			return nil, err
		}
	}

	if inputSearchProduct.Category != "" {
		category, err = s.user_service.GetCategoryIdByName(inputSearchProduct.Category)

		if err != nil {
			log.Printf(err.Error())
			return nil, err
		}
	}

	var ctx context.Context

	result, err := s.product_service.GetData(ctx, inputSearchProduct.SKU, inputSearchProduct.Title, brand.ID, category.ID)

	if err != nil {
		log.Printf(err.Error())
		return result, err
	}

	return result, nil

}

func (s *productService) UpdateProduct(productId string, input entity.ProductInput) error {
	product := model.Product{
		Title:       input.Title,
		Description: input.Description,
		SKU:         input.SKU,
		Price:       input.Price,
		CategoryId:  input.CategoryID,
		InStock:     input.InStock,
	}

	err := s.product_service.UpdateProduct(productId, product)

	if err != nil {
		log.Printf(err.Error())
		return err
	}

	return nil
}

func (s *productService) DeleteProduct(input entity.ProductInput) error {
	err := s.product_service.DeleteProduct(input.ID)

	if err != nil {
		return err
	}

	return nil
}
