package account

import (
	"fmt"

	"github.com/labstack/echo/v4"

	model "jank.com/jank_blog/internal/model/account"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/account/dto"
	"jank.com/jank_blog/pkg/serve/mapper"
	"jank.com/jank_blog/pkg/vo/account"
)

// CreateRole 创建角色
func CreateRole(req *dto.CreateRoleRequest, c echo.Context) (*account.RoleVo, error) {
	role := &model.Role{
		Code:        req.Code,
		Description: req.Description,
	}

	if err := mapper.CreateRole(role); err != nil {
		utils.BizLogger(c).Errorf("创建角色失败: %v", err)
		return nil, fmt.Errorf("创建角色失败: %v", err)
	}

	vo, err := utils.MapModelToVO(role, &account.RoleVo{})
	if err != nil {
		utils.BizLogger(c).Errorf("角色创建时映射 vo 失败: %v", err)
		return nil, fmt.Errorf("角色创建时映射 vo 失败: %v", err)
	}

	return vo.(*account.RoleVo), nil
}

// UpdateRole 更新角色
func UpdateRole(req *dto.UpdateRoleRequest, c echo.Context) (*account.RoleVo, error) {
	role, err := mapper.GetRoleByID(req.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取角色失败: %v", err)
		return nil, fmt.Errorf("获取角色失败: %v", err)
	}

	role.Code = req.Code
	role.Description = req.Description

	if err := mapper.UpdateRole(role); err != nil {
		utils.BizLogger(c).Errorf("更新角色失败: %v", err)
		return nil, fmt.Errorf("更新角色失败: %v", err)
	}

	vo, err := utils.MapModelToVO(role, &account.RoleVo{})
	if err != nil {
		utils.BizLogger(c).Errorf("角色更新时映射 vo 失败: %v", err)
		return nil, fmt.Errorf("角色更新时映射 vo 失败: %v", err)
	}

	return vo.(*account.RoleVo), nil
}

// DeleteRole 删除角色
func DeleteRole(req *dto.DeleteRoleRequest, c echo.Context) error {
	if err := mapper.DeleteRoleSoftly(req.ID); err != nil {
		utils.BizLogger(c).Errorf("删除角色失败: %v", err)
		return fmt.Errorf("删除角色失败: %v", err)
	}

	return nil
}

// ListRoles 获取所有角色
func ListRoles(c echo.Context) ([]*account.RoleVo, error) {
	roles, err := mapper.GetAllRoles()
	if err != nil {
		utils.BizLogger(c).Errorf("获取所有角色失败: %v", err)
		return nil, fmt.Errorf("获取所有角色失败: %v", err)
	}

	var roleVos []*account.RoleVo
	for _, role := range roles {
		roleVo, err := utils.MapModelToVO(role, &account.RoleVo{})
		if err != nil {
			utils.BizLogger(c).Errorf("获取所有角色时映射 vo 失败: %v", err)
			return nil, fmt.Errorf("获取所有时映射 vo 失败: %v", err)
		}
		roleVos = append(roleVos, roleVo.(*account.RoleVo))
	}

	return roleVos, nil
}
