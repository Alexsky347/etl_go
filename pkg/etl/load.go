package etl

import (
    "database/sql"
    "encoding/csv"
    "fmt"
    "log"
    "os"
    "strings"

    _ "github.com/lib/pq"
)


func LoadToDB() {
    // Connect to the database
    db, err := sql.Open("postgres", PostgresConnStr)
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }
    defer db.Close()

    // Truncate the table
    _, err = db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", TableName))
    if err != nil {
        log.Fatal("Error truncating table:", err)
    }
    fmt.Println("Table truncated successfully")

    // Open the CSV file
    file, err := os.Open(CSVFilePath)
    if err != nil {
        log.Fatal("Error opening CSV file:", err)
    }
    defer file.Close()

    // Create a new CSV reader
    r := csv.NewReader(file)

    // Read all records
    records, err := r.ReadAll()
    if err != nil {
        log.Fatal("Error reading CSV file:", err)
    }

    // Prepare the bulk insert statement
    var sb strings.Builder
    sb.WriteString(fmt.Sprintf("INSERT INTO %s VALUES ", TableName))
    for i, record := range records {
        sb.WriteString("(")
        for j, field := range record {
            sb.WriteString(fmt.Sprintf("'%s'", field))
            if j < len(record)-1 {
                sb.WriteString(", ")
            }
        }
        sb.WriteString(")")
        if i < len(records)-1 {
            sb.WriteString(", ")
        }
    }
    insertQuery := sb.String()

    // Execute the bulk insert
    _, err = db.Exec(insertQuery)
    if err != nil {
        log.Fatal("Error executing bulk insert:", err)
    }
    fmt.Println("Bulk insert completed successfully")
}