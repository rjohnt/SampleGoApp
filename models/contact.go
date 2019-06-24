package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/rjohnt/SampleGoApp/utils"
)

type Contact struct {
	gorm.Model
	Name string `json:"name"`
	Phone string `json:"phone"`
	UserId uint `json:"user_id"`
}

/*
	This struct function validates the required params in the request body.
*/
func (contact *Contact) Validate() (map[string]interface{}, bool) {
	if contact.Name == "" {
		return utils.Message(false, "Contact name should be on the payload."), false
	}

	if contact.Phone == "" {
		return utils.Message(false, "Phone number should be on the payload."), false
	}

	if contact.UserId < 0 {
		return utils.Message(false, "User is not recognized."), false
	}

	return utils.Message(true, "success"), true
}

func (contact *Contact) Create() map[string]interface{} {
	if response, ok := contact.Validate(); !ok {
		return response
	}

	GetDB().Create(contact)

	response := utils.Message(true, "success")
	response["contact"] = contact
	return response
}

func GetContacts(id uint) []*Contact {
	contacts := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("id = ?", id).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return contacts
}