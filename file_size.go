package gox

import (
	`fmt`
	`strconv`
	`strings`
)

// 配置格式 "10G100M100K"
// 单位：B
// 1G = 1*1024*1024*1024
// 1M = 1*1024*1024
// 1K = 1*1024
var (
	sizeMap = map[units]int64{
		g: 1024 * 1024 * 1024,
		m: 1024 * 1024,
		k: 1024,
	}
)

const (
	g units = "g"
	m units = "m"
	k units = "k"
)

type (
	// 默认大小:B
	FileSize string
	// 计量单位
	units string
)

// 返回多少B
func (fs *FileSize) Size() (num int64) {
	var (
		str string
		sm  = make(map[units]int)
		i   int
	)
	str = fs.toString()
	for k := range sizeMap {
		i, str = fs.toInt(str, k)
		sm[k] = i
	}
	if len(str) != 0 {
		panic(fmt.Sprintf("配置错误:%v", str))
	}

	for k, i := range sm {
		num = num + sizeMap[k]*int64(i)
	}

	return
}

func (fs *FileSize) toInt(str string, u units) (i int, ret string) {
	var err error
	strs := strings.SplitN(str, u.toString(), 2)

	if len(strs) == 0 {
		return
	}

	if len(strs) == 1 {
		ret = strs[0]
		return
	}

	if i, err = strconv.Atoi(strs[0]); nil != err {
		panic(fmt.Sprintf("配置错误:%v", err))
	} else {
		return i, strs[1]
	}
	return 0, ""
}

func (fs *FileSize) toString() string {
	return strings.TrimSpace(strings.ToLower(string(*fs)))
}

func (u *units) toString() string {
	return string(*u)
}
