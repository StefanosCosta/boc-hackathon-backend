package controllers

import (
	"OpenBankingAPI/config"
	corebanking "OpenBankingAPI/corebanking"
	dt "OpenBankingAPI/data"
	"OpenBankingAPI/database"
	"OpenBankingAPI/models"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type BusinessUser struct {
	Id             uint   `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       []byte `json:"-"`
	Role           string `json:"role"`
	HasAccess      bool   `json:"has_access"`
	SubscriptionId string `json:"subscription_id"`
	ClientId       string `json:"client_id"`
}

func GetBusinessCustomers(c *fiber.Ctx) error {
	var (
		user models.User
		err  error
	)
	if user, err = authenticateToken(c); err != nil {
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	if user.Role != "Auditor" {
		return errors.Errorf("User with id %d is not authorised to view business owners", user.Id)
	}

	var businessUsers []models.User
	err = dt.ByRole(&businessUsers, "Business")
	if err != nil {
		if errors.Is(errors.Cause(err), gorm.ErrRecordNotFound) {
			return c.JSON(businessUsers)
		} else {
			return err
		}
	}
	var bUsers []BusinessUser

	for _, bUser := range businessUsers {
		var access []models.AuditorRequests
		_ = dt.ByAuditorAndBusinessIdRequest(&access, user.Id, uint(bUser.Id))
		var newBUser BusinessUser
		newBUser.Id = bUser.Id
		newBUser.Email = bUser.Email
		newBUser.Name = bUser.Name
		newBUser.Role = bUser.Role
		if len(access) > 0 {
			newBUser.HasAccess = access[0].Accept
			newBUser.SubscriptionId = access[0].SubscriptionID
			newBUser.ClientId = config.ClientID
		}
		bUsers = append(bUsers, newBUser)
	}

	return c.JSON(bUsers)
}

func RequestAccess(c *fiber.Ctx) error {
	var (
		user models.User
		err  error
	)
	if user, err = authenticateToken(c); err != nil {
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	if user.Role != "Auditor" {
		return errors.Errorf("User with id %d is not authorised to view business owners", user.Id)
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
	err = dt.ByAuditorAndBusinessIdRequest(&access, user.Id, uint(userID))
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
			return errors.Errorf("Already Requested access")
		}
	}

	var token string

	if token, err = corebanking.GetToken(config.BaseURI, config.ClientID, config.ClientSecret); err != nil {
		log.Fatalf("Error: %v", err)
	} else {
		fmt.Println("Token:", token)
	}
	var subscriptionID string
	if subscriptionID, err = corebanking.CreateSubscription(config.BaseURI, token); err != nil {
		log.Fatalf("Error: %v", err)
	} else {
		fmt.Println("Subscription ID:", subscriptionID)
	}

	accessRequest := models.AuditorRequests{AuditorId: user.Id, BusinessId: uint(userID), SubscriptionID: subscriptionID}
	database.DB.Create(&accessRequest)

	templateFile := "Templates/RequestAccessEmail.html"
	templateContent, err := ioutil.ReadFile(templateFile)
	if err != nil {
		fmt.Println("Error reading template file:", err)
		return err
	}

	body := strings.ReplaceAll(string(templateContent), "[Auditor name]", user.Name)
	var businessClient models.User
	if err := dt.ById(&businessClient, strconv.FormatInt(userID, 10)); err != nil {
		return err
	}

	subject := "Email Access Authorisation"
	recipient := businessClient.Email // REPLACE WITH ACTUAL
	sender := "openbanking@finovex.com"

	SendEmail(subject, body, recipient, sender)

	return nil
}

func GetBalances(c *fiber.Ctx) error {
	var (
		token               string
		accounts            []models.AccountJSON
		accountTransactions models.AccountTransactions
		accountBalances     []models.AccountBalances
		err                 error
		user                models.User
		subscriptionStatus  string
	)

	if user, err = authenticateToken(c); err != nil {
		return nil
	}

	if user.Role != "Auditor" {
		return errors.Errorf("User with id %d is not authorised to view business owners", user.Id)
	}

	if token, err = corebanking.GetToken(config.BaseURI, config.ClientID, config.ClientSecret); err != nil {
		log.Fatalf("Error: %v", err)
		return nil
	} else {
		fmt.Println("Token:", token)
	}

	businessId := c.Params("id")

	var accessRequest models.AuditorRequests
	err = dt.ByAuditorAndBusinessID(&accessRequest, user.Id, businessId)
	if err != nil {
		if errors.Is(errors.Cause(err), gorm.ErrRecordNotFound) {
			return c.JSON(accessRequest)
		} else {
			return err
		}
	}
	subscriptionID := accessRequest.SubscriptionID

	if subscriptionStatus, err = corebanking.GetSubscriptionStatus(config.BaseURI, token, subscriptionID); err != nil {
		log.Fatalf("Error: %v", err)
		return nil
	}

	if subscriptionStatus != "ACTV" {
		subscriptionID = "Subid000001-1696681762806"
	}

	if accounts, err = corebanking.GetAccounts(config.BaseURI, token, subscriptionID); err != nil {
		log.Fatalf("Error: %v", err)
		return nil
	}

	var accountBalanceResponse []models.AccountBalancesResponse
	for _, account := range accounts {
		accountID := account.AccountID
		if accountBalances, err = corebanking.GetAccountBalances(config.BaseURI, token, subscriptionID, accountID); err != nil {
			log.Fatalf("Error: %v", err)
			return nil
		}

		if accountTransactions, err = corebanking.GetAccountTransactions(config.BaseURI, token, subscriptionID, accountID); err != nil {
			log.Fatalf("Error: %v", err)
			return nil
		}

		var accountBalance models.AccountBalancesResponse

		accountBalance.Balance = corebanking.CurrentBalance(accountBalances[0].Balances) - corebanking.SumOfTransactionsOfPreviousYear(accountTransactions)
		accountBalance.Balance = math.Round(accountBalance.Balance*100) / 100
		accountBalance.AccountNumber = account.AccountID
		accountBalance.AccountType = account.AccountType
		accountBalance.Currency = account.Currency
		accountBalance.MaturityDate = account.MaturityDate
		currentTime := time.Now()
		lastDayOfPreviousYear := time.Date(currentTime.Year()-1, time.December, 31, 0, 0, 0, 0, currentTime.Location())
		accountBalance.BalanceDate = lastDayOfPreviousYear.Format("02-01-2006")
		accountBalanceResponse = append(accountBalanceResponse, accountBalance)
	}
	return c.JSON(accountBalanceResponse)
}
