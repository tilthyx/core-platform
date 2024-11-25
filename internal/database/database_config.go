package database

import (
	"encoding/json"
	"fmt"
	"github.com/tilthyx/core-platform/internal/config"
	"os"
)

func RetrieveDatabaseConnections(databaseConfigJson string) (*config.DatabaseConfig, error) {
	file, err := os.Open(databaseConfigJson)
	if err != nil {
		return nil, fmt.Errorf("failed to open database config file: %w", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	var databaseConfig config.DatabaseConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&databaseConfig); err != nil {
		return nil, fmt.Errorf("failed to decode database config: %w", err)
	}

	return &databaseConfig, nil
}
