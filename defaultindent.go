package jsonhelper

import (
	"io"
)

// https://docs.openstack.org/doc-contrib-guide/json-conv.html
var (
	// Default prefix (see [json.Indent]). Assign desired value here, if needed.
	DefaultPrefix = ""
	// Default indent (see [json.Indent]). Assign desired value here, if needed.
	DefaultIndent = "    "
)

// DefaultRW reads JSON data from [io.Reader] and writes default indented JSON data to [io.Writer].
func DefaultRW(r io.Reader, w io.Writer) error {
	return IndentRW(r, w, DefaultPrefix, DefaultIndent)
}

// DefaultStr returns the default indented form of JSON string 'j'.
func DefaultStr(j string) (string, error) {
	return IndentStr(j, DefaultPrefix, DefaultIndent)
}

// DefaultAny returns the default indented JSON encoding of 'v' as string.
func DefaultAny(v any) (string, error) {
	return IndentAny(v, DefaultPrefix, DefaultIndent)
}
