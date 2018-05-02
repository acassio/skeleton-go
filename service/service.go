package service

import (
	"encoding/json"
	"fmt"

	"github.com/acassio/base-app-go/model"
)

//GetUser obtem obj user
func GetUser() (js []byte) {

	user := model.User{Name: "João", Type: "guest"}

	js, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Erro conversao de dados json: " + err.Error())
	}

	return
}

//CreateUser salva e obtem obj user
func CreateUser(user model.User) (js []byte) {

	fmt.Println("Creating user: ", user)

	js, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Erro conversao de dados json: " + err.Error())
	}

	return
}
