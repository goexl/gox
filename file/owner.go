package file

type owner struct {
	uid int
	gid int
}

// NewOwner 创建拥有者
func NewOwner(uid int, gid int) *owner {
	return &owner{
		uid: uid,
		gid: gid,
	}
}
