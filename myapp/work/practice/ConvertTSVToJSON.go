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
    "fmt"
    "log"
    "os"
)

/**
 * Function: main
 * ------------------------------
 * @brief アプリのエントリポイント
 * @details コマンドライン引数を解析し、TSV → JSON 変換を実行する。
 * @return void
 * @retval void 何も返さない（処理のみ実行）
 */
func main() {
    if len(os.Args) < 3 {
        log.Fatal("usage: app <input.tsv> <output.json>")
    }

    var input = os.Args[1]
    var output = os.Args[2]

    if err := ConvertTSVFileToJSONFile(input, output); err != nil {
        log.Fatalf("convert failed: %v", err)
    }
    log.Println("conversion completed:", output)
}
