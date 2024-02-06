package controllers

import (
	"OpenBankingAPI/config"
	"OpenBankingAPI/corebanking"
	dt "OpenBankingAPI/data"
	"OpenBankingAPI/database"
	"OpenBankingAPI/models"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func ProvideAccess(c *fiber.Ctx) error {
	var (
		err error
	)
	if _, err = authenticateToken(c); err != nil {
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	var data map[string]int64

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	userID := data["id"]
	if err != nil {
		return err
	}
	var access []models.AuditorRequests
	err = dt.ByRequestId(&access, uint(userID))
	if err != nil {
		if !errors.Is(errors.Cause(err), gorm.ErrRecordNotFound) {
			return c.JSON(fiber.Map{
				"message": "Server Error",
			})
		}
	}
	if len(access) > 0 {
		if access[0].Accept {
			return errors.Errorf("Already have access")
		} else {
			access[0].Accept = true
			database.DB.Save(&access)

		}
	} else {
		return errors.Errorf("No request was made")
	}

	return nil
}

type AuditorRequestsPayload struct {
	RequestId      uint   `json:"request_id"`
	AuditorId      uint   `json:"auditor_id"`
	AuditorEmail   string `json:"email"`
	Accept         bool   `json:"accepted"`
	SubscriptionID string `json:"subscription_id"`
}

func MapAuditorRequests(acReq []models.AuditorRequests) ([]AuditorRequestsPayload, error) {
	var (
		auditorRequests []AuditorRequestsPayload
	)
	for _, req := range acReq {
		var user models.User
		err := dt.ById(&user, strconv.Itoa(int(req.AuditorId)))
		if err != nil {
			return auditorRequests, err
		}
		auditorRequests = append(auditorRequests,
			AuditorRequestsPayload{RequestId: req.ID, AuditorId: req.AuditorId,
				AuditorEmail:   user.Email,
				Accept:         req.Accept,
				SubscriptionID: req.SubscriptionID})

	}
	return auditorRequests, nil
}

func GetAuditorsAccessRequests(c *fiber.Ctx) error {
	var (
		user models.User
		err  error
	)
	if user, err = authenticateToken(c); err != nil {
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	var access []models.AuditorRequests
	err = dt.ByBusinessIdRequest(&access, user.Id)
	if err != nil {
		if !errors.Is(errors.Cause(err), gorm.ErrRecordNotFound) {
			return c.JSON(fiber.Map{
				"message": "Server Error",
			})
		}
	}
	auditorRequests, err := MapAuditorRequests(access)
	if err != nil {
		return err
	}

	return c.JSON(auditorRequests)
}

//func MapAccountJSONtoAccount(accountsJSON []models.AccountJSON, subscriptionID string) []models.Account {
//	var (
//		accounts []models.Account
//	)
//	for _, acc := range accountsJSON {
//		var dbAccount models.Account
//		dbAccount.AccountID = acc.AccountID
//		dbAccount.AccountAlias = acc.AccountAlias
//		dbAccount.AccountName = acc.AccountName
//		dbAccount.AccountType = acc.AccountType
//		dbAccount.BankID = acc.BankID
//		dbAccount.Currency = acc.Currency
//		dbAccount.IBAN = acc.IBAN
//		dbAccount.InfoTimeStamp = acc.InfoTimeStamp
//		dbAccount.LastPaymentDate = acc.LastPaymentDate
//		dbAccount.MaturityDate = acc.MaturityDate
//		accounts = append(accounts, dbAccount)
//	}
//	return accounts
//
//}

func storeAccountData(subscriptionID string) error {
	var token string
	var err error
	var accounts []models.AccountJSON

	if token, err = corebanking.GetToken(config.BaseURI, config.ClientID, config.ClientSecret); err != nil {
		//log.Fatalf("Error: %v", err)#
		return err
	} else {

		fmt.Println("Token:", token)
	}

	// if subscriptionID, err = corebanking.CreateSubscription(config.BaseURI, token); err != nil {
	// 	log.Fatalf("Error: %v", err)
	// } else {
	// 	fmt.Println("Subscription ID:", subscriptionID)
	// }

	subscriptionID = "Subid000001-1696681762806"
	// accountID = "351092345676"

	if accounts, err = corebanking.GetAccounts(config.BaseURI, token, subscriptionID); err != nil {
		log.Fatalf("Error: %v", err)
	} else {
		database.DB.Create(accounts)
	}

	for _, account := range accounts {
		accountID := account.AccountID
		fmt.Println("Account ID:", accountID)

		if accountBalances, err := corebanking.GetAccountBalances(config.BaseURI, token, subscriptionID, accountID); err != nil {
			log.Fatalf("Error: %v", err)
		} else {
			fmt.Println("Account Balances:", accountBalances)
		}

		if accountTransactions, err := corebanking.GetAccountTransactions(config.BaseURI, token, subscriptionID, accountID); err != nil {
			log.Fatalf("Error: %v", err)
		} else {
			fmt.Println("Account Transactions:", accountTransactions)
		}
	}
	return nil
}
