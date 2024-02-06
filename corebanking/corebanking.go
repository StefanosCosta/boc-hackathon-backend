package corebanking

import (
	"OpenBankingAPI/models"
	"OpenBankingAPI/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func GetToken(baseURI, clientID, clientSecret string) (string, error) {
	url := baseURI + "/oauth2/token"
	method := "POST"

	payload := fmt.Sprintf("client_id=%s&client_secret=%s&grant_type=client_credentials&scope=TPPOAuth2Security", clientID, clientSecret)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(payload))

	if err != nil {
		return "", err
	}

	req.Header.Add("x-client-certificate", "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")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", "TS013b36ab=0179594e11530f14ff6c39ec7222f89f895c4261a58fb51be61bb5dda3b5026acbd858e003f52cfcc44451d2f9604c6b5b5509e2a7b36a0815bc19f61103d77a96697aa2b1; de2a657d1673ca26a0e0abed5da67a83=5649bec0e006d3b80f439de2aa4ba934")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// Parse the response JSON
	var tokenData map[string]interface{}
	json.Unmarshal(body, &tokenData)
	// Print the access token
	fmt.Println("Access Token:", tokenData["access_token"])
	fmt.Println("Error:", tokenData["error"])

	accessToken, ok := tokenData["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("access_token is not a string")
	}

	if tokenData["error"] != nil {
		err = fmt.Errorf("%v", tokenData["error"])
	}

	return accessToken, err
}

func CreateSubscription(baseURI, token string) (string, error) {
	url := baseURI + "/v1/subscriptions"
	method := "POST"

	payload := strings.NewReader(`{
	 "accounts": {
		"transactionHistory": true,
		"balance": true,
		"details": true,
		"checkFundsAvailability": true
	  },
	  "payments": {
		"limit": 99999999,
		"currency": "EUR",
		"amount": 999999999
	  }
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("originUserId", "50520222")
	req.Header.Add("timeStamp", "1696687049")
	req.Header.Add("journeyId", "0a02fc86-a019-48d1-a661-36b9a435dcad")
	req.Header.Add("x-client-certificate", "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")
	req.Header.Add("Cookie", "TS013b36ab=0179594e11ab88ea14f5b8a2af53b8762c8d2ffa9251842ced965ff41c58bcc2500ece06fc19c25dba513c80efd4d300b7856422b061a563b86e86f0cb013618deb59fc23a; de2a657d1673ca26a0e0abed5da67a83=5649bec0e006d3b80f439de2aa4ba934")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Parse the response JSON
	var subscriptionData map[string]interface{}
	json.Unmarshal(body, &subscriptionData)

	subscriptionID, ok := subscriptionData["subscriptionId"].(string)
	if !ok {
		return "", fmt.Errorf("subscriptionId is not a string")
	}

	if subscriptionData["error"] != nil {
		err = fmt.Errorf("%v", subscriptionData["error"])
	}

	return subscriptionID, err
}

func GetAccounts(baseURI, token, subscriptionID string) ([]models.AccountJSON, error) {
	url := baseURI + "/v1/accounts"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("subscriptionId", subscriptionID)
	req.Header.Add("originUserId", "50520222")
	req.Header.Add("journeyId", "8c3f34b6-aecc-4454-a0c6-8aaca235bd3b")
	req.Header.Add("timeStamp", "1696692025")
	req.Header.Add("x-client-certificate", "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")
	req.Header.Add("Cookie", "TS013b36ab=0179594e119081fbb635e4d68afcbd27572e3f3d02577d74df8c964496ab33abb48613fa4914dea20fcf71b727bd4e9f8a7669db9860bf678da9896f9cebacfb5bb8203a10; de2a657d1673ca26a0e0abed5da67a83=5649bec0e006d3b80f439de2aa4ba934")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Parse the response JSON
	var accounts []models.AccountJSON
	json.Unmarshal(body, &accounts)

	if accounts == nil {
		return nil, fmt.Errorf("accounts is nil")
	}

	return accounts, nil
}

func GetAccountBalances(baseURI, token, subscriptionID string, accountID string) ([]models.AccountBalances, error) {
	url := baseURI + "/v1/accounts/" + accountID + "/balance"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("subscriptionId", subscriptionID)
	req.Header.Add("originUserId", "50520222")
	req.Header.Add("journeyId", "8c3f34b6-aecc-4454-a0c6-8aaca235bd3b")
	req.Header.Add("timeStamp", "1696692025")
	req.Header.Add("x-client-certificate", "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")
	req.Header.Add("Cookie", "TS013b36ab=0179594e119081fbb635e4d68afcbd27572e3f3d02577d74df8c964496ab33abb48613fa4914dea20fcf71b727bd4e9f8a7669db9860bf678da9896f9cebacfb5bb8203a10; de2a657d1673ca26a0e0abed5da67a83=5649bec0e006d3b80f439de2aa4ba934")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Parse the response JSON
	var accountBalances []models.AccountBalances
	json.Unmarshal(body, &accountBalances)

	if accountBalances == nil {
		return nil, fmt.Errorf("accountBalances is nil")
	}

	return accountBalances, nil
}

func GetAccountTransactions(baseURI, token, subscriptionID string, accountID string) (models.AccountTransactions, error) {
	url := baseURI + "/v1/accounts/" + accountID + "/statement"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return models.AccountTransactions{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("subscriptionId", subscriptionID)
	req.Header.Add("originUserId", "50520222")
	req.Header.Add("journeyId", "8c3f34b6-aecc-4454-a0c6-8aaca235bd3b")
	req.Header.Add("timeStamp", "1696692025")
	req.Header.Add("x-client-certificate", "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")
	req.Header.Add("Cookie", "TS013b36ab=0179594e119081fbb635e4d68afcbd27572e3f3d02577d74df8c964496ab33abb48613fa4914dea20fcf71b727bd4e9f8a7669db9860bf678da9896f9cebacfb5bb8203a10; de2a657d1673ca26a0e0abed5da67a83=5649bec0e006d3b80f439de2aa4ba934")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return models.AccountTransactions{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return models.AccountTransactions{}, err
	}

	// Parse the response JSON
	var accountTransactions models.AccountTransactions
	json.Unmarshal(body, &accountTransactions)

	return accountTransactions, nil
}

func GetSubscriptionStatus(baseURI, token, subscriptionID string) (string, error) {
	url := baseURI + "/v1/subscriptions/" + subscriptionID
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("originUserId", "50520222")
	req.Header.Add("timeStamp", "1696749669")
	req.Header.Add("journeyId", "10b794a2-a518-44c2-9fc0-239c4a477787")
	req.Header.Add("x-client-certificate", "MIIH4TCCBcmgAwIBAgIUVjRXOiJ9y9zF+zhFqE1XIvpuYuIwDQYJKoZIhvcNAQELBQAwgYUxCzAJBgNVBAYTAklUMRgwFgYDVQQKDA9JbmZvQ2VydCBTLnAuQS4xIzAhBgNVBAsMGldTQSBUcnVzdCBTZXJ2aWNlIFByb3ZpZGVyMTcwNQYDVQQDDC5JbmZvQ2VydCBPcmdhbml6YXRpb24gVmFsaWRhdGlvbiBTSEEyNTYgLSBDQSAzMB4XDTIxMTAyMjEzNDcxM1oXDTIzMTAyMjAwMDAwMFowgfwxEzARBgsrBgEEAYI3PAIBAxMCQ1kxHTAbBgNVBA8MFFByaXZhdGUgT3JnYW5pemF0aW9uMRgwFgYDVQRhDA9QU0RDWS1DQkMtSEUxNjUxCzAJBgNVBAYTAkNZMSowKAYDVQQKDCFCQU5LIE9GIENZUFJVUyBQVUJMSUMgQ09NUEFOWSBMVEQxEDAOBgNVBAgMB05JQ09TSUExEDAOBgNVBAcMB05pY29zaWExGDAWBgNVBAsMD1JFVEFJTCBESVZJU0lPTjEOMAwGA1UEBRMFSEUxNjUxJTAjBgNVBAMMHGFwaXMtc2VjdXJlLmJhbmtvZmN5cHJ1cy5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDFahBqfOZ/gYuFVha3A6/Z/PXjVR9G88ztvjMAAR6zXzgj/VKnMi811ukk5Gv8JhHO004BSVYvuIVmX1aytSeQWhMblsEp/Q07pMjCplDCJxtV7vBtIm5E4aNZ172vYIoSiIcFbbBpF771ZfuwT47uA6UZc1y2te3hRgFGrB8C/jPOx/1MRPHS56vH3w8xyqbrEkK5ByOkztsTJ7xkILisLLhKN0ovyLXLcXAbSOqH+5jKjsTpvqaJFjUkCYdAbC9V+ecPbwsuoqu4oVn5DtJUhzs3HKp5ty+Xa7nJ/ShnaWvlbmlfXfNk/EmZxHaLm8RMreUbiZYab06FHDn+sZZDAgMBAAGjggLOMIICyjBxBggrBgEFBQcBAQRlMGMwLAYIKwYBBQUHMAGGIGh0dHA6Ly9vY3NwLm92Y2EuY2EzLmluZm9jZXJ0Lml0MDMGCCsGAQUFBzAChidodHRwOi8vY2VydC5pbmZvY2VydC5pdC9jYTMvb3ZjYS9DQS5jcnQwOgYDVR0fBDMwMTAvoC2gK4YpaHR0cDovL2NybC5pbmZvY2VydC5pdC9jYTMvb3ZjYS9DUkwwMS5jcmwwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMIH3BggrBgEFBQcBAwSB6jCB5zAIBgYEAI5GAQEwCwYGBACORgEDAgEUMBMGBgQAjkYBBjAJBgcEAI5GAQYDMD8GBgQAjkYBBTA1MDMWLWh0dHBzOi8vd3d3LmZpcm1hLmluZm9jZXJ0Lml0L3BkZi9QS0ktU1NMLnBkZhMCZW4weAYGBACBmCcCMG4wTDARBgcEAIGYJwEBDAZQU1BfQVMwEQYHBACBmCcBAgwGUFNQX1BJMBEGBwQAgZgnAQMMBlBTUF9BSTARBgcEAIGYJwEEDAZQU1BfSUMMFkNlbnRyYWwgQmFuayBvZiBDeXBydXMMBkNZLUNCQzBmBgNVHSAEXzBdMAkGBwQAi+xAAQQwUAYHK0wkAQEtBDBFMEMGCCsGAQUFBwIBFjdodHRwOi8vd3d3LmZpcm1hLmluZm9jZXJ0Lml0L2RvY3VtZW50YXppb25lL21hbnVhbGkucGhwMA4GA1UdDwEB/wQEAwIFoDAnBgNVHREEIDAeghxhcGlzLXNlY3VyZS5iYW5rb2ZjeXBydXMuY29tMB8GA1UdIwQYMBaAFAcL9d6GcvxHreaSNOQviuahp7laMB0GA1UdDgQWBBRpue2+nnlK/7a7QgzbUrbg6EDVljAfBgVngQwDAQQWMBQTA1BTRBMCQ1kMCUNCQy1IRTE2NTANBgkqhkiG9w0BAQsFAAOCAgEAZ5TJa1xf3VC76zmLQNsngZPCSg90h38HI3SiGTehxBn5xdFYhAaA8XhlR5BQBAwQraD+qDpogq2fwF+ciA/urFEuSzuY8X3tpULsQimgkEX6456TcCwl6NkF4UUcD2fcKEmHvKnHmyFFGSoVU9gdQRu+5znm/bfYLcjiG79xjEfkM68OgTTx2z+F5tSnz0p3T6hHKz9l+lWexpvcqB8y34ZT6casYRSdhTL6/D/hSCxbSAXETXz7xGKGSgcTpFlmPRIdgoS1AbSECe2A2pt5C8EoLZ6ajmrtB4BjWj2OcYFVNgIVDDwwYw8HYWBcEbXAgTWUs8m9M3iIw3XsDjjmZ1xv87FVC46OedfikncO1usIw/E91AaebRtJt2POtGmNEJfgmuOQQCUCuwWqT+WXHULPpHWd7q8efB5gm+6imTbhSsFKr5QTOMvsy95sFB/2brJQUip5+Zw3bEsjDxgzMSd2FzTT4eMTEZTyAkfS2lNngmyFU3fJbJwJ4gxyYF4+9L/6r2syjsid6+FIASg3u9iVuscpwZgdQTuXYlseBBDofFl5nNC/DikbV2lJ1HDcxbqLNj/P7CKlX+SAthJRRzv6U7LwvdFZ3EcgOQzwZCgWYHxY6LqOMdI094m0/65bWaNObSQNU3PyYFJmB+EF2+m6wNFZvyLsW/QDzWShOj4=")
	req.Header.Add("Cookie", "TS013b36ab=0179594e11362fd6dc2dee34f554bf59482c76b7416f6e04741f186533b9ab6b0fd3f53fa31c875bc3b748858d5ab8bd11f3331408f4e32a585f960b58d798178078ac8904; de2a657d1673ca26a0e0abed5da67a83=5649bec0e006d3b80f439de2aa4ba934")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Println(string(body))

	// Parse the response JSON
	var subscriptionData []map[string]interface{}
	json.Unmarshal(body, &subscriptionData)

	if len(subscriptionData) > 0 {
		subscriptionStatus, ok := subscriptionData[0]["status"].(string)
		if !ok {
			return "", fmt.Errorf("subscription status is not a string")
		}

		if subscriptionData[0]["error"] != nil {
			err = fmt.Errorf("%v", subscriptionData[0]["error"])
		}
		return subscriptionStatus, err
	} else {
		return "", fmt.Errorf("subscriptionData is empty")
	}

}

func SumOfTransactionsOfPreviousYear(accountTransactions models.AccountTransactions) float64 {
	var sum float64 = 0
	var tDate time.Time
	var err error

	for _, transaction := range accountTransactions.Transactions {

		if tDate, err = utils.ConvertDateStringToTime(transaction.ValueDate); err != nil {
			log.Fatalf("Error: %v", err)
		}
		//if transaction time after the beginning of this year then sum the transaction amount
		if tDate.After(time.Date(time.Now().Year()-1, 12, 31, 0, 0, 0, 0, time.UTC)) {
			if transaction.DcInd == "CREDIT" {
				sum += transaction.TransactionAmount.Amount
			} else {
				sum -= transaction.TransactionAmount.Amount
			}
		}
	}

	return sum
}

func CurrentBalance(accountBalances []models.Balance) float64 {
	for _, balance := range accountBalances {
		if balance.BalanceType == "CURRENT" {
			return balance.Amount
		}
	}
	return 0
}
