package main

func main() {

	// Clear existing transactions
	clearTransactionCollection("transactionsDB")

	// Create and start api server
	server := getServer()
	server.Logger.Fatal(server.Start(":5000"))

}
