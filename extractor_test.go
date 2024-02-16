package ddtags_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/webdestroya/ddtags"
)

type dummyTagStruct struct {
	TagA         string  `ddtag:"tag_a"`
	TagB         string  `ddtag:"tag_b"`
	EmptyTag     string  `ddtag:"empty"`
	StrPtr       *string `ddtag:"strptr"`
	IgnoredTag   string  `ddtag:"-"`
	NonTaggedVal string
	MapType      map[string]any `ddtag:"maptype"`

	IntP  *int  `ddtag:"intp"`
	BoolP *bool `ddtag:"boolp"`
}

type simpleTagStruct struct {
	TagA string `ddtag:"taga"`
	TagB string `ddtag:"tagb"`
	TagC string `ddtag:"tagc"`
}

type fullTagStruct struct {
	Str  string  `ddtag:"str"`
	StrP *string `ddtag:"strp"`

	Bool  bool  `ddtag:"bool"`
	BoolP *bool `ddtag:"boolp"`

	Int   int   `ddtag:"int"`
	Int8  int8  `ddtag:"int8"`
	Int16 int16 `ddtag:"int16"`
	Int32 int32 `ddtag:"int32"`
	Int64 int64 `ddtag:"int64"`

	IntP   *int   `ddtag:"intp"`
	Int8P  *int8  `ddtag:"int8p"`
	Int16P *int16 `ddtag:"int16p"`
	Int32P *int32 `ddtag:"int32p"`
	Int64P *int64 `ddtag:"int64p"`

	Uint   uint   `ddtag:"uint"`
	Uint8  uint8  `ddtag:"uint8"`
	Uint16 uint16 `ddtag:"uint16"`
	Uint32 uint32 `ddtag:"uint32"`
	Uint64 uint64 `ddtag:"uint64"`

	UintP   *uint   `ddtag:"uintp"`
	Uint8P  *uint8  `ddtag:"uint8p"`
	Uint16P *uint16 `ddtag:"uint16p"`
	Uint32P *uint32 `ddtag:"uint32p"`
	Uint64P *uint64 `ddtag:"uint64p"`

	Float32 float32 `ddtag:"float32"`
	Float64 float64 `ddtag:"float64"`

	Float32P *float32 `ddtag:"float32p"`
	Float64P *float64 `ddtag:"float64p"`

	Float64Fmt float64 `ddtag:"float64fmt,fmt=%.9f"`
	IntFmt     int     `ddtag:"intfmt,fmt=0x%08x"`
}

func TestExtract(t *testing.T) {
	t.Run("basic strings", func(t *testing.T) {
		struc := &dummyTagStruct{
			TagA:         "avalue",
			TagB:         "bvalue",
			NonTaggedVal: "ignored",
			IgnoredTag:   "ignored2",
			MapType:      map[string]any{"foo": "bar"},
			StrPtr:       ptr("testing"),
		}

		tags := ddtags.Extract(struc)
		require.NotNil(t, tags)
		require.Contains(t, tags, "tag_a:avalue")
		require.Contains(t, tags, "tag_b:bvalue")
		require.Contains(t, tags, "strptr:testing")
		require.Len(t, tags, 3)
	})

	t.Run("expected tags", func(t *testing.T) {
		tables := []struct {
			value *dummyTagStruct
			exp   []string
		}{

			{
				value: &dummyTagStruct{IntP: ptr(int(123))},
				exp:   []string{"intp:123"},
			},
		}

		for tableNum, table := range tables {
			t.Run(fmt.Sprintf("table_%02d", tableNum), func(t *testing.T) {
				tags := ddtags.Extract(table.value)
				require.IsType(t, []string{}, tags)
				require.Len(t, tags, len(table.exp))
				if len(table.exp) > 0 {
					for _, v := range table.exp {
						require.Contains(t, tags, v)
					}
				}
			})
		}
	})

	t.Run("bad types", func(t *testing.T) {
		tables := []struct {
			value any
		}{
			{nil},
			{int(123)},
			{dummyTagStruct{}},
			{ptr(int(1234))},
		}

		for tableNum, table := range tables {
			t.Run(fmt.Sprintf("table_%02d", tableNum), func(t *testing.T) {
				tags := ddtags.Extract(table.value)
				require.IsType(t, []string{}, tags)
				require.Len(t, tags, 0)
			})
		}
	})
}

func TestExtractFull(t *testing.T) {
	tags := ddtags.Extract(buildFull())
	expected := []string{
		"str:strval",
		"strp:strp",

		"bool:true",
		"boolp:false",

		"int:1",
		"int8:1",
		"int16:1",
		"int32:1",
		"int64:1",

		"intp:0",
		"int8p:0",
		"int16p:0",
		"int32p:0",
		"int64p:0",

		"uint:1",
		"uint8:1",
		"uint16:1",
		"uint32:1",
		"uint64:1",

		"uintp:0",
		"uint8p:0",
		"uint16p:0",
		"uint32p:0",
		"uint64p:0",

		"float32:1.30000",
		"float64:1.30000",

		"float32p:1.30000",
		"float64p:1.30000",

		"float64fmt:0.333333333",
		"intfmt:0x000004d2",
	}

	require.ElementsMatch(t, tags, expected)

}

func BenchmarkExtract(b *testing.B) {

	tags := &fullTagStruct{
		Str:   "strv",
		StrP:  ptr("strpv"),
		Bool:  true,
		BoolP: ptr(false),

		Float64Fmt: (float64(1) / float64(3)),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ddtags.Extract(tags)
	}
}

func BenchmarkExtractSimple(b *testing.B) {

	tags := &simpleTagStruct{
		TagA: "foo",
		TagB: "bar",
		TagC: "baz",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ddtags.Extract(tags)
	}
}

func BenchmarkExtractFull(b *testing.B) {

	tags := buildFull()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ddtags.Extract(tags)
	}
}

func ptr[T any](v T) *T {
	return &v
}

func buildFull() *fullTagStruct {
	return &fullTagStruct{
		Str:        "strval",
		Bool:       true,
		StrP:       ptr("strp"),
		BoolP:      ptr(false),
		Int:        int(1),
		Int8:       int8(1),
		Int16:      int16(1),
		Int32:      int32(1),
		Int64:      int64(1),
		IntP:       ptr(int(0)),
		Int8P:      ptr(int8(0)),
		Int16P:     ptr(int16(0)),
		Int32P:     ptr(int32(0)),
		Int64P:     ptr(int64(0)),
		Uint:       uint(1),
		Uint8:      uint8(1),
		Uint16:     uint16(1),
		Uint32:     uint32(1),
		Uint64:     uint64(1),
		UintP:      ptr(uint(0)),
		Uint8P:     ptr(uint8(0)),
		Uint16P:    ptr(uint16(0)),
		Uint32P:    ptr(uint32(0)),
		Uint64P:    ptr(uint64(0)),
		Float32:    float32(1.3),
		Float64:    float64(1.3),
		Float32P:   ptr(float32(1.3)),
		Float64P:   ptr(float64(1.3)),
		Float64Fmt: (float64(1) / float64(3)),
		IntFmt:     1234,
	}
}
