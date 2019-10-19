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

package transactionServices

import (
	"encoding/json"
	"log"
	"net/http"
	"transactionServices/extraction"
)

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	transactionText := extraction.GetTransactionTextFromRequest(r)

	analyseEntitiesResponse, err := extraction.AnalyseEntitiesInText(&transactionText)

	if err != nil {
		log.Fatalf("Failed to analyse entities: %v", err)
	}

	transaction := extraction.CreateTransactionFromAnalyseEntitiesResponse(analyseEntitiesResponse)

	_ = json.NewEncoder(w).Encode(transaction)
}

func SaveTransaction(w http.ResponseWriter, r *http.Request) {
	tx := extraction.GetTransactionFromFromHttpRequest(r)
}
