// Copyright 2020 Bjerk AS
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
