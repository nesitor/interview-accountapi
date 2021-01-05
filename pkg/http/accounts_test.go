package http

import (
	"github.com/nesitor/interview-accountapi-form3/pkg/config"
	"github.com/nesitor/interview-accountapi-form3/pkg/models"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAccountsRepository_Create_Ok(t *testing.T) {
	repository := newRepository()

	account := models.NewAccount("ad27e265-9605-4b4b-a0e5-3003ea9cc4de", "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")

	account.Attributes = models.AccountAttributes{
		Country:          "GB",
		BaseCurrency:     "GBP",
		BankId:           "400300",
		BankIdCode:       "GBDSC",
		Bic:              "NWBKGB22",
		AlternativeNames: nil,
	}

	received, err := repository.Create(account)
	assert.NoError(t, err)
	assert.Equal(t, account, received)
}

func TestAccountsRepository_Get_Ok(t *testing.T) {
	repository := newRepository()

	expected := &models.Account{
		Id:             "ad27e265-9605-4b4b-a0e5-3003ea9cc4de",
		OrganizationId: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Type:           "accounts",
		Attributes: models.AccountAttributes{
			Country:          "GB",
			BaseCurrency:     "GBP",
			BankId:           "400300",
			BankIdCode:       "GBDSC",
			Bic:              "NWBKGB22",
			AlternativeNames: nil,
		},
	}

	account, err := repository.Get("ad27e265-9605-4b4b-a0e5-3003ea9cc4de")
	assert.NoError(t, err)
	assert.Equal(t, expected, account)
}

func TestAccountsRepository_List_Ok(t *testing.T) {
	repository := newRepository()

	account := &models.Account{
		Id:             "ad27e265-9605-4b4b-a0e5-3003ea9cc4de",
		OrganizationId: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Type:           "accounts",
		Attributes: models.AccountAttributes{
			Country:          "GB",
			BaseCurrency:     "GBP",
			BankId:           "400300",
			BankIdCode:       "GBDSC",
			Bic:              "NWBKGB22",
			AlternativeNames: nil,
		},
	}

	expected := make([]*models.Account, 1)
	expected[0] = account

	accounts, err := repository.List(2, 1)
	assert.NoError(t, err)
	assert.Equal(t, expected, accounts)
}

func TestAccountsRepository_Delete_Ok(t *testing.T) {
	repository := newRepository()

	err := repository.Delete("ad27e265-9605-4b4b-a0e5-3003ea9cc4de", 0)
	assert.NoError(t, err)
}

func newRepository() *accountsRepository {
	serverUrl := os.Getenv("TEST_SERVER_URL")
	configuration := config.Configuration{BaseUrl: serverUrl}
	return NewAccountsRepository(configuration)
}
