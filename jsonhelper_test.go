package jsonhelper

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestMarshalIndentStr(t *testing.T) {
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
				v:      complex(1.2, 3.4),
				prefix: "",
				indent: "  ",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MarshalIndentStr(tt.args.v, tt.args.prefix, tt.args.indent)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalIndentStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MarshalIndentStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarshalIndentStrMust(t *testing.T) {
	type is struct {
		I int
		S string
	}
	type args struct {
		v      any
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
				v:      is{1, "one"},
				prefix: "",
				indent: "  ",
			},
			want: `{
  "I": 1,
  "S": "one"
}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MarshalIndentStrMust(tt.args.v, tt.args.prefix, tt.args.indent); got != tt.want {
				t.Errorf("MarshalIndentStrMust() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormat(t *testing.T) {
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
			if err := Format(tt.args.r, w, tt.args.prefix, tt.args.indent); (err != nil) != tt.wantErr {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got := w.String(); got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatStrMust(t *testing.T) {
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
			if got := FormatStrMust(tt.args.json, tt.args.prefix, tt.args.indent); got != tt.want {
				t.Errorf("FormatStrMust() = %v, want %v", got, tt.want)
			}
		})
	}
}
