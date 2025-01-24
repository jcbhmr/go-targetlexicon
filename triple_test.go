package targetlexicon

import "testing"
import "github.com/stretchr/testify/assert"

func TestTripleParseErrors(t *testing.T) {
	v, err := TripleFromStr("")
	assert.NotNil(t, err)
	v, err = TripleFromStr("foo")
	assert.NotNil(t, err)
	v, err = TripleFromStr("unknown-unknown-foo")
	assert.NotNil(t, err)
	v, err = TripleFromStr("unknown-unknown-unknown-foo")
	assert.NotNil(t, err)
	v, err = TripleFromStr("unknown-unknown-unknown-unknown-foo")
	assert.NotNil(t, err)
	v, err = TripleFromStr("unknown-unknown-unknown-unknown-unknown-foo")
	assert.NotNil(t, err)
	_ = v
}

func TestTripleDefaults(t *testing.T) {
	v, err := TripleFromStr("unknown-unknown-unknown")
	assert.Nil(t, err)
	assert.Equal(t, TripleUnknown(), v)
	v, err = TripleFromStr("unknown-unknown-unknown-unknown")
	assert.Nil(t, err)
	assert.Equal(t, TripleUnknown(), v)
	v, err = TripleFromStr("unknown-unknown-unknown-unknown-unknown")
	assert.Nil(t, err)
	assert.Equal(t, TripleUnknown(), v)
	_ = v
}

func TestTripleUnknownProperties(t *testing.T) {
	v, ok := TripleUnknown().Endianness()
	assert.Equal(t, ok, false)
	v, ok = TripleUnknown().PointerWidth()
	assert.Equal(t, ok, false)
	v, ok = TripleUnknown().DefaultCallingConvention()
	assert.Equal(t, ok, false)
	_ = v
}

func TestTripleAppleCallingConvention(t *testing.T) {
	for triple := range []string{
		"aarch64-apple-darwin",
		"aarch64-apple-ios",
		"aarch64-apple-ios-macabi",
		"aarch64-apple-tvos",
		"aarch64-apple-watchos",
	} {
		triple2, err := TripleFromStr(triple)
		assert.Nil(t, err)
		d, ok := triple2.DefaultCallingConvention()
		assert.True(t, ok)
		assert.Equal(t, CallingConventionAppleAArch64, d)
	}
}

func TestTripleP32ABI(t *testing.T) {
	for triple := range []string{
		"x86_64-unknown-linux-gnux32",
		"aarch64_be-unknown-linux-gnu_ilp32",
		"aarch64-unknown-linux-gnu_ilp32",
	} {
		triple2, err := TripleFromStr(triple)
		assert.Nil(t, err)
		p, ok := triple2.PointerWidth()
		assert.True(t, ok)
		assert.Equal(t, PointerWidthU32, p)
	}

	for triple := range []string{
		"x86_64-unknown-linux-gnu",
		"aarch64_be-unknown-linux-gnu",
		"aarch64-unknown-linux-gnu",
	} {
		triple2, err := TripleFromStr(triple)
		assert.Nil(t, err)
		p, ok := triple2.PointerWidth()
		assert.True(t, ok)
		assert.Equal(t, PointerWidthU64, p)
	}
}
