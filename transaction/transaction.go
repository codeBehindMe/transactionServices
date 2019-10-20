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

import (
	"strconv"
	"strings"
	"time"
)

type Transaction struct {
	Location      string
	Amount        string
	NumericAmount float32
	NotifiedTime  time.Time
	UnixEpoch     int64
}

func New(location, dollarAmount string) Transaction {
	// FIXME: Design needs to be revised.
	amount, _ := strconv.ParseFloat(strings.Trim(dollarAmount, "$"), 32)
	amount32 := float32(amount)

	notifiedTime := time.Now().Round(0)

	tx := Transaction{
		Location:      location,
		Amount:        dollarAmount,
		NumericAmount: amount32,
		NotifiedTime:  notifiedTime,
		UnixEpoch:     notifiedTime.Unix(),
	}

	return tx
}

