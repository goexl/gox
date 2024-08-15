package gox

// Incomparable 不可比较
// 包含本结构的结构体不能用于排序和比较大小
type Incomparable [0]func()
