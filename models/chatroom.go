package models

import (
	"github.com/macococo/go-gamereviews/modules"
	"github.com/macococo/go-gamereviews/utils"
	"strconv"
	"time"
)

type Chatroom struct {
	Id          int
	Code        string
	MaxMembers  int
	EndDatetime *time.Time
	Model
}

type ChatroomManager struct {
}

func (this *ChatroomManager) TableName() string {
	return "t_chatroom"
}

func (this *ChatroomManager) ListCacheKey(page int) string {
	return "chatroom_list_" + strconv.Itoa(page)
}

func (this *ChatroomManager) CountCacheKey() string {
	return "chatroom_count"
}

func (this *ChatroomManager) AddTable() {
	DbMap.AddTableWithName(Chatroom{}, this.TableName()).SetKeys(true, "Id").ColMap("Code").SetNotNull(true)
}

func (this *ChatroomManager) Count() int64 {
	return this.stats().Count
}

func (this *ChatroomManager) stats() (stats *ModelStats) {
	countKey := this.CountCacheKey()
	json := modules.AppCache.Get(countKey)
	if json == nil {
		count, err := DbMap.SelectInt("SELECT count(*) FROM " + this.TableName())
		utils.HandleError(err)

		stats = &ModelStats{Count: count}
		modules.AppCache.Put(countKey, utils.ToJsonBytes(&stats))
	} else {
		stats = &ModelStats{}
		utils.FromJsonBytes(json, &stats)
	}

	return stats
}

func (this *ChatroomManager) Find(pagination *Pagination) []*Chatroom {
	var chatrooms []*Chatroom
	_, err := DbMap.Select(&chatrooms, "SELECT * FROM "+this.TableName()+" LIMIT ?, ?", (pagination.Page-1)*pagination.Length, pagination.Length)
	utils.HandleError(err)

	return chatrooms
}

func (this *ChatroomManager) Create(chatroom *Chatroom) *Chatroom {
	utils.HandleError(DbMap.Insert(chatroom))
	modules.AppCache.Delete(this.CountCacheKey())

	return chatroom
}
