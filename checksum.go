package gox

const (
	// ChecksumTypeCrc32 Crc32
	ChecksumTypeCrc32 ChecksumType = "crc32"
	// ChecksumTypeCrc64 Crc64
	ChecksumTypeCrc64 ChecksumType = "crc64"
	// ChecksumTypeSha1 Crc64
	ChecksumTypeSha1 ChecksumType = "sha1"
	// ChecksumTypeSha256 Crc64
	ChecksumTypeSha256 ChecksumType = "sha256"
)

type (
	// ChecksumType 数据校验合类型
	ChecksumType string

	// Checksum 数据校验
	Checksum struct {
		// Sum 校验字符串
		Sum string `json:"sum" validate:"required"`
		// Type 类型
		Type ChecksumType `json:"type" validate:"required,oneof=crc32 crc64 sha1 sha256"`
	}
)
