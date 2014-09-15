package main

import (
	"github.com/macococo/go-gamereviews/controllers"
	"net/http"
)

func initRouter() {
	http.HandleFunc("/api/user/list", controllers.UserListController)
	http.HandleFunc("/api/user/create", controllers.UserCreateController)

	http.ListenAndServe(":8080", nil)
}

func main() {
	initRouter()
}
