package handlers

import (
	"go-echo-app/common"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) GetCustomers(c echo.Context) error {
	var customers []common.Customer
	if err := h.DB.Find(&customers).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, customers)
}

func (h *Handler) CreateCustomer(c echo.Context) error {
	var customer common.Customer
	if err := c.Bind(&customer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.DB.Create(&customer).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, customer)
}
