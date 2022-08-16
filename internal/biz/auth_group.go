package biz

import (
	"time"
)

type AuthGroup struct {
	ID          int64            `bun:"id,pk,autoincrement"`
	Name        string           `bun:"name,nullzero,notnull,unique"`
	Permissions []AuthPermission `bun:"m2m:auth_group_permissions,join:AuthGroup=AuthPermission"`
	Users       []AuthUser       `bun:"m2m:auth_user_groups,join:AuthGroup=AuthUser"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

type AuthGroupPermission struct {
	GroupID   int64      `bun:"group_id,pk"`
	AuthGroup *AuthGroup `bun:"rel:belongs-to,join:group_id=id"`

	PermissionID   int64           `bun:"permission_id,pk"`
	AuthPermission *AuthPermission `bun:"rel:belongs-to,join:permission_id=id"`
}
