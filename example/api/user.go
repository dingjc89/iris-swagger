package api

import (
	"github.com/kataras/iris/v12"
	"strconv"
)

type UserController struct {
}

type User struct {
	ID      int    `json:"id"`
	Account string `json:"account"`
	Name    string `json:"name"`
	Pass    string `json:"pass"`
}

// Detail get user detail by user id
// @Description get user detail information
// @Summary user information
// @Param user_id query string true "user id"
// @Success 200  {object} api.User `{"id":1, "account":"kevin", "name":"ding", "pass":"123123"}`
// @Router /detail [get]
func (u *UserController) Detail(ctx iris.Context) {
	userId := ctx.FormValue("user_id")
	userID, _ := strconv.Atoi(userId)
	ctx.JSON(&User{ID: userID, Account: "kevin", Name: "ding", Pass: "123123"})
}
