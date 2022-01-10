package file_test

import (
	`testing`

	`github.com/storezhang/gox/file`
)

func TestCopy(t *testing.T) {
	if err := file.Copy(`../testdata/file/copy`, `../testdata/file/test`); nil != err {
		t.FailNow()
	}
}
