package controllers

import (
	"github.com/macococo/go-gamereviews/models"
	"github.com/macococo/go-gamereviews/utils"
	"net/http"
)

func ChatroomListController(w http.ResponseWriter, r *http.Request) {
	page := utils.GetParamInt(r, "page", 1)

	chatroomManager := models.ChatroomManager{}
	pagination := models.NewPagination(page, models.PAGINATION_DEFAULT_LENGTH, chatroomManager.Count())
	chatrooms := chatroomManager.Find(pagination)

	bytes := utils.ToJsonBytes(&chatrooms)

	utils.WriteJsonBytes(w, bytes)
}
