package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/iotaledger/hive.go/timedexecutor"
	uuid "github.com/satori/go.uuid"
)


type Transaction struct {
	ID        string `json:"id,omitempty"`
	Value     int    `json:"value,omitempty"`
	Timestamp int    `json:"timestamp,omitempty"`
	Receiver  string `json:"receiver,omitempty"`
	Confirmed bool   `json:"confirmed,omitempty"`
	Sender    string `json:"sender,omitempty"`
}

func generateTransactionData() Transaction {
	rand.Seed(time.Now().UnixNano())
	transaction := Transaction{
		ID:        uuid.NewV4().String(),
		Value:     rand.Intn(500),
		Timestamp: int(time.Now().Unix()),
		Receiver:  uuid.NewV4().String(),
		Confirmed: false,
		Sender:    uuid.NewV4().String(),
	}

	return transaction
}


func addNewTransations() {
	var timedExecutor  *timedexecutor.TimedExecutor = timedexecutor.New(1)
    t := time.Now()
	for {
		timedExecutor.ExecuteAt(func() {
			transaction := generateTransactionData()
			jsonData, err := json.Marshal(transaction)
			if err != nil {
				panic(err)
			}

			req, err := http.NewRequest("POST", "http://webapi:5000/api/transaction", bytes.NewBuffer(jsonData))
			if err != nil {
				log.Fatalf("Error Occurred. %+v", err)
			}

			req.Header.Set("Content-Type", "application/json; charset=UTF-8")
			client := &http.Client{}
			response, err := client.Do(req)
			if err != nil {
				log.Fatalf("Error sending request to API endpoint. %+v", err)
			} else {
				// Confirm transation
				confirmTransation(transaction.ID)
			}

            fmt.Println(response)

		}, t)
        t = t.Add(time.Minute * 1)
		time.Sleep(time.Second * 50)
    }
}

func confirmTransation(id string) {
	var timedExecutor  *timedexecutor.TimedExecutor = timedexecutor.New(1)

	timedExecutor.ExecuteAt(func() {
		uri := "http://webapi:5000/api/transaction/"+id
		req, err := http.NewRequest("PUT", uri, nil)
		if err != nil {
			log.Fatalf("Error Occurred. %+v", err)
		}

		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		client := &http.Client{}
		response, err := client.Do(req)
		if err != nil {
			log.Fatalf("Error sending request to API endpoint. %+v", err)
		}
		fmt.Println(response)
	}, time.Now().Add(time.Second * 10))

}

func main() {
    addNewTransations()
}
