package rest

import (
	"fmt"
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/module"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
)

// UserHandler contains the function of handler for domain user
type CustomerHandler interface {
	Customers(c *gin.Context)
	CustomerById(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	CreateCustomer(c *gin.Context)
	DeleteCustomer(c *gin.Context)
}
type customerHandler struct {
	Serv  module.CustomerUsecase
	CSRF []byte
}


// CustomerInit is to initialize the rest handler for domain customer
func CustomerInit(userCase module.CustomerUsecase, cs []byte) CustomerHandler {
	return &customerHandler{Serv: userCase, CSRF: cs}
}

func (ah *customerHandler) Customers(c *gin.Context) {
	fmt.Println("cust get started!!")
	customers, _ := ah.Serv.Get()
	fmt.Println("data", customers)
	c.JSON(http.StatusOK, gin.H{"customers": customers})
}
func (ah *customerHandler) CustomerById(c *gin.Context) {
	param := c.Param("id")
	id,err:=uuid.FromString(param)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input!"})
		return
	}
	custom, err := ah.Serv.GetById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"customers": custom})

}
func (ah *customerHandler) UpdateCustomer(c *gin.Context) {
	var custom *model.Customer
	param := c.Param("id")
	id,err:=uuid.FromString(param)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input!"})
		return
	}
	er := c.ShouldBindJSON(&custom)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}
	   custom.ID = id
	_, err = ah.Serv.Update(custom)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"customers": custom})
}
func (ah *customerHandler) CreateCustomer(c *gin.Context) {
	var customer model.Customer
	c.BindJSON(&customer)
	_, err := ah.Serv.Create(customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer registration failed! !"})
		return
	}
	custom, err := ah.Serv.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer registration failed! !", "statusCode": http.StatusInternalServerError})
	}
	c.JSON(http.StatusOK, gin.H{"account": custom})
}
func (ah *customerHandler) DeleteCustomer(c *gin.Context) {
	param := c.Param("id")
	id,err:=uuid.FromString(param)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input!"})
		return
	}
	err = ah.Serv.Delete(id)
	fmt.Println("line 141", err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "statusCode": http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, gin.H{"delete": "successed", "statusCode": http.StatusOK})
}
