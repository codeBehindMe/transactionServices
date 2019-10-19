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

/*
Note: You will need to set GOOGLE_APPLICATION_CREDENTIALS path if you're
running this locally. If you're using cloud build, it should have necessary
IAM roles to use Cloud NLP.
*/

package extraction

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCloudNLP(t *testing.T) {
	testText := "You spend $10 at home"
	AEResponse, err := AnalyseEntitiesInText(&testText)

	if err != nil {
		panic(err)
	}
	for _, e := range AEResponse.GetEntities() {
		fmt.Println(e.Name)
	}
}

func TestGetTransactionTextFromRequest(t *testing.T) {
	jsonPayload := []byte(`{"TransactionText":"You spent $13.40 at Maccas"}`)
	expectedResponse := "You spent $13.40 at Maccas"

	req, err := http.NewRequest("POST", "/tdecode", bytes.NewBuffer(jsonPayload))
	if err != nil {
		t.Fatalf("Failed test: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	tDecodeHandler := func(w http.ResponseWriter, r *http.Request) {
		textOutput := GetTransactionTextFromRequest(r)
		_, _ = io.WriteString(w, textOutput)
	}

	reqResponse := httptest.NewRecorder()

	handler := http.HandlerFunc(tDecodeHandler)

	handler.ServeHTTP(reqResponse, req)

	if reqResponse.Body.String() != expectedResponse {
		t.Errorf("Returned unexpected body: got %v \n want: %v", reqResponse.Body.String(), expectedResponse)
	}
}

func TestCreateTransactionFromAnalyseEntitiesResponse(t *testing.T) {
	testText := "You spent $10 at home"
	AEResponse, err := AnalyseEntitiesInText(&testText)

	targetTx := Transaction{
		Location:      "home",
		Amount:        "$10",
		NumericAmount: 10,
		NotifiedTime:  time.Now(),
	}

	if err != nil {
		panic(err)
	}
	outputTx := CreateTransactionFromAnalyseEntitiesResponse(AEResponse)

	fmt.Println(outputTx.Location)

	if outputTx.Location != targetTx.Location {
		t.Errorf("Incorrect location in transaction. Expected \n %v found \n %v",
			targetTx.Location, outputTx.Location)
	}

	if outputTx.Amount != targetTx.Amount {
		t.Errorf("Incorrect amount in transaction. Expected \n %v found \n %v",
			targetTx.Amount, outputTx.Amount)
	}

	if outputTx.NumericAmount != targetTx.NumericAmount {
		t.Errorf("Incorrect NumericAmount in transaction. Expected \n %v found \n %v",
			targetTx.NumericAmount, outputTx.NumericAmount)
	}
}

func TestGetTransactionFromFromHttpRequest(t *testing.T) {
	ntfTime := time.Now()
	tx := Transaction{
		Location:      "home",
		Amount:        "$10.00",
		NumericAmount: 10,
		NotifiedTime:  ntfTime,
		UnixEpoch:     ntfTime.Unix(),
	}

	// When you Jsonise only the wall clock time is kept, so need to get rid of it in the original transaction so the
	// returned result once it comes back through http request is the same.
	// Comment this line out to see what happens.
	// FIXME: Is there a cleaner way to do this?
	tx.NotifiedTime = tx.NotifiedTime.Round(0)
	jsonPayload, err := json.Marshal(&tx)

	if err != nil {
		t.Fatalf("Failed test with error: %v", err)
	}

	req, err := http.NewRequest("POST", "/txdecode", bytes.NewBuffer(jsonPayload))

	if err != nil {
		t.Fatalf("Failed test error: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	tDecodeHandler := func(w http.ResponseWriter, r *http.Request) {
		txOutput := GetTransactionFromFromHttpRequest(r)

		if txOutput != tx {
			t.Errorf("Expected: %v \n :got: %v", tx, txOutput)
		}
	}

	handler := http.HandlerFunc(tDecodeHandler)
	handler.ServeHTTP(nil, req)
}
