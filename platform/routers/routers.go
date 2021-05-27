package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rileyr/middleware"
	"github.com/rileyr/middleware/wares"
)

type Routers interface {
	Serve()
}

type Router struct {
	Method      string
	Path        string
	Handler     func(handle *gin.Context)
	MiddleWares []middleware.Middleware
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
	for _, router := range r.routers {
		if router.MiddleWares == nil {
			server.Handle(router.Method, router.Path, router.Handler)
		} else {
			s := middleware.NewStack()
			for _, middle := range router.MiddleWares {
				s.Use(middle)
			}
			s.Use(wares.RequestID)
			s.Use(wares.Logging)
			server.Handle(router.Method, router.Path, router.Handler)
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
