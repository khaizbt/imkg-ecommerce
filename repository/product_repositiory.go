package repository

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/khaizbt/imkg-ecommerce/helper"
	"github.com/khaizbt/imkg-ecommerce/model"
	"github.com/olivere/elastic"
)

type ElasticRepository interface {
	GetData(ctx context.Context, sku, title string, brand, category int) ([]model.Product, error)
	InsertData(post model.Product) error
	UpdateProduct(productId string, post model.Product) error
	GetDataByProductId(productId string) (model.Product, error)
	DeleteProduct(productId string) error
	UpdateStock(productId string, qty int) error
}

func (r *elasticRepo) InsertData(post model.Product) error {
	productId, _ := helper.GenerateNumber(6)
	product := map[string]interface{}{
		"id":          productId,
		"slug":        post.Slug,
		"title":       post.Title,
		"sku":         post.SKU,
		"description": post.Description,
		"price":       post.Price,
		"in_stock":    post.InStock,
		"user_id":     post.UserID,
		"stock":       post.Stock,
		"category_id": post.CategoryID,
		"brand_id":    post.BrandID,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
	}

	_, err := r.Client.Index().Index("products").Type("_doc").Id(productId).BodyJson(product).Do(context.TODO())

	if err != nil {
		return err
	}

	return nil
}

func (r *elasticRepo) GetData(ctx context.Context, sku, title string, brand, category int) ([]model.Product, error) {
	var results []model.Product

	query := elastic.NewMatchAllQuery()
	bq := elastic.NewBoolQuery().Must(query)
	// bq := elastic.NewMatchQuery("sku", sku)

	if title != "" {

		bq.Must(elastic.NewMatchQuery("title", title))
	}

	if brand > 0 {
		bq.Must(elastic.NewMatchQuery("brand_id", brand))
	}

	if category > 0 {
		bq.Must(elastic.NewMatchQuery("category", category))
	}

	if sku != "" {
		bq.Must(elastic.NewMatchQuery("sku", sku))

	}

	search, err := r.Client.Search().Index("products").Type("_doc").Query(bq).Do(context.TODO())

	if err != nil {
		return results, err
	}

	for _, hit := range search.Hits.Hits {
		var data model.Product
		err := json.Unmarshal(*hit.Source, &data)

		if err != nil {
			return results, err
		}

		results = append(results, data)
	}

	return results, nil
}

func (r *repository) GetBrandIdByName(name string) (model.ProductBrand, error) {
	var result model.ProductBrand

	err := r.db.Where("name = ?", name).First(&result).Error

	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *repository) GetBrandById(brandId int) (model.ProductBrand, error) {
	var result model.ProductBrand

	err := r.db.Where("id = ?", brandId).First(&result).Error

	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *repository) GetCategoryIdByName(name string) (model.ProductCategory, error) {
	var result model.ProductCategory

	err := r.db.Where("name = ?", name).First(&result).Error

	if err != nil {
		return result, err
	}

	return result, nil
}
func (r *repository) GetCategoryById(categoryId int) (model.ProductCategory, error) {
	var result model.ProductCategory

	err := r.db.Where("id = ?", categoryId).First(&result).Error

	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *elasticRepo) UpdateProduct(productId string, post model.Product) error {
	product := map[string]interface{}{
		"title":       post.Title,
		"sku":         post.SKU,
		"description": post.Description,
		"price":       post.Price,
		"in_stock":    post.InStock,
		"updated_at":  time.Now(),
	}

	_, err := r.Client.Update().Index(os.Getenv("ELASTIC_INDEX")).Type(os.Getenv("ELASTIC_TYPE")).Id(productId).Doc(product).Do(context.TODO())

	if err != nil {
		return err
	}

	return nil
}

func (r *elasticRepo) GetDataByProductId(productId string) (model.Product, error) {
	var result model.Product

	query := elastic.NewMatchQuery("_id", productId)

	search, err := r.Client.Search().Index(os.Getenv("ELASTIC_INDEX")).Type(os.Getenv("ELASTIC_TYPE")).Query(query).Do(context.TODO())

	if err != nil {
		return result, err
	}

	for _, hit := range search.Hits.Hits {
		var data model.Product
		err := json.Unmarshal(*hit.Source, &data)

		if err != nil {
			return result, err
		}

		result.ID = data.ID
		result.Title = data.Title
		result.UserID = data.UserID
		result.InStock = data.InStock
		result.Stock = data.Stock
		result.BrandID = data.BrandID
		result.CategoryID = data.CategoryID
		result.Price = data.Price
		result.SKU = data.SKU
		result.Description = data.Description
		result.Slug = data.Slug
		result.CreatedAt = data.CreatedAt
	}

	return result, nil
}

func (r *elasticRepo) DeleteProduct(productId string) error {
	_, err := r.Client.Delete().Index(os.Getenv("ELASTIC_INDEX")).Type(os.Getenv("ELASTIC_TYPE")).Id(productId).Do(context.TODO())

	if err != nil {
		return err
	}

	return nil
}

func (r *elasticRepo) UpdateStock(productId string, qty int) error {
	product := map[string]interface{}{
		"stock": qty,
	}

	_, err := r.Client.Update().Index("ELASTIC_INDEX").Type(os.Getenv("ELASTIC_TYPE")).Id(productId).Doc(product).Do(context.TODO())

	if err != nil {
		return err
	}

	return nil
}

//TODO Stock dibuat di Elastic Aja, Mongo Khusus dipakai untuk Cart, Transaction
