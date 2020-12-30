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

func TestAccountsRepository_Get_Error(t *testing.T) {
	repository := newRepository()

	_, err := repository.Get("")
	assert.Error(t, err)
}

func newRepository() *accountsRepository {
	serverUrl := os.Getenv("TEST_SERVER_URL")
	configuration := config.Configuration{BaseUrl: serverUrl}
	return NewAccountsRepository(configuration)
}
