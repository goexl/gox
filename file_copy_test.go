package gox

import (
	`testing`
)

func TestCopyFile(t *testing.T) {
	if _, err := CopyFile("./zip.go", "../test"); nil != err {
		t.Errorf(" CopyFile err: %v", err)
	}
	if _, err := CopyFile("./zip.go", "../test/1.go"); nil != err {
		t.Errorf(" CopyFile err: %v", err)
	}

	//if _,err:= 	CopyFile("./","../test1");nil != err{
	//	t.Errorf(" CopyFile err: %v", err)
	//}
}
