package loader

import (
    "encoding/csv"
    "os"
)

type Country struct {
    Code      string
    NameJP    string
    NameJPS   string
    CapitalJP string
    NameEN    string
    NameENS   string
    CapitalEN string
    Lat       string
    Lon       string
}

func LoadCountries(path string) ([]Country, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    r := csv.NewReader(f)
    r.Comma = '\t'
    r.FieldsPerRecord = -1

    rows, err := r.ReadAll()
    if err != nil {
        return nil, err
    }

    var list []Country
    for i, row := range rows {
        if i == 0 {
            continue
        }
        list = append(list, Country{
            Code:      row[0],
            NameJP:    row[1],
            NameJPS:   row[2],
            CapitalJP: row[3],
            NameEN:    row[4],
            NameENS:   row[5],
            CapitalEN: row[6],
            Lat:       row[7],
            Lon:       row[8],
        })
    }
    return list, nil
}
