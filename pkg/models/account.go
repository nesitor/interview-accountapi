package models

type Account struct {
	Id             string            `json:"id"`
	Type           string            `json:"type"`
	OrganizationId string            `json:"organization_id"`
	Attributes     AccountAttributes `json:"attributes"`
}

type AccountAttributes struct {
	Country                 string   `json:"country"`
	BaseCurrency            string   `json:"id"`
	BankId                  int      `json:"bank_id"`
	BankIdCode              string   `json:"bank_id_code"`
	AccountNumber           string   `json:"account_number"`
	Bic                     string   `json:"bic"`
	Iban                    string   `json:"iban"`
	CustomerId              int      `json:"customer_id"`
	Name                    []string `json:"name"`
	AlternativeNames        []string `json:"alternative_name"`
	AccountClassification   string   `json:"account_classification"`
	JointAccount            bool     `json:"joint_account"`
	AccountMatchingOptOut   bool     `json:"account_matching_opt_out"`
	SecondaryIdentification string   `json:"secondary_identification"`
	Switched                bool     `json:"switched"`
	Status                  string   `json:"status"`
}

func NewAccount() *Account {
	return &Account{}
}
