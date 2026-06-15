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
 * Function: ConvertTSVFileToJSONFile
 * ------------------------------
 * @brief TSV ファイルを読み込み JSON ファイルとして書き出す
 * @param inputPath 入力 TSV ファイルパス
 * @param outputPath 出力 JSON ファイルパス
 * @return error
 * @retval nil 正常終了
 * @retval error 入力ファイルが開けない／ヘッダ読み込み失敗／書き込み失敗など
 */
func ConvertTSVFileToJSONFile(inputPath, outputPath string) error {
    inputFile, err := os.Open(inputPath)
    if err != nil {
        return err
    }
    defer inputFile.Close()

    reader := createTSVReader(inputFile)

    header, err := readHeaderRow(reader)
    if err != nil {
        return err
    }

    outputFile, err := os.Create(outputPath)
    if err != nil {
        return err
    }
    defer outputFile.Close()

    return processDataRowsToWriter(reader, header, outputFile)
}

/**
 * Function: ConvertTSVToJSON
 * ------------------------------
 * @brief TSV を読み込み JSON を標準出力へ書き出す
 * @param filePath 入力 TSV ファイルパス
 * @return error
 * @retval nil 正常終了
 * @retval error ファイルオープン失敗／ヘッダ読み込み失敗／行読み込み失敗など
 */
func ConvertTSVToJSON(filePath string) error {
    inputFile, err := openTSVFile(filePath)
    if err != nil {
        return err
    }
    defer inputFile.Close()

    reader := createTSVReader(inputFile)

    headerRow, err := readHeaderRow(reader)
    if err != nil {
        return err
    }

    return processDataRows(reader, headerRow)
}
