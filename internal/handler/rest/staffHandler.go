package rest

import (
	"fmt"
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/module"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

//  StaffHandler contains the function of handler for domain staff
type StaffHandler interface {
	Staffs(c *gin.Context)
	StaffById(c *gin.Context)
	UpdateStaff(c *gin.Context)
	CreateStaff(c *gin.Context)
	DeleteStaff(c *gin.Context)
}
type staffHandler struct {
	Serv  module.StaffUsecase
	CSRF []byte
}

//  StaffInit is to initialize the rest handler for domain staff
func  StaffInit(userCase module.StaffUsecase, cs []byte)  StaffHandler {
	return &staffHandler{Serv: userCase, CSRF: cs}
}

func (ah *staffHandler)  Staffs(c *gin.Context) {
	staffs, _ := ah.Serv.Get()
	fmt.Println("data", staffs)
	c.JSON(http.StatusOK, gin.H{"staffs": staffs})
}
func (ah *staffHandler)  StaffById(c *gin.Context) {
	fmt.Println("start handler")
	param := c.Param("id")
	id,err:=uuid.FromString(param)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type conversion Failed!"})
		return
	}
	staff, err := ah.Serv.GetById(id)
	if err != nil {
		fmt.Println("line 54", err)
		fmt.Println("data=", staff)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"staff": staff})

}
func (ah *staffHandler) UpdateStaff(c *gin.Context) {
	var staff *model.Staff
	param := c.Param("id")
	id,err:=uuid.FromString(param)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type conversion Failed!"})
		return
	}
	er := c.ShouldBindJSON(&staff)
	staff.ID = id
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}
	_, err = ah.Serv.Update(staff)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"staff": staff})
}
func (ah *staffHandler) CreateStaff(c *gin.Context) {
	var staff model.Staff
	c.BindJSON(&staff)
	_, err := ah.Serv.Create(staff)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "new staff member registration failed! !"})
		return
	}
	staffs, err := ah.Serv.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "new staff member registration failed! !", "statusCode": http.StatusInternalServerError})
	}
	c.JSON(http.StatusOK, gin.H{"staff": staffs})
}
func (ah *staffHandler) DeleteStaff(c *gin.Context) {
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
