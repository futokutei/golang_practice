package program

import (
	"strings"
	"testing"
)

type User struct {
	Name    string   `json:"user_name"`
	Age     int      `json:"age"`
	IsAdmin bool     `json:"is_admin"`
	Tags    []string `json:"tags"`
}

type SimpleStruct struct {
	Title string
}

func TestToYAML(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		want    string
		wantErr bool
	}{
		{
			name: "Success with basic struct and tags",
			input: User{
				Name:    "Alice",
				Age:     30,
				IsAdmin: true,
				Tags:    []string{"golang", "dev"},
			},
			want:    "user_name: \"Alice\"\nage: 30\nis_admin: true\ntags: \n   - \"golang\"\n   - \"dev\"\n",
			wantErr: false,
		},
		{
			name: "Success with pointer to struct",
			input: &User{
				Name:    "Bob",
				Age:     25,
				IsAdmin: false,
				Tags:    []string{"qa"},
			},
			want:    "user_name: \"Bob\"\nage: 25\nis_admin: false\ntags: \n   - \"qa\"\n",
			wantErr: false,
		},
		{
			name: "Success without json tags (uses field name)",
			input: SimpleStruct{
				Title: "Hello Go",
			},
			want:    "Title: \"Hello Go\"",
			wantErr: false,
		},
		{
			name:    "Error when passing non-struct (int)",
			input:   42,
			want:    "",
			wantErr: true,
		},
		{
			name:    "Error when passing non-struct (string)",
			input:   "some string",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToYAML(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if strings.TrimSpace(got) != strings.TrimSpace(tt.want) {
				t.Errorf("ToYAML() \ngot:  %q\nwant: %q", got, tt.want)
			}
		})
	}
}
