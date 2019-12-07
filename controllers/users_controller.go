package controllers

import (
	"net/http"
	"strconv"

	"encoding/json"

	"github.com/waytkheming/golang-mvc/services"
	"github.com/waytkheming/golang-mvc/utils"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userIDParam := req.URL.Query().Get("user_id")
	userID, err := strconv.ParseInt(userIDParam, 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "userid must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write([]byte(jsonValue))
		return
	}
	user, apiErr := services.GetUser(userID)
	if err != nil {
		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write([]byte(jsonValue))
		return
	}

	//return user to client
	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}
