package handler

import (
	"github.com/acassio/base-app-go/model"
	"github.com/acassio/base-app-go/service"
	"gopkg.in/macaron.v1"
)

func GetUser(ctx *macaron.Context) string {

	jsonResp := service.GetUser()

	ctx.Header().Set("Content-Type", "application/json")
	ctx.WriteHeader(200)

	return string(jsonResp)
}

func CreateUser(user model.User, ctx *macaron.Context) string {

	response := service.CreateUser(user)

	ctx.Header().Set("Content-Type", "application/json")
	ctx.WriteHeader(200)

	return string(response)
}
