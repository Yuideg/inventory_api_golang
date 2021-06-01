package rest

import (
	"fmt"
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/module"
	"github.com/Yideg/inventory_app/pkg/comparision"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
)

// OrderHandler contains the function of handler for domain order
type OrderHandler interface {
	Orders(c *gin.Context)
	OrderById(c *gin.Context)
	UpdateOrder(c *gin.Context)
	CreateOrder(c *gin.Context)
	DeleteOrder(c *gin.Context)
}
type orderHandler struct {
	Serv  module.OrderUsecase
	CSRF []byte
}

// OrderInit is to initialize the rest handler for domain order
func OrderInit(userCase module.OrderUsecase, cs []byte) OrderHandler {
	return &orderHandler{Serv: userCase, CSRF: cs}
}

func (ah *orderHandler) Orders(c *gin.Context) {
	var newOrder[]model.Order
	orders, _ := ah.Serv.GetOrders()
	for _,v:= range orders{
		ExpirationTime:=v.ExpiredOn
		if comparision.CheckExpirationTime(ExpirationTime) {
			newOrder=append(newOrder,v)
		}
		continue
	}
	if len(newOrder)==0 {
		c.JSON(http.StatusOK, gin.H{"orders": "No Active Order Found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"orders": newOrder})
}
func (ah *orderHandler) OrderById(c *gin.Context) {
	param := c.Param("id")
	id,err:=uuid.FromString(param)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input!"})
		return
	}
	order, err := ah.Serv.GetOrderByID(id)
	if err != nil {
		fmt.Println("line 54", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ExpirationTime:=order.ExpiredOn
	if comparision.CheckExpirationTime(ExpirationTime) {
		c.JSON(http.StatusOK, gin.H{"orders": order})
		return
	}
	c.JSON(http.StatusOK, gin.H{"orders": "No Active Order Found. Order May be Expired!"})

}
func (ah *orderHandler) UpdateOrder(c *gin.Context) {
	var order *model.Order
	param := c.Param("id")
	id,err:=uuid.FromString(param)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input!"})
		return
	}
	er := c.ShouldBindJSON(&order)
	order.ID = id
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}
	_, err = ah.Serv.UpdateOrder(order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"orders": order})
}
func (ah *orderHandler) CreateOrder(c *gin.Context) {
	var order model.Order
	c.BindJSON(&order)
	_, err := ah.Serv.CreateOrder(order)
	if err != nil {
		fmt.Println("line 79",err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "create orders failed! !"})
		return
	}
	ord, err := ah.Serv.GetOrders()
	if err != nil {
		fmt.Println("line 85",err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "create order failed! !", "statusCode": http.StatusInternalServerError})
	}
	c.JSON(http.StatusOK, gin.H{"orders": ord})
}
func (ah *orderHandler) DeleteOrder(c *gin.Context) {
	param := c.Param("id")
	id,err:=uuid.FromString(param)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input!"})
		return
	}
	err = ah.Serv.DeleteOrder(id)
	fmt.Println("line 141", err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "statusCode": http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, gin.H{"delete": "successed", "statusCode": http.StatusOK})
}
