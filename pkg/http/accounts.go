package http

import (
	"encoding/json"
	"github.com/nesitor/interview-accountapi-form3/internal/http"
	"github.com/nesitor/interview-accountapi-form3/pkg/config"
	"github.com/nesitor/interview-accountapi-form3/pkg/models"
	"strconv"
)

type AccountsRepository struct {
	configuration config.Configuration
	basePath      string
}

func (ar *AccountsRepository) Get(id string) (*models.Account, error) {
	var account *models.Account
	entireUrl := ar.configuration.BaseUrl + ar.basePath + "/" + id

	request := http.NewRequest()
	response, err := request.Get(entireUrl)
	if err != nil {
		return account, err
	}

	model, _ := json.Marshal(response)
	err = json.Unmarshal(model, &account)
	if err != nil {
		return account, err
	}

	return account, nil
}

func (ar *AccountsRepository) List(page int, size int) ([]*models.Account, error) {
	account := make([]*models.Account, 0, size)
	entireUrl := ar.configuration.BaseUrl + ar.basePath + "?page[number]=" + strconv.Itoa(page) + "&page[size]=" + strconv.Itoa(size)

	request := http.NewRequest()
	response, err := request.Get(entireUrl)
	if err != nil {
		return account, err
	}

	model, _ := json.Marshal(response)
	err = json.Unmarshal(model, &account)
	if err != nil {
		return account, err
	}

	return account, nil
}

func (ar *AccountsRepository) Create(newAccount *models.Account) (*models.Account, error) {
	var account *models.Account
	entireUrl := ar.configuration.BaseUrl + ar.basePath

	request := http.NewRequest()
	response, err := request.Post(entireUrl, newAccount)
	if err != nil {
		return account, err
	}

	model, _ := json.Marshal(response)
	err = json.Unmarshal(model, &account)
	if err != nil {
		return account, err
	}

	return account, nil
}

func (ar *AccountsRepository) Delete(id string, version int) error {
	entireUrl := ar.configuration.BaseUrl + ar.basePath + "/" + id + "?version=" + strconv.Itoa(version)

	request := http.NewRequest()
	err := request.Delete(entireUrl)
	if err != nil {
		return err
	}

	return nil
}

func NewAccountsRepository(configuration config.Configuration) *AccountsRepository {
	return &AccountsRepository{
		configuration: configuration,
	}
}
