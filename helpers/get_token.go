package helpers

import (
	"time"

	apiclient "github.com/bjerkio/tripletex-go/client"
	"github.com/bjerkio/tripletex-go/client/session"
)

var token string

// CreateToken retrieves a new tripletex token for 24 hours
func CreateToken(consumerToken string, employeeToken string) (string, error) {

	// if there is a token already made, use it.
	if token != "" {
		return token, nil
	}

	expirationDate := 24 * time.Hour
	client := apiclient.Default
	sessionReq := &session.TokenSessionCreateCreateParams{
		ConsumerToken:  consumerToken,
		EmployeeToken:  employeeToken,
		ExpirationDate: time.Now().Add(expirationDate).Format("2006-01-02"),
	}

	res, err := client.Session.TokenSessionCreateCreate(sessionReq.WithTimeout(10 * time.Second))

	if err != nil {
		return "", err
	}

	token = res.Payload.Value.Token

	return token, nil
}
