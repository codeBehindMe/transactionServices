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
	"log"
	"strconv"
	"strings"
	"time"
)

const Version = "transactionv3"

type Transaction struct {
	TransactionVersion string
	Location           string
	Amount             string
	NumericAmount      float32
	TxNotifyUnixEpoch  int64
}

func New(location, dollarAmount string) Transaction {
	// FIXME: Design needs to be revised.
	amount, err := strconv.ParseFloat(strings.Trim(dollarAmount, "$"), 32)

	if err != nil {
		log.Fatalf("Error when parsing amount: %v", err)
	}
	amount32 := float32(amount)

	notifiedTime := time.Now().Round(0)

	tx := Transaction{
		TransactionVersion: Version,
		Location:           location,
		Amount:             dollarAmount,
		NumericAmount:      amount32,
		TxNotifyUnixEpoch:  notifiedTime.Unix(),
	}

	return tx
}
