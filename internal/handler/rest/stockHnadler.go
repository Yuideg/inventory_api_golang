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
	Serv    module.StockUsecae
	supServ module.SupplierUsecase
	CSRF    []byte
}

// StockInit is to initialize the rest handler for domain stco
func StockInit(userCase module.StockUsecae, sup module.SupplierUsecase, cs []byte) StockHandler {
	return &stockHandler{Serv: userCase, supServ: sup, CSRF: cs}
}

func (ah *stockHandler) Stocks(c *gin.Context) {
	var newStock []model.Stock
	stocks, _ := ah.Serv.GetStocks()
	for _, stock := range stocks {
		sup, err := ah.supServ.GetSupplierBySupplierID(stock.SupplierID)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err})
			return
		}
		stock.Supplier_List = sup
		newStock = append(newStock, stock)
	}
	c.JSON(http.StatusOK, gin.H{"stock-list": newStock})
}
func (ah *stockHandler) StockById(c *gin.Context) {
	param := c.Param("id")
	id, err := uuid.FromString(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type conversion Failed!"})
		return
	}
	stock, err := ah.Serv.GetStockByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"stock": stock})

}
func (ah *stockHandler) UpdateStock(c *gin.Context) {
	var stock *model.Stock
	param := c.Param("id")
	id, err := uuid.FromString(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type conversion Failed!"})
		return
	}
	er := c.ShouldBindJSON(&stock)
	stock.ID = id
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}
	_, err = ah.Serv.UpdateStock(stock)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"updated-stock": stock})
}
func (ah *stockHandler) CreateStock(c *gin.Context) {
	var stock model.Stock
	err := c.BindJSON(&stock)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	_, err = ah.Serv.CreateStock(stock)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "create new stock failed! !"})
		return
	}
	stocks, err := ah.Serv.GetStocks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "create new stock failed! !", "statusCode": http.StatusInternalServerError})
	}
	c.JSON(http.StatusOK, gin.H{"new-stock": stocks})
}
func (ah *stockHandler) DeleteStock(c *gin.Context) {
	param := c.Param("id")
	id, err := uuid.FromString(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type conversion Failed!"})
		return
	}
	err = ah.Serv.DeleteStock(id)
	fmt.Println("line 141", err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "statusCode": http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK})
}
