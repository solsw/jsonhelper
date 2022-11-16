// Package jsonhelper contains 'json' package helpers.
package jsonhelper

import (
	"encoding/json"
	"io"
	"strings"
)

// MarshalMust is like json.Marshal but panics in case of error.
func MarshalMust(v any) []byte {
	bb, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return bb
}

// MarshalIndentMust is like json.MarshalIndent but panics in case of error.
// (See json.MarshalIndent for 'prefix' and 'indent' usage.)
func MarshalIndentMust(v any, prefix, indent string) []byte {
	bb, err := json.MarshalIndent(v, prefix, indent)
	if err != nil {
		panic(err)
	}
	return bb
}

// MarshalIndentStr formats 'v' to string.
// (See json.MarshalIndent for 'prefix' and 'indent' usage.)
func MarshalIndentStr(v any, prefix, indent string) (string, error) {
	bb, err := json.MarshalIndent(v, prefix, indent)
	if err != nil {
		return "", err
	}
	return string(bb), nil
}

// MarshalIndentStrMust is like MarshalIndentStr but panics in case of error.
func MarshalIndentStrMust(v any, prefix, indent string) string {
	s, err := MarshalIndentStr(v, prefix, indent)
	if err != nil {
		panic(err)
	}
	return s
}

// Format reads JSON-encoded data from 'r', then writes formatted data to 'w'.
// (See json.MarshalIndent for 'prefix' and 'indent' usage.)
func Format(r io.Reader, w io.Writer, prefix, indent string) error {
	var v any
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		return err
	}
	e := json.NewEncoder(w)
	e.SetIndent(prefix, indent)
	return e.Encode(v)
}

// FormatStr formats JSON-encoded 'json' to string.
// (See json.MarshalIndent for 'prefix' and 'indent' usage.)
func FormatStr(json, prefix, indent string) (string, error) {
	var b strings.Builder
	if err := Format(strings.NewReader(json), &b, prefix, indent); err != nil {
		return "", err
	}
	return b.String(), nil
}

// FormatStrMust is like FormatStr but panics in case of error.
func FormatStrMust(json, prefix, indent string) string {
	s, err := FormatStr(json, prefix, indent)
	if err != nil {
		panic(err)
	}
	return s
}
