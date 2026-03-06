package main

import (
	"log/slog"
	"net/http"
	"pizza-tracker-go/internal/models"

	"github.com/gin-gonic/gin"
)

type OrderFormData struct {
	PizzaTypes  []string
	PizzaSizes  []string
	PizzaCrusts []string
	PizzaAddOns []string
}

type OrderRequest struct {
	Name         string   `json:"name" binding:"required,min=3,max=64"`
	Phone        string   `json:"phone" binding:"required,e164"`
	Address      string   `json:"address" binding:"required,min=5,max=256"`
	PizzaSizes   []string `json:"sizes" binding:"required,dive,valid_pizza_size"`
	PizzaTypes   []string `json:"types" binding:"required,dive,valid_pizza_type"`
	PizzaCrusts  []string `json:"crusts" binding:"required,dive,valid_pizza_crust"`
	PizzaAddOns  []string `json:"addons" binding:"dive,valid_pizza_addon"`
	Instructions []string `json:"instructions" binding:"max=256"`
}

func (h *Handler) ServeNewOrderForm(c *gin.Context) {
	c.HTML(http.StatusOK, "order.tmpl", OrderFormData{
		PizzaTypes:  models.PizzaTypes,
		PizzaSizes:  models.PizzaSizes,
		PizzaCrusts: models.PizzaCrusts,
		PizzaAddOns: models.PizzaAddOns,
	})
}

func (h *Handler) HandleNewOrderPost(c *gin.Context) {
	var form OrderRequest
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderItems := make([]models.OrderItem, len(form.PizzaSizes))
	for i := range orderItems {
		orderItems[i] = models.OrderItem{
			PizzaType:    form.PizzaTypes[i],
			PizzaSize:    form.PizzaSizes[i],
			PizzaCrust:   form.PizzaCrusts[i],
			AddOns:       form.PizzaAddOns[i],
			Instructions: form.Instructions[i],
		}
	}

	order := models.Order{
		CustomerName: form.Name,
		Phone:        form.Phone,
		Address:      form.Address,
		Status:       models.OrderStatuses[0],
		Items:        orderItems,
	}

	if err := h.orders.CreateOrder(&order); err != nil {
		slog.Error("Failed to create order", "error", err)
		c.String(http.StatusInternalServerError, "Something went wrong")
		return
	}

	slog.Info("Order created", "orderID", order.ID, "customer", order.CustomerName)
	c.Redirect(http.StatusSeeOther, "/customer/"+order.ID)

}

func (h *Handler) serveCustomer(c *gin.Context) {
	orderID := c.Param("id")
	if orderID == "" {
		c.String(http.StatusBadRequest, "order ID is required")
		return
	}
	order, err := h.orders.GetOrderByID(orderID)
	if err != nil {
		c.String(http.StatusNotFound, "order not found")
		return
	}
	c.HTML(http.StatusOK, "customer.tmpl", gin.H{
		"Order": order,
	})
}
