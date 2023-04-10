package morse

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/serialt/lancet/internal"
)

var tests = []struct {
	input     string // 输入值
	separator string // 分隔符
	output    string // 期望值
}{
	{"", "", ""},
	{"1", "/", ".----"},
	{"F", "/", "..-."},
	{"dongle", "|", "-..|---|-.|--.|.-..|."},
	{"SOS", "/", ".../---/..."},
}

func TestMorseEncode(t *testing.T) {
	assert := internal.NewAssert(t, "TestMorseEncode")
	for index, test := range tests {
		dst, err := Encode(string2bytes(test.input), test.separator)

		t.Run(fmt.Sprintf("test_%d", index), func(t *testing.T) {

			assert.IsNil(err)
			assert.Equal(test.output, dst)
		})
	}
}

func TestDecode(t *testing.T) {
	assert := internal.NewAssert(t, "TestMorseDecode")
	for index, test := range tests {
		dst, err := Decode(string2bytes(test.output), test.separator)

		t.Run(fmt.Sprintf("test_%d", index), func(t *testing.T) {
			assert.IsNil(err)
			assert.Equal(strings.ToLower(test.input), dst)
		})
	}
}

func TestMorseError(t *testing.T) {
	assert := internal.NewAssert(t, "TestMorseError")
	_, err1 := Encode([]byte("hello world"), "/")
	assert.Equal(errors.New("can't contain spaces"), err1)

	_, err2 := Decode([]byte("hello world"), "/")
	assert.Equal(fmt.Errorf("unknown character hello world"), err2)
}
