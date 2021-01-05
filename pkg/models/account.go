package models

type Account struct {
	Id             string            `json:"id"`
	Type           string            `json:"type,omitempty"`
	OrganizationId string            `json:"organisation_id"`
	Version        int               `json:"version,omitempty"`
	Attributes     AccountAttributes `json:"attributes"`
}

type AccountAttributes struct {
	Country                 string   `json:"country"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	BankId                  string   `json:"bank_id,omitempty"`
	BankIdCode              string   `json:"bank_id_code,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	CustomerId              int      `json:"customer_id,omitempty"`
	Name                    []string `json:"name,omitempty"`
	AlternativeNames        []string `json:"alternative_name,omitempty"`
	AccountClassification   string   `json:"account_classification,omitempty"`
	JointAccount            bool     `json:"joint_account,omitempty"`
	AccountMatchingOptOut   bool     `json:"account_matching_opt_out,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Switched                bool     `json:"switched,omitempty"`
	Status                  string   `json:"status,omitempty"`
}

func NewAccount(id string, organizationId string) *Account {
	return &Account{
		Type:           "accounts",
		Id:             id,
		OrganizationId: organizationId,
		Attributes:     AccountAttributes{},
	}
}
