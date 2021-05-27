package rest

import (
	"fmt"
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/module"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

// RoleHandler contains the function of handler for domain user
type RoleHandler interface {
	Roles(c *gin.Context)
	RoleById(c *gin.Context)
	UpdateRole(c *gin.Context)
	CreateRole(c *gin.Context)
	DeleteRole(c *gin.Context)
}
type roleHandler struct {
	Serv  module.RoleServices
	CSRF []byte
}

// RoleInit is to initialize the rest handler for domain role
func RoleInit(userCase module.RoleServices, cs []byte) RoleHandler {
	return &roleHandler{Serv: userCase, CSRF: cs}
}

func (ah *roleHandler) Roles(c *gin.Context) {
	roles, _ := ah.Serv.Get()
	fmt.Println("data", roles)
	c.JSON(http.StatusOK, gin.H{"roles": roles})
}
func (ah *roleHandler) RoleById(c *gin.Context) {
	param := c.Param("id")
	id,err:=uuid.FromString(param)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type conversion Failed!"})
		return
	}
	role, err := ah.Serv.GetById(id)
	if err != nil {
		fmt.Println("line 54", err)
		fmt.Println("data=", role)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"role": role})

}
func (ah *roleHandler) UpdateRole(c *gin.Context) {
	var role *model.Role
	param := c.Param("id")
	id,err:=uuid.FromString(param)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type conversion Failed!"})
		return
	}
	er := c.ShouldBindJSON(&role)
	role.ID = id
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}
	_, err = ah.Serv.Update(role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"roles": role})
}
func (ah *roleHandler) CreateRole(c *gin.Context) {
	var role model.Role
	c.BindJSON(&role)
	_, err := ah.Serv.Create(role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "create role failed! !"})
		return
	}
	rol, err := ah.Serv.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "create role failed! !", "statusCode": http.StatusInternalServerError})
	}
	c.JSON(http.StatusOK, gin.H{"role": rol})
}
func (ah *roleHandler) DeleteRole(c *gin.Context) {
	param := c.Param("id")
	id,err:=uuid.FromString(param)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type conversion Failes!"})
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
