package service

import (
	"context"
	"errors"
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
		GetProduct(inputSearchProduct entity.SearchInput) ([]map[string]interface{}, error)
		UpdateProduct(productId string, input entity.ProductInput) error
		DeleteProduct(input entity.ProductInput) error
		DetailProduct(input entity.IdUserInput) (map[string]interface{}, error)
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
		InStock:     input.InStock,
		UserID:      input.UserID,
		BrandID:     input.BrandID,
		CategoryID:  input.CategoryID,
		Image:       input.Image,
		// Slug:        input.Slug,
		CreatedAt: time.Time{},
	}

	product.Slug = slug.Make(input.Title)

	err := s.product_service.InsertData(product)

	if err != nil {
		log.Printf(err.Error())
		return err
	}

	return nil
}

func (s *productService) GetProduct(inputSearchProduct entity.SearchInput) ([]map[string]interface{}, error) {
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

	var products []map[string]interface{}
	if err != nil {
		log.Printf(err.Error())
		return products, err
	}

	for _, dataProduct := range result {
		var product map[string]interface{}

		getDataBrand, _ := s.user_service.GetBrandById(dataProduct.BrandID)
		dataProduct.BrandName = getDataBrand.Name

		getDataCategory, _ := s.user_service.GetCategoryById(dataProduct.CategoryID)
		dataProduct.CategoryName = getDataCategory.Name

		product = map[string]interface{}{
			"id":           dataProduct.ID,
			"brand_name":   getDataBrand.Name,
			"product_name": getDataCategory.Name,
			"slug":         dataProduct.Slug,
			"title":        dataProduct.Title,
			"price":        dataProduct.Price,
			"qty":          dataProduct.Stock,
			"image":        dataProduct.Image,
			"created_at":   dataProduct.CreatedAt,
		}

		products = append(products, product)

	}

	return products, nil

}

func (s *productService) UpdateProduct(productId string, input entity.ProductInput) error {
	product := model.Product{
		Title:       input.Title,
		Description: input.Description,
		SKU:         input.SKU,
		Price:       input.Price,
		CategoryID:  input.CategoryID,
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

func (s *productService) DetailProduct(input entity.IdUserInput) (map[string]interface{}, error) {

	dataProduct, err := s.product_service.GetDataByProductId(input.ProductId)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if dataProduct.ID == "" {
		return nil, errors.New("product not found")
	}

	getDataBrand, _ := s.user_service.GetBrandById(dataProduct.BrandID)
	dataProduct.BrandName = getDataBrand.Name

	getDataCategory, _ := s.user_service.GetCategoryById(dataProduct.CategoryID)
	dataProduct.CategoryName = getDataCategory.Name

	product := map[string]interface{}{
		"id":            dataProduct.ID,
		"brand_name":    getDataBrand.Name,
		"category_name": getDataCategory.Name,
		"product_name":  dataProduct.Title,
		"description":   dataProduct.Description,
		"sku":           dataProduct.SKU,
		"slug":          dataProduct.Slug,
		"title":         dataProduct.Title,
		"price":         dataProduct.Price,
		"qty":           dataProduct.Stock,
		"image":         dataProduct.Image,
		"created_at":    dataProduct.CreatedAt,
	}

	return product, err
}
