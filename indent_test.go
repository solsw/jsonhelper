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
				r:      strings.NewReader(`{"i":1,"s":"one"}`),
				prefix: "",
				indent: "  ",
			},
			want: "{\n  \"i\": 1,\n  \"s\": \"one\"\n}",
		},
		{name: "2",
			args: args{
				r:      strings.NewReader(`[{"i":1,"s":"one"}]`),
				prefix: "",
				indent: "  ",
			},
			want: "[\n  {\n    \"i\": 1,\n    \"s\": \"one\"\n  }\n]",
		},
		{name: "3",
			args: args{
				r:      strings.NewReader(`[{"i":1,"s":"one"},{"i":2,"s":"two"}]`),
				prefix: "",
				indent: "  ",
			},
			want: "[\n  {\n    \"i\": 1,\n    \"s\": \"one\"\n  },\n  {\n    \"i\": 2,\n    \"s\": \"two\"\n  }\n]",
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
		j      string
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
				j:      `{"i":1,"s":"one"}`,
				prefix: "",
				indent: "  ",
			},
			want: "{\n  \"i\": 1,\n  \"s\": \"one\"\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := IndentStr(tt.args.j, tt.args.prefix, tt.args.indent)
			if got != tt.want {
				t.Errorf("IndentStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndentAny(t *testing.T) {
	type args struct {
		v      any
		prefix string
		indent string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "1",
			args: args{
				v: struct {
					I int    `json:"i"`
					S string `json:"s"`
				}{I: 1, S: "one"},
				prefix: "",
				indent: "  ",
			},
			want: "{\n  \"i\": 1,\n  \"s\": \"one\"\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IndentAny(tt.args.v, tt.args.prefix, tt.args.indent)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndentAny() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IndentAny() = %v, want %v", got, tt.want)
			}
		})
	}
}
