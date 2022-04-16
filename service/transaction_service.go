package service

import (
	"errors"
	"log"
	"time"

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
		//Cart
		AddToCart(input entity.CartInput) error
		ListCart(userId int) ([]map[string]interface{}, error)
		UpdateCart(input entity.CartInput, userId int) error
		DeleteCart(input entity.CartInput) error

		//Transaction
		Checkout(input entity.Checkout) error
		ListTransaction(userId int) ([]model.Transaction, error)
		// ListSalesTransaction() ([]model.Transaction, error)
		UpdateStatus(input entity.TransactionStatus) error
		ListSalesTransaction() ([]map[string]interface{}, error)
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

func (s *transactionService) ListCart(userId int) ([]map[string]any, error) {
	getDataCart, err := s.transaction_service.ListCart(userId)

	if err != nil {
		return nil, err
	}
	var carts []map[string]interface{}

	//Map Data Cart
	for _, dataCart := range getDataCart {
		var cart map[string]interface{}
		detailProduct, _ := s.product_service.GetDataByProductId(dataCart.ProductID)

		dataCart.ProductName = detailProduct.Title

		cart = map[string]interface{}{
			"_id":          dataCart.ID,
			"product_id":   dataCart.ProductID,
			"product_name": dataCart.ProductName,
			"qty":          dataCart.Qty,
			"created_at":   dataCart.CreatedAt,
		}

		carts = append(carts, cart)
	}

	return carts, nil

}

func (s *transactionService) UpdateCart(input entity.CartInput, userId int) error {
	cartData, err := s.transaction_service.DetailCart(input.CartId)

	if err != nil {
		return err
	}

	if cartData.UserID != userId {
		return errors.New("Cart Not Found")
	}

	getDataProduct, err := s.product_service.GetDataByProductId(cartData.ProductID)

	if err != nil {
		log.Printf(err.Error())
		return err
	}

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

//TODO jika Checkout maka mapping data id cart yang di checkout lalu delete 1-1
func (s *transactionService) Checkout(input entity.Checkout) error {
	var cartData model.Transaction
	var cartList []model.TransactionDetail

	var carts []string
	for _, cartId := range input.CheckoutDetail {

		dataCart, err := s.transaction_service.DetailCart(cartId.CartId)

		var cart model.TransactionDetail

		cart.ProductID = dataCart.ProductID
		cart.Qty = dataCart.Qty
		if err != nil {
			log.Printf(err.Error())
			return err
		}

		cartList = append(cartList, cart)

		getProductData, err := s.product_service.GetDataByProductId(dataCart.ProductID)

		if err != nil {
			log.Printf(err.Error())
			return err
		}

		if getProductData.Stock < dataCart.Qty { //Pengecekan Stock apakah masih tersedia apa tidak
			return errors.New("Stock Habis!")
		}

		qtyNow := getProductData.Stock - dataCart.Qty

		err = s.product_service.UpdateStock(dataCart.ProductID, qtyNow) //Stock di Master Data dikurangi

		if err != nil {
			log.Printf(err.Error())
			return err
		}

		//Mapping Slice Cart to delete
		carts = append(carts, cartId.CartId)

	}
	//Delete Cart
	err := s.transaction_service.DeleteBulkCart(carts)

	if err != nil {
		log.Println(err)
		return err
	}

	cartData.PaymentType = input.PaymentType
	cartData.TransactionDetail = cartList
	cartData.Address = model.Address(input.Address)
	cartData.UserID = input.UserID
	cartData.CreatedAt = time.Now()
	cartData.UpdatedAt = time.Now()

	err = s.transaction_service.Checkout(cartData)

	if err != nil {
		log.Printf(err.Error())
		return err
	}

	return nil
}

func (s *transactionService) ListTransaction(userId int) ([]model.Transaction, error) {
	listTransaction, err := s.transaction_service.ListTransaction(userId)

	if err != nil {
		log.Println(err.Error())
		return listTransaction, err
	}

	return listTransaction, nil
}

func (s *transactionService) ListSalesTransaction() ([]map[string]interface{}, error) {
	checkouts := []map[string]interface{}{}
	listTransaction, err := s.transaction_service.ListSalesTransaction()

	if err != nil {
		log.Println(err.Error())
		return checkouts, err
	}

	//Mapping Data
	for _, transaction := range listTransaction {
		var checkout map[string]interface{}
		var details []model.TransactionDetail

		for _, transaction_detail := range transaction.TransactionDetail {
			detail := model.TransactionDetail{}

			detailProduct, _ := s.product_service.GetDataByProductId(transaction_detail.ProductID)
			detail.ProductName = detailProduct.Title
			detail.Qty = transaction_detail.Qty
			detail.ProductID = transaction_detail.ProductID
			details = append(details, detail)
		}

		//Mapping Address
		// var address model.Address
		getCity, _ := s.user_service.GetCityById(transaction.Address.CityId)

		transaction.Address.CityName = getCity.CityName

		getProvince, _ := s.user_service.GetProvinceById(transaction.Address.ProvinceId)

		transaction.Address.ProvinceName = getProvince.ProvinceName

		checkout = map[string]interface{}{
			"_id":                transaction.ID,
			"user_id":            transaction.UserID,
			"transaction_detail": details,
			"address":            transaction.Address,
			// "payment_type":       post.PaymentType,
			"status":     "Dibayar",
			"courier":    nil,
			"created_at": time.Time{},
			"updated_at": time.Time{},
		}

		checkouts = append(checkouts, checkout)
	}

	return checkouts, nil
}

func (s *transactionService) UpdateStatus(input entity.TransactionStatus) error {
	err := s.transaction_service.UpdateTransaction(input.TransactionId, "status", input.Status)

	if err != nil {
		log.Printf(err.Error())
		return err
	}

	if input.Status == "Dikirim" {
		err := s.transaction_service.UpdateTransaction(input.TransactionId, "courier", input.Courier)

		if err != nil {
			log.Printf(err.Error())

			return err
		}
	}

	return nil
}

//TODO pesanan Dikonfirmasi, Dikirim, Diterima
