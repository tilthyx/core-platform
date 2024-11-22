package main

import (
	"fmt"
	"github.com/tilthyx/core-platform/internal/database"
)

func main() {
	// Reading database json config
	databaseConfig, err := database.Database("./configs/database-sample.json")
	if err != nil {
		panic(err)
	}
	// Exibir os dados do banco de dados
	fmt.Printf("Configuração do Banco de Dados:\n")
	fmt.Printf("Host: %s\nPort: %d\nUser: %s\nDatabase: %s\n",
		databaseConfig.Database.Host, databaseConfig.Database.Port, databaseConfig.Database.User, databaseConfig.Database.Database)

	// Exibir as tabelas
	fmt.Println("Tabelas:")
	for _, table := range databaseConfig.Database.Tables {
		fmt.Printf("- %s\n", table.Table)
	}
}
