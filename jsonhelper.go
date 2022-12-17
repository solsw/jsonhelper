// Package jsonhelper contains 'json' package helpers.
package jsonhelper

import (
	"encoding/json"
	"io"
	"strings"
)

// IndentRW reads JSON-encoded data from 'r' and writes indented data to 'w'.
// (See json.Indent for 'prefix' and 'indent' usage.)
func IndentRW(r io.Reader, w io.Writer, prefix, indent string) error {
	var v any
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		return err
	}
	e := json.NewEncoder(w)
	e.SetIndent(prefix, indent)
	return e.Encode(v)
}

// IndentStr returns indented form of the JSON-encoded 'json'.
// (See json.Indent for 'prefix' and 'indent' usage.)
func IndentStr(json, prefix, indent string) (string, error) {
	var b strings.Builder
	if err := IndentRW(strings.NewReader(json), &b, prefix, indent); err != nil {
		return "", err
	}
	return b.String(), nil
}
