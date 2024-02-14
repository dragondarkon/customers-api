package apis

import (
	"dragondarkon/customers-api/database"
	"dragondarkon/customers-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	DB database.CustomersRepository
}

type Resp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	customer := model.Customer{}
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, Resp{
			Code:    "500",
			Message: err.Error(),
		})
		return
	}

	if _, err := h.DB.CreateCustomer(customer); err != nil {
		c.JSON(http.StatusInternalServerError, Resp{
			Code:    "500",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, &customer)
}

func (h *CustomerHandler) ListCustomer(c *gin.Context) {
	customers, err := h.DB.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Resp{
			Code:    "500",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, customers)
}

func (h *CustomerHandler) GetCustomer(c *gin.Context) {
	id := c.Param("id")
	customer, err := h.DB.FindOne(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Resp{
			Code:    "500",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	customer := model.Customer{}
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, Resp{
			Code:    "500",
			Message: err.Error(),
		})
		return
	}

	if _, err := h.DB.UpdateCustomer(customer); err != nil {
		c.JSON(http.StatusNotFound, Resp{
			Code:    "404",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	if err := h.DB.DeleteCustomer(id); err != nil {
		c.JSON(http.StatusInternalServerError, Resp{
			Code:    "500",
			Message: err.Error(),
		})
		return
	}
	c.Status(http.StatusNoContent)
}
