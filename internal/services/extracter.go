package services

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
    Client struct {
        ClientID struct {
            Type   string `json:"type"`
            MaxLen int    `json:"maxLen"`
        } `json:"client_id"`
        Num1 struct {
            Type   string `json:"type"`
            MaxLen int    `json:"maxLen"`
        } `json:"num1"`
        Num2 struct {
            Type   string `json:"type"`
            MaxLen int    `json:"maxLen"`
        } `json:"num2"`
        Num3 struct {
            Type   string `json:"type"`
            MaxLen int    `json:"maxLen"`
        } `json:"num3"`
        Filiale struct {
            Type   string `json:"type"`
            MaxLen int    `json:"maxLen"`
        } `json:"filiale_id"`
    } `json:"client"`
}

func getConfigFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func getDataFile(path string) (*os.File, error) {
	data, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func Extract(configPath string, csvPath string) ([]string, [][]string, error) {
    configData, err := getConfigFile(configPath)
    if err != nil {
        fmt.Println("Error reading config file:", err)
        return nil, nil, err
    }

    var config Config
    err = json.Unmarshal(configData, &config)
    if err != nil {
        return nil, nil, err
    }

    // Open the CSV file
    file, err := getDataFile(csvPath)
    if err != nil {
        return nil, nil, err
    }
    defer file.Close()
    

    // Create a new CSV reader
    r := csv.NewReader(file)
    r.Comma = separator

    // Read all records
    records, err := r.ReadAll()
    if err != nil {
        return nil, nil, err
    }

    // Extract headers and data
    headers := records[0]
    data := records[1:]

	return headers, data, nil
}