package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"os"
	"strings"
	"github.com/rjohnt/SampleGoApp/utils"
	"golang.org/x/crypto/bcrypt"
)

// JWT Claims Struct
type Token struct {
	UserId uint
	jwt.StandardClaims
}

// User Account Struct
type Account struct {
	gorm.Model
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token";sql:"-"`
}

// Validate Incoming User Details
func (account *Account) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(account.Email, "@") {
		return utils.Message(false, "Email address is required."), false
	}

	if len(account.Password) < 6 {
		return utils.Message(false, "Password is required."), false
	}

	// Validate Email Uniqueness
	temp := &Account{}

	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "Connection error.  Please retry."), false
	}

	if temp.Email != "" {
		return utils.Message(false, "Email address already in use by another user."), false
	}

	return utils.Message(false, "Requirement passed."), true
}

func (account *Account) Create() (map[string]interface{}) {
	if response, ok := account.Validate(); !ok {
		return response
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)

	if account.ID <= 0 {
		return utils.Message(false, "Failed to create account, Connection Error.")
	}

	// Create new JWT token for newly registered account.
	tokenStruct := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenStruct)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" // Delete password.

	response := utils.Message(true, "Account has been created.")
	response["account"] = account
	return response
}

func Login(email, password string) map[string]interface{} {
	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "Email address not found.")
		}
		return utils.Message(false, "Connection error.  Please retry.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return utils.Message(false, "Invalid login credentials.  Please try again.")
	}

	account.Password = ""

	tokenStruct := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenStruct)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	response := utils.Message(true, "Logged In.")
	response["account"] = account
	return response
}

func GetUser(userId uint) *Account {
	account := &Account{}
	GetDB().Table("accounts").Where("id = ?", userId).First(account)
	if account.Email == "" {
		return nil
	}

	account.Password = ""
	return account
}