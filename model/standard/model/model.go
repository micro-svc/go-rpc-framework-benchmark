package model

var (
	Example = example()
)

func example() *Message {
	sub1 := SubMessage{
		Field01: 123456.789,
		Field02: 1234567890,
		Field03: true,
		Field04: "string example",
		Field05: []byte("bytes example"),
		Field06: Enum_EnumC,
		Field07: []float64{123456.789, 223456.789, 323456.789},
		Field08: []int64{1234567890, 2234567890, 3234567890},
		Field09: []bool{true, false, true},
		Field10: []string{"string example 1", "string example 2", "string example 3"},
		Field11: [][]byte{[]byte("bytes example 1"), []byte("bytes example 2"), []byte("bytes example 3")},
		Field12: []Enum{Enum_EnumA, Enum_EnumB, Enum_EnumC},
	}

	msg := &Message{
		Field01: 123456.789,
		Field02: 1234567890,
		Field03: true,
		Field04: "string example",
		Field05: []byte("bytes example"),
		Field06: Enum_EnumC,
		Field07: []float64{123456.789, 223456.789, 323456.789},
		Field08: []int64{1234567890, 2234567890, 3234567890},
		Field09: []bool{true, false, true},
		Field10: []string{"string example 1", "string example 2", "string example 3"},
		Field11: [][]byte{[]byte("bytes example 1"), []byte("bytes example 2"), []byte("bytes example 3")},
		Field12: []Enum{Enum_EnumA, Enum_EnumB, Enum_EnumC},
		Field13: &sub1,
	}

	return msg
}
