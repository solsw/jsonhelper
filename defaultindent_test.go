package jsonhelper

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestDefaultRW(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "1e",
			args: args{
				r: &strings.Reader{},
			},
			wantErr: true,
		},
		{name: "1",
			args: args{
				r: strings.NewReader(`{"i":1,"s":"one"}`),
			},
			want: "{\n    \"i\": 1,\n    \"s\": \"one\"\n}",
		},
		{name: "2",
			args: args{
				r: strings.NewReader(`[{"i":1,"s":"one"}]`),
			},
			want: "[\n    {\n        \"i\": 1,\n        \"s\": \"one\"\n    }\n]",
		},
		{name: "3",
			args: args{
				r: strings.NewReader(`[{"i":1,"s":"one"},{"i":2,"s":"two"}]`),
			},
			want: "[\n    {\n        \"i\": 1,\n        \"s\": \"one\"\n    },\n" +
				"    {\n        \"i\": 2,\n        \"s\": \"two\"\n    }\n]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := DefaultRW(tt.args.r, w); (err != nil) != tt.wantErr {
				t.Errorf("DefaultRW() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got := w.String(); got != tt.want {
				t.Errorf("DefaultRW() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultStr(t *testing.T) {
	type args struct {
		j string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1",
			args: args{
				j: `{"i":1,"s":"one"}`,
			},
			want: "{\n    \"i\": 1,\n    \"s\": \"one\"\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := DefaultStr(tt.args.j)
			if got != tt.want {
				t.Errorf("DefaultStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultAny(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1",
			args: args{
				v: struct {
					I int    `json:"i"`
					S string `json:"s"`
				}{I: 1, S: "one"},
			},
			want: "{\n    \"i\": 1,\n    \"s\": \"one\"\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := DefaultAny(tt.args.v)
			if got != tt.want {
				t.Errorf("DefaultAny() = %v, want %v", got, tt.want)
			}
		})
	}
}
