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
	"errors"
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

const authRefernceID = "5194620877fd2a6cb2ea948682b171f5"

func NewAuthenticator() Authenticator {
	return Authenticator{isAuthenticated: false}
}

func datastoreKey(id int64) *datastore.Key {
	return datastore.IDKey(keyType, id, nil)
}

func (a *Authenticator) Authenticate(clientKey string) error {

	if clientKey == authRefernceID {
		a.isAuthenticated = true
		return nil
	}
	a.isAuthenticated = false

	return errors.New("not authorised")
}
