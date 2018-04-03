package db

import (
	"github.com/asdine/storm"
	"github.com/dnlo/web/simpleLogin/model"
	"log"
	_ "github.com/asdine/storm/q"
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

func AuthUser(name string, password string) error {
	var user model.User

	err := db.One("Name", name, &user)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

func GetUser(name string) (*model.User, error) {
	var user model.User

	err := db.One("Name", name, &user)
	if err != nil {
		return nil, err
	}
	
	return &user, nil
}

func GetUserList() (users []model.User) {
	db.All(&users)
	return
}

func CreateUser(name, password string, admin bool) (*model.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	
	newUser := model.User{Name: name, Password: string(hash), Admin: admin}
	err = db.Save(&newUser)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
