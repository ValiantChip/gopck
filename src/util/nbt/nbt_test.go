package nbt_test

import (
	"testing"

	. "github.com/ValiantChip/gopck/src/util/nbt"
)

var nbtTests = []struct {
	in  any
	out string
}{
	{true, "true"},
	{false, "false"},
	{int8(1), "1b"},
	{int16(1), "1s"},
	{int32(1), "1"},
	{int64(1), "1l"},
	{float32(1.0), "1f"},
	{float32(3.56), "3.56f"},
	{float64(3.56), "3.56"},
	{"test", "\"test\""},
	{[]any{1, 2, 3}, "[1,2,3]"},
	{map[string]any{"a": 1, "b": false}, "{\"a\":1,\"b\":false}"},
	{map[string]any{"a": int32(27), "b": "false"}, "{\"a\":27,\"b\":\"false\"}"},
	{ByteArray{1, 2, 3}, "[B;1b,2b,3b]"},
	{IntArray{1, 2, 3}, "[I;1,2,3]"},
	{LongArray{1, 2, 3}, "[L;1l,2l,3l]"},
}

func TestParse(t *testing.T) {
	for _, tt := range nbtTests {
		t.Run(tt.out, func(t *testing.T) {
			s, err := Parse(tt.in)
			if err != nil {
				t.Error(err)
			}

			if s != tt.out {
				t.Errorf("got %s, want %s", s, tt.out)
			}
		},
		)
	}
}
