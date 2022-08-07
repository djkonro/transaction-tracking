package main


import (
	"sync"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
)


var (
	server     *echo.Echo
	serverOnce sync.Once
)


func getServer() *echo.Echo {
    serverOnce.Do(func() {
		// Initializing api server instance
		server = echo.New()

		server.Use(middleware.Logger())
		server.Use(middleware.Recover())
	
		server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		}))

		setupRoute(server)
	})

	return server
}


func setupRoute(server *echo.Echo) {
    t := Transaction{}
	t.Database = "transactionsDB"

	server.GET("/api/transactions", t.GetTransactions)
	server.POST("/api/transaction", t.NewTransaction)
	server.PUT("/api/transaction/:transactionID", t.UpdateTransaction)
	server.GET("/api/transaction/:transactionID", t.GetTransactionById)
	server.DELETE("/api/transaction/:transactionID", t.DeleteTransaction)
}
