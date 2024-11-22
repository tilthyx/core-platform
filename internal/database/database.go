package database

import (
	"encoding/json"
	"fmt"
	"github.com/tilthyx/core-platform/internal/config"
	"os"
)

// Database lê o arquivo de configuração e retorna os dados decodificados
func Database(configFile string) (*config.Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}(file)

	// Verificar se o arquivo está vazio
	info, err := file.Stat()
	if err != nil {
		fmt.Printf("Erro ao obter informações do arquivo: %v\n", err)
		return nil, err
	}
	if info.Size() == 0 {
		fmt.Println("O arquivo está vazio.")
		return nil, fmt.Errorf("config file is empty")
	}

	// Decodificar o JSON para a estrutura correta
	var configData config.Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&configData); err != nil {
		fmt.Printf("Error decoding file: %v\n", err)
		return nil, err
	}

	return &configData, nil
}
