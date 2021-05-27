package rest

import (
	"fmt"
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/module"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ProductHandler contains the function of handler for domain product
type ProductHandler interface {
	Products(c *gin.Context)
	ProductById(c *gin.Context)
	UpdateProduct(c *gin.Context)
	CreateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}
type productHandler struct {
	Serv  module.ProductUsecase
	CSRF []byte
}

// ProductInit is to initialize the rest handler for domain product
func ProductInit(userCase module.ProductUsecase, cs []byte) ProductHandler {
	return &productHandler{Serv: userCase, CSRF: cs}
}

func (ah *productHandler) Products(c *gin.Context) {
	products, _ := ah.Serv.Get()
	fmt.Println("data", products)
	c.JSON(http.StatusOK, gin.H{"products": products})
}
func (ah *productHandler) ProductById(c *gin.Context) {
	fmt.Println("start handler")
	param := c.Param("id")
	//there is somthing to be done since the id is a type isbn data type not string but for temporary use as string
	//
	//
	role, err := ah.Serv.GetById(param)
	if err != nil {
		fmt.Println("line 54", err)
		fmt.Println("data=", role)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"product": role})

}
func (ah *productHandler) UpdateProduct(c *gin.Context) {
	var product *model.Product
	id:= c.Param("id")
	if id==""{
		c.JSON(http.StatusBadRequest, gin.H{"error":"Input id not provided"})
		return
	}
	er := c.ShouldBindJSON(&product)
	product.Barcode =id
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}
	_, err:= ah.Serv.Update(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}
func (ah *productHandler) CreateProduct(c *gin.Context) {
	var product model.Product
	c.BindJSON(&product)
	_, err := ah.Serv.Create(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": " product registration failed! !"})
		return
	}
	prod, err := ah.Serv.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": " product registration failed! !", "statusCode": http.StatusInternalServerError})
	}
	c.JSON(http.StatusOK, gin.H{"product": prod})
}
func (ah *productHandler) DeleteProduct(c *gin.Context) {
	param := c.Param("id")
	if param ==""{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input id not provided! ", "statusCode": http.StatusNotAcceptable})
		return
	}
	err := ah.Serv.Delete(param)
	fmt.Println("line 141", err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "statusCode": http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, gin.H{"delete": "successed", "statusCode": http.StatusOK})
}
