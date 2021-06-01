package rest

import (
	"fmt"
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/module"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

// SupplierHandler contains the function of handler for domain supplier
type SupplierHandler interface {
	Suppliers(c *gin.Context)
	SupplierById(c *gin.Context)
	UpdateSupplier(c *gin.Context)
	CreateSupplier(c *gin.Context)
	DeleteSupplier(c *gin.Context)
}
type supplierHandler struct {
	Serv  module.SupplierUsecase
	CSRF []byte
}

// SupplierInit is to initialize the rest handler for domain supplier
func SupplierInit(userCase module.SupplierUsecase, cs []byte) SupplierHandler {
	return &supplierHandler{Serv: userCase, CSRF: cs}
}

func (ah *supplierHandler) Suppliers(c *gin.Context) {
	suppliers, _ := ah.Serv.GetSuppliers()
	c.JSON(http.StatusOK, gin.H{"supplier-list": suppliers})
}
func (ah *supplierHandler) SupplierById(c *gin.Context) {
	param := c.Param("id")
	id,err:=uuid.FromString(param)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type conversion Failed!"})
		return
	}
	supplier, err := ah.Serv.GetSupplierByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"supplier": supplier})

}
func (ah *supplierHandler) UpdateSupplier(c *gin.Context) {
	var supplier *model.Supplier
	param := c.Param("id")
	id,err:=uuid.FromString(param)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type conversion Failed!"})
		return
	}
	er := c.ShouldBindJSON(&supplier)
	supplier.ID = id
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}
	_, err = ah.Serv.UpdateSupplier(supplier)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"supplier": supplier})
}
func (ah *supplierHandler) CreateSupplier(c *gin.Context) {
	var supplier model.Supplier
	c.BindJSON(&supplier)
	_, err := ah.Serv.CreateSupplier(supplier)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "create new supplier failed! !"})
		return
	}
	suppliers, err := ah.Serv.GetSuppliers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "create new supplier failed! !", "statusCode": http.StatusInternalServerError})
	}
	c.JSON(http.StatusOK, gin.H{"new-supplier": suppliers})
}
func (ah *supplierHandler) DeleteSupplier(c *gin.Context) {
	param := c.Param("id")
	id,err:=uuid.FromString(param)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type conversion Failed!"})
		return
	}
	err = ah.Serv.DeleteSupplier(id)
	fmt.Println("line 141", err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "statusCode": http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, gin.H{"delete": "successed", "statusCode": http.StatusOK})
}
