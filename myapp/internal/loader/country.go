package loader

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Country struct {
	Code       string
	NameJP     string
	NameJPS    string
	CapitalJP  string
	NameEN     string
	NameENS    string
	CapitalEN  string
	Lat        string
	Lon        string
	IsNorthern bool
}

func LoadCountries(path string) ([]Country, error) {
	// ファイルを開く
	var file *os.File
	var err error

	file, err = os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// CSV リーダーの準備
	var reader *csv.Reader
	reader = csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1

	// 全行を読み込む
	var rows [][]string
	rows, err = reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// 結果を入れるスライス
	var countries []Country

	// 行を1つずつ処理
	for index, row := range rows {
		if index == 0 {
			continue // ヘッダー行をスキップ
		}

		// 緯度をfloatに変換
		var latitude float64
		latitude, err = strconv.ParseFloat(row[7], 64)
		if err != nil {
			// 緯度が不正フォーマットの場合はデフォルト値 0 を使用
			latitude = 0
		}

		// Country構造体を追加
		var country Country
		country = Country{
			Code:       row[0],
			NameJP:     row[1],
			NameJPS:    row[2],
			CapitalJP:  row[3],
			NameEN:     row[4],
			NameENS:    row[5],
			CapitalEN:  row[6],
			Lat:        row[7],
			Lon:        row[8],
			IsNorthern: latitude > 0,
		}

		countries = append(countries, country)
	}

	return countries, nil
}

func (c Country) Hemisphere() string {
	if c.IsNorthern {
		return "北"
	}
	return "南"
}
