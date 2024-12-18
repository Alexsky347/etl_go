package etl

import (
    "strconv"
)

type Record struct {
    ClientID string `json:"client_id"`
    Num1     string `json:"num1"`
    Num2     string `json:"num2"`
    Num3     string `json:"num3"`
    Filiale  int    `json:"filiale_id"`
}

func MapDataToConfig(data [][]string, headers []string) []Record, error {
    var mappedRecords []Record
    for _, record := range data {
        var rec Record
        for i, field := range record {
            header := headers[i]
            switch header {
            case "client_id":
                rec.ClientID = field
            case "num1":
                rec.Num1 = field
            case "num2":
                rec.Num2 = field
            case "num3":
                rec.Num3 = field
            case "filiale_id":
                filiale, err := strconv.Atoi(field)
                if err != nil {
                    return nil, err
                }
                rec.Filiale = filiale
            }
        }
        mappedRecords = append(mappedRecords, rec)
    }
    return mappedRecords, nil
}