package impl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/internal/utils/logger"
	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/serve/mapper"
	"github.com/Done-0/jank/pkg/serve/service"
	"github.com/Done-0/jank/pkg/vo"
)

// RBACServiceImpl RBAC服务实现
type RBACServiceImpl struct {
	rbacMapper mapper.RBACMapper
	userMapper mapper.UserMapper
}

// NewRBACService 创建RBAC服务实例
func NewRBACService(rbacMapper mapper.RBACMapper, userMapper mapper.UserMapper) service.RBACService {
	return &RBACServiceImpl{
		rbacMapper: rbacMapper,
		userMapper: userMapper,
	}
}

// CreatePermission 创建权限策略
func (s *RBACServiceImpl) CreatePermission(c *app.RequestContext, req *dto.CreatePermissionRequest) (*vo.PolicyOpResponse, error) {
	allRoles, err := s.rbacMapper.ListRoles(c)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get system roles: %v", err)
		return nil, fmt.Errorf("failed to get system roles: %w", err)
	}

	roleExists := false
	for _, rolePolicy := range allRoles {
		if rolePolicy.V0 == req.Role {
			roleExists = true
			break
		}
	}
	if !roleExists {
		logger.BizLogger(c).Warnf("role does not exist in system, skipping policy addition: role[%s]", req.Role)
		return nil, fmt.Errorf("role does not exist in system: %s", req.Role)
	}

	existingPolicies, err := s.rbacMapper.GetRolePermissions(c, req.Role)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get existing policies for role: role[%s]: %v", req.Role, err)
		return nil, fmt.Errorf("failed to check existing policies: %w", err)
	}

	for _, existingPolicy := range existingPolicies {
		if existingPolicy.V1 == req.Resource && existingPolicy.V2 == req.Action {
			logger.BizLogger(c).Warnf("policy already exists, skipping addition: role[%s], resource[%s], action[%s]", req.Role, req.Resource, req.Action)
			return nil, fmt.Errorf("policy already exists: role '%s', resource '%s', action '%s'", req.Role, req.Resource, req.Action)
		}
	}

	_, err = s.rbacMapper.CreatePermission(c, req.Name, req.Description, req.Role, req.Resource, req.Action)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to add policy: role[%s], resource[%s], action[%s]: %v", req.Role, req.Resource, req.Action, err)
		return nil, fmt.Errorf("failed to add policy: %w", err)
	}

	logger.BizLogger(c).Infof("successfully added policy: name[%s], role[%s], resource[%s], action[%s]", req.Name, req.Role, req.Resource, req.Action)
	return &vo.PolicyOpResponse{
		Success:     true,
		Name:        req.Name,
		Description: req.Description,
		Role:        req.Role,
		Resource:    req.Resource,
		Action:      req.Action,
		Message:     "Policy added successfully",
	}, nil
}

// DeletePermission 删除权限策略
func (s *RBACServiceImpl) DeletePermission(c *app.RequestContext, req *dto.DeletePermissionRequest) (*vo.PolicyOpResponse, error) {
	allRoles, err := s.rbacMapper.ListRoles(c)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get system roles: %v", err)
		return nil, fmt.Errorf("failed to get system roles: %w", err)
	}

	roleExists := false
	for _, rolePolicy := range allRoles {
		if rolePolicy.V0 == req.Role {
			roleExists = true
			break
		}
	}
	if !roleExists {
		logger.BizLogger(c).Errorf("role does not exist in system: role[%s]", req.Role)
		return nil, fmt.Errorf("role '%s' does not exist in system", req.Role)
	}

	existingPolicies, err := s.rbacMapper.GetRolePermissions(c, req.Role)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get existing policies for role: role[%s]: %v", req.Role, err)
		return nil, fmt.Errorf("failed to check existing policies: %w", err)
	}

	policyExists := false
	for _, existingPolicy := range existingPolicies {
		if existingPolicy.V1 == req.Resource && existingPolicy.V2 == req.Action {
			policyExists = true
			break
		}
	}
	if !policyExists {
		logger.BizLogger(c).Warnf("policy does not exist, skipping removal: role[%s], resource[%s], action[%s]", req.Role, req.Resource, req.Action)
		return nil, fmt.Errorf("policy does not exist: role '%s', resource '%s', action '%s'", req.Role, req.Resource, req.Action)
	}

	success, err := s.rbacMapper.DeletePermission(c, req.Role, req.Resource, req.Action)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to remove policy: role[%s], resource[%s], action[%s]: %v", req.Role, req.Resource, req.Action, err)
		return nil, fmt.Errorf("failed to remove policy: %w", err)
	}
	if !success {
		logger.BizLogger(c).Warnf("policy removal had no effect: role[%s], resource[%s], action[%s]", req.Role, req.Resource, req.Action)
		return nil, fmt.Errorf("policy removal failed: no matching policy found")
	}

	logger.BizLogger(c).Infof("successfully removed policy: role[%s], resource[%s], action[%s]", req.Role, req.Resource, req.Action)
	return &vo.PolicyOpResponse{
		Success:  true,
		Role:     req.Role,
		Resource: req.Resource,
		Action:   req.Action,
		Message:  "Policy removed successfully",
	}, nil
}

// ListPermissions 获取所有权限
func (s *RBACServiceImpl) ListPermissions(c *app.RequestContext) (*vo.PermissionListResponse, error) {
	policies, err := s.rbacMapper.ListPermissions(c)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to list permissions: %v", err)
		return nil, fmt.Errorf("failed to list permissions: %w", err)
	}

	var permissions []vo.PermissionResponse
	for _, policy := range policies {
		permissions = append(permissions, vo.PermissionResponse{
			Name:        policy.V3,
			Description: policy.V4,
			Resource:    policy.V1,
			Action:      policy.V2,
		})
	}

	logger.BizLogger(c).Infof("successfully listed permissions: count[%d]", len(permissions))
	return &vo.PermissionListResponse{
		List:  permissions,
		Total: int64(len(permissions)),
	}, nil
}

// AssignPermission 为角色分配权限
func (s *RBACServiceImpl) AssignPermission(c *app.RequestContext, req *dto.AssignPermissionRequest) (*vo.PolicyOpResponse, error) {
	roleExists, err := s.rbacMapper.RoleExists(c, req.Role)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to check role existence: role[%s]: %v", req.Role, err)
		return nil, fmt.Errorf("failed to check role existence: %w", err)
	}
	if !roleExists {
		logger.BizLogger(c).Errorf("role does not exist: role[%s]", req.Role)
		return nil, fmt.Errorf("role '%s' does not exist", req.Role)
	}

	permissionExists, err := s.rbacMapper.PermissionExists(c, req.Resource, req.Action)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to check permission existence: resource[%s], action[%s]: %v", req.Resource, req.Action, err)
		return nil, fmt.Errorf("failed to check permission existence: %w", err)
	}
	if !permissionExists {
		logger.BizLogger(c).Errorf("permission does not exist: resource[%s], action[%s]", req.Resource, req.Action)
		return nil, fmt.Errorf("permission does not exist: resource '%s', action '%s'", req.Resource, req.Action)
	}

	allPermissions, err := s.rbacMapper.ListPermissions(c)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get existing permissions: %v", err)
		return nil, fmt.Errorf("failed to get existing permissions: %w", err)
	}

	var permissionName, permissionDesc string
	for _, perm := range allPermissions {
		if perm.V1 == req.Resource && perm.V2 == req.Action {
			permissionName = perm.V3
			permissionDesc = perm.V4
			break
		}
	}

	if permissionName == "" {
		permissionName = fmt.Sprintf("%s:%s", req.Resource, req.Action)
		permissionDesc = fmt.Sprintf("Permission for %s on %s", req.Action, req.Resource)
		logger.BizLogger(c).Warnf("no existing permission definition found, using default: resource[%s], action[%s]", req.Resource, req.Action)
	}

	_, err = s.rbacMapper.CreatePermission(c, permissionName, permissionDesc, req.Role, req.Resource, req.Action)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to assign permission: role[%s], resource[%s], action[%s]: %v", req.Role, req.Resource, req.Action, err)
		return nil, fmt.Errorf("failed to assign permission: %w", err)
	}

	logger.BizLogger(c).Infof("successfully assigned permission: role[%s], name[%s], resource[%s], action[%s]", req.Role, permissionName, req.Resource, req.Action)
	return &vo.PolicyOpResponse{
		Success:     true,
		Name:        permissionName,
		Description: permissionDesc,
		Role:        req.Role,
		Resource:    req.Resource,
		Action:      req.Action,
		Message:     "Permission assigned successfully",
	}, nil
}

// RevokePermission 撤销角色权限
func (s *RBACServiceImpl) RevokePermission(c *app.RequestContext, req *dto.RevokePermissionRequest) (*vo.PolicyOpResponse, error) {
	roleExists, err := s.rbacMapper.RoleExists(c, req.Role)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to check role existence: role[%s]: %v", req.Role, err)
		return nil, fmt.Errorf("failed to check role existence: %w", err)
	}
	if !roleExists {
		logger.BizLogger(c).Errorf("role does not exist: role[%s]", req.Role)
		return nil, fmt.Errorf("role '%s' does not exist", req.Role)
	}

	success, err := s.rbacMapper.DeletePermission(c, req.Role, req.Resource, req.Action)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to revoke permission: role[%s], resource[%s], action[%s]: %v", req.Role, req.Resource, req.Action, err)
		return nil, fmt.Errorf("failed to revoke permission: %w", err)
	}
	if !success {
		logger.BizLogger(c).Warnf("permission revocation had no effect: role[%s], resource[%s], action[%s]", req.Role, req.Resource, req.Action)
		return nil, fmt.Errorf("permission revocation failed: no matching policy found")
	}

	logger.BizLogger(c).Infof("successfully revoked permission: role[%s], resource[%s], action[%s]", req.Role, req.Resource, req.Action)
	return &vo.PolicyOpResponse{
		Success:  true,
		Role:     req.Role,
		Resource: req.Resource,
		Action:   req.Action,
		Message:  "Permission revoked successfully",
	}, nil
}

// CreateRole 创建角色
func (s *RBACServiceImpl) CreateRole(c *app.RequestContext, req *dto.CreateRoleRequest) (*vo.CreateRoleResponse, error) {
	exists, err := s.rbacMapper.RoleExists(c, req.Role)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to check role existence: %v", err)
		return nil, fmt.Errorf("failed to check role existence: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("role already exists: %s", req.Role)
	}

	_, err = s.rbacMapper.CreatePermission(c, req.Name, req.Description, req.Role, req.Resource, req.Action)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to create role: role[%s]: %v", req.Role, err)
		return nil, fmt.Errorf("failed to create role: %w", err)
	}

	logger.BizLogger(c).Infof("successfully created role: name[%s], role[%s]", req.Name, req.Role)
	return &vo.CreateRoleResponse{
		Success:     true,
		Name:        req.Name,
		Description: req.Description,
		Role:        req.Role,
		Resource:    req.Resource,
		Action:      req.Action,
		Message:     "Role created successfully",
	}, nil
}

// DeleteRole 删除角色 - 主流RBAC最佳实践
func (s *RBACServiceImpl) DeleteRole(c *app.RequestContext, req *dto.DeleteRoleRequest) (*vo.DeleteRoleResponse, error) {
	exists, err := s.rbacMapper.RoleExists(c, req.Role)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to check role existence: %v", err)
		return nil, fmt.Errorf("failed to check role existence: %w", err)
	}
	if !exists {
		logger.BizLogger(c).Warnf("role does not exist: %s", req.Role)
		return nil, fmt.Errorf("role does not exist: %s", req.Role)
	}

	allUsers, err := s.rbacMapper.ListUsers(c)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get users for role cleanup: %v", err)
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	userRevokeCount := 0
	for _, userPolicy := range allUsers {
		hasRole, _ := s.rbacMapper.UserHasRole(c, userPolicy.V0, req.Role)
		if hasRole {
			if revoked, _ := s.rbacMapper.RevokeRole(c, userPolicy.V0, req.Role); revoked {
				userRevokeCount++
				logger.BizLogger(c).Infof("revoked role from user: user[%s], role[%s]", userPolicy.V0, req.Role)
			}
		}
	}

	rolePermissions, err := s.rbacMapper.GetRolePermissions(c, req.Role)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get role permissions: %v", err)
		return nil, fmt.Errorf("failed to get role permissions: %w", err)
	}

	permissionDeleteCount := 0
	for _, perm := range rolePermissions {
		if deleted, _ := s.rbacMapper.DeletePermission(c, perm.V0, perm.V1, perm.V2); deleted {
			permissionDeleteCount++
			logger.BizLogger(c).Infof("deleted permission: role[%s], resource[%s], action[%s]", perm.V0, perm.V1, perm.V2)
		}
	}

	if userRevokeCount == 0 && permissionDeleteCount == 0 {
		logger.BizLogger(c).Warnf("role deletion had no effect: role[%s]", req.Role)
		return nil, fmt.Errorf("role deletion failed: no associated data found")
	}

	return &vo.DeleteRoleResponse{
		Success: true,
		Role:    req.Role,
		Message: fmt.Sprintf("Role deleted successfully (users: %d, permissions: %d)", userRevokeCount, permissionDeleteCount),
	}, nil
}

// ListRoles 获取所有角色
func (s *RBACServiceImpl) ListRoles(c *app.RequestContext) (*vo.RoleListResponse, error) {
	policies, err := s.rbacMapper.ListRoles(c)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to list roles: %v", err)
		return nil, fmt.Errorf("failed to list roles: %w", err)
	}

	var roles []string
	for _, policy := range policies {
		roles = append(roles, policy.V0)
	}

	logger.BizLogger(c).Infof("successfully listed roles: count[%d]", len(roles))
	return &vo.RoleListResponse{
		List:  roles,
		Total: int64(len(roles)),
	}, nil
}

// GetRolePermissions 获取角色权限
func (s *RBACServiceImpl) GetRolePermissions(c *app.RequestContext, req *dto.GetRolePermissionsRequest) (*vo.RolePermissionsResponse, error) {
	exists, err := s.rbacMapper.RoleExists(c, req.Role)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to check role existence: %v", err)
		return nil, fmt.Errorf("failed to check role existence: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("role does not exist: %s", req.Role)
	}

	policies, err := s.rbacMapper.GetRolePermissions(c, req.Role)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get role permissions: role[%s]: %v", req.Role, err)
		return nil, fmt.Errorf("failed to get role permissions: %w", err)
	}

	var permissions []vo.PermissionResponse
	for _, policy := range policies {
		permissions = append(permissions, vo.PermissionResponse{
			Resource:    policy.V1, // 资源路径存储在 V1 字段
			Action:      policy.V2, // 操作方法存储在 V2 字段
			Name:        policy.V3, // 权限名称存储在 V3 字段
			Description: policy.V4, // 权限描述存储在 V4 字段
		})
	}

	logger.BizLogger(c).Infof("successfully got role permissions: role[%s], count[%d]", req.Role, len(permissions))
	return &vo.RolePermissionsResponse{
		Role:        req.Role,
		Permissions: permissions,
		Total:       int64(len(permissions)),
	}, nil
}

// AssignRole 为用户分配角色
func (s *RBACServiceImpl) AssignRole(c *app.RequestContext, req *dto.AssignRoleRequest) (*vo.RoleOpResponse, error) {
	userID, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		logger.BizLogger(c).Errorf("invalid user ID format: %s", req.UserID)
		return nil, fmt.Errorf("invalid user ID format: %s", req.UserID)
	}

	user, err := s.userMapper.GetUserByID(c, userID)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get user: userID[%s]: %v", req.UserID, err)
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		logger.BizLogger(c).Errorf("user does not exist: userID[%s]", req.UserID)
		return nil, fmt.Errorf("user does not exist: %s", req.UserID)
	}

	roleExists, err := s.rbacMapper.RoleExists(c, req.Role)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to check role existence: role[%s]: %v", req.Role, err)
		return nil, fmt.Errorf("failed to check role existence: %w", err)
	}
	if !roleExists {
		logger.BizLogger(c).Errorf("role does not exist: role[%s]", req.Role)
		return nil, fmt.Errorf("role does not exist: %s", req.Role)
	}

	hasRole, err := s.rbacMapper.UserHasRole(c, req.UserID, req.Role)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to check user role: userID[%s], role[%s]: %v", req.UserID, req.Role, err)
		return nil, fmt.Errorf("failed to check user role: %w", err)
	}
	if hasRole {
		logger.BizLogger(c).Warnf("user already has role: userID[%s], role[%s]", req.UserID, req.Role)
		return nil, fmt.Errorf("user already has role: %s", req.Role)
	}

	_, err = s.rbacMapper.AssignRole(c, req.UserID, req.Role)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to assign role: userID[%s], role[%s]: %v", req.UserID, req.Role, err)
		return nil, fmt.Errorf("failed to assign role: %w", err)
	}

	logger.BizLogger(c).Infof("successfully assigned role: userID[%s], role[%s]", req.UserID, req.Role)
	return &vo.RoleOpResponse{
		Success: true,
		UserID:  req.UserID,
		Role:    req.Role,
		Message: "Role assigned successfully",
	}, nil
}

// RevokeRole 撤销用户角色
func (s *RBACServiceImpl) RevokeRole(c *app.RequestContext, req *dto.RevokeRoleRequest) (*vo.RoleOpResponse, error) {
	userID, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		logger.BizLogger(c).Errorf("invalid user ID format: %s", req.UserID)
		return nil, fmt.Errorf("invalid user ID format: %s", req.UserID)
	}

	user, err := s.userMapper.GetUserByID(c, userID)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get user: userID[%s]: %v", req.UserID, err)
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		logger.BizLogger(c).Errorf("user does not exist: userID[%s]", req.UserID)
		return nil, fmt.Errorf("user does not exist: %s", req.UserID)
	}

	hasRole, err := s.rbacMapper.UserHasRole(c, req.UserID, req.Role)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to check user role: userID[%s], role[%s]: %v", req.UserID, req.Role, err)
		return nil, fmt.Errorf("failed to check user role: %w", err)
	}
	if !hasRole {
		logger.BizLogger(c).Warnf("user does not have role: userID[%s], role[%s]", req.UserID, req.Role)
		return nil, fmt.Errorf("user does not have role: %s", req.Role)
	}

	success, err := s.rbacMapper.RevokeRole(c, req.UserID, req.Role)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to revoke role: userID[%s], role[%s]: %v", req.UserID, req.Role, err)
		return nil, fmt.Errorf("failed to revoke role: %w", err)
	}
	if !success {
		logger.BizLogger(c).Warnf("role revocation had no effect: userID[%s], role[%s]", req.UserID, req.Role)
		return nil, fmt.Errorf("role revocation failed: no matching role found")
	}

	logger.BizLogger(c).Infof("successfully revoked role: userID[%s], role[%s]", req.UserID, req.Role)
	return &vo.RoleOpResponse{
		Success: true,
		UserID:  req.UserID,
		Role:    req.Role,
		Message: "Role revoked successfully",
	}, nil
}

// GetUserRoles 获取用户角色
func (s *RBACServiceImpl) GetUserRoles(c *app.RequestContext, req *dto.GetUserRolesRequest) (*vo.UserRolesResponse, error) {
	userID, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		logger.BizLogger(c).Errorf("invalid user ID format: %s", req.UserID)
		return nil, fmt.Errorf("invalid user ID format: %s", req.UserID)
	}

	user, err := s.userMapper.GetUserByID(c, userID)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get user: userID[%s]: %v", req.UserID, err)
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		logger.BizLogger(c).Errorf("user does not exist: userID[%s]", req.UserID)
		return nil, fmt.Errorf("user does not exist: %s", req.UserID)
	}

	policies, err := s.rbacMapper.GetUserRoles(c, req.UserID)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get user roles: userID[%s]: %v", req.UserID, err)
		return nil, fmt.Errorf("failed to get user roles: %w", err)
	}

	var roles []string
	for _, policy := range policies {
		roles = append(roles, policy.V1)
	}

	logger.BizLogger(c).Infof("successfully got user roles: userID[%s], count[%d]", req.UserID, len(roles))
	return &vo.UserRolesResponse{
		UserID: req.UserID,
		Roles:  roles,
	}, nil
}

// CheckPermission 权限检查
func (s *RBACServiceImpl) CheckPermission(c *app.RequestContext, req *dto.CheckPermissionRequest) (*vo.CheckResponse, error) {
	userID, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		logger.BizLogger(c).Errorf("invalid user ID format: %s", req.UserID)
		return nil, fmt.Errorf("invalid user ID format: %s", req.UserID)
	}

	user, err := s.userMapper.GetUserByID(c, userID)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get user: userID[%s]: %v", req.UserID, err)
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		logger.BizLogger(c).Errorf("user does not exist: userID[%s]", req.UserID)
		return nil, fmt.Errorf("user does not exist: %s", req.UserID)
	}

	allowed, err := s.rbacMapper.CheckPermission(c, req.UserID, req.Resource, req.Action)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to check permission: userID[%s], resource[%s], action[%s]: %v", req.UserID, req.Resource, req.Action, err)
		return nil, fmt.Errorf("failed to check permission: %w", err)
	}

	var reason string
	if allowed {
		reason = fmt.Sprintf("用户 %s 有权限访问资源 %s 进行 %s 操作", req.UserID, req.Resource, req.Action)
	} else {
		userRoles, _ := s.rbacMapper.GetUserRoles(c, req.UserID)
		if len(userRoles) == 0 {
			reason = fmt.Sprintf("用户 %s 没有分配任何角色，无法访问资源 %s 进行 %s 操作", req.UserID, req.Resource, req.Action)
		} else {
			roleNames := make([]string, len(userRoles))
			for i, role := range userRoles {
				roleNames[i] = role.V1
			}
			reason = fmt.Sprintf("用户 %s 的角色 [%s] 没有权限访问资源 %s 进行 %s 操作", req.UserID, strings.Join(roleNames, ", "), req.Resource, req.Action)
		}
	}

	logger.BizLogger(c).Infof("permission check result: userID[%s], resource[%s], action[%s], allowed[%t], reason[%s]", req.UserID, req.Resource, req.Action, allowed, reason)
	return &vo.CheckResponse{
		Allowed: allowed,
		Reason:  reason,
	}, nil
}
