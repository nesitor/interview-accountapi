package pkg

import (
	"github.com/nesitor/interview-accountapi-form3/internal/accounts"
	"github.com/nesitor/interview-accountapi-form3/pkg/models"
)

type Accounts struct {
	repository accounts.Repository
}

func (acc *Accounts) Fetch(id string) (*models.Account, error) {
	return acc.repository.Get(id)
}

func (acc *Accounts) List(page int, size int) ([]*models.Account, error) {
	return acc.repository.List(page, size)
}

func (acc *Accounts) Create(account interface{}) (*models.Account, error) {
	return acc.repository.Create(account)
}

func (acc *Accounts) Delete(id string, version int) error {
	return acc.repository.Delete(id, version)
}

func NewAccounts(repository accounts.Repository) *Accounts {
	return &Accounts{repository: repository}
}
