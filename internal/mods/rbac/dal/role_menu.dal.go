package dal

import (
	"context"

	"github.com/LyricTian/gin-admin/v10/internal/mods/rbac/schema"
	"github.com/LyricTian/gin-admin/v10/internal/utils"
	"github.com/LyricTian/gin-admin/v10/pkg/errors"
	"gorm.io/gorm"
)

// Get role menu storage instance
func GetRoleMenuDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return utils.GetDB(ctx, defDB).Model(new(schema.RoleMenu))
}

// Role permissions for RBAC
type RoleMenu struct {
	DB *gorm.DB
}

// Query rolemenus from the database based on the provided parameters and options.
func (a *RoleMenu) Query(ctx context.Context, params schema.RoleMenuQueryParam, opts ...schema.RoleMenuQueryOptions) (*schema.RoleMenuQueryResult, error) {
	var opt schema.RoleMenuQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	db := GetRoleMenuDB(ctx, a.DB)
	if v := params.RoleID; len(v) > 0 {
		db = db.Where("role_id = ?", v)
	}
	if v := params.MenuID; len(v) > 0 {
		db = db.Where("menu_id = ?", v)
	}

	var list schema.RoleMenus
	pageResult, err := utils.WrapPageQuery(ctx, db, params.PaginationParam, opt.QueryOptions, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	queryResult := &schema.RoleMenuQueryResult{
		PageResult: pageResult,
		Data:       list,
	}
	return queryResult, nil
}

// Get the specified role menu from the database.
func (a *RoleMenu) Get(ctx context.Context, id string, opts ...schema.RoleMenuQueryOptions) (*schema.RoleMenu, error) {
	var opt schema.RoleMenuQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	item := new(schema.RoleMenu)
	ok, err := utils.FindOne(ctx, GetRoleMenuDB(ctx, a.DB).Where("id=?", id), opt.QueryOptions, item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item, nil
}

// Exist checks if the specified role menu exists in the database.
func (a *RoleMenu) Exists(ctx context.Context, id string) (bool, error) {
	ok, err := utils.Exists(ctx, GetRoleMenuDB(ctx, a.DB).Where("id=?", id))
	return ok, errors.WithStack(err)
}

// Create a new role menu.
func (a *RoleMenu) Create(ctx context.Context, item *schema.RoleMenu) error {
	result := GetRoleMenuDB(ctx, a.DB).Create(item)
	return errors.WithStack(result.Error)
}

// Update the specified role menu in the database.
func (a *RoleMenu) Update(ctx context.Context, item *schema.RoleMenu) error {
	result := GetRoleMenuDB(ctx, a.DB).Where("id=?", item.ID).Select("*").Omit("created_at").Updates(item)
	return errors.WithStack(result.Error)
}

// Delete the specified role menu from the database.
func (a *RoleMenu) Delete(ctx context.Context, id string) error {
	result := GetRoleMenuDB(ctx, a.DB).Where("id=?", id).Delete(new(schema.RoleMenu))
	return errors.WithStack(result.Error)
}
