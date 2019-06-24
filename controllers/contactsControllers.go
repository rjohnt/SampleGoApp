package controllers

import (
	"encoding/json"
	"github.com/rjohnt/SampleGoApp/models"
	"github.com/rjohnt/SampleGoApp/utils"
	"net/http"
)

var CreateContact = func(writer http.ResponseWriter, request *http.Request) {
	user := request.Context().Value("user").(uint)
	contact := &models.Contact{}

	err := json.NewDecoder(request.Body).Decode(contact)
	if err != nil {
		utils.Respond(writer, utils.Message(false, "Error while decoding request body."))
		return
	}

	contact.UserId = user
	response := contact.Create()
	utils.Respond(writer, response)
}

var GetContacts = func(writer http.ResponseWriter, request *http.Request) {
	userId := request.Context().Value("user").(uint)
	data := models.GetContacts(userId)
	response := utils.Message(true, "success")
	response["data"] = data
	utils.Respond(writer, response)
}
