package services

import (
	"ddd/aggregate"
	"testing"

	"github.com/google/uuid"
)

func init_products(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "Healthy Beverage", 1.2)

	if err != nil {
		t.Fatal(err)
	}

	peanuts, err := aggregate.NewProduct("Peanuts", "Snacks Beverage", 0.99)

	if err != nil {
		t.Fatal(err)
	}

	wine, err := aggregate.NewProduct("Wine", "Some drink", 0.99)

	if err != nil {
		t.Fatal(err)
	}

	return []aggregate.Product{
		beer, peanuts, wine,
	}
}

func TestOrder_NewOrderService(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregate.NewCustomer("Percy")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(cust.GetId(), order)
	if err != nil {
		t.Error(err)
	}
}
