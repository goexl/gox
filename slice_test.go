package gox_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/goexl/gox"
)

type UserSlice struct {
	User     gox.Slice[User]  `json:"user,omitempty"`
	Users    gox.Slice[User]  `json:"users,omitempty"`
	UserPtr  gox.Slice[*User] `json:"user_ptr,omitempty"`
	UsersPtr gox.Slice[*User] `json:"users_ptr,omitempty"`
}

func receiveSlice[T any](data gox.Slice[T]) {
	fmt.Println(data)
}

func receiveIntSlice(data []int) {
	fmt.Println(data)
}

func TestNewSlice(t *testing.T) {
	receiveSlice(gox.NewSlice(1))
	receiveIntSlice(gox.NewSlice(2))
}

func TestSliceJSON(t *testing.T) {
	slice := new(UserSlice)
	if bytes, rfe := os.ReadFile("testdata/json/user_slice.json"); nil != rfe {
		t.Errorf("读取文件内容出错，%v", rfe)
	} else if ue := json.Unmarshal(bytes, slice); nil != ue {
		t.Errorf("反序列化JSON出错，%v", ue)
	} else {
		checkUserSlice(t, slice)
	}
}

func checkUserSlice(t *testing.T, slice *UserSlice) {
	if 1 != slice.User.Length() {
		t.Error("User字段反序列化后不是只有一个元素")
	} else if "storezhang" != slice.User[0].Name && 39 != slice.User[0].Age {
		t.Error("User字段反序列化后字段值不正确")
	} else if 1 != slice.UserPtr.Length() {
		t.Error("User指针字段反序列化后不是只有一个元素")
	} else if "store" != slice.UserPtr[0].Name && 19 != slice.UserPtr[0].Age {
		t.Error("User指针字段反序列化后字段值不正确")
	} else if 2 != slice.Users.Length() {
		t.Error("Users字段反序列化后字段长度正确")
	} else if 3 != slice.UsersPtr.Length() {
		t.Error("Users指针字段反序列化后字段长度正确")
	}
}
