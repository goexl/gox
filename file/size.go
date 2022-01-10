package file

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
		gb: 1024 * 1024 * 1024,
		mb: 1024 * 1024,
		kb: 1024,
	}
)

const (
	gb units = "gb"
	mb units = "mb"
	kb units = "kb"
)

type (
	// Size 默认大小:B
	Size string
	// 计量单位
	units string
)

// Size 返回多少字节
func (s *Size) Size() (num int64) {
	var (
		str string
		sm  = make(map[units]int)
		i   int
	)
	str = s.toString()
	for k := range sizeMap {
		i, str = s.toInt(str, k)
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

func (s *Size) toInt(str string, u units) (i int, ret string) {
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

func (s *Size) toString() string {
	return strings.TrimSpace(strings.ToLower(string(*s)))
}

func (u *units) toString() string {
	return string(*u)
}
