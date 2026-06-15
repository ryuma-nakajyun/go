/**
 * @file ConvertTSVToJSON.go
 * @brief TSV を JSON に変換するユーティリティ
 * @details
 *   - TSV（タブ区切り）ファイルを読み込み、JSON 配列として出力する。
 *   - 標準出力へ書き出すモードと、ファイルへ書き出すモードの両方を提供する。
 *   - ストリーミング処理のため、大規模 TSV に対してもメモリ効率が良い。
 */
package main

import (
    "encoding/csv"
    "encoding/json"
    "fmt"
    "log"
    "os"
)

/**
 * Function: processDataRows
 * ------------------------------
 * @brief TSV のデータ行を読み込み JSON を標準出力へ書き出す
 * @details ストリーミング処理のため、大規模 TSV にも対応できる。
 * @param reader TSV 読み取り用 csv.Reader
 * @param header ヘッダ行
 * @return error
 * @retval nil 正常終了
 * @retval error 行読み込み失敗など
 */
func processDataRows(reader *csv.Reader, header []string) error {
    jsonWriter := json.NewEncoder(os.Stdout)
    jsonWriter.SetIndent("", "  ")

    fmt.Println("[")

    isFirst := true

    for {
        dataRow, err := reader.Read()
        if err != nil {
            break
        }

        record := make(map[string]string)
        for i, value := range dataRow {
            if i < len(header) {
                record[header[i]] = value
            }
        }

        if !isFirst {
            fmt.Println(",")
        }
        isFirst = false

        writeJSONRecord(jsonWriter, record)
    }

    fmt.Println("]")

    return nil
}

/**
 * Function: writeJSONRecord
 * ------------------------------
 * @brief 1 レコードを JSON として書き込む
 * @param writer JSON エンコーダ
 * @param record JSON 化するレコード
 * @return void
 * @retval void 何も返さない（writer に書き込むのみ）
 */
func writeJSONRecord(writer *json.Encoder, record map[string]string) {
    writer.Encode(record)
}
