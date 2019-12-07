package services

import (
	"github.com/waytkheming/golang-mvc/domain"
	"github.com/waytkheming/golang-mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
