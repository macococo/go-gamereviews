package controllers

import (
	"github.com/macococo/go-gamereviews/models"
	"github.com/macococo/go-gamereviews/utils"
	"net/http"
	"time"
)

func ChatroomCreateController(w http.ResponseWriter, r *http.Request) {
	chatroomManager := models.ChatroomManager{}

	// 1 hour
	endDatetime := time.Now().Add(time.Duration(1) * time.Hour)

	chatroom := models.Chatroom{Code: utils.GetGUID(), MaxMembers: 5, EndDatetime: &endDatetime}
	chatroomManager.Create(&chatroom)

	utils.WriteJson(w, &chatroom)
}
