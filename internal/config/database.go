package config

type Config struct {
	Database struct {
		Host           string `json:"host"`
		Port           int    `json:"port"`
		User           string `json:"user"`
		Password       string `json:"password"`
		Database       string `json:"db"`
		DatasetDestiny string `json:"dataset_destiny"`
		Tables         []struct {
			Table string `json:"table"`
		} `json:"tables"`
	} `json:"database"`
}
