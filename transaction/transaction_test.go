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

package transaction

import "testing"

func TestNewTransaction(t *testing.T) {
	tx := New("Home", "$2.20")

	if tx.Location != "Home" {
		t.Errorf("Incorrect location")
	}
	if tx.Amount != "$2.20" {
		t.Errorf("Incorrect Amount")
	}
	if tx.NumericAmount != 2.2 {
		t.Errorf("Incorrect Numeric Aount")
	}
}

