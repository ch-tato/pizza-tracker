package main

import (
	"fmt"
	"net/http"
	"pizza-tracker-go/internal/models"

	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Error string
}

type AdminDashboardData struct {
	Username string
	Orders   []models.Order
	Statuses []string
}

func (h *Handler) HandleLoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", LoginData{})
}

func (h *Handler) HandleLoginPost(c *gin.Context) {
	var form struct {
		Username string `form:"username" binding:"required,min=3,max=50"`
		Password string `form:"password" binding:"required,min=8"`
	}

	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusOK, "login.tmpl", LoginData{
			Error: "invalid input: " + err.Error(),
		})
		return
	}

	user, err := h.users.AuthenticateUser(form.Username, form.Password)
	if err != nil {
		c.HTML(http.StatusOK, "login.tmpl", LoginData{
			Error: "invalid credentials",
		})
		return
	}

	SetSessionValue(c, "userID", fmt.Sprintf("%v", user.ID))
	SetSessionValue(c, "username", user.Username)

	c.Redirect(http.StatusSeeOther, "/admin")
}

func (h *Handler) HandleLogout(c *gin.Context) {
	if err := ClearSession(c); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Redirect(http.StatusSeeOther, "/login")
}

func (h *Handler) ServeAdminDashboard(c *gin.Context) {
	orders, err := h.orders.GetAllOrders()
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to get orders: "+err.Error())
		return
	}

	username := GetSessionString(c, "username")
	c.HTML(http.StatusOK, "admin.tmpl", AdminDashboardData{
		Username: username,
		Orders:   orders,
		Statuses: models.OrderStatuses,
	})
}

func (h *Handler) HandleOrderPut(c *gin.Context) {
	id := c.Param("id")
	newStatus := c.PostForm("status")
	if err := h.orders.UpdateOrderStatus(id, newStatus); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Redirect(http.StatusSeeOther, "/admin")
}

func (h *Handler) HandleOrderDelete(c *gin.Context) {
	id := c.Param("id")
	if err := h.orders.DeleteOrder(id); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Redirect(http.StatusSeeOther, "/admin")
}
