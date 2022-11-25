package gox

const (
	// CasePositionNone 无
	CasePositionNone casePosition = iota
	// CasePositionHead 首
	CasePositionHead
	// CasePositionTail 尾
	CasePositionTail
	// CasePositionAll 全部
	CasePositionAll
)

type casePosition uint8
