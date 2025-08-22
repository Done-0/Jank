/**
 * RBAC 权限管理主界面组件
 * 统一的角色和权限管理，遵循 shadcn UI 设计规范和 Twitter 极简风格
 */

import { useState, useMemo } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Badge } from "@/components/ui/badge";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter,
} from "@/components/ui/dialog";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import {
  Search,
  Plus,
  MoreHorizontal,
  Eye,
  Trash2,
  Settings,
  Shield,
  UserPlus,
  Crown,
} from "lucide-react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { toast } from "sonner";

import {
  useRoles,
  usePermissions,
  useRolePermissions,
  useCreateRole,
  useDeleteRole,
  useCreatePermission,
  useAssignPermission,
} from "@/hooks/use-rbac";
import { useUsers } from "@/hooks/use-user";
import { RESOURCE_GROUPS, RBAC_ACTION } from "@/constants/rbac";

// ===== 类型定义 =====

interface RBACMainContentProps {
  selectedSection: string;
  searchTerm: string;
  onSearchChange: (term: string) => void;
}

// ===== 主组件 =====

export function RBACMainContent({
  selectedSection,
  searchTerm,
  onSearchChange,
}: RBACMainContentProps) {
  // ===== 对话框状态 =====
  const [isCreateRoleOpen, setIsCreateRoleOpen] = useState(false);
  const [isCreatePermissionOpen, setIsCreatePermissionOpen] = useState(false);
  const [isAssignRoleOpen, setIsAssignRoleOpen] = useState(false);
  const [isAssignPermissionOpen, setIsAssignPermissionOpen] = useState(false);
  const [isViewRolePermissionsOpen, setIsViewRolePermissionsOpen] =
    useState(false);
  const [selectedRole, setSelectedRole] = useState<string>("");

  // ===== 表单状态 =====
  const [newRoleName, setNewRoleName] = useState("");
  const [newRoleDescription, setNewRoleDescription] = useState("");
  const [newRoleResource, setNewRoleResource] = useState("");
  const [newRoleAction, setNewRoleAction] = useState("");
  const [newPermissionName, setNewPermissionName] = useState("");
  const [newPermissionDescription, setNewPermissionDescription] = useState("");
  const [newPermissionResource, setNewPermissionResource] = useState("");
  const [newPermissionAction, setNewPermissionAction] = useState("");
  const [newPermissionRole, setNewPermissionRole] = useState("");
  const [assignUserId, setAssignUserId] = useState("");
  const [assignRoleId, setAssignRoleId] = useState("");
  const [assignPermissionRole, setAssignPermissionRole] = useState("");
  const [assignPermissionResource, setAssignPermissionResource] = useState("");
  const [assignPermissionAction, setAssignPermissionAction] = useState("");

  // ===== 数据查询 Hooks =====
  const { data: rolesResponse, isLoading: rolesLoading } = useRoles();
  const roles = useMemo(() => rolesResponse?.list || [], [rolesResponse]);

  const { data: permissionsResponse, isLoading: permissionsLoading } =
    usePermissions();
  const permissions = useMemo(
    () => permissionsResponse?.list || [],
    [permissionsResponse]
  );

  // 角色权限查询
  const { data: rolePermissionsData, isLoading: rolePermissionsLoading } =
    useRolePermissions(selectedRole, {
      enabled: !!selectedRole,
    });

  const rolePermissions = useMemo(
    () => (rolePermissionsData as any)?.permissions || [],
    [rolePermissionsData]
  );
  const rolePermissionsTotal = useMemo(
    () => (rolePermissionsData as any)?.total || 0,
    [rolePermissionsData]
  );

  // ===== 操作 Hooks =====
  const createRoleMutation = useCreateRole();
  const deleteRoleMutation = useDeleteRole();
  const createPermissionMutation = useCreatePermission();
  const assignPermissionMutation = useAssignPermission();

  // ===== 工具函数 =====

  // 根据资源路径获取资源信息
  const getResourceInfo = (resourcePath: string) => {
    for (const group of RESOURCE_GROUPS) {
      const resource = group.resources.find((r) => r.value === resourcePath);
      if (resource) {
        return {
          name: resource.name,
          description: resource.description,
          groupName: group.name,
          groupIcon: group.icon,
        };
      }
    }
    return {
      name: resourcePath,
      description: "系统资源",
      groupName: "其他",
      groupIcon: Shield,
    };
  };

  // ===== 事件处理函数 =====

  // 查看角色权限
  const handleViewRolePermissions = (role: string) => {
    setSelectedRole(role);
    setIsViewRolePermissionsOpen(true);
  };

  // 删除角色
  const handleDeleteRole = (role: string) => {
    if (window.confirm(`确定要删除角色 "${role}" 吗？`)) {
      deleteRoleMutation.mutate({ role });
    }
  };

  // 权限信息提示
  const handlePermissionInfo = () => {
    toast.info('请在"角色权限管理"页面中编辑权限，以确保操作准确性');
  };

  // 删除权限提示
  const handleDeletePermission = () => {
    toast.error('请在"角色权限管理"页面中删除权限，以确保操作准确性');
  };

  // 创建角色
  const handleCreateRole = () => {
    if (
      !newRoleName.trim() ||
      !newRoleResource.trim() ||
      !newRoleAction.trim()
    ) {
      toast.error("请填写完整的角色信息");
      return;
    }

    createRoleMutation.mutate({
      name: newRoleName,
      description: newRoleDescription || "",
      role: newRoleName,
      resource: newRoleResource,
      action: newRoleAction as any,
    });

    if (!createRoleMutation.isPending) {
      setIsCreateRoleOpen(false);
      setNewRoleName("");
      setNewRoleDescription("");
      setNewRoleResource("");
      setNewRoleAction("");
    }
  };

  // 创建权限
  const handleCreatePermission = () => {
    if (
      !newPermissionRole.trim() ||
      !newPermissionResource.trim() ||
      !newPermissionAction.trim()
    ) {
      toast.error("请填写完整的权限信息，包括角色");
      return;
    }

    createPermissionMutation.mutate({
      name: newPermissionName,
      description: newPermissionDescription,
      role: newPermissionRole,
      resource: newPermissionResource,
      action: newPermissionAction as any,
    });

    // 成功后清理表单
    if (!createPermissionMutation.isPending) {
      setIsCreatePermissionOpen(false);
      setNewPermissionName("");
      setNewPermissionDescription("");
      setNewPermissionResource("");
      setNewPermissionAction("");
      setNewPermissionRole("");
    }
  };

  // 分配角色（对话框中的）
  const handleAssignRoleDialog = () => {
    if (!assignUserId.trim() || !assignRoleId.trim()) {
      toast.error("请填写完整的分配信息");
      return;
    }

    toast.info("角色分配功能开发中...");

    // 清理表单
    setIsAssignRoleOpen(false);
    setAssignUserId("");
    setAssignRoleId("");
  };

  // 分配权限
  const handleAssignPermission = () => {
    if (
      !assignPermissionRole.trim() ||
      !assignPermissionResource.trim() ||
      !assignPermissionAction.trim()
    ) {
      toast.error("请填写完整的权限分配信息");
      return;
    }

    assignPermissionMutation.mutate({
      role: assignPermissionRole,
      resource: assignPermissionResource,
      action: assignPermissionAction as any,
    });
  };

  // ===== 数据过滤 =====
  const filteredRoles = useMemo(() => {
    if (!searchTerm) return roles;
    return roles.filter((role: string) =>
      role.toLowerCase().includes(searchTerm.toLowerCase())
    );
  }, [roles, searchTerm]);

  const filteredPermissions = useMemo(() => {
    if (!searchTerm) return permissions;
    return permissions.filter(
      (permission: any) =>
        permission.resource?.toLowerCase().includes(searchTerm.toLowerCase()) ||
        permission.action?.toLowerCase().includes(searchTerm.toLowerCase())
    );
  }, [permissions, searchTerm]);

  // ===== 渲染函数 =====

  // 渲染角色列表
  const renderRolesList = () => (
    <div className="h-full flex flex-col">
      {/* 顶部操作栏 */}
      <div className="px-4 py-4 border-b">
        <div className="flex items-center justify-between gap-4">
          <div className="relative flex-1 sm:flex-initial sm:w-80">
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
            <Input
              placeholder="搜索角色..."
              className="pl-9 h-10 rounded-full"
              value={searchTerm}
              onChange={(e) => onSearchChange(e.target.value)}
            />
          </div>
          <Button
            onClick={() => setIsCreateRoleOpen(true)}
            className="h-10 px-4 shrink-0 rounded-full"
          >
            <Plus className="mr-2 h-4 w-4" />
            <span className="hidden sm:inline">创建角色</span>
            <span className="sm:hidden">创建</span>
          </Button>
        </div>
      </div>

      {/* 角色列表区域 */}
      <div className="flex-1 flex flex-col overflow-hidden w-full">
        {rolesLoading ? (
          <div className="flex items-center justify-center h-full">
            <div className="text-center">
              <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary mx-auto mb-4"></div>
              <p className="text-sm text-muted-foreground">加载中...</p>
            </div>
          </div>
        ) : filteredRoles.length === 0 ? (
          <div className="flex-1 overflow-auto p-4">
            <div className="text-center py-12 border-2 border-dashed border-muted-foreground/30 rounded-lg">
              <Crown className="h-12 w-12 text-muted-foreground mx-auto mb-4" />
              <h3 className="text-lg font-medium mb-2">暂无角色</h3>
              <p className="text-muted-foreground mb-4">暂无任何角色数据</p>
              <p className="text-sm text-muted-foreground">
                <span className="hidden sm:inline">
                  点击上方&ldquo;创建角色&rdquo;按钮开始创建
                </span>
                <span className="sm:hidden">点击上方按钮开始创建</span>
              </p>
            </div>
          </div>
        ) : (
          <div className="flex-1 overflow-y-scroll scrollbar-hidden">
            <div className="divide-y divide-border">
              {filteredRoles.map((role: string, index: number) => (
                <div
                  key={index}
                  className="px-4 py-4 hover:bg-accent/50 transition-colors"
                >
                  <div className="flex flex-col gap-3">
                    <div className="flex items-start justify-between gap-3">
                      <h3 className="font-medium text-lg line-clamp-2 flex-1 min-w-0">
                        {role}
                      </h3>
                      <Badge variant="secondary" className="shrink-0">
                        角色
                      </Badge>
                    </div>

                    <div className="flex items-center justify-end gap-3">
                      <DropdownMenu>
                        <DropdownMenuTrigger asChild>
                          <Button
                            variant="ghost"
                            size="sm"
                            className="h-8 w-8 p-0 rounded-full"
                          >
                            <MoreHorizontal className="h-4 w-4" />
                          </Button>
                        </DropdownMenuTrigger>
                        <DropdownMenuContent align="end" className="w-40">
                          <DropdownMenuItem
                            onClick={() => handleViewRolePermissions(role)}
                            className="cursor-pointer"
                          >
                            <Eye className="mr-2 h-4 w-4" />
                            查看权限
                          </DropdownMenuItem>
                          <DropdownMenuSeparator />
                          <DropdownMenuItem
                            onClick={() => handleDeleteRole(role)}
                            className="cursor-pointer text-destructive focus:text-destructive"
                          >
                            <Trash2 className="mr-2 h-4 w-4" />
                            删除角色
                          </DropdownMenuItem>
                        </DropdownMenuContent>
                      </DropdownMenu>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        )}
      </div>
    </div>
  );

  // 渲染权限列表
  const renderPermissionsList = () => (
    <div className="h-full flex flex-col">
      {/* 顶部操作栏 */}
      <div className="px-4 py-4 border-b">
        <div className="flex items-center justify-between gap-4">
          <div className="relative flex-1 sm:flex-initial sm:w-80">
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
            <Input
              placeholder="搜索权限..."
              className="pl-9 h-10 rounded-full"
              value={searchTerm}
              onChange={(e) => onSearchChange(e.target.value)}
            />
          </div>
          <Button
            onClick={() => setIsCreatePermissionOpen(true)}
            className="h-10 px-4 shrink-0 rounded-full"
          >
            <Plus className="mr-2 h-4 w-4" />
            <span className="hidden sm:inline">创建权限</span>
            <span className="sm:hidden">创建</span>
          </Button>
        </div>
      </div>

      {/* 权限列表区域 */}
      <div className="flex-1 flex flex-col overflow-hidden w-full">
        {permissionsLoading ? (
          <div className="flex items-center justify-center h-full">
            <div className="text-center">
              <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary mx-auto mb-4"></div>
              <p className="text-sm text-muted-foreground">加载中...</p>
            </div>
          </div>
        ) : filteredPermissions.length === 0 ? (
          <div className="flex-1 overflow-auto p-4">
            <div className="text-center py-12 border-2 border-dashed border-muted-foreground/30 rounded-lg">
              <Shield className="h-12 w-12 text-muted-foreground mx-auto mb-4" />
              <h3 className="text-lg font-medium mb-2">暂无权限</h3>
              <p className="text-muted-foreground mb-4">暂无任何权限数据</p>
              <p className="text-sm text-muted-foreground">
                点击上方&ldquo;创建权限&rdquo;按钮开始创建
              </p>
            </div>
          </div>
        ) : (
          <div className="flex-1 overflow-y-scroll scrollbar-hidden">
            <div className="divide-y divide-border">
              {filteredPermissions.map((permission: any, index: number) => {
                const resourceInfo = getResourceInfo(permission.resource);

                return (
                  <div
                    key={index}
                    className="px-4 py-4 hover:bg-accent/50 transition-colors"
                  >
                    <div className="flex flex-col gap-3">
                      <div className="flex items-start justify-between gap-3">
                        <h3 className="font-medium text-lg line-clamp-2 flex-1 min-w-0">
                          {permission.name ||
                            resourceInfo?.name ||
                            permission.resource}
                        </h3>
                        <Badge variant="secondary" className="shrink-0">
                          {permission.action}
                        </Badge>
                      </div>

                      <div className="min-h-[1.25rem]">
                        <p className="text-sm text-muted-foreground line-clamp-2">
                          {permission.description ||
                            resourceInfo?.description ||
                            permission.resource}
                        </p>
                      </div>

                      <div className="flex items-center justify-between gap-3">
                        <div className="flex items-center gap-2 text-xs text-muted-foreground">
                          <div className="flex items-center gap-1.5">
                            <div className="w-1.5 h-1.5 rounded-full bg-primary" />
                            <span>{resourceInfo?.groupName || "权限"}</span>
                          </div>
                        </div>

                        <DropdownMenu>
                          <DropdownMenuTrigger asChild>
                            <Button
                              variant="ghost"
                              size="sm"
                              className="h-8 w-8 p-0 rounded-full"
                            >
                              <MoreHorizontal className="h-4 w-4" />
                            </Button>
                          </DropdownMenuTrigger>
                          <DropdownMenuContent align="end" className="w-40">
                            <DropdownMenuItem
                              onClick={handlePermissionInfo}
                              className="cursor-pointer"
                            >
                              <Eye className="mr-2 h-4 w-4" />
                              查看详情
                            </DropdownMenuItem>
                            <DropdownMenuSeparator />
                            <DropdownMenuItem
                              onClick={handleDeletePermission}
                              className="cursor-pointer text-destructive focus:text-destructive"
                            >
                              <Trash2 className="mr-2 h-4 w-4" />
                              删除权限
                            </DropdownMenuItem>
                          </DropdownMenuContent>
                        </DropdownMenu>
                      </div>
                    </div>
                  </div>
                );
              })}
            </div>
          </div>
        )}
      </div>
    </div>
  );

  // 渲染角色权限管理
  const renderRolePermissions = () => (
    <div className="h-full flex flex-col">
      {/* 顶部操作栏 */}
      <div className="px-4 py-4 border-b">
        <div className="flex items-center justify-between gap-4">
          <div className="relative flex-1 sm:flex-initial sm:w-80">
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
            <Input
              placeholder="搜索角色..."
              className="pl-9 h-10 rounded-full"
              value={searchTerm}
              onChange={(e) => onSearchChange(e.target.value)}
            />
          </div>
          <div className="flex items-center gap-2">
            <Badge variant="outline" className="text-xs">
              {roles.length} 个角色
            </Badge>
          </div>
        </div>
      </div>

      {/* 角色权限列表区域 */}
      <div className="flex-1 flex flex-col overflow-hidden w-full">
        {rolesLoading ? (
          <div className="flex items-center justify-center h-full">
            <div className="text-center">
              <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary mx-auto mb-4"></div>
              <p className="text-sm text-muted-foreground">加载中...</p>
            </div>
          </div>
        ) : filteredRoles.length === 0 ? (
          <div className="flex-1 overflow-auto p-4">
            <div className="text-center py-12 border-2 border-dashed border-muted-foreground/30 rounded-lg">
              <Settings className="h-12 w-12 text-muted-foreground mx-auto mb-4" />
              <h3 className="text-lg font-medium mb-2">暂无角色</h3>
              <p className="text-muted-foreground mb-4">暂无任何角色数据</p>
              <p className="text-sm text-muted-foreground">
                <span className="hidden sm:inline">
                  请先在角色管理中创建角色
                </span>
                <span className="sm:hidden">请先创建角色</span>
              </p>
            </div>
          </div>
        ) : (
          <div className="flex-1 overflow-y-scroll scrollbar-hidden">
            <div className="divide-y divide-border">
              {filteredRoles.map((role: string, index: number) => (
                <div
                  key={index}
                  className="px-4 py-4 hover:bg-accent/50 transition-colors"
                >
                  <div className="flex flex-col gap-3">
                    <div className="flex items-start justify-between gap-3">
                      <h3 className="font-medium text-lg line-clamp-2 flex-1 min-w-0">
                        {role}
                      </h3>
                      <Badge variant="secondary" className="shrink-0">
                        角色
                      </Badge>
                    </div>

                    <div className="flex items-center justify-end gap-3">
                      <DropdownMenu>
                        <DropdownMenuTrigger asChild>
                          <Button
                            variant="ghost"
                            size="sm"
                            className="h-8 w-8 p-0 rounded-full"
                          >
                            <MoreHorizontal className="h-4 w-4" />
                          </Button>
                        </DropdownMenuTrigger>
                        <DropdownMenuContent align="end" className="w-40">
                          <DropdownMenuItem
                            onClick={() => handleViewRolePermissions(role)}
                            className="cursor-pointer"
                          >
                            <Settings className="mr-2 h-4 w-4" />
                            配置权限
                          </DropdownMenuItem>
                        </DropdownMenuContent>
                      </DropdownMenu>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        )}
      </div>
    </div>
  );

  // 获取用户列表数据
  const { data: usersData, isLoading: usersLoading } = useUsers({
    page_no: 1,
    page_size: 100,
    keyword: searchTerm || undefined,
  });

  // 用户列表（使用useMemo优化）
  const users = useMemo(() => usersData?.list || [], [usersData?.list]);

  // 过滤用户列表
  const filteredUsers = useMemo(() => {
    if (!searchTerm) return users;
    return users.filter(
      (user) =>
        user.nickname.toLowerCase().includes(searchTerm.toLowerCase()) ||
        user.email.toLowerCase().includes(searchTerm.toLowerCase()) ||
        user.roles.some((role) =>
          role.toLowerCase().includes(searchTerm.toLowerCase())
        )
    );
  }, [users, searchTerm]);

  // 渲染用户角色管理
  const renderUserRoles = () => (
    <div className="h-full flex flex-col">
      {/* 顶部操作栏 */}
      <div className="px-4 py-4 border-b">
        <div className="flex items-center justify-between gap-4">
          <div className="relative flex-1 sm:flex-initial sm:w-80">
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
            <Input
              placeholder="搜索用户..."
              className="pl-9 h-10 rounded-full"
              value={searchTerm}
              onChange={(e) => onSearchChange(e.target.value)}
            />
          </div>
          <div className="flex items-center gap-2">
            <Badge variant="outline" className="text-xs">
              {filteredUsers.length} 个用户
            </Badge>
          </div>
        </div>
      </div>

      {/* 用户列表区域 */}
      <div className="flex-1 flex flex-col overflow-hidden w-full">
        {usersLoading ? (
          <div className="flex items-center justify-center h-full">
            <div className="text-center">
              <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary mx-auto mb-4"></div>
              <p className="text-sm text-muted-foreground">加载中...</p>
            </div>
          </div>
        ) : filteredUsers.length === 0 ? (
          <div className="flex-1 overflow-auto p-4">
            <div className="text-center py-12 border-2 border-dashed border-muted-foreground/30 rounded-lg">
              <UserPlus className="h-12 w-12 text-muted-foreground mx-auto mb-4" />
              <h3 className="text-lg font-medium mb-2">暂无用户</h3>
              <p className="text-muted-foreground mb-4">没有找到匹配的用户</p>
              <p className="text-sm text-muted-foreground">
                <span className="hidden sm:inline">尝试调整搜索条件</span>
                <span className="sm:hidden">调整搜索条件</span>
              </p>
            </div>
          </div>
        ) : (
          <div className="flex-1 overflow-y-scroll scrollbar-hidden">
            <div className="divide-y divide-border">
              {filteredUsers.map((user) => (
                <div
                  key={user.id}
                  className="px-4 py-4 hover:bg-accent/50 transition-colors"
                >
                  <div className="flex flex-col gap-3">
                    <div className="flex items-start justify-between gap-3">
                      <div className="flex-1 min-w-0">
                        <h3 className="font-medium text-lg line-clamp-2">
                          {user.nickname}
                        </h3>
                        <p className="text-sm text-muted-foreground line-clamp-1">
                          {user.email}
                        </p>
                      </div>
                      <Badge variant="secondary" className="shrink-0">
                        {user.roles.length} 个角色
                      </Badge>
                    </div>

                    <div className="flex items-center justify-between gap-3">
                      <div className="flex flex-wrap gap-2 flex-1">
                        {user.roles.length === 0 ? (
                          <span className="text-sm text-muted-foreground">
                            暂无角色
                          </span>
                        ) : (
                          user.roles.map((role, roleIndex) => (
                            <Badge
                              key={roleIndex}
                              variant="outline"
                              className="text-xs"
                            >
                              <Shield className="mr-1 h-3 w-3" />
                              {role}
                            </Badge>
                          ))
                        )}
                      </div>
                      <DropdownMenu>
                        <DropdownMenuTrigger asChild>
                          <Button
                            variant="ghost"
                            size="sm"
                            className="h-8 w-8 p-0 rounded-full"
                          >
                            <MoreHorizontal className="h-4 w-4" />
                          </Button>
                        </DropdownMenuTrigger>
                        <DropdownMenuContent align="end" className="w-40">
                          <DropdownMenuItem
                            onClick={() =>
                              toast.info(
                                `管理用户 ${user.nickname} 的角色功能开发中...`
                              )
                            }
                            className="cursor-pointer"
                          >
                            <UserPlus className="mr-2 h-4 w-4" />
                            分配角色
                          </DropdownMenuItem>
                          <DropdownMenuSeparator />
                          <DropdownMenuItem
                            onClick={() =>
                              toast.info(
                                `查看用户 ${user.nickname} 详情功能开发中...`
                              )
                            }
                            className="cursor-pointer text-muted-foreground"
                          >
                            <Eye className="mr-2 h-4 w-4" />
                            查看详情
                          </DropdownMenuItem>
                        </DropdownMenuContent>
                      </DropdownMenu>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        )}
      </div>
    </div>
  );

  // ===== 主渲染 =====
  return (
    <div className="h-full">
      {/* 内容区域 */}
      {selectedSection === "roles" && renderRolesList()}
      {selectedSection === "permissions" && renderPermissionsList()}
      {selectedSection === "role-permissions" && renderRolePermissions()}
      {selectedSection === "user-roles" && renderUserRoles()}

      {/* 创建角色对话框 */}
      <Dialog open={isCreateRoleOpen} onOpenChange={setIsCreateRoleOpen}>
        <DialogContent className="sm:max-w-md">
          <DialogHeader>
            <DialogTitle>创建角色</DialogTitle>
          </DialogHeader>
          <div className="space-y-4">
            <div className="space-y-4">
              <div>
                <Label htmlFor="roleName">角色名称</Label>
                <Input
                  id="roleName"
                  value={newRoleName}
                  onChange={(e) => setNewRoleName(e.target.value)}
                  placeholder="输入角色名称"
                />
              </div>
              <div>
                <Label htmlFor="roleDescription">
                  描述 <span className="text-muted-foreground">(可选)</span>
                </Label>
                <Input
                  id="roleDescription"
                  value={newRoleDescription}
                  onChange={(e) => setNewRoleDescription(e.target.value)}
                  placeholder="输入角色描述"
                />
              </div>
            </div>
            <div className="border-t pt-4">
              <h4 className="font-medium mb-4">初始权限设置</h4>
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <Label htmlFor="roleResource">资源</Label>
                  <Select
                    value={newRoleResource}
                    onValueChange={setNewRoleResource}
                  >
                    <SelectTrigger>
                      <SelectValue placeholder="选择资源" />
                    </SelectTrigger>
                    <SelectContent className="max-h-80">
                      {RESOURCE_GROUPS.map((group) => (
                        <div key={group.key}>
                          <div className="px-2 py-1.5 text-xs font-semibold text-muted-foreground bg-muted/30 sticky top-0">
                            {group.name}
                          </div>
                          {group.resources.map((resource) => (
                            <SelectItem
                              key={resource.value}
                              value={resource.value}
                              className="pl-4"
                            >
                              <div className="flex flex-col">
                                <span className="font-medium">
                                  {resource.name}
                                </span>
                                <span className="text-xs text-muted-foreground">
                                  {resource.description}
                                </span>
                              </div>
                            </SelectItem>
                          ))}
                        </div>
                      ))}
                    </SelectContent>
                  </Select>
                </div>
                <div>
                  <Label htmlFor="roleAction">操作</Label>
                  <Select
                    value={newRoleAction}
                    onValueChange={setNewRoleAction}
                  >
                    <SelectTrigger>
                      <SelectValue placeholder="选择操作" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value={RBAC_ACTION.GET}>
                        {RBAC_ACTION.GET}
                      </SelectItem>
                      <SelectItem value={RBAC_ACTION.POST}>
                        {RBAC_ACTION.POST}
                      </SelectItem>
                      <SelectItem value={RBAC_ACTION.ALL}>
                        {RBAC_ACTION.ALL}
                      </SelectItem>
                    </SelectContent>
                  </Select>
                </div>
              </div>
            </div>
          </div>
          <DialogFooter>
            <Button
              variant="outline"
              onClick={() => setIsCreateRoleOpen(false)}
            >
              取消
            </Button>
            <Button
              onClick={handleCreateRole}
              disabled={
                createRoleMutation.isPending ||
                !newRoleName.trim() ||
                !newRoleResource ||
                !newRoleAction
              }
            >
              {createRoleMutation.isPending ? "创建中..." : "创建角色"}
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* 创建权限对话框 */}
      <Dialog
        open={isCreatePermissionOpen}
        onOpenChange={setIsCreatePermissionOpen}
      >
        <DialogContent className="sm:max-w-md">
          <DialogHeader>
            <DialogTitle>创建新权限</DialogTitle>
          </DialogHeader>
          <div className="space-y-4">
            <div>
              <Label htmlFor="permissionRole">关联角色 *</Label>
              <Select
                value={newPermissionRole}
                onValueChange={setNewPermissionRole}
              >
                <SelectTrigger>
                  <SelectValue placeholder="选择角色" />
                </SelectTrigger>
                <SelectContent>
                  {roles.map((role: string) => (
                    <SelectItem key={role} value={role}>
                      {role}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
            <div>
              <Label htmlFor="permissionName">权限名称</Label>
              <Input
                id="permissionName"
                value={newPermissionName}
                onChange={(e) => setNewPermissionName(e.target.value)}
                placeholder="输入权限名称"
              />
            </div>
            <div>
              <Label htmlFor="permissionDescription">权限描述</Label>
              <Input
                id="permissionDescription"
                value={newPermissionDescription}
                onChange={(e) => setNewPermissionDescription(e.target.value)}
                placeholder="输入权限描述"
              />
            </div>
            <div>
              <Label htmlFor="permissionResource">资源</Label>
              <Select
                value={newPermissionResource}
                onValueChange={setNewPermissionResource}
              >
                <SelectTrigger>
                  <SelectValue placeholder="选择资源" />
                </SelectTrigger>
                <SelectContent className="max-h-80">
                  {RESOURCE_GROUPS.map((group) => (
                    <div key={group.key}>
                      <div className="px-2 py-1.5 text-xs font-semibold text-muted-foreground bg-muted/30 sticky top-0">
                        {group.name}
                      </div>
                      {group.resources.map((resource) => (
                        <SelectItem
                          key={resource.value}
                          value={resource.value}
                          className="pl-4"
                        >
                          <div className="flex flex-col">
                            <span className="font-medium">{resource.name}</span>
                            <span className="text-xs text-muted-foreground">
                              {resource.description}
                            </span>
                          </div>
                        </SelectItem>
                      ))}
                    </div>
                  ))}
                </SelectContent>
              </Select>
            </div>
            <div>
              <Label htmlFor="permissionAction">操作</Label>
              <Select
                value={newPermissionAction}
                onValueChange={setNewPermissionAction}
              >
                <SelectTrigger>
                  <SelectValue placeholder="选择操作" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value={RBAC_ACTION.GET}>
                    {RBAC_ACTION.GET}
                  </SelectItem>
                  <SelectItem value={RBAC_ACTION.POST}>
                    {RBAC_ACTION.POST}
                  </SelectItem>
                  <SelectItem value={RBAC_ACTION.ALL}>
                    {RBAC_ACTION.ALL}
                  </SelectItem>
                </SelectContent>
              </Select>
            </div>
          </div>
          <DialogFooter>
            <Button
              variant="outline"
              onClick={() => setIsCreatePermissionOpen(false)}
            >
              取消
            </Button>
            <Button
              onClick={handleCreatePermission}
              disabled={createPermissionMutation.isPending}
            >
              {createPermissionMutation.isPending ? "创建中..." : "创建"}
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* 分配角色对话框 */}
      <Dialog open={isAssignRoleOpen} onOpenChange={setIsAssignRoleOpen}>
        <DialogContent className="sm:max-w-md">
          <DialogHeader>
            <DialogTitle>分配角色</DialogTitle>
          </DialogHeader>
          <div className="space-y-4">
            <div>
              <Label htmlFor="assignUserId">用户ID</Label>
              <Input
                id="assignUserId"
                value={assignUserId}
                onChange={(e) => setAssignUserId(e.target.value)}
                placeholder="输入用户ID"
              />
            </div>
            <div>
              <Label htmlFor="assignRoleId">角色</Label>
              <Select value={assignRoleId} onValueChange={setAssignRoleId}>
                <SelectTrigger>
                  <SelectValue placeholder="选择角色" />
                </SelectTrigger>
                <SelectContent>
                  {roles.map((role: string) => (
                    <SelectItem key={role} value={role}>
                      {role}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
          </div>
          <DialogFooter>
            <Button
              variant="outline"
              onClick={() => setIsAssignRoleOpen(false)}
            >
              取消
            </Button>
            <Button onClick={handleAssignRoleDialog}>分配</Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* 分配权限对话框 */}
      <Dialog
        open={isAssignPermissionOpen}
        onOpenChange={setIsAssignPermissionOpen}
      >
        <DialogContent className="sm:max-w-md">
          <DialogHeader>
            <DialogTitle>分配权限</DialogTitle>
          </DialogHeader>
          <div className="space-y-4">
            <div>
              <Label htmlFor="assignPermissionRole">角色</Label>
              <Select
                value={assignPermissionRole}
                onValueChange={setAssignPermissionRole}
              >
                <SelectTrigger>
                  <SelectValue placeholder="选择角色" />
                </SelectTrigger>
                <SelectContent>
                  {roles.map((role: string) => (
                    <SelectItem key={role} value={role}>
                      {role}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
            <div>
              <Label htmlFor="assignPermissionResource">资源</Label>
              <Select
                value={assignPermissionResource}
                onValueChange={setAssignPermissionResource}
              >
                <SelectTrigger>
                  <SelectValue placeholder="选择资源" />
                </SelectTrigger>
                <SelectContent className="max-h-80">
                  {RESOURCE_GROUPS.map((group) => (
                    <div key={group.key}>
                      <div className="px-2 py-1.5 text-xs font-semibold text-muted-foreground bg-muted/30 sticky top-0">
                        {group.name}
                      </div>
                      {group.resources.map((resource) => (
                        <SelectItem
                          key={resource.value}
                          value={resource.value}
                          className="pl-4"
                        >
                          <div className="flex flex-col">
                            <span className="font-medium">{resource.name}</span>
                            <span className="text-xs text-muted-foreground">
                              {resource.description}
                            </span>
                          </div>
                        </SelectItem>
                      ))}
                    </div>
                  ))}
                </SelectContent>
              </Select>
            </div>
            <div>
              <Label htmlFor="assignPermissionAction">操作</Label>
              <Select
                value={assignPermissionAction}
                onValueChange={setAssignPermissionAction}
              >
                <SelectTrigger>
                  <SelectValue placeholder="选择操作" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value={RBAC_ACTION.GET}>
                    {RBAC_ACTION.GET}
                  </SelectItem>
                  <SelectItem value={RBAC_ACTION.POST}>
                    {RBAC_ACTION.POST}
                  </SelectItem>
                  <SelectItem value={RBAC_ACTION.ALL}>
                    {RBAC_ACTION.ALL}
                  </SelectItem>
                </SelectContent>
              </Select>
            </div>
          </div>
          <DialogFooter>
            <Button
              variant="outline"
              onClick={() => setIsAssignPermissionOpen(false)}
            >
              取消
            </Button>
            <Button
              onClick={handleAssignPermission}
              disabled={assignPermissionMutation.isPending}
            >
              {assignPermissionMutation.isPending ? "分配中..." : "分配"}
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      {/* 查看角色权限对话框 */}
      <Dialog
        open={isViewRolePermissionsOpen}
        onOpenChange={setIsViewRolePermissionsOpen}
      >
        <DialogContent className="sm:max-w-2xl">
          <DialogHeader>
            <DialogTitle className="flex items-center gap-2">
              <Shield className="h-5 w-5" />
              {selectedRole} 的权限
              {rolePermissionsTotal > 0 && (
                <Badge variant="secondary" className="ml-2">
                  {rolePermissionsTotal}
                </Badge>
              )}
            </DialogTitle>
          </DialogHeader>
          <div className="space-y-4">
            {rolePermissionsLoading ? (
              <div className="flex items-center justify-center py-8">
                <div className="text-center">
                  <div className="animate-spin rounded-full h-6 w-6 border-b-2 border-primary mx-auto mb-2"></div>
                  <p className="text-sm text-muted-foreground">加载中...</p>
                </div>
              </div>
            ) : rolePermissions.length === 0 ? (
              <div className="text-center py-8 border-2 border-dashed border-muted-foreground/30 rounded-lg">
                <Shield className="h-8 w-8 text-muted-foreground mx-auto mb-3" />
                <h4 className="font-medium mb-1">暂无权限</h4>
                <p className="text-sm text-muted-foreground">
                  该角色还没有分配任何权限
                </p>
              </div>
            ) : (
              <div className="space-y-2 max-h-96 overflow-y-auto">
                {rolePermissions.map((permission: any, index: number) => {
                  const resourceInfo = getResourceInfo(permission.resource);
                  const IconComponent = resourceInfo.groupIcon;

                  return (
                    <div
                      key={index}
                      className="px-4 py-4 hover:bg-accent/50 transition-colors border rounded-lg"
                    >
                      <div className="flex flex-col gap-3">
                        <div className="flex items-start justify-between gap-3">
                          <div className="flex items-start gap-3 flex-1 min-w-0">
                            <div className="w-8 h-8 rounded-full bg-primary/10 flex items-center justify-center shrink-0 mt-0.5">
                              <IconComponent className="h-4 w-4 text-primary" />
                            </div>
                            <div className="flex-1 min-w-0">
                              <h4 className="font-medium text-base line-clamp-1">
                                {resourceInfo.name}
                              </h4>
                              <p className="text-sm text-muted-foreground line-clamp-2 mt-1">
                                {resourceInfo.description}
                              </p>
                            </div>
                          </div>
                          {permission.action && (
                            <Badge
                              variant="outline"
                              className="text-xs shrink-0"
                            >
                              {permission.action}
                            </Badge>
                          )}
                        </div>

                        <div className="flex items-center justify-between gap-3">
                          <div className="flex items-center gap-4 text-xs text-muted-foreground">
                            <div className="flex items-center gap-1.5">
                              <div className="w-1.5 h-1.5 rounded-full bg-blue-500" />
                              <span>{resourceInfo.groupName}</span>
                            </div>
                            <div className="flex items-center gap-1.5">
                              <div className="w-1.5 h-1.5 rounded-full bg-green-500" />
                              <span>资源: {permission.resource}</span>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                  );
                })}
              </div>
            )}
          </div>
          <DialogFooter>
            <Button
              variant="outline"
              onClick={() => setIsViewRolePermissionsOpen(false)}
              className="w-full"
            >
              关闭
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
}
