package services

const (
    // Kafka Configuration
    KafkaServer  = "localhost:9092"
    KafkaTopic   = "test-topic"
    KafkaGroupId = "product-service"
    
    // File paths
    CSVFilePath     = "data/data.csv"
    ConfigPath      = "config.json"
    
    // Database Configuration
    PostgresConnStr = "user=youruser dbname=yourdb sslmode=disable password=yourpassword"
    TableName       = "your_table"
    
    // Other Configuration
    OutputModeBatch = false
	separator  = ';'
)