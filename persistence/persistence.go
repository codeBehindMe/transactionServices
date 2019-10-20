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

package persistence

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"log"
	"transactionServices/transaction"
)

const kind = "transactionv2"

func SaveToDatabase(t *transaction.Transaction, projectId string) {
	ctx := context.Background()

	dsClient, err := datastore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Error occurred when trying to create data store client: %v", err)
	}

	taskKey := datastore.NameKey(kind, "", nil)

	_, err = dsClient.Put(ctx, taskKey, t)
	if err != nil {
		log.Fatalf("Failed to save transaction: %v", err)
	}
	// FIXME: Printing instead of logging
	fmt.Printf("Saved %v",taskKey)
}

func AddTransactionToBudget(t *transaction.Transaction) {
	log.Fatalf("Not Implemented.")
}
