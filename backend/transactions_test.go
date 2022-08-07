package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)


var transaction = Transaction{
		
	ID:        "55555b50-174a-414d-a1af-9bf4eee4bbdd",
	Value:     1000,
	Timestamp: 1698723400,
	Receiver:  "a1a80b50-174a-414d-a1af-9bf4eee4ccaa",
	Confirmed: true,
	Sender:    "6c5ddebf-e168-4603-96af-9af8cc114e35",
}

func setup() {
	transactionsCollection, ctx := getTransactionCollection("testDB")

	_, err := transactionsCollection.InsertOne(ctx, transaction)
	if err != nil {
		panic(err)
	}
}

func TestNewTransaction(t *testing.T) {
	out, err := json.Marshal(transaction)
    if err != nil {
        panic (err)
    }

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(out))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Transaction{Database: "testDB"}

	// Assertions
	if assert.NoError(t, h.NewTransaction(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Transaction successfully added!\n", rec.Body.String())
	}
	clearTransactionCollection("testDB")
}

func TestGetTransactions(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Transaction{Database: "testDB"}
	setup()

	out, err := json.Marshal(transaction)
    if err != nil {
        panic (err)
    }

	// Assertions
	if assert.NoError(t, h.GetTransactions(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), string(out))
	}
	clearTransactionCollection("testDB")
}

func TestGetTransactionById(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/transaction/:transactionID")
	c.SetParamNames("transactionID")
	c.SetParamValues("55555b50-174a-414d-a1af-9bf4eee4bbdd")
	h := &Transaction{Database: "testDB"}
	setup()

	out, err := json.Marshal(transaction)
    if err != nil {
        panic (err)
    }

	// Assertions
	if assert.NoError(t, h.GetTransactions(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), string(out))
	}
	clearTransactionCollection("testDB")
}

func TestUpdateTransaction(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &Transaction{Database: "testDB"}
	setup()

	// Assertions
	if assert.NoError(t, h.UpdateTransaction(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Update successful")
	}
	clearTransactionCollection("testDB")
}

func TestDeleteTransaction(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/transaction/:transactionID")
	c.SetParamNames("transactionID")
	c.SetParamValues("55555b50-174a-414d-a1af-9bf4eee4bbdd")
	h := &Transaction{Database: "testDB"}
	setup()

	// Assertions
	if assert.NoError(t, h.DeleteTransaction(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Delete successful")
	}
	clearTransactionCollection("testDB")
}
