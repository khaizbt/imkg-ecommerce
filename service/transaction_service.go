package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/khaizbt/imkg-ecommerce/entity"
	"github.com/khaizbt/imkg-ecommerce/model"
	"github.com/khaizbt/imkg-ecommerce/repository"
)

type (
	transactionService struct {
		product_service     repository.ElasticRepository
		user_service        repository.UserRepository
		transaction_service repository.TransactionRepository
	}

	TransactionService interface {
		AddToCart(input entity.CartInput) error
		ListCart(userId int) ([]model.Cart, error)
		UpdateCart(input entity.CartInput, userId int) error
		DeleteCart(input entity.CartInput) error
	}
)

func NewTransactionService(productRepository repository.ElasticRepository, userService repository.UserRepository, trancationService repository.TransactionRepository) *transactionService {
	return &transactionService{productRepository, userService, trancationService}
}

func (s *transactionService) AddToCart(input entity.CartInput) error {
	getDataProduct, err := s.product_service.GetDataByProductId(input.ProductID)

	if err != nil {
		log.Printf(err.Error())
		return err
	}

	if getDataProduct.Stock < input.Qty {
		return errors.New("Stock Kosong!")
	}

	cart := model.Cart{
		ProductID: input.ProductID,
		Qty:       input.Qty,
		UserID:    input.UserID,
	}

	err = s.transaction_service.AddToCart(cart)

	if err != nil {
		log.Printf(err.Error())
		return err
	}

	return nil
}

func (s *transactionService) ListCart(userId int) ([]model.Cart, error) {
	getDataCart, err := s.transaction_service.ListCart(userId)

	if err != nil {
		return nil, err
	}

	return getDataCart, nil

}

func (s *transactionService) UpdateCart(input entity.CartInput, userId int) error {
	cartData, err := s.transaction_service.DetailCart(input.CartId)

	if err != nil {
		return err
	}

	if cartData.UserID != userId {
		return errors.New("Cart Not Found")
	}

	fmt.Println(cartData)
	getDataProduct, err := s.product_service.GetDataByProductId(cartData.ProductID)

	if err != nil {
		log.Printf(err.Error())
		return err
	}

	fmt.Println(getDataProduct)

	if getDataProduct.Stock < input.Qty {
		return errors.New("Stock Kosong!")
	}

	err = s.transaction_service.UpdateCart(input.CartId, input.Qty)

	if err != nil {
		return err
	}

	return nil
}

func (s *transactionService) DeleteCart(input entity.CartInput) error {
	cartData, err := s.transaction_service.DetailCart(input.CartId)

	if err != nil {
		return err
	}

	if cartData.UserID != input.UserID {
		return errors.New("Cart Not Found")
	}

	err = s.transaction_service.DeleteCart(input.CartId)

	if err != nil {
		log.Printf(err.Error())
		return err
	}

	return nil
}

//TODO jika add Cart dan ada data Product IDnya maka update qty jika tidak ada maka create baru
