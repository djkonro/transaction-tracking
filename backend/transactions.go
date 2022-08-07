package main

import (
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/labstack/echo"
)

type Transaction struct {
	ID        string `bson:"id,omitempty"`
	Value     int    `bson:"value,omitempty"`
	Timestamp int    `bson:"timestamp,omitempty"`
	Receiver  string `bson:"receiver,omitempty"`
	Confirmed bool   `bson:"confirmed,omitempty"`
	Sender    string `bson:"sender,omitempty"`
	Database string
}

type Response struct {
	Message      string        `json:"message,omitempty"`
	Error        error         `json:"error,omitempty"`
	Transactions []Transaction `json:"transactions,omitempty"`
}

type DateRange struct {
	fromDateTime string
	toDateTime   string
}

func (t *Transaction) GetTransactions(c echo.Context) error {
	var transactions []Transaction
	transactionsCollection, ctx := getTransactionCollection(t.Database)

	cursor, err := transactionsCollection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &transactions); err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, Response{Transactions: transactions})
}

func (t *Transaction) NewTransaction(c echo.Context) error {
	transactionsCollection, ctx := getTransactionCollection(t.Database)

	transaction := &Transaction{}
	if err := c.Bind(transaction); err != nil {
		return c.JSON(http.StatusBadRequest, Response{Error: err})
	}
	_, err := transactionsCollection.InsertOne(ctx, transaction)
	if err != nil {
		panic(err)
	}

	return c.String(http.StatusOK, "Transaction successfully added!\n")
}

func (t *Transaction) UpdateTransaction(c echo.Context) error {
	transactionsCollection, ctx := getTransactionCollection(t.Database)
	transactionId := c.Param("transactionID")

	filter := bson.D{{"id", transactionId}}
	update := bson.D{
		{"$set", bson.D{{"confirmed", true}}},
	}

	_, err := transactionsCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, Response{Message: "Update successful"})
}

func (t Transaction) GetTransactionById(c echo.Context) error {
	var transaction Transaction
	transactionsCollection, ctx := getTransactionCollection(t.Database)
	transactionId := c.Param("transactionID")

	err := transactionsCollection.FindOne(ctx, bson.D{{"id", transactionId}}).Decode(&transaction)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, transaction)
}

func (t *Transaction) GetTransactionByDate(c echo.Context) error {
	var transaction Transaction
	transactionsCollection, ctx := getTransactionCollection(t.Database)

	var dateRange DateRange
	if err := c.Bind(dateRange); err != nil {
		return c.JSON(http.StatusBadRequest, Response{Error: err})
	}

	layout := "01/30/2006 3:04:55 PM"
	fromTime, _ := time.Parse(layout, dateRange.fromDateTime)
	fromTimestamp := fromTime.Unix()
	toTime, _ := time.Parse(layout, dateRange.toDateTime)
	toTimestamp := toTime.Unix()

	err := transactionsCollection.FindOne(ctx, bson.M{
		"timestamp": bson.M{
			"$gt": fromTimestamp,
			"$lt": toTimestamp,
		},
	}).Decode(&transaction)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, transaction)
}

func (t *Transaction) DeleteTransaction(c echo.Context) error {
	transactionsCollection, ctx := getTransactionCollection(t.Database)
	transactionId := c.Param("transactionID")

	_, err := transactionsCollection.DeleteOne(ctx, bson.D{{"id", transactionId}})
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, Response{Message: "Delete successful"})
}
