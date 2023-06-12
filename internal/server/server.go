package server

import (
	"employees/internal/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

type server struct {
	router  *gin.Engine
	handler handler.Handler
}

func newServer(handler handler.Handler) *server {
	server := &server{
		router:  gin.Default(),
		handler: handler,
	}

	server.configureRouter()

	return server
}

func (server *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	server.router.ServeHTTP(writer, request)
}

func (server *server) configureRouter() {
	auth := server.router.Group("/auth")
	{
		auth.POST("/signUp", server.handler.AuthHandler.SignUp)
		auth.POST("/signIn", server.handler.AuthHandler.SignIn)
	}

	api := server.router.Group("/api", server.handler.AuthMiddleware.Autorize)
	{
		api.POST("/employee", server.handler.EmployeeHandler.CreateEmployee)
		api.POST("/setDepartment", server.handler.EmployeeHandler.SetEmployeeDepartment)
		api.GET("/employee/:id", server.handler.EmployeeHandler.GetEmployee)
		api.GET("/employees", server.handler.EmployeeHandler.GetAllEmployees)
		api.PUT("/employee/:id", server.handler.EmployeeHandler.UpdateEmployee)
		api.DELETE("/employee/:id", server.handler.EmployeeHandler.DeleteEmployee)

		api.POST("/department", server.handler.DepartmentHandler.CreateDepartment)
		api.GET("/department/:id", server.handler.DepartmentHandler.GetDepartment)
		api.GET("/departments", server.handler.DepartmentHandler.GetAllDepartments)
		api.PUT("/department/:id", server.handler.DepartmentHandler.UpdateDepartment)
		api.DELETE("/department/:id", server.handler.DepartmentHandler.DeleteDepartment)
	}
}
