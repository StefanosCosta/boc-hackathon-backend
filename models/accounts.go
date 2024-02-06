package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
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
	SubscriptionID  string `json:"subscriptionID"`
}

type Transaction struct {
	gorm.Model
	AuditorId  uint       `gorm:"many2one:user_id;" json:"auditor_id"`
	BusinessId uint       `json:"business_id"`
	Key        string     `json:"key"`
	ExpiresAt  *time.Time `json:"expires_at"`
}

// Account Number, Account Type, Currency, Balance, Interest Rate, Maturity Date
type AccountBalancesResponse struct {
	AccountNumber string  `json:"accountNumber"`
	AccountType   string  `json:"accountType"`
	Currency      string  `json:"currency"`
	Balance       float64 `json:"balance"`
	InterestRate  float64 `json:"interestRate"`
	MaturityDate  string  `json:"maturityDate"`
	BalanceDate   string  `json:"balanceDate"`
}
