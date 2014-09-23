package models

import (
	"github.com/macococo/go-gamereviews/utils"
)

type User struct {
	Id   int
	Name string
	Type int
	Model
}

type UserManager struct {
}

func (this *UserManager) TableName() string {
	return "t_user"
}

func (this *UserManager) AddTable() {
	DbMap.AddTableWithName(User{}, this.TableName()).SetKeys(true, "Id")
}

func (this *UserManager) Count(t int) int64 {
	count, err := DbMap.SelectInt("SELECT count(*) FROM "+this.TableName()+" WHERE type = ?", t)
	utils.HandleError(err)

	return count
}

func (this *UserManager) Find(t int, pagination *Pagination) []*User {
	var users []*User
	_, err := DbMap.Select(&users, "SELECT * FROM "+this.TableName()+" WHERE type = ? LIMIT ?, ?", t, (pagination.Page-1)*pagination.Length, pagination.Length)
	utils.HandleError(err)

	return users
}

func (this *UserManager) Create(user *User) *User {
	utils.HandleError(DbMap.Insert(user))
	return user
}
