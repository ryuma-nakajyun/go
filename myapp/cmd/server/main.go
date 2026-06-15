package main

import (
    "html/template"
    "log"
    "net/http"

    "myapp/internal/loader"
)

func main() {
    tmpl := template.Must(template.ParseFiles("web/index.html"))

    // 動的ページ（TSV → HTML）
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data, err := loader.LoadCountries("internal/data/r0711world_utf8.tsv")
        if err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        tmpl.Execute(w, data)
    })

    // 静的ファイル
    fs := http.FileServer(http.Dir("web"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    log.Println("HTTPS server started :8443")

    // 証明書と秘密鍵を正しく指定
    err := http.ListenAndServeTLS(
        ":8443",
        "/home/ryuma/dev/go/myapp/cert/192.168.3.250.pem",        // ← 証明書
        "/home/ryuma/dev/go/myapp/cert/192.168.3.250-key.pem",    // ← 秘密鍵
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
}
