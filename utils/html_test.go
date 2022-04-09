// Copyright (c) 2022, Daniel M. Jaén
// All rights reserved.

package utils

import(
	"testing"
	"strings"
)

func Test_GenHTML(t *testing.T) {
	got := strings.Split(GenHTML(&HTMLOptions{
		Lang: "es",
		Charset: "UTF-8",
		Title: "Prueba",
		CSS: "body { background-color: #303030; }",
		WASMP: "",
		GOII: "ovni_goii.js",
	}), "\n")[1]

	expected := "<html lang=\"es\">"

	if got != expected {
		t.Errorf("\nExpected: %s\nGot: %s\n", expected, got)
	}
}