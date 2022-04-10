// Copyright (c) 2022, Daniel M. Jaén
// All rights reserved.

package utils

import(
	"path/filepath"
	"runtime"
	"fmt"
)

// Generates a small HTML template, which adds the scripts for WebAssembly.
func GenHTML(opt *HTMLOptions) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="%s">
<head>
<meta charset="%s">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>%s</title>
<style>
%s
</style>
<script src="%s"></script>
<script src="%s"></script>
</head>
<body>
</body>
</html>`, opt.Lang, opt.Charset, opt.Title, opt.CSS, func(path string) string {
		if path == "" {
			return filepath.Join(runtime.GOROOT(), "misc", "wasm", "wasm_exec.js")
		}
		return path
	}(opt.WASMP), func(path string) string {
		if path == "" {
			// ovni go instance initializer
			return "main.wasm"
		}
		return path
	}(opt.GOII))
}

// Pass the options to generate the HTML template.
type HTMLOptions struct {
	Lang 		string
	Charset string
	Title 	string
	CSS			string
	WASMP 	string // wasm_exec.js path
	GOII 		string // path to Go Instance Initializer
}