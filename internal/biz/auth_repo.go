package biz

import (
	"context"

	"github.com/uptrace/bun"
)

type AuthRepo interface {
	// UserHasPerm 用户是否有权限
	UserHasPerm(ctx context.Context, db bun.IDB, uid, perm, act string) (bool, error)
	// UserGetPermissions 获取用户权限
	UserGetPermissions(ctx context.Context, db bun.IDB, uid string) ([]AuthPermission, error)
}
