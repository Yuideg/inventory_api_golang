package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Routers interface {
	Serve()
}

type Router struct {
	Method      string
	Path        string
	Authenticated  func(handle *gin.Context)
	Authorized  func(handle *gin.Context)
	Handler     func(handle *gin.Context)

}
type routing struct {
	host    string
	port    string
	routers []Router
}

func NewRouting(host, port string, routers []Router) Routers {
	return &routing{
		host,
		port,
		routers,
	}
}

func (r *routing) Serve() {
	server := gin.Default()
	api:=server.Group("api/v1")
	for _, router := range r.routers {
		if router.Authenticated == nil && router.Authorized==nil {
			fmt.Println("login start line 39")
			api.Handle(router.Method, router.Path , router.Handler)
		} else {
			api.Handle(router.Method, router.Path,router.Authenticated, router.Authorized, router.Handler)
		}
	}
	addr := fmt.Sprintf("%s:%s", r.host, r.port)
	fmt.Println(addr)
	err := server.Run(addr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("server started at %s:%s", r.host, r.port)
}
