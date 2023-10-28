package managers

import (
	"github.com/kreaeutersalz/fantasymma/database" // import the package that defines the database variable
	"github.com/kreaeutersalz/fantasymma/models"
)

func GetAccount(id string) (error, *models.Account) {
	err, account := database.GetAccount(id) // call the GetAccount function from the database package
	if err == nil {
		account.Password = ""
	}
	return err, account
}

func CreateAccount(account *models.Account) (error, *models.Account) {
	// hash password
	// check lenght of password
	// check if email is already in use
	// check if email is valid
	account.GroupIds = []int64{}
	account.Role = "user"
	err, account := database.CreateAccount(account) // call the CreateAccount function from the database package
	if err != nil {
		return err, nil
	}
	account.Password = ""
	return nil, account
}
