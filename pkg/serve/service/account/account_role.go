package account

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/account/dto"
	"jank.com/jank_blog/pkg/serve/mapper"
	"jank.com/jank_blog/pkg/vo/account"
	"strconv"
)

// AssignRoleToAcc 为用户分配角色
func AssignRoleToAcc(req *dto.AssignRoleRequest, c echo.Context) error {
	if err := mapper.AssignRoleToAcc(req.AccountID, req.RoleID); err != nil {
		utils.BizLogger(c).Errorf("为用户分配角色失败: %v", err)
		return fmt.Errorf("为用户分配角色失败: %v", err)
	}
	return nil
}

// RemoveRoleFromAcc 移除用户角色
func RemoveRoleFromAcc(req *dto.AssignRoleRequest, c echo.Context) error {
	if err := mapper.DeleteRoleFromAccSoftly(req.AccountID, req.RoleID); err != nil {
		utils.BizLogger(c).Errorf("移除用户角色失败: %v", err)
		return fmt.Errorf("移除用户角色失败: %v", err)
	}

	return nil
}

// UpdateRoleForAcc 更新用户角色
func UpdateRoleForAcc(req *dto.AssignRoleRequest, c echo.Context) error {
	if err := mapper.UpdateRoleForAcc(req.AccountID, req.RoleID); err != nil {
		utils.BizLogger(c).Errorf("更新用户角色失败: %v", err)
		return fmt.Errorf("更新用户角色失败: %v", err)
	}

	return nil
}

// GetRolesByAcc 获取用户的所有角色
func GetRolesByAcc(req *dto.GetRolesByAccRequest, c echo.Context) ([]*account.RoleVo, error) {
	roles, err := mapper.GetRolesByAccountID(strconv.FormatInt(req.AccountID, 10))
	if err != nil {
		utils.BizLogger(c).Errorf("获取用户角色失败: %v", err)
		return nil, fmt.Errorf("获取用户角色失败: %v", err)
	}

	var roleVos []*account.RoleVo
	for _, role := range roles {
		roleVo, err := utils.MapModelToVO(role, &account.RoleVo{})
		if err != nil {
			utils.BizLogger(c).Errorf("角色更新时映射 vo 失败: %v", err)
			return nil, fmt.Errorf("角色更新时映射 vo 失败: %v", err)
		}
		roleVos = append(roleVos, roleVo.(*account.RoleVo))
	}

	return roleVos, nil
}
