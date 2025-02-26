package service

import (
	"fmt"

	"github.com/labstack/echo/v4"

	model "jank.com/jank_blog/internal/model/account"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/account/dto"
	"jank.com/jank_blog/pkg/serve/mapper"
	"jank.com/jank_blog/pkg/vo/account"
)

// CreatePermission 创建权限
func CreatePermission(req *dto.CreatePermissionRequest, c echo.Context) (*account.PermissionVo, error) {
	permission := &model.Permission{
		Code:        req.Code,
		Description: req.Description,
	}

	if err := mapper.CreatePermission(permission); err != nil {
		utils.BizLogger(c).Errorf("创建权限失败: %v", err)
		return nil, fmt.Errorf("创建权限失败: %v", err)
	}

	vo, err := utils.MapModelToVO(permission, &account.PermissionVo{})
	if err != nil {
		utils.BizLogger(c).Errorf("权限创建时映射 vo 失败: %v", err)
		return nil, fmt.Errorf("权限创建时映射 vo 失败: %v", err)
	}

	return vo.(*account.PermissionVo), nil
}

// UpdatePermission 更新权限
func UpdatePermission(req *dto.UpdatePermissionRequest, c echo.Context) (*account.PermissionVo, error) {
	permission, err := mapper.GetPermissionByID(req.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取权限失败: %v", err)
		return nil, fmt.Errorf("获取权限失败: %v", err)
	}

	permission.Code = req.Code
	permission.Description = req.Description

	if err := mapper.UpdatePermission(permission); err != nil {
		utils.BizLogger(c).Errorf("更新权限失败: %v", err)
		return nil, fmt.Errorf("更新权限失败: %v", err)
	}

	vo, err := utils.MapModelToVO(permission, &account.PermissionVo{})
	if err != nil {
		utils.BizLogger(c).Errorf("权限更新时映射 vo 失败: %v", err)
		return nil, fmt.Errorf("权限更新时映射 vo 失败: %v", err)
	}

	return vo.(*account.PermissionVo), nil
}

// DeletePermission 删除权限
func DeletePermission(req *dto.DeletePermissionRequest, c echo.Context) error {
	if err := mapper.DeletePermissionSoftly(req.ID); err != nil {
		utils.BizLogger(c).Errorf("删除权限失败: %v", err)
		return fmt.Errorf("删除权限失败: %v", err)
	}

	return nil
}

// ListPermissions 获取所有权限
func ListPermissions(c echo.Context) ([]*account.PermissionVo, error) {
	permissions, err := mapper.GetAllPermissions()
	if err != nil {
		utils.BizLogger(c).Errorf("获取所有权限失败: %v", err)
		return nil, fmt.Errorf("获取所有权限失败: %v", err)
	}

	var permissionVos []*account.PermissionVo
	for _, permission := range permissions {
		permissionVo, err := utils.MapModelToVO(permission, &account.PermissionVo{})
		if err != nil {
			utils.BizLogger(c).Errorf("获取所有权限时映射 vo 失败: %v", err)
			return nil, fmt.Errorf("获取所有权限时映射 vo 失败: %v", err)
		}
		permissionVos = append(permissionVos, permissionVo.(*account.PermissionVo))
	}

	return permissionVos, nil
}
