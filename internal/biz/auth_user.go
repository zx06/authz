package biz

import (
	"time"
)

type AuthUser struct {
	ID          int64            `bun:"id,pk,autoincrement"`
	UID         string           `bun:"uid,nullzero,notnull,unique"`
	Groups      []AuthGroup      `bun:"m2m:auth_user_groups,join:AuthUser=AuthGroup"`
	Permissions []AuthPermission `bun:"m2m:auth_user_permissions,join:AuthUser=AuthPermission"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

type AuthUserGroup struct {

	UserID int64     `bun:"user_id,pk"`
	AuthUser   *AuthUser `bun:"rel:belongs-to,join:user_id=id"`

	GroupID int64      `bun:"group_id,pk"`
	AuthGroup  *AuthGroup `bun:"rel:belongs-to,join:group_id=id"`
}

type AuthUserPermission struct {
	UserID int64     `bun:"user_id,pk"`
	AuthUser   *AuthUser `bun:"rel:belongs-to,join:user_id=id"`

	PermissionID int64           `bun:"permission_id,pk"`
	AuthPermission   *AuthPermission `bun:"rel:belongs-to,join:permission_id=id"`
}
