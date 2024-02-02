package app

import (
	"fmt"
	"testing"
)

func Test_setEncoreAppID(t *testing.T) {
	tests := []struct {
		data         []byte
		id           string
		commentLines []string
		want         string
	}{
		{
			data:         []byte(`{}`),
			id:           "foo",
			commentLines: []string{"bar"},
			want: `{
	"$schema": "https://encore.dev/schemas/appfile.schema.json",
	// bar
	"id": "foo",
}
`,
		},
		{
			data:         []byte(``),
			id:           "foo",
			commentLines: []string{"bar"},
			want: `{
	"$schema": "https://encore.dev/schemas/appfile.schema.json",
	// bar
	"id": "foo",
}
`,
		},
		{
			data: []byte(`{
	// foo
	"id": "test",
}`),
			id:           "foo",
			commentLines: []string{"bar", "baz"},
			want: `{
	"$schema": "https://encore.dev/schemas/appfile.schema.json",
	// bar
	// baz
	"id": "foo",
}
`,
		},
		{
			data: []byte(`{
	"$schema": "https://encore.dev/AN-OLD-SCHEMA",
	"some_other_field": true,
	// foo
	"id": "test",
}`),
			id:           "foo",
			commentLines: []string{"bar", "baz"},
			want: `{
	"$schema":          "https://encore.dev/schemas/appfile.schema.json",
	"some_other_field": true,
	// bar
	// baz
	"id": "foo",
}
`,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got, err := setEncoreAppID(tt.data, tt.id, tt.commentLines)
			if err != nil {
				t.Fatal(err)
			}

			gotStr := string(got)
			if gotStr != tt.want {
				t.Errorf("setEncoreAppID() = %q, want %q", gotStr, tt.want)
			}
		})
	}
}
