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

package extraction

import (
	lang "cloud.google.com/go/language/apiv1"
	"context"
	"encoding/json"
	langpb "google.golang.org/genproto/googleapis/cloud/language/v1"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type TransactionText struct {
	TransactionText string
}

type Transaction struct {
	Location      string
	Amount        string
	NumericAmount float32
	NotifiedTime  time.Time
	UnixEpoch     int64
}

func NewTransaction(location, dollarAmount string) Transaction {
	// FIXME: Design needs to be revised.
	amount, _ := strconv.ParseFloat(strings.Trim(dollarAmount, "$"), 32)
	amount32 := float32(amount)

	tx := Transaction{
		Location:      location,
		Amount:        dollarAmount,
		NumericAmount: amount32,
		NotifiedTime:  time.Now(),
		UnixEpoch:     0,
	}
	tx.NotifiedTime = tx.NotifiedTime.Round(0)
	tx.UnixEpoch = tx.NotifiedTime.Unix()

	return tx
}

var ctx = context.Background()

// Gets the transaction text from the incoming request.
func GetTransactionTextFromRequest(r *http.Request) string {
	var txt TransactionText

	err := json.NewDecoder(r.Body).Decode(&txt)

	if err != nil {
		log.Fatalf(
			"Failed to extract transaction text from request: %v", err)
	}

	return txt.TransactionText
}

func AnalyseEntitiesInText(text *string) (*langpb.AnalyzeEntitiesResponse, error) {

	nlpClient, err := lang.NewClient(ctx)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return nlpClient.AnalyzeEntities(ctx,
		&langpb.AnalyzeEntitiesRequest{
			Document: &langpb.Document{
				Type: langpb.Document_PLAIN_TEXT,
				Source: &langpb.Document_Content{
					Content: *text,
				},
			},
			EncodingType: langpb.EncodingType_UTF8,
		})

}

func CreateTransactionFromAnalyseEntitiesResponse(
	aeResponse *langpb.AnalyzeEntitiesResponse) Transaction {

	var transaction Transaction

	for _, e := range aeResponse.GetEntities() {
		switch e.Type {
		case langpb.Entity_PRICE:
			transaction.Amount = e.Name
		case langpb.Entity_LOCATION:
			transaction.Location = e.Name
		case langpb.Entity_OTHER:
			transaction.Location = e.Name
		case langpb.Entity_ORGANIZATION:
			transaction.Location = e.Name
		case langpb.Entity_NUMBER:
			f, err := strconv.ParseFloat(e.Name, 32)
			if err != nil {
				panic(err)
			}
			transaction.NumericAmount = float32(f)
		}
	}

	transaction.NotifiedTime = time.Now()
	transaction.UnixEpoch = transaction.NotifiedTime.Unix()

	return transaction
}

func GetTransactionFromFromHttpRequest(r *http.Request) Transaction {
	var tx Transaction

	err := json.NewDecoder(r.Body).Decode(&tx)
	if err != nil {
		log.Fatalf("Failed to extract transaction from request: %v", err)
	}

	return tx
}
