/*
   transactionServices
   Copyright (C) 2019  aarontillekeratne

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

/*
  Author: aarontillekeratne
  Contact: github.com/codeBehindMe
*/

package authentication

import (
	"cloud.google.com/go/datastore"
	"context"
	"errors"
	"log"
)

const HeaderKey = "Authorization"
const projectID = "test-trapezitam"
const keyType = "key"

type Authenticator struct {
	isAuthenticated bool
}

type AuthKey struct {
	AuthKey string
	KeyType string
}

const authRefernceID = 5637476211228672
func NewAuthenticator() Authenticator {
	return Authenticator{isAuthenticated: false}
}

func datastoreKey(id int64) *datastore.Key {
	return datastore.IDKey(keyType, id, nil)
}

func (a *Authenticator) Authenticate(clientKey string)  error{
	ctx := context.Background()

	dsClient, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Error occured while trying to create datastore client: %v", err)
	}

	k := datastoreKey(authRefernceID)

	authKey := &AuthKey{}

	err = dsClient.Get(ctx, k, authKey)
	if err != nil {
		log.Fatalf("Failed to get authentication clientKey: %v", err)
	}

	if authKey.AuthKey == clientKey {
		a.isAuthenticated = true
		return nil
	}
	a.isAuthenticated = false

	return errors.New("not authorised")
}
