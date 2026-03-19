package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `<!DOCTYPE html>
<html>
<head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><title>Hello</title>
<style>*{margin:0;padding:0}body{background:#1a1a2e;color:#e0e0e0;font-family:monospace;display:flex;align-items:center;justify-content:center;min-height:100vh;text-align:center}.msg{font-size:2rem}</style>
</head>
<body><div class="msg">Hello from Shipwright!</div></body>
</html>`)
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	fmt.Println("listening on :3000")
	http.ListenAndServe(":3000", nil)
}
