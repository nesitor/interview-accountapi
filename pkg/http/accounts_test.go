package http

import (
	"github.com/nesitor/interview-accountapi-form3/pkg/config"
	"github.com/nesitor/interview-accountapi-form3/pkg/models"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAccountsRepository_Get_Ok(t *testing.T) {
	repository := newRepository()

	expected := models.Account{
		Id:             "",
		OrganizationId: "",
		Attributes: models.AccountAttributes{
			Country:                 "",
			BaseCurrency:            "",
			BankId:                  0,
			BankIdCode:              "",
			AccountNumber:           "",
			Bic:                     "",
			Iban:                    "",
			CustomerId:              0,
			Name:                    nil,
			AlternativeNames:        nil,
			AccountClassification:   "",
			JointAccount:            false,
			AccountMatchingOptOut:   false,
			SecondaryIdentification: "",
			Switched:                false,
			Status:                  "",
		},
	}

	account, err := repository.Get("")
	assert.NoError(t, err)
	assert.Equal(t, expected, account)
}

func TestAccountsRepository_List_Ok(t *testing.T) {
	repository := newRepository()

	account := models.Account{
		Id:             "",
		OrganizationId: "",
		Attributes: models.AccountAttributes{
			Country:                 "",
			BaseCurrency:            "",
			BankId:                  0,
			BankIdCode:              "",
			AccountNumber:           "",
			Bic:                     "",
			Iban:                    "",
			CustomerId:              0,
			Name:                    nil,
			AlternativeNames:        nil,
			AccountClassification:   "",
			JointAccount:            false,
			AccountMatchingOptOut:   false,
			SecondaryIdentification: "",
			Switched:                false,
			Status:                  "",
		},
	}
	expected := make([]*models.Account, 1)
	expected = append(expected, &account)

	accounts, err := repository.List(1, 1)
	assert.NoError(t, err)
	assert.Equal(t, expected, accounts)
}

func TestAccountsRepository_Create_Ok(t *testing.T) {
	repository := newRepository()

	account := models.Account{
		Id:             "",
		OrganizationId: "",
		Attributes: models.AccountAttributes{
			Country:                 "",
			BaseCurrency:            "",
			BankId:                  0,
			BankIdCode:              "",
			AccountNumber:           "",
			Bic:                     "",
			Iban:                    "",
			CustomerId:              0,
			Name:                    nil,
			AlternativeNames:        nil,
			AccountClassification:   "",
			JointAccount:            false,
			AccountMatchingOptOut:   false,
			SecondaryIdentification: "",
			Switched:                false,
			Status:                  "",
		},
	}

	received, err := repository.Create(&account)
	assert.NoError(t, err)
	assert.Equal(t, account, received)
}

func TestAccountsRepository_Delete_Ok(t *testing.T) {
	repository := newRepository()

	err := repository.Delete("", 1)
	assert.NoError(t, err)
}

func newRepository() *accountsRepository {
	serverUrl := os.Getenv("TEST_SERVER_URL")
	configuration := config.Configuration{BaseUrl: serverUrl}
	return NewAccountsRepository(configuration)
}
