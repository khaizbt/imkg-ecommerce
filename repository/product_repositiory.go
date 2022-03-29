package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/khaizbt/imkg-ecommerce/helper"
	"github.com/khaizbt/imkg-ecommerce/model"
	"github.com/olivere/elastic"
	"time"
)

type ElasticRepository interface {
	GetData(ctx context.Context, sku, title string, brand, category int) ([]model.Product, error)
	InsertData(post model.Product) error
	UpdateProduct(productId string, post model.Product) error
}

func (r *elasticRepo) InsertData(post model.Product) error {
	product := map[string]interface{}{
		"title":       post.Title,
		"sku":         post.SKU,
		"description": post.Description,
		"price":       post.Price,
		"in_stock":    post.InStock,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"user_id":     post.UserID,
	}
	productId, _ := helper.GenerateNumber(6)

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

	if title != "" {
		bq.Must(elastic.NewMatchQuery("title", title))
	}

	if brand > 0 {
		fmt.Println(brand)
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

func (r *repository) GetCategoryIdByName(name string) (model.ProductCategory, error) {
	var result model.ProductCategory

	err := r.db.Where("name = ?", name).First(&result).Error

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
		"user_id":     post.UserID,
	}

	_, err := r.Client.Update().Index("products").Type("_doc").Id(productId).Doc(product).Do(context.TODO())

	if err != nil {
		return err
	}

	return nil
}

//TODO Stock dibuat di Elastic Aja, Mongo Khusus dipakai untuk Cart, Transaction
