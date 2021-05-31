package rbac

import (
"fmt"
_ "github.com/dgrijalva/jwt-go"
"github.com/gin-gonic/gin"
"github.com/zpatrick/rbac"
"net/http"
"time"
)
var rbac_map=map[string]string{
	"/api/v1/stocks":"list stocks",
	"/api/v1/login":"login -permission",
	"/api/v1/users":"list all users",

}
func GetAction(targer_id string ) string {
	return rbac_map[targer_id]
}

func GetPermission() []rbac.Role {

	roles := []rbac.Role{
		{
			RoleID: "ADMIN",
			Permissions: []rbac.Permission{
				rbac.NewGlobPermission("list stocks", "/api/v1/stocks"),
			},
		},
		{
			RoleID: "USER",
			Permissions: []rbac.Permission{
				rbac.NewGlobPermission("login -permission", "/api/v1/login"),
			    rbac.NewGlobPermission("list stocks", "/api/v1/stocks"),
				rbac.NewGlobPermission("list all users", "/api/v1/users"),
			},
		},
	}
	return roles
}











func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Token")

		if len(tokenHeader) == 0 {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Token header is missing"})
			return
		}

		//token, err := jwt.Parse(tokenHeader, func(t *jwt.Token) (interface{}, error) {
		//	// Don't forget to validate the alg is what you expect:
		//	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		//		return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		//	}
		//	if int64(t.Claims["exp"].(float64)) < time.Now().Unix() {
		//		return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		//	}
		//	return []byte(config.Config.Key), nil
		//})

		//if err != nil || !token.Valid {
		//	c.Abort()
		//	c.JSON(http.StatusUnauthorized, gin.H{"Error": err.Error()})
		//	return
		//}

		//c.Set("UserID", int64(token.Claims["id"].(float64)))
		//c.Set("AuthToken", token)

	}
}


//HandlerFunc
func indexHandler(c *gin.Context)  {
	fmt.Println("index")
	c.JSON(http.StatusOK, gin.H{
		"msg": "index",
	})
}

//Define a middleware
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	// timing
	start := time.Now()
	c.Next() //Call the subsequent processing function

	//c.Abort() //Prevent calling subsequent processing functions
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Printf("m1 out ...")
}

func m2(c *gin.Context)  {
	fmt.Println("m2 in ...")
	c.Next()  //Call the subsequent processing function
	fmt.Println("m2 out ...")
}

func main()  {
	r := gin.Default()

	//Click GET to view the source code discovery
	//GET(relativePath string, handlers ...HandlerFunc) IRoutes In addition, ... means that multiple functions of HandlerFunc type can be passed

	//r.GET("/index", m1, indexHandler)
	//r.GET("/shop", m1, func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "shop",
	//	})
	//})
	//r.GET("/user", m1,  func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "user",
	//	})
	//})

	// Click Use to view the source code to find Use(middleware ...HandlerFunc) IRoutes
	r.Use(m1, m2)   //Global registration middleware function m1 m2

	/*Access/index
	  Before executing indexHandler, execute the registered middleware. The general execution printing order is: m1 in -> m2 in -> index -> m2 out -> m1 out
	*/
	r.GET("/index", indexHandler)
	r.GET("/shop",  func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "shop",
		})
	})
	r.GET("/user",  func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "user",
		})
	})

	r.Run(":9090")
}

