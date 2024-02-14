package database

import (
	"dragondarkon/customers-api/model"
	"log"
)

type CustomersRepository interface {
	CreateCustomer(customer model.Customer) (model.Customer, error)
	FindAll() ([]*model.Customer, error)
	FindOne(id string) (model.Customer, error)
	UpdateCustomer(customer model.Customer) (model.Customer, error)
	DeleteCustomer(id string) error
}

func (db *Database) CreateCustomer(customer model.Customer) (model.Customer, error) {
	if result := db.Create(customer); result.Error != nil {
		log.Fatal("error:" + result.Error.Error())
		return customer, result.Error
	}
	return customer, nil
}

func (db *Database) FindAll() ([]*model.Customer, error) {
	var customers []*model.Customer
	if result := db.Find(&customers); result.Error != nil {
		log.Fatal("error:" + result.Error.Error())
		return customers, result.Error
	}
	return customers, nil
}
func (db *Database) FindOne(id string) (model.Customer, error) {
	var customer model.Customer
	if result := db.Find(&customer, id); result.Error != nil {
		log.Fatal("error:" + result.Error.Error())
		return customer, result.Error
	}
	return customer, nil
}

func (db *Database) UpdateCustomer(customer model.Customer) (model.Customer, error) {
	if result := db.Updates(&customer); result.Error != nil {
		log.Fatal("error:" + result.Error.Error())
		return customer, result.Error
	}
	return customer, nil
}

func (db *Database) DeleteCustomer(id string) error {
	if result := db.Delete(&model.Customer{}, id); result.Error != nil {
		log.Fatal("error:" + result.Error.Error())
		return result.Error
	}
	return nil
}
