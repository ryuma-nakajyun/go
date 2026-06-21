package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"myapp/internal/loader"
	"myapp/internal/logging"
)

func main() {
	// ロガー初期化
	if err := logging.InitLogger(); err != nil {
		log.Fatal(err)
	}
	defer logging.CloseLogger()

	// ログ出力テスト
	logging.Info("アプリケーション起動")
	logging.Debug("デバッグ値: x=%d", 123)
	logging.Warn("警告: %s", "設定値が不正です")
	logging.Error("エラー発生: %v", fmt.Errorf("sample error"))
	fmt.Println("ログ出力完了。log/ ディレクトリを確認してください。")

	// ------------------------------
	// 1. HTMLテンプレートを読み込む
	// ------------------------------
	var tmpl *template.Template
	var err error

	tmpl, err = template.ParseFiles("web/index.html")
	if err != nil {
		log.Fatal("テンプレート読み込みエラー:", err)
	}

	// ------------------------------
	// 2. 動的ページ（TSV → HTML）
	// ------------------------------
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TSV を読み込む
		var data []loader.Country
		data, err = loader.LoadCountries("internal/data/r0711world_utf8.tsv")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// HTMLに埋め込んで出力
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	})

	// 静的ファイル（CSS/JS/画像）
	var fs http.Handler
	fs = http.FileServer(http.Dir("web"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// HTTPS サーバーを起動
	log.Println("HTTPS server started :8443")
	err = http.ListenAndServeTLS(
		":8443",
		"/home/ryuma/dev/go/myapp/cert/192.168.3.250.pem",     // 証明書
		"/home/ryuma/dev/go/myapp/cert/192.168.3.250-key.pem", // 秘密鍵
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
}
