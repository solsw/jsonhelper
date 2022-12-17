package jsonhelper

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestIndentRW(t *testing.T) {
	type args struct {
		r      io.Reader
		prefix string
		indent string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "1e",
			args: args{
				r:      &strings.Reader{},
				prefix: "",
				indent: "  ",
			},
			wantErr: true,
		},
		{name: "1",
			args: args{
				r:      strings.NewReader(`{"i":1,"s":"string"}`),
				prefix: "",
				indent: "  ",
			},
			want: "{\n  \"i\": 1,\n  \"s\": \"string\"\n}\n",
		},
		{name: "2",
			args: args{
				r:      strings.NewReader(`[{"i":1,"s":"one"}]`),
				prefix: "",
				indent: "  ",
			},
			want: "[\n  {\n    \"i\": 1,\n    \"s\": \"one\"\n  }\n]\n",
		},
		{name: "3",
			args: args{
				r:      strings.NewReader(`[{"i":1,"s":"one"},{"i":2,"s":"two"}]`),
				prefix: "",
				indent: "  ",
			},
			want: "[\n  {\n    \"i\": 1,\n    \"s\": \"one\"\n  },\n  {\n    \"i\": 2,\n    \"s\": \"two\"\n  }\n]\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := IndentRW(tt.args.r, w, tt.args.prefix, tt.args.indent); (err != nil) != tt.wantErr {
				t.Errorf("IndentRW() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got := w.String(); got != tt.want {
				t.Errorf("IndentRW() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndentStr(t *testing.T) {
	type args struct {
		json   string
		prefix string
		indent string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1",
			args: args{
				json:   `{"i":1,"s":"string"}`,
				prefix: "",
				indent: "  ",
			},
			want: "{\n  \"i\": 1,\n  \"s\": \"string\"\n}\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := IndentStr(tt.args.json, tt.args.prefix, tt.args.indent); got != tt.want {
				t.Errorf("IndentStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
