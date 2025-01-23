package jsonhelper

import (
	"bytes"
	"encoding/json"
	"io"
)

// IndentRW reads JSON data from [io.Reader] and writes indented data to [io.Writer].
// (See [json.Indent] for 'prefix' and 'indent' usage.)
func IndentRW(r io.Reader, w io.Writer, prefix, indent string) error {
	bb, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := json.Indent(&buf, bb, prefix, indent); err != nil {
		return err
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		return err
	}
	return nil
}

// IndentStr returns the indented form of JSON string 'j'.
// (See [json.Indent] for 'prefix' and 'indent' usage.)
func IndentStr(j, prefix, indent string) (string, error) {
	var buf bytes.Buffer
	if err := json.Indent(&buf, []byte(j), prefix, indent); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// DefaultStr returns the default indented form of JSON string 'j'.
func DefaultStr(j string) (string, error) {
	// https://docs.openstack.org/doc-contrib-guide/json-conv.html
	return IndentStr(j, "", "    ")
}

// IndentAny returns the indented JSON encoding of 'v' as string.
// (See [json.Indent] for 'prefix' and 'indent' usage.)
func IndentAny(v any, prefix, indent string) (string, error) {
	bb, err := json.MarshalIndent(v, prefix, indent)
	if err != nil {
		return "", err
	}
	return string(bb), nil
}

// DefaultAny returns the default indented JSON encoding of 'v' as string.
func DefaultAny(v any) (string, error) {
	// https://docs.openstack.org/doc-contrib-guide/json-conv.html
	return IndentAny(v, "", "    ")
}
