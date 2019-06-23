package controllers

import (
	"encoding/json"
	"github.com/rjohnt/SampleGoApp/models"
	"github.com/rjohnt/SampleGoApp/utils"
	"net/http"
)

var CreateAccount = func(writer http.ResponseWriter, request *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(request.Body).Decode(account)
	if err != nil {
		utils.Respond(writer, utils.Message(false, "Invalid Request"))
		return
	}

	response := account.Create()
	utils.Respond(writer, response)
}

var Authenticate = func(writer http.ResponseWriter, request *http.Request) {
	account := models.Account{}
	err := json.NewDecoder(request.Body).Decode(account)
	if err != nil {
		utils.Respond(writer, utils.Message(false, "Invalid request."))
		return
	}

	response := models.Login(account.Email, account.Password)
	utils.Respond(writer, response)
}