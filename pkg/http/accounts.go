package http

import (
	"encoding/json"
	"fmt"
	"github.com/nesitor/interview-accountapi-form3/internal/http"
	"github.com/nesitor/interview-accountapi-form3/pkg/config"
	"github.com/nesitor/interview-accountapi-form3/pkg/models"
	"strconv"
)

type accountsRepository struct {
	configuration config.Configuration
	basePath      string
}

func (ar *accountsRepository) Get(id string) (*models.Account, error) {
	account := DataObjectStructure{}
	entireUrl := ar.configuration.BaseUrl + ar.basePath + "/" + id

	request := http.NewClient()
	response, err := request.Get(entireUrl)
	if err != nil {
		return account.Data, err
	}

	err = json.Unmarshal([]byte(response), &account)
	if err != nil {
		return account.Data, err
	}

	return account.Data, nil
}

func (ar *accountsRepository) List(page int, size int) ([]*models.Account, error) {
	accounts := DataListStructure{}
	entireUrl := ar.configuration.BaseUrl + ar.basePath + "?page[number]=" + strconv.Itoa(page) + "&page[size]=" + strconv.Itoa(size)

	request := http.NewClient()
	response, err := request.Get(entireUrl)
	if err != nil {
		return accounts.Data, err
	}
	fmt.Println(response)
	err = json.Unmarshal([]byte(response), &accounts)
	if err != nil {
		return accounts.Data, err
	}

	return accounts.Data, nil
}

func (ar *accountsRepository) Create(newAccount *models.Account) (*models.Account, error) {
	account := DataObjectStructure{}
	newData := DataObjectStructure{Data: newAccount}
	entireUrl := ar.configuration.BaseUrl + ar.basePath

	request := http.NewClient()
	response, err := request.Post(entireUrl, newData)
	if err != nil {
		return account.Data, err
	}

	err = json.Unmarshal([]byte(response), &account)
	if err != nil {
		return account.Data, err
	}

	return account.Data, nil
}

func (ar *accountsRepository) Delete(id string, version int) error {
	entireUrl := ar.configuration.BaseUrl + ar.basePath + "/" + id + "?version=" + strconv.Itoa(version)

	request := http.NewClient()
	err := request.Delete(entireUrl)
	if err != nil {
		return err
	}

	return nil
}

func NewAccountsRepository(configuration config.Configuration) *accountsRepository {
	return &accountsRepository{
		configuration: configuration,
		basePath:      "/accounts",
	}
}
