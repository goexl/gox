package gox_test

import (
	"encoding/json"
	"testing"

	"github.com/goexl/gox"
)

func TestBitmapJSON(t *testing.T) {
	tests := []struct {
		enabled []uint
	}{
		{[]uint{1}},
		{[]uint{100}},
		{[]uint{1000000}},
		{[]uint{99, 843, 80879}},
	}
	for index, test := range tests {
		original := gox.NewBitmap()
		original.Set(test.enabled[0], test.enabled[1:]...)

		decoded := gox.NewBitmap()
		if data, me := json.Marshal(original); nil != me {
			t.Errorf("第%d个测试未通过，编码出错：%v", index+1, me)
		} else if ue := json.Unmarshal(data, decoded); nil != ue {
			t.Errorf("第%d个测试未通过，解码出错：%v", index+1, ue)
		} else if !decoded.All(test.enabled[0], test.enabled[1:]...) {
			t.Errorf("第%d个测试未通过，编解码后数据不一致", index+1)
		} else {
			t.Logf("第%d个测试编码结果：%s", index+1, string(data))
		}
	}
}
