package rest

import (
	"fmt"
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/module"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

// StockHandler contains the function of handler for domain stock
type StockHandler interface {
	Stocks(c *gin.Context)
	StockById(c *gin.Context)
	UpdateStock(c *gin.Context)
	CreateStock(c *gin.Context)
	DeleteStock(c *gin.Context)
}
type stockHandler struct {
	Serv  module.StockUsecae
	CSRF []byte
}

// StockInit is to initialize the rest handler for domain stco
func StockInit(userCase module.StockUsecae, cs []byte) StockHandler {
	return &stockHandler{Serv: userCase, CSRF: cs}
}

func (ah *stockHandler) Stocks(c *gin.Context) {
	stocks, _ := ah.Serv.Get()
	c.JSON(http.StatusOK, gin.H{"stock-list": stocks})
}
func (ah *stockHandler) StockById(c *gin.Context) {
	param := c.Param("id")
	id,err:=uuid.FromString(param)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type conversion Failed!"})
		return
	}
	stock, err := ah.Serv.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"stock": stock})

}
func (ah *stockHandler) UpdateStock(c *gin.Context) {
	var stock *model.Stock
	param := c.Param("id")
	id,err:=uuid.FromString(param)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type conversion Failed!"})
		return
	}
	er := c.ShouldBindJSON(&stock)
	stock.ID = id
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}
	_, err = ah.Serv.Update(stock)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"updated-stock": stock})
}
func (ah *stockHandler) CreateStock(c *gin.Context) {
	var stock model.Stock
	c.BindJSON(&stock)
	_, err := ah.Serv.Create(stock)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "create new stock failed! !"})
		return
	}
	stocks, err := ah.Serv.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "create new stock failed! !", "statusCode": http.StatusInternalServerError})
	}
	c.JSON(http.StatusOK, gin.H{"new-stock": stocks})
}
func (ah *stockHandler) DeleteStock(c *gin.Context) {
	param := c.Param("id")
	id,err:=uuid.FromString(param)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type conversion Failed!"})
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
