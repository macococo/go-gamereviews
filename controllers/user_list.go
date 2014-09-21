package controllers

import (
	"github.com/macococo/go-gamereviews/models"
	"github.com/macococo/go-gamereviews/modules"
	"github.com/macococo/go-gamereviews/utils"
	"net/http"
	"strconv"
)

func UserListController(w http.ResponseWriter, r *http.Request) {
	page := utils.GetParamInt(r, "page", 1)
	t := utils.GetParamInt(r, "type", 1)

	key := "user_list_" + strconv.Itoa(t) + "_" + strconv.Itoa(page)
	json := modules.AppCache.Get(key)
	if json == nil {
		userManager := models.UserManager{}
		pagination := models.NewPagination(page, models.PAGINATION_DEFAULT_LENGTH, userManager.Count(t))
		users := userManager.Find(t, pagination)

		bytes := utils.ToJsonBytes(&users)
		modules.AppCache.Put(key, bytes)

		utils.WriteJsonBytes(w, bytes)
	} else {
		utils.WriteJsonBytes(w, json)
	}

}
