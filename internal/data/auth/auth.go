package auth

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/zx06/authz/internal/biz"
)

var _ biz.AuthRepo = (*authRepo)(nil)

type authRepo struct {
}

func NewAuthRepo() biz.AuthRepo {
	return &authRepo{}
}

// UserGetPermissions implements biz.AuthRepo
func (ar *authRepo) UserGetPermissions(ctx context.Context, db bun.IDB, uid string) ([]biz.AuthPermission, error) {
	var res []biz.AuthPermission
	apMap := make(map[int64]biz.AuthPermission)
	u := new(biz.AuthUser)
	err := db.NewSelect().
		Model(u).
		Relation("Groups").
		Relation("Permissions").
		Relation("Groups.Permissions").
		Where("uid = ?", uid).
		Limit(1).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Printf("\n%#v\n", u)
	for _, p := range u.Permissions {
		apMap[p.ID] = p
	}
	for _, g := range u.Groups {
		for _, p := range g.Permissions {
			apMap[p.ID] = p
		}
	}
	for _, p := range apMap {
		res = append(res, p)
	}
	return res, err

}

// UserHasPerm implements biz.AuthRepo
func (ar *authRepo) UserHasPerm(ctx context.Context, db bun.IDB, uid string, perm string, act string) (bool, error) {
	panic("unimplemented")
}
