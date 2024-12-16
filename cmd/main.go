package main

import (
	extract "etl-go/internal/services"
	transform "etl-go/internal/services"
    "encoding/json"
	"fmt"
	"time"
)

const (
    csv_path   = "gen_data/data/small-data.csv"
    config_path = "internal/conf/config.json"
)



func main() {
    // Start time
    start := time.Now()
    headers, data, err := extract.Extract(config_path, csv_path)
    if err != nil {
        fmt.Println()
    }
    // Convert the mapped records to JSON
    mappedRecordsJSON, err := json.Marshal(transform.MapDataToConfig(data, headers))
    if err != nil {
        fmt.Println("Error converting mapped records to JSON:", err)
        return
    }

    // Print the JSON array of objects
    fmt.Println(string(mappedRecordsJSON))

    // End time
    end := time.Now()

    // Calculate duration
    duration := end.Sub(start)
    fmt.Printf("Duration: %v\n", duration)
}