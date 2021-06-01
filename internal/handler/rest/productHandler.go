package rest

import (
	"fmt"
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/module"
	"github.com/Yideg/inventory_app/pkg/comparision"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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
	var items []model.Product
	products, _ := ah.Serv.GetProducts()
	for _,item:= range products{
		ExpirationTime:=item.ExpiredOn
		if comparision.CheckExpirationTime(ExpirationTime) {
			items=append(items,item)
		}
		continue
	}
	if len(items)==0 {
		c.JSON(http.StatusOK, gin.H{"products": "No Active Product Found. Product may be expired!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": items})
}
func (ah *productHandler) ProductById(c *gin.Context) {
	param := c.Param("id")
	id,err:=uuid.FromString(param)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type conversion Failed!"})
		return
	}
	item, err := ah.Serv.GetProductsByID(id)
	if err != nil {
		fmt.Println("line 54", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ExpirationTime:=item.ExpiredOn
	if comparision.CheckExpirationTime(ExpirationTime) {
		c.JSON(http.StatusOK, gin.H{"Products": item})
		return
	}
	c.JSON(http.StatusOK, gin.H{"orders": "No Active Product Found. Product May be Expired!"})
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
	_, err:= ah.Serv.UpdateProduct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}
func (ah *productHandler) CreateProduct(c *gin.Context) {
	var product model.Product
	c.BindJSON(&product)
	_, err := ah.Serv.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": " product registration failed! !"})
		return
	}
	prod, err := ah.Serv.GetProducts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": " product registration failed! !", "statusCode": http.StatusInternalServerError})
	}
	c.JSON(http.StatusOK, gin.H{"product": prod})
}
func (ah *productHandler) DeleteProduct(c *gin.Context) {
	param := c.Param("id")
	id,err:=uuid.FromString(param)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type conversion Failed!"})
		return
	}
	err = ah.Serv.DeleteProduct(id)
	fmt.Println("line 141", err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "statusCode": http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, gin.H{"delete": "successed", "statusCode": http.StatusOK})
}
