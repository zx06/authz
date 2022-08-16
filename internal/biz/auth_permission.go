package biz

import (
	"time"

)

type AuthPermission struct {
	ID          int64  `bun:"id,pk,autoincrement"`
	Description string `bun:"description,nullzero,notnull"`
	Resource    string `bun:"resource,nullzero,notnull,unique"`
	Action      string `bun:"resource,nullzero,notnull"`

	Users  []AuthUser  `bun:"m2m:auth_user_permissions,join:AuthPermission=AuthUser"`
	Groups []AuthGroup `bun:"m2m:auth_group_permissions,join:AuthPermission=AuthGroup"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
