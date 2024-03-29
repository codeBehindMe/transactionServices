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

import "testing"

func TestAuthenticationSuccessfulIsClean(t *testing.T){
	a := NewAuthenticator()

	err := a.Authenticate("41f7561f8f654c99826b4c0f07dea2b1")
	if err != nil {
		t.Fail()
	}
}

func TestIncorrectAuthenticationCausesError(t *testing.T){
	a := NewAuthenticator()

	err := a.Authenticate("incorrectkey")
	if err == nil || a.isAuthenticated == true{
		t.Fail()
	}
}