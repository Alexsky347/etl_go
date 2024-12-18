package main

import (
	"encoding/json"
	"etl-go/pkg/config"
	"etl-go/pkg/etl"
	"log"
	"time"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	cfg, err := config.LoadConfig()
	if err != nil {
		errorLog.Fatalf("Failed to load configuration: %v", err)
	}

	logger.Init(cfg.Environment)

	db, err := database.New(cfg.Dsn)
	if err != nil {
		errorLog.Fatalf("Cannot connect to database %s", err)
	}
	defer db.Close()

	go func() {
		// Start time
		start := time.Now()

		headers, data, err := extract.Extract(cfg.ConfigPath, cfg.CSVPath)
		if err != nil {
			errorLog.Fatalf("Error extracting data: %s", err)
			return
		}
		// Convert the mapped records to JSON
		mappedRecordsJSON, err := json.Marshal(transform.MapDataToConfig(data, headers))
		if err != nil {
			errorLog.Fatalf("Error converting mapped records to JSON: %s", err)
			return
		}

		// Print the JSON array of objects
		logInfo.Printf(string(mappedRecordsJSON))
		// Log duration
		logInfo.Printf("Duration: %v\n", time.Now().Sub(start))
	}()

	select {}
}
