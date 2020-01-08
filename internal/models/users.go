package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     *string `gorm:"column:name;"`
	Email    *string `gorm:"column:email;"`
	Password *string `gorm:"column:password;"`
	Token    *string `gorm:"column:token;"`
}

func NewUser() *User {
	return &User{}
}

func (*User) TableName() string {
	return "users"
}

// AddUser to add
func (MysqlDBConn *Mysqldb) AddUser(name, email, password, token *string) error {
	user := &User{
		Name:     name,
		Email:    email,
		Password: password,
		Token:    token,
	}
	db := MysqlDBConn.db
	e := db.Debug().Create(user)
	if e.Error != nil {
		return errors.New("DB Operation Failed")
	}
	return nil
}

func createTable(db *gorm.DB) error {
	e := db.Debug().CreateTable(&User{})
	if e.Error != nil {
		return errors.New("DB Operation Failed")
	}
	return nil
}

func (MysqlDBConn *Mysqldb) HasTable() error {
	db := MysqlDBConn.db
	e := db.HasTable("users")
	if !e {
		createTable(db)
	}
	return nil
}

//GetUser get
func (MysqlDBConn *Mysqldb) GetUser(email, password *string) (*User, error) {
	db := MysqlDBConn.db
	user := &User{}
	e := db.Table("users").Where("email = ?", *email).Find(user)
	if e.Error != nil {
		return nil, errors.New("Not Found")
	}
	if *user.Password != *password {
		return nil, errors.New("Unauthorized")
	}
	return user, nil
}

func (MysqlDBConn *Mysqldb) GetUserByEmail(email *string) (*User, error) {
	db := MysqlDBConn.db
	user := &User{}
	e := db.Table("users").Where("email = ?", *email).Find(user)
	if e.Error != nil {
		if e.Error.Error() == "record not found" {
			return nil, errors.New("Not Found")
		} else {
			return nil, e.Error
		}
	}
	return user, nil
}

func (MysqlDBConn *Mysqldb) Update(email, name, password *string) error {
	db := MysqlDBConn.db
	user := &User{}
	e := db.Model(&user).Updates(User{Name: name, Password: password})
	if e.Error != nil {
		return errors.New("Not Found")
	}
	if *user.Password != *password {
		return errors.New("Unauthorized")
	}
	return nil
}
