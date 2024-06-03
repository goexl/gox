package size_test

import (
	"encoding/json"
	"testing"

	"github.com/goexl/gox/internal/size"
)

type bytesJson struct {
	Bytes size.Bytes `json:"bytes"`
}

func TestParseBytes(t *testing.T) {
	tests := []struct {
		in     string
		expect size.Bytes
	}{
		{in: "1m", expect: size.BytesMB},
		{in: "1g", expect: size.BytesGB},
		{in: "1g 1m", expect: size.BytesGB + size.BytesMB},
	}

	for _, test := range tests {
		actual, _ := size.ParseBytes(test.in)
		if actual != test.expect {
			t.Fatalf("期望：%v，实际：%v", test.expect, actual)
		}
	}
}

func TestUnmarshalBytes(t *testing.T) {
	tests := []struct {
		in     string
		expect bytesJson
	}{
		{in: `{"bytes": "1m"}`, expect: bytesJson{Bytes: size.BytesMB}},
		{in: `{"bytes": "1g"}`, expect: bytesJson{Bytes: size.BytesGB}},
		{in: `{"bytes": "1g 1m"}`, expect: bytesJson{Bytes: size.BytesGB + size.BytesMB}},
	}

	for _, test := range tests {
		bytes := new(bytesJson)
		if err := json.Unmarshal([]byte(test.in), bytes); nil != err {
			t.Fatalf("错误：%v", err)
		} else if bytes.Bytes != test.expect.Bytes {
			t.Fatalf("期望：%v，实际：%v", test.expect, bytes.Bytes)
		}
	}
}

func TestMarshalSize(t *testing.T) {
	tests := []struct {
		in     bytesJson
		expect string
	}{
		{in: bytesJson{Bytes: size.BytesMB}, expect: `{"bytes":"1m"}`},
		{in: bytesJson{Bytes: size.BytesGB}, expect: `{"bytes":"1g"}`},
		{in: bytesJson{Bytes: size.BytesGB + size.BytesMB}, expect: `{"bytes":"1g 1m"}`},
	}

	for _, test := range tests {
		if bytes, err := json.Marshal(test.in); nil != err {
			t.Fatalf("错误：%v", err)
		} else if string(bytes) != test.expect {
			t.Fatalf("期望：%v，实际：%v", test.expect, string(bytes))
		}
	}
}

func TestFormatSize(t *testing.T) {
	tests := []struct {
		in     size.Bytes
		expect string
	}{
		{in: size.BytesMB, expect: "1m"},
		{in: size.BytesGB, expect: "1g"},
		{in: size.BytesGB + size.BytesMB, expect: "1g 1m"},
	}

	for _, test := range tests {
		actual := test.in.Formatter().Format()
		if actual != test.expect {
			t.Fatalf("期望：%v，实际：%v", test.expect, actual)
		}
	}
}
