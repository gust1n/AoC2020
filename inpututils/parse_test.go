package inpututils_test

import (
	"testing"

	"github.com/gust1n/AoC2020/inpututils"
	"gotest.tools/assert"
)

func TestParseInts(t *testing.T) {
	cases := []struct {
		input   []byte
		want    []int
		wantErr bool
	}{
		{
			input: []byte(`
11
22
33`),
			want:    []int{11, 22, 33},
			wantErr: false,
		},
		{
			input: []byte(`
5511
4422
3355

91928
`),
			want:    []int{5511, 4422, 3355, 91928},
			wantErr: false,
		},
		{
			input: []byte(`
hej
jag
kodar
fint
`),
			want:    nil,
			wantErr: true,
		},
	}

	for _, tc := range cases {
		got, gotErr := inpututils.ParseInts(tc.input)
		assert.DeepEqual(t, got, tc.want)
		assert.Equal(t, gotErr != nil, tc.wantErr)
	}
}
