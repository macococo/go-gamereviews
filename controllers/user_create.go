package controllers

import (
	"github.com/macococo/go-gamereviews/models"
	"github.com/macococo/go-gamereviews/utils"
	"math/rand"
	"net/http"
)

func UserCreateController(w http.ResponseWriter, r *http.Request) {
	userManager := models.UserManager{}
	user := models.User{Type: rand.Intn(3)}
	userManager.Create(&user)

	utils.WriteJson(w, &user)
}
