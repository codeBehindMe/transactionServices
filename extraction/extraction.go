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
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	"transactionServices/transaction"
)

type transactionText struct {
	TransactionText string
}

var _ctx = context.Background()

// Gets the transaction text from the incoming request.
func GetTransactionTextFromRequest(r *http.Request) string {
	var txt transactionText

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Could not read bytes from request body: %v", err)
	}
	requestBody := string(bodyBytes)
	log.Printf("Recieved body: %v", requestBody)
	err = json.Unmarshal(bodyBytes,&txt)

	if err != nil {
		log.Fatalf(
			"Failed to extract transaction text from request: %v", err)
	}

	return txt.TransactionText
}

func AnalyseEntitiesInText(text *string) (*langpb.AnalyzeEntitiesResponse, error) {

	nlpClient, err := lang.NewClient(_ctx)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return nlpClient.AnalyzeEntities(_ctx,
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
	aeResponse *langpb.AnalyzeEntitiesResponse) transaction.Transaction {

	var tx transaction.Transaction

	for _, e := range aeResponse.GetEntities() {
		switch e.Type {
		case langpb.Entity_PRICE:
			tx.Amount = e.Name
		case langpb.Entity_LOCATION:
			tx.Location = e.Name
		case langpb.Entity_OTHER:
			tx.Location = e.Name
		case langpb.Entity_ORGANIZATION:
			tx.Location = e.Name
		case langpb.Entity_NUMBER:
			f, err := strconv.ParseFloat(e.Name, 32)
			if err != nil {
				panic(err)
			}
			tx.NumericAmount = float32(f)
		}
	}

	tx.TxNotifyUnixEpoch = time.Now().Unix()
	tx.TransactionVersion = transaction.Version

	txFromCon := transaction.New(tx.Location,tx.Amount)

	return txFromCon
}

func GetTransactionFromFromHttpRequest(r *http.Request) transaction.Transaction {
	var tx transaction.Transaction

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Fatalf("Error decoding body %v", err)
	}

	requestBody := string(bodyBytes)
	log.Printf("Recieved body: %v", requestBody)

	err = json.Unmarshal(bodyBytes,&tx)
	if err != nil {
		log.Fatalf("Failed to extract transaction from request: %v", err)
	}

	return tx
}
