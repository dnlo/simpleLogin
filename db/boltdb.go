package db

import (
	"github.com/asdine/storm"
	"simpleLogin/model"
	"log"
	"github.com/asdine/storm/q"
	"golang.org/x/crypto/bcrypt"
)

var db *storm.DB

func InitDB() {
	var err error
	db, err = storm.Open("accounts.db")
	if err != nil {
		log.Fatal(err)
	}
}

func GetUser(name string, password string) (*model.User, error) {
	var user model.User

	query := db.Select(q.Eq("Name", name))
	err := query.First(&user)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(name string, password string) (*model.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	newUser := model.User{Name: name, Password: string(hash)}
	err = db.Save(&newUser)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
