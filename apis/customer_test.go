package apis_test

import (
	"bytes"
	"dragondarkon/customers-api/apis"
	"dragondarkon/customers-api/model"
	"dragondarkon/customers-api/router"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock repository
type mockDB struct{}

func (db *mockDB) CreateCustomer(customer model.Customer) (model.Customer, error) {
	return customer, nil
}

func (db *mockDB) FindAll() ([]*model.Customer, error) {
	return []*model.Customer{
		{
			ID:   "1",
			Name: "user01",
			Age:  18,
		}, {
			ID:   "2",
			Name: "user02",
			Age:  18,
		}, {
			ID:   "3",
			Name: "user03",
			Age:  18,
		},
	}, nil
}

func (db *mockDB) FindOne(id string) (model.Customer, error) {
	return model.Customer{
		ID:   "1",
		Name: "user01",
		Age:  18,
	}, nil
}

func (db *mockDB) UpdateCustomer(customer model.Customer) (model.Customer, error) {
	return model.Customer{
		ID:   "1",
		Name: "user01",
		Age:  18,
	}, nil
}

func (db *mockDB) DeleteCustomer(id string) error {
	return nil
}

// mock error repository
type mockErrDB struct{}

func (db *mockErrDB) CreateCustomer(customer model.Customer) (model.Customer, error) {
	return customer, errors.New("error: cannot create customer")
}

func (db *mockErrDB) FindAll() ([]*model.Customer, error) {
	return nil, errors.New("error: list customer error")
}

func (db *mockErrDB) FindOne(id string) (model.Customer, error) {
	return model.Customer{}, errors.New("error: find customer error")
}

func (db *mockErrDB) UpdateCustomer(customer model.Customer) (model.Customer, error) {
	return model.Customer{}, errors.New("error: update customer error")
}

func (db *mockErrDB) DeleteCustomer(id string) error {
	return errors.New("error: delete customer error")
}

func Test_Get_Customers_Success(t *testing.T) {
	expected := `[{"id":"1","name":"user01","age":18},{"id":"2","name":"user02","age":18},{"id":"3","name":"user03","age":18}]`
	request := httptest.NewRequest("GET", "/api/customers", nil)
	write := httptest.NewRecorder()

	router.CustomersRouter{
		CustomersAPIs: apis.CustomerHandler{DB: &mockDB{}},
	}.Route().ServeHTTP(write, request)
	actual := strings.Trim(write.Body.String(), "\n")

	assert.Equal(t, http.StatusOK, write.Code)
	assert.NotNil(t, write.Body)
	assert.Equal(t, expected, actual)
}

func Test_Get_Customers_Fail(t *testing.T) {
	expected := `{"code":"500","message":"error: list customer error"}`
	request := httptest.NewRequest("GET", "/api/customers", nil)
	write := httptest.NewRecorder()

	router.CustomersRouter{
		CustomersAPIs: apis.CustomerHandler{DB: &mockErrDB{}},
	}.Route().ServeHTTP(write, request)
	actual := strings.Trim(write.Body.String(), "\n")

	assert.Equal(t, http.StatusInternalServerError, write.Code)
	assert.NotNil(t, write.Body)
	assert.Equal(t, expected, actual)
}

func Test_Get_Customer_By_ID_Success(t *testing.T) {
	expected := `{"id":"1","name":"user01","age":18}`
	request := httptest.NewRequest("GET", "/api/customers/1234", nil)
	write := httptest.NewRecorder()

	router.CustomersRouter{
		CustomersAPIs: apis.CustomerHandler{DB: &mockDB{}},
	}.Route().ServeHTTP(write, request)
	actual := strings.Trim(write.Body.String(), "\n")

	assert.Equal(t, http.StatusOK, write.Code)
	assert.NotNil(t, write.Body)
	assert.Equal(t, expected, actual)
}

func Test_Get_Customer_By_ID_Fail(t *testing.T) {
	expected := `{"code":"500","message":"error: find customer error"}`
	request := httptest.NewRequest("GET", "/api/customers/1234", nil)
	write := httptest.NewRecorder()

	router.CustomersRouter{
		CustomersAPIs: apis.CustomerHandler{DB: &mockErrDB{}},
	}.Route().ServeHTTP(write, request)
	actual := strings.Trim(write.Body.String(), "\n")

	assert.Equal(t, http.StatusInternalServerError, write.Code)
	assert.NotNil(t, write.Body)
	assert.Equal(t, expected, actual)
}

func Test_Create_Customer_Success(t *testing.T) {
	expected := `{"id":"4","name":"user04","age":18}`
	customer := model.Customer{
		ID:   "4",
		Name: "user04",
		Age:  18,
	}
	requestJSON, _ := json.Marshal(customer)
	request := httptest.NewRequest("POST", "/api/customers", bytes.NewBuffer(requestJSON))
	write := httptest.NewRecorder()

	router.CustomersRouter{
		CustomersAPIs: apis.CustomerHandler{DB: &mockDB{}},
	}.Route().ServeHTTP(write, request)
	actual := strings.Trim(write.Body.String(), "\n")

	assert.Equal(t, http.StatusCreated, write.Code)
	assert.NotNil(t, write.Body)
	assert.Equal(t, expected, actual)
}

func Test_Create_Customer_Fail(t *testing.T) {
	expected := `{"code":"500","message":"error: cannot create customer"}`
	customer := model.Customer{
		ID:   "4",
		Name: "user04",
		Age:  18,
	}
	requestJSON, _ := json.Marshal(customer)
	request := httptest.NewRequest("POST", "/api/customers", bytes.NewBuffer(requestJSON))
	write := httptest.NewRecorder()

	router.CustomersRouter{
		CustomersAPIs: apis.CustomerHandler{DB: &mockErrDB{}},
	}.Route().ServeHTTP(write, request)
	actual := strings.Trim(write.Body.String(), "\n")

	assert.Equal(t, http.StatusInternalServerError, write.Code)
	assert.NotNil(t, write.Body)
	assert.Equal(t, expected, actual)
}

func Test_Update_Customer_Success(t *testing.T) {
	expected := `{"id":"1","name":"user01","age":20}`
	customer := model.Customer{
		ID:   "1",
		Name: "user01",
		Age:  20,
	}
	requestJSON, _ := json.Marshal(customer)
	request := httptest.NewRequest("PUT", "/api/customers", bytes.NewBuffer(requestJSON))
	write := httptest.NewRecorder()

	router.CustomersRouter{
		CustomersAPIs: apis.CustomerHandler{DB: &mockDB{}},
	}.Route().ServeHTTP(write, request)
	actual := strings.Trim(write.Body.String(), "\n")

	assert.Equal(t, http.StatusOK, write.Code)
	assert.NotNil(t, write.Body)
	assert.Equal(t, expected, actual)
}

func Test_Update_Customer_Fail(t *testing.T) {
	expected := `{"code":"404","message":"error: update customer error"}`
	customer := model.Customer{
		ID:   "1",
		Name: "user01",
		Age:  20,
	}
	requestJSON, _ := json.Marshal(customer)
	request := httptest.NewRequest("PUT", "/api/customers", bytes.NewBuffer(requestJSON))
	write := httptest.NewRecorder()

	router.CustomersRouter{
		CustomersAPIs: apis.CustomerHandler{DB: &mockErrDB{}},
	}.Route().ServeHTTP(write, request)
	actual := strings.Trim(write.Body.String(), "\n")

	assert.Equal(t, http.StatusNotFound, write.Code)
	assert.NotNil(t, write.Body)
	assert.Equal(t, expected, actual)
}

func Test_Delete_Customer_Success(t *testing.T) {
	request, _ := http.NewRequest("DELETE", "/api/customers/1", nil)
	write := httptest.NewRecorder()
	router.CustomersRouter{
		CustomersAPIs: apis.CustomerHandler{DB: &mockDB{}},
	}.Route().ServeHTTP(write, request)

	assert.Equal(t, http.StatusNoContent, write.Code)
}

func Test_Delete_Customer_Fail(t *testing.T) {
	expected := `{"code":"500","message":"error: delete customer error"}`
	request, _ := http.NewRequest("DELETE", "/api/customers/1", nil)
	write := httptest.NewRecorder()
	router.CustomersRouter{
		CustomersAPIs: apis.CustomerHandler{DB: &mockErrDB{}},
	}.Route().ServeHTTP(write, request)
	actual := strings.Trim(write.Body.String(), "\n")

	assert.Equal(t, http.StatusInternalServerError, write.Code)
	assert.NotNil(t, write.Body)
	assert.Equal(t, expected, actual)
}
