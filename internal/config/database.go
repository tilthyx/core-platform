package config

type DatabaseConfig struct {
	Connection Connection `json:"connection"`
	Dataset    Dataset    `json:"dataset"`
	Tables     []Table    `json:"tables"`
}

type Connection struct {
	DatabaseHost     string `json:"database_host"`
	DatabasePort     int64  `json:"database_port"`
	DatabaseUser     string `json:"database_user"`
	DatabasePassword string `json:"database_password"`
	DatabaseSchema   string `json:"database_schema"`
}

type Dataset struct {
	OriginType         string `json:"origin_type"`
	DatasetName        string `json:"dataset_name"`
	DatasetModule      string `json:"dataset_module"`
	DatasetDescription string `json:"dataset_description"`
}

type Table struct {
	Table string `json:"table"`
}
