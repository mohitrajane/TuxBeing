package main

import (
	"github.com/zserge/webview"
)

func ui(d chan int) {

	webView := webview.New(webview.Settings{
		Title:                  "My test web view app",
		URL:                    "http://192.168.199.201",
		Width:                  1000,
		Height:                 800,
		Resizable:              true,
		Debug:                  true,
		ExternalInvokeCallback: nil,
	})

	webView.SetFullscreen(true)
	webView.Run()
	d <- 1
}
