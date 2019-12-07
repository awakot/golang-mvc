package domain

import (
	"fmt"

	"net/http"

	"github.com/waytkheming/golang-mvc/utils"
)

var (
	users = map[int64]*User{
		0001: &User{ID: 0001, FirstNname: "hoge", LastName: "murakawa", Email: "hoge@example.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	user := users[userId]
	if user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}
