package models

import (
	"github.com/macococo/go-gamereviews/utils"
)

func init() {
	DbMap.AddTableWithName(User{}, "t_user").SetKeys(true, "Id")
}

type User struct {
	Id   int
	Name string
	Type int
}

type UserManager struct {
}

func (this *UserManager) Count(t int) int64 {
	count, err := DbMap.SelectInt("SELECT count(*) FROM t_user WHERE type = ?", t)
	utils.HandleError(err)

	return count
}

func (this *UserManager) Find(t int, pagination *Pagination) []*User {
	var users []*User
	_, err := DbMap.Select(&users, "SELECT * FROM t_user WHERE type = ? LIMIT ?, ?", t, (pagination.Page-1)*pagination.Length, pagination.Length)
	utils.HandleError(err)

	return users
}

func (this *UserManager) Create(user *User) *User {
	utils.HandleError(DbMap.Insert(user))
	return user
}
