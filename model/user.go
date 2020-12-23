package model

import (
	"github.com/astaxie/beego/logs"
	"strconv"
	"strings"
)

type User struct {
	username string
	password string
	age      int
	sex      string
}

func NewUser() *User {
	return &User{}
}

func (u *User) SetUsername(username string) {
	u.username = username
}
func (u *User) SetPassword(password string) {
	u.password = password
}
func (u *User) SetAge(age int) {
	u.age = age
}
func (u *User) SetSex(sex string) {
	u.sex = sex
}

func (u *User) GetUsername() string {
	return u.username
}
func (u *User) GetPassword() string {
	return u.password
}
func (u *User) GetAge() int {
	return u.age
}
func (u *User) GetSex() string {
	return u.sex
}

// 格式化输出数据信息
func (u *User) ToString() string {
	return strings.Join(u.ToStringSlice(), ",")
}

func (u *User) ToStringSlice() []string {
	return []string{u.username, u.password, strconv.Itoa(u.age), u.sex}
}

func (u *User) Save() bool {
	err := rwdata("user", u)
	if err != nil {
		logs.Error("user 保存失败")
		return false
	}
	return true
}

func (u *User) All() (datas PrimaryModels, err error) {
	logs.Debug("user all")
	datas, err = rfdata("user", "username")
	if err != nil {
		logs.Error("user 获取失败")
	}
	return
}
