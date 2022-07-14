package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() customerRepositoryMock {
	customers := []Customer{
		{CustomerID: 1001, Name: "Ashish", DateOfBirth: "2002-08-01", City: "New Delhi", ZipCode: "1559", Status: 1},
		{CustomerID: 1002, Name: "Jhon", DateOfBirth: "2002-06-01", City: "London", ZipCode: "20000", Status: 1},
	}

	return customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}

	return nil, errors.New("customer not found")
}