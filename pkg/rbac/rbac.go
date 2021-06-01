package rbac

import (
	_ "github.com/dgrijalva/jwt-go"
	"github.com/zpatrick/rbac"
)
var rbac_map=map[string]string{
	"/api/v1/login":"login to system",
	"/api/v1/users":"list all users",
	"/api/v1/orders":"list orders",
	"/api/v1/role":"list all roles",
	"/api/v1/stocks":"list stocks",
	"/api/v1/products":"list products",
	"/api/v1/suppliers":"list all suppliers",
}
func GetAction(targer_id string ) string {
	return rbac_map[targer_id]
}
//Admin permission
func GetPermission(roletype string)rbac.Role {

	roles := []rbac.Role{
		{
			RoleID: "ADMIN",
			Permissions: []rbac.Permission{
				rbac.NewGlobPermission("list stocks", "/api/v1/stocks"),
				rbac.NewGlobPermission("login to system", "/api/v1/login"),
				rbac.NewGlobPermission("list orders", "/api/v1/orders"),
				rbac.NewGlobPermission("list all users", "/api/v1/users"),
				rbac.NewGlobPermission("list roles", "/api/v1/role"),
				rbac.NewGlobPermission("list all suppliers", "/api/v1/suppliers"),
				rbac.NewGlobPermission("list products", "/api/v1/products"),
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
	if roletype=="ADMIN" {
		return roles[0]
	}
	return roles[1]
}