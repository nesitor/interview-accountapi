package http

import (
	"github.com/nesitor/interview-accountapi-form3/pkg/models"
)

type DataObjectStructure struct {
	Data  *models.Account  `json:"data"`
	Links ObjectPagination `json:"links,omitempty"`
}

type ObjectPagination struct {
	Self string `json:"self"`
}

type DataListStructure struct {
	Data  []*models.Account `json:"data"`
	Links ListPagination    `json:"links"`
}

type ListPagination struct {
	First    string `json:"first"`
	Last     string `json:"last"`
	Previous string `json:"previous,omitempty"`
	Next     string `json:"next,omitempty"`
	Self     string `json:"self"`
}
