package main

import (
	"fmt"
	"github.com/nesitor/interview-accountapi-form3/pkg"
	"github.com/nesitor/interview-accountapi-form3/pkg/config"
	"github.com/nesitor/interview-accountapi-form3/pkg/http"
	"github.com/nesitor/interview-accountapi-form3/pkg/models"
	"log"
	"os"
)

func main() {

	configuration := config.Configuration{
		BaseUrl: os.Getenv("SERVER_URL"),
	}

	repository := http.NewAccountsRepository(configuration)

	accounts := pkg.NewAccounts(repository)

	account, err := accounts.Fetch("")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account)

	accountsList, err := accounts.List(1, 5)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(accountsList)

	newAccount := models.Account{}
	account, err = accounts.Create(&newAccount)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account)

	err = accounts.Delete("", 0)
	if err != nil {
		log.Fatalln(err)
	}
}
