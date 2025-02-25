package account

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"

	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/account/dto"
	"jank.com/jank_blog/pkg/serve/mapper"
	"jank.com/jank_blog/pkg/vo/account"
)

// AssignPermissionToRole 为角色分配权限
func AssignPermissionToRole(req *dto.AssignPermissionRequest, c echo.Context) error {
	if err := mapper.AssignPermissionToRole(req.RoleID, req.PermissionID); err != nil {
		utils.BizLogger(c).Errorf("为角色分配权限失败: %v", err)
		return fmt.Errorf("为角色分配权限失败: %v", err)
	}

	return nil
}

// RemovePermissionFromRole 移除角色权限
func RemovePermissionFromRole(req *dto.AssignPermissionRequest, c echo.Context) error {
	if err := mapper.DeletePermissionFromRoleSoftly(req.RoleID, req.PermissionID); err != nil {
		utils.BizLogger(c).Errorf("移除角色权限失败: %v", err)
		return fmt.Errorf("移除角色权限失败: %v", err)
	}

	return nil
}

// UpdatePermissionForRole 更新角色权限
func UpdatePermissionForRole(req *dto.AssignPermissionRequest, c echo.Context) error {
	if err := mapper.UpdatePermissionForRole(req.RoleID, req.PermissionID); err != nil {
		utils.BizLogger(c).Errorf("更新角色权限失败: %v", err)
		return fmt.Errorf("更新角色权限失败: %v", err)
	}

	return nil
}

// GetPermissionsByRole 获取角色的所有权限
func GetPermissionsByRole(req *dto.GetPermissionsByRoleRequest, c echo.Context) ([]*account.PermissionVo, error) {
	permissions, err := mapper.GetPermissionsByRoleID(strconv.FormatInt(req.RoleID, 10))
	if err != nil {
		utils.BizLogger(c).Errorf("获取角色权限失败: %v", err)
		return nil, fmt.Errorf("获取角色权限失败: %v", err)
	}

	var permissionVos []*account.PermissionVo
	for _, permission := range permissions {
		permissionVo, err := utils.MapModelToVO(permission, &account.PermissionVo{})
		if err != nil {
			utils.BizLogger(c).Errorf("权限更新时映射 vo 失败: %v", err)
			return nil, fmt.Errorf("权限更新时映射 vo 失败: %v", err)
		}
		permissionVos = append(permissionVos, permissionVo.(*account.PermissionVo))
	}

	return permissionVos, nil
}
