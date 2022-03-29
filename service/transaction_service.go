package service

import (
	"github.com/khaizbt/imkg-ecommerce/entity"
	"github.com/khaizbt/imkg-ecommerce/model"
	"github.com/khaizbt/imkg-ecommerce/repository"
	"log"
)

type (
	transactionService struct {
		product_service     repository.ElasticRepository
		user_service        repository.UserRepository
		transaction_service repository.TransactionRepository
	}

	TransactionService interface {
		AddToCart(input entity.CartInput) error
	}
)

func NewTransactionService(productRepository repository.ElasticRepository, userService repository.UserRepository, trancationService repository.TransactionRepository) *transactionService {
	return &transactionService{productRepository, userService, trancationService}
}

func (s *transactionService) AddToCart(input entity.CartInput) error {
	cart := model.Cart{
		ItemName: "Gateron Blue Switch",
		Seller:   "Admin",
		Qty:      input.Qty,
		UserID:   input.UserID,
	}

	err := s.transaction_service.AddToCart(cart)

	if err != nil {
		log.Printf(err.Error())
		return err
	}

	return nil
}

//TODO jika add Cart dan ada data Product IDnya maka update qty jikatidak ada maka create baru
