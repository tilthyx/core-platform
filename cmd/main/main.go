package main

import (
	"fmt"
	"github.com/tilthyx/core-platform/internal/database"
)

func main() {
	// Reading database json config
	databaseConn, err := database.RetrieveDatabaseConnections("./configs/database-sample.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(databaseConn)
}
