package models

type Balance struct {
	Amount      float64 `json:"amount"`
	BalanceType string  `json:"balanceType"`
}

type AccountJSON struct {
	BankID          string `json:"bankId"`
	AccountID       string `json:"accountId"`
	AccountAlias    string `json:"accountAlias"`
	AccountType     string `json:"accountType"`
	AccountName     string `json:"accountName"`
	IBAN            string `json:"IBAN"`
	Currency        string `json:"currency"`
	InfoTimeStamp   string `json:"infoTimeStamp"`
	MaturityDate    string `json:"maturityDate"`
	LastPaymentDate string `json:"lastPaymentDate"`
}

type TransactionJSON struct {
	ID                string `json:"id"`
	DcInd             string `json:"dcInd"`
	TransactionAmount struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	} `json:"transactionAmount"`
	Description     string `json:"description"`
	PostingDate     string `json:"postingDate"`
	ValueDate       string `json:"valueDate"`
	TransactionType string `json:"transactionType"`
}

type AccountTransactions struct {
	Account      AccountJSON       `json:"account"`
	Transactions []TransactionJSON `json:"transaction"`
}

type AccountBalances struct {
	AccountJSON
	Balances []Balance `json:"balances"`
}
