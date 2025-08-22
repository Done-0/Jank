/**
 * 用户管理主内容组件
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
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Search, MoreHorizontal, Crown, User } from "lucide-react";
import { toast } from "sonner";

import { useUpdateUserRole } from "@/hooks/use-user";
import { useRoles } from "@/hooks/use-rbac";
import type { UserItem } from "@/types/user";

interface UsersContentProps {
  selectedRole: string | null;
  searchQuery: string;
  allUsers: UserItem[];
  isLoading: boolean;
  onSearchChange: (query: string) => void;
}

export function UsersContent({
  selectedRole,
  searchQuery,
  allUsers,
  isLoading,
  onSearchChange,
}: UsersContentProps) {
  // ===== 对话框状态 =====
  const [isRoleDialogOpen, setIsRoleDialogOpen] = useState(false);
  const [selectedUser, setSelectedUser] = useState<UserItem | null>(null);

  // ===== 表单状态 =====
  const [editUserRole, setEditUserRole] = useState("");

  // ===== 数据查询 =====
  const { data: rolesData } = useRoles();
  const roles = rolesData?.list || [];

  // ===== Mutations =====

  const updateUserRoleMutation = useUpdateUserRole({
    onSuccess: () => {
      toast.success("用户角色更新成功");
      setIsRoleDialogOpen(false);
    },
    onError: (error) => {
      toast.error(`角色更新失败: ${error.message}`);
    },
  });

  // ===== 过滤用户 =====
  const filteredUsers = useMemo(() => {
    return allUsers.filter((user: UserItem) => {
      const matchesSearch =
        user.nickname?.toLowerCase().includes(searchQuery.toLowerCase()) ||
        user.email?.toLowerCase().includes(searchQuery.toLowerCase());
      const matchesRole = !selectedRole || user.roles.includes(selectedRole);
      return matchesSearch && matchesRole;
    });
  }, [allUsers, searchQuery, selectedRole]);

  // ===== 事件处理函数 =====
  const handleAssignRole = (user: UserItem) => {
    setSelectedUser(user);
    setEditUserRole(user.roles[0] || "");
    setIsRoleDialogOpen(true);
  };

  const handleSaveRole = () => {
    if (!selectedUser) return;

    updateUserRoleMutation.mutate({
      id: selectedUser.id,
      role: editUserRole,
    });
  };

  return (
    <div className="flex-1 flex flex-col h-full">
      {/* 顶部操作栏 */}
      <div className="px-4 py-4 border-b">
        <div className="flex items-center justify-between gap-4">
          <div className="relative flex-1 sm:flex-initial sm:w-80">
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
            <Input
              placeholder="搜索用户..."
              className="pl-9 h-10 rounded-full"
              value={searchQuery}
              onChange={(e) => onSearchChange(e.target.value)}
            />
          </div>
        </div>
      </div>

      {/* 用户列表区域 */}
      <div className="flex-1 flex flex-col overflow-hidden w-full">
        {isLoading ? (
          <div className="flex items-center justify-center h-full">
            <div className="text-center">
              <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary mx-auto mb-4"></div>
              <p className="text-sm text-muted-foreground">加载中...</p>
            </div>
          </div>
        ) : filteredUsers.length > 0 ? (
          <div className="flex-1 overflow-y-scroll scrollbar-hidden">
            <div className="divide-y divide-border">
              {filteredUsers.map((user) => (
                <div
                  key={user.id}
                  className="px-4 py-4 hover:bg-accent/50 transition-colors"
                >
                  <div className="flex flex-col gap-3">
                    <div className="flex items-start justify-between gap-3">
                      <div className="flex items-center gap-3 flex-1 min-w-0">
                        <div className="w-10 h-10 rounded-full bg-primary/10 flex items-center justify-center ring-1 ring-border flex-shrink-0">
                          <span className="text-sm font-medium text-primary">
                            {user.nickname.charAt(0).toUpperCase()}
                          </span>
                        </div>
                        <div className="min-w-0 flex-1">
                          <h3 className="font-medium text-lg line-clamp-1">
                            {user.nickname}
                          </h3>
                          <p className="text-sm text-muted-foreground line-clamp-1">
                            {user.email}
                          </p>
                        </div>
                      </div>
                      <div className="flex items-center gap-2 flex-shrink-0">
                        {user.roles.map((role, index) => (
                          <Badge
                            key={index}
                            variant={
                              role === "super_admin" ? "default" : "secondary"
                            }
                            className="shrink-0"
                          >
                            {role === "super_admin" ? "超级管理员" : role}
                          </Badge>
                        ))}
                      </div>
                    </div>

                    <div className="flex items-center justify-between gap-3">
                      <div className="flex items-center gap-2 text-xs text-muted-foreground">
                        <div className="flex items-center gap-1.5">
                          <div className="w-1.5 h-1.5 rounded-full bg-slate-400" />
                          <span>
                            创建时间: {new Date().toLocaleDateString()}
                          </span>
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
                            onClick={() => handleAssignRole(user)}
                            className="cursor-pointer"
                          >
                            <Crown className="mr-2 h-4 w-4" />
                            分配角色
                          </DropdownMenuItem>
                        </DropdownMenuContent>
                      </DropdownMenu>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        ) : (
          <div className="flex-1 overflow-auto p-4">
            <div className="text-center py-12 border-2 border-dashed border-muted-foreground/30 rounded-lg">
              <User className="h-12 w-12 text-muted-foreground mx-auto mb-4" />
              <h3 className="text-lg font-medium mb-2">暂无用户</h3>
              <p className="text-muted-foreground mb-4">
                {searchQuery
                  ? `未找到包含“${searchQuery}”的用户`
                  : "暂无任何用户内容"}
              </p>
              <p className="text-sm text-muted-foreground">暂无用户数据</p>
            </div>
          </div>
        )}
      </div>

      {/* 分配角色对话框 */}
      <Dialog open={isRoleDialogOpen} onOpenChange={setIsRoleDialogOpen}>
        <DialogContent className="sm:max-w-md">
          <DialogHeader>
            <DialogTitle>分配角色</DialogTitle>
          </DialogHeader>
          <div className="space-y-4 py-4">
            <div className="space-y-2">
              <Label htmlFor="role">用户角色</Label>
              <Select value={editUserRole} onValueChange={setEditUserRole}>
                <SelectTrigger>
                  <SelectValue placeholder="选择角色" />
                </SelectTrigger>
                <SelectContent>
                  {roles.map((role, index) => (
                    <SelectItem key={index} value={role}>
                      {role}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
            {selectedUser && (
              <div className="text-sm text-muted-foreground">
                用户: {selectedUser.nickname} ({selectedUser.email})
              </div>
            )}
          </div>
          <DialogFooter>
            <Button
              variant="outline"
              onClick={() => setIsRoleDialogOpen(false)}
            >
              取消
            </Button>
            <Button
              onClick={handleSaveRole}
              disabled={
                !editUserRole.trim() || updateUserRoleMutation.isPending
              }
            >
              {updateUserRoleMutation.isPending ? "保存中..." : "保存"}
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
}
