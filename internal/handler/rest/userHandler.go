package rest

import (
	"fmt"
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/module"
	"github.com/Yideg/inventory_app/pkg/comparision"
	"github.com/Yideg/inventory_app/pkg/encreyption"
	"github.com/Yideg/inventory_app/pkg/jwt"
	"github.com/Yideg/inventory_app/pkg/matchstr"
	"github.com/Yideg/inventory_app/pkg/rbac"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

// UserHandler contains the function of handler for domain user
type UserHandler interface {
	UsersHandler(*gin.Context)
	GetUserByIDHandler(*gin.Context)
	UpdateUserHandler(*gin.Context)
	CreateUserHandler(*gin.Context)
	DeleteUserHandler(*gin.Context)
	Authenticated(c *gin.Context)
	Authorized(c *gin.Context)
	LoginUserHandler(c *gin.Context)
}
type userHandler struct {
	Serv      module.UserUsecase
	rolserv   module.RoleServices
	orderServ module.OrderUsecase
	CSRF      []byte
}

// UserInit is to initialize the rest handler for domain user
func UserInit(userCase module.UserUsecase, rol module.RoleServices, oServ module.OrderUsecase, cs []byte) UserHandler {
	return &userHandler{Serv: userCase, rolserv: rol, orderServ: oServ, CSRF: cs}
}
func (uph *userHandler) Authenticated(c *gin.Context) {
	err := jwt.TokenValid(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "user hasn't logged in yet")
		c.Abort()
		return
	}
	c.Next()

}
func (uph *userHandler) Authorized(c *gin.Context) {
	var user1 model.User
	uid, err := jwt.ExtractTokenID(c.Request)

	uuid, err := uuid.FromString(uid)
	user1.ID = uuid
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Type conversion error")
		c.Abort()
		return
	}
	user, err := uph.Serv.GetUserByID(uuid)
	fmt.Println("user role id =", user.RoleID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "user doesn't exit")
		c.Abort()
		return
	}
	roleDB, er := uph.rolserv.GetRoleByID(user.RoleID)
	fmt.Println("full path=", c.Request.URL)
	fmt.Println("role=", roleDB)
	if er != nil {
		c.JSON(http.StatusUnauthorized, "it seems not role exist")
		c.Abort()
		return
	}

	roles := rbac.GetPermission(roleDB.Name)
	for _, _ = range roles.Permissions {
		target_id := c.Request.URL.String()
		for _, target_value := range []string{target_id} {
			fmt.Println("target id", target_id)
			action := rbac.GetAction(target_id)
			fmt.Println("permission=", action, target_value)
			canDo, _ := roles.Can(action, target_value)
			fmt.Printf("Can  %s? %t\n", action, canDo)
			if canDo {
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, "sorry,You don't have enough permission!")
				c.Abort()
				return

			}

		}

	}
	return
}
func (ah *userHandler) UsersHandler(c *gin.Context) {
	fmt.Println("cust get started!!")
	var allUser []model.User
	users, _ := ah.Serv.GetUsers()
	uid, err := jwt.ExtractTokenID(c.Request)

	uuid, err := uuid.FromString(uid)
	if err != nil {
		fmt.Println("line 109 ",err)
		c.JSON(http.StatusUnauthorized, "Type conversion error")
		c.Abort()
		return
	}
	currentlyLoggedUser,err:=ah.Serv.GetUserByID(uuid)
	if err != nil {
		fmt.Println("line 116 ",err)
		c.JSON(http.StatusUnauthorized, "User Fetching Error")
		c.Abort()
		return
	}
	loggedUserRole,err:=ah.rolserv.GetRoleByID(currentlyLoggedUser.RoleID)
	if err != nil {
		fmt.Println("line 123 ",err)
		c.JSON(http.StatusUnauthorized, "User role Fetching Error")
		c.Abort()
		return
	}
	for _, user := range users {
		getOrder, _ := ah.orderServ.GetOrderByUserID(user.ID)
		getRole, _ := ah.rolserv.GetRoleByID(user.RoleID)
		user.Role = getRole
		//user password is vissible only for admins
		fmt.Println("logged user =",currentlyLoggedUser.Role)
		if loggedUserRole.Name=="USER"{
			user.Password = ""
		}
		user.Orders = getOrder
		allUser = append(allUser, user)

	}
	c.JSON(http.StatusOK, gin.H{"users": allUser})
}
func (ah *userHandler) GetUserByIDHandler(c *gin.Context) {
	param := c.Param("id")
	id, err := uuid.FromString(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input!"})
		return
	}
	user, err := ah.Serv.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	getOrder, _ := ah.orderServ.GetOrderByUserID(user.ID)
	getRole, _ := ah.rolserv.GetRoleByID(user.RoleID)
	user.Orders = getOrder
	user.Role = getRole
	c.JSON(http.StatusOK, gin.H{"user": user})
}
func (ah *userHandler) UpdateUserHandler(c *gin.Context) {
	var user *model.User
	param := c.Param("id")
	id, err := uuid.FromString(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input!"})
		return
	}
	er := c.ShouldBindJSON(&user)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}
	user.ID = id
	user.Password = encreyption.HashPassword(c, user.Password)
	if !matchstr.EmailMatcher(user.Email) && !matchstr.PhoneMatcher(user.Phone) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input Found!"})
		return
	}
	_, err = ah.Serv.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"customers": user})
}
func (ah *userHandler) CreateUserHandler(c *gin.Context) {
	var user model.User
	er := c.ShouldBindJSON(&user)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}
	user.Password = encreyption.HashPassword(c, user.Password)
	if !matchstr.EmailMatcher(user.Email) && !matchstr.PhoneMatcher(user.Phone) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input Found!"})
		return
	}
	_, err := ah.Serv.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user registration failed! !"})
		return
	}
	custom, err := ah.Serv.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user registration failed! !", "statusCode": http.StatusInternalServerError})
	}
	c.JSON(http.StatusOK, gin.H{"account": custom})
}
func (ah *userHandler) DeleteUserHandler(c *gin.Context) {
	param := c.Param("id")
	id, err := uuid.FromString(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Identifier Found!"})
		return
	}
	err = ah.Serv.DeleteUser(id)
	fmt.Println("line 141", err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "statusCode": http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK})
}

func (uph *userHandler) LoginUserHandler(c *gin.Context) {
	var user *model.User
	var token string
	var expiry int64
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println("line 115")
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error(), "statusCode": http.StatusBadRequest})
		c.Abort()
	}
	userDb, errs := uph.Serv.GetUserByID(user.ID)
	if errs != nil {
		fmt.Println("line 123")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "statusCode": http.StatusNotFound})
		c.Abort()
	}
	er := comparision.ComparePassword([]byte(user.Password), []byte(userDb.Password))
	if er == bcrypt.ErrMismatchedHashAndPassword {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error(), "statusCode": http.StatusUnprocessableEntity})
		c.Abort()
	}
	token, err = jwt.CreateToken(userDb.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error(), "statusCode": http.StatusUnprocessableEntity})
		c.Abort()
	}
	expiry = time.Now().Add(time.Minute * 30).Unix()
	expiryString := strconv.Itoa(int(expiry))
	c.Header("Content-Type", "application/json")
	c.Header("token", token)
	c.Header("expiry_date", expiryString)
	return
}
