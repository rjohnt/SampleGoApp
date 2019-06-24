package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rjohnt/SampleGoApp/models"
	"github.com/rjohnt/SampleGoApp/utils"
	"net/http"
	"strconv"
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

var GetContactsFor = func(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(writer, utils.Message(false, "There was an error in your request."))
		return
	}

	data := models.GetContacts(uint(id))
	response := utils.Message(true, "success")
	response["data"] = data
	utils.Respond(writer, response)
}
