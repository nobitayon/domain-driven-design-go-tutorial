package services

import (
	"ddd/aggregate"
	"ddd/domain/customer"
	"ddd/domain/customer/memory"
	"ddd/domain/product"
	prodmem "ddd/domain/product/memory"
	"log"

	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	// loop through all the Cfgs and apply them
	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

// WithCustomerRepository applies a customer repository to the OrderService
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	// Return a function that matches the orderconfiguration alias
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmem.New()
		for _, p := range products {
			if err := pr.Add(p); err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, prodctsIDs []uuid.UUID) (float64, error) {
	// fetch the customer
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0.0, err
	}
	// Get each Product, Ouchie no product repository
	var products []aggregate.Product
	var total float64

	for _, id := range prodctsIDs {
		p, err := o.products.GetByID(id)

		if err != nil {
			return 0.0, err
		}
		products = append(products, p)
		total += p.GetPrice()
	}
	log.Printf("Customer: %s has ordered %d products", c.GetId(), len(products))

	return total, nil
}
