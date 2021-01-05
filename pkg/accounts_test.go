package pkg

import (
	"fmt"
	"github.com/nesitor/interview-accountapi-form3/pkg/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestRepository struct {
	err error
}

func (tr *TestRepository) Get(id string) (*models.Account, error) {
	if tr.err != nil {
		return &models.Account{}, tr.err
	}

	account := models.Account{
		Id: id,
	}

	return &account, nil
}

func (tr *TestRepository) List(page int, size int) ([]*models.Account, error) {
	accounts := make([]*models.Account, 0, size)

	if tr.err != nil {
		return accounts, tr.err
	}

	account := models.Account{
		Id: "123",
	}
	accounts = append(accounts, &account)
	return accounts, nil
}

func (tr *TestRepository) Create(newAccount *models.Account) (*models.Account, error) {
	if tr.err != nil {
		return &models.Account{}, tr.err
	}

	return newAccount, nil
}

func (tr *TestRepository) Delete(id string, version int) error {
	if tr.err != nil {
		return tr.err
	}

	return nil
}

func NewTestRepository(err error) *TestRepository {
	repository := TestRepository{
		err: err,
	}
	return &repository
}

func TestAccounts_Fetch_Ok(t *testing.T) {
	repository := NewTestRepository(nil)

	api := NewAccounts(repository)

	id := "test123"
	account, err := api.Fetch(id)
	assert.NoError(t, err)
	assert.Equal(t, id, account.Id)
}

func TestAccounts_Fetch_Error(t *testing.T) {
	repository := NewTestRepository(fmt.Errorf("some error"))

	api := NewAccounts(repository)

	id := "test123"
	_, err := api.Fetch(id)
	assert.Error(t, err)
}

func TestAccounts_List_Ok(t *testing.T) {
	repository := NewTestRepository(nil)

	api := NewAccounts(repository)

	accounts, err := api.List(1, 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(accounts))
}

func TestAccounts_List_Error(t *testing.T) {
	repository := NewTestRepository(fmt.Errorf("some error"))

	api := NewAccounts(repository)

	_, err := api.List(1, 1)
	assert.Error(t, err)
}

func TestAccounts_Create_Ok(t *testing.T) {
	repository := NewTestRepository(nil)

	api := NewAccounts(repository)

	acc := models.Account{
		Id: "test123",
	}
	account, err := api.Create(&acc)
	assert.NoError(t, err)
	assert.Equal(t, acc.Id, account.Id)
}

func TestAccounts_Create_Error(t *testing.T) {
	repository := NewTestRepository(fmt.Errorf("some error"))

	api := NewAccounts(repository)

	acc := models.Account{
		Id: "test123",
	}
	_, err := api.Create(&acc)
	assert.Error(t, err)
}

func TestAccounts_Delete_Ok(t *testing.T) {
	repository := NewTestRepository(nil)

	api := NewAccounts(repository)

	id := "test123"
	err := api.Delete(id, 1)
	assert.NoError(t, err)
}

func TestAccounts_Delete_Error(t *testing.T) {
	repository := NewTestRepository(fmt.Errorf("some error"))

	api := NewAccounts(repository)

	id := "test123"
	err := api.Delete(id, 1)
	assert.Error(t, err)
}
