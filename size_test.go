package gox_test

import (
	"encoding/json"
	"testing"

	"github.com/goexl/gox"
)

type sizeJson struct {
	Size gox.Size `json:"size"`
}

func TestParseSize(t *testing.T) {
	tests := []struct {
		in     string
		expect gox.Size
	}{
		{in: "1m", expect: gox.SizeMB},
		{in: "1g", expect: gox.SizeGB},
		{in: "1g 1m", expect: gox.SizeGB + gox.SizeMB},
	}

	for _, test := range tests {
		actual, _ := gox.ParseSize(test.in)
		if actual != test.expect {
			t.Fatalf("期望：%v，实际：%v", test.expect, actual)
		}
	}
}

func TestUnmarshalSize(t *testing.T) {
	tests := []struct {
		in     string
		expect sizeJson
	}{
		{in: `{"size": "1m"}`, expect: sizeJson{Size: gox.SizeMB}},
		{in: `{"size": "1g"}`, expect: sizeJson{Size: gox.SizeGB}},
		{in: `{"size": "1g 1m"}`, expect: sizeJson{Size: gox.SizeGB + gox.SizeMB}},
	}

	for _, test := range tests {
		size := new(sizeJson)
		if err := json.Unmarshal([]byte(test.in), size); nil != err {
			t.Fatalf("错误：%v", err)
		} else if size.Size != test.expect.Size {
			t.Fatalf("期望：%v，实际：%v", test.expect, size.Size)
		}
	}
}

func TestMarshalSize(t *testing.T) {
	tests := []struct {
		in     sizeJson
		expect string
	}{
		{in: sizeJson{Size: gox.SizeMB}, expect: `{"size":"1m"}`},
		{in: sizeJson{Size: gox.SizeGB}, expect: `{"size":"1g"}`},
		{in: sizeJson{Size: gox.SizeGB + gox.SizeMB}, expect: `{"size":"1g 1m"}`},
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
		in     gox.Size
		expect string
	}{
		{in: gox.SizeMB, expect: "1m"},
		{in: gox.SizeGB, expect: "1g"},
		{in: gox.SizeGB + gox.SizeMB, expect: "1g 1m"},
	}

	for _, test := range tests {
		actual := test.in.Formatter().Format()
		if actual != test.expect {
			t.Fatalf("期望：%v，实际：%v", test.expect, actual)
		}
	}
}
