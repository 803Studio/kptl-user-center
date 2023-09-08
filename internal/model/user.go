package model

const (
	RoleBoss  byte = 0
	RoleAdmin byte = 1
	RoleUser  byte = 2
	RoleBaned byte = 3
)

type UserAccount struct {
	Id   uint32
	WxId string
	Role byte
}

var keys = []string{"id", "wxid", "role"}

func (*UserAccount) Keys() []string {
	return keys
}

func (u *UserAccount) Values() []any {
	return []any{
		u.Id,
		u.WxId,
		u.Role,
	}
}

func (u *UserAccount) PtrVec() []any {
	return []any{
		&u.Id,
		&u.WxId,
		&u.Role,
	}
}
