package accounts

import (
	"github.com/nesitor/interview-accountapi-form3/pkg/models"
)

type Repository interface {
	Get(id string) (*models.Account, error)
	List(page int, size int) ([]*models.Account, error)
	Create(newAccount *models.Account) (*models.Account, error)
	Delete(id string, version int) error
}
