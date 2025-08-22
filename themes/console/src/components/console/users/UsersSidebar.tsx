/**
 * 用户管理侧边栏组件
 */

import { Users, Crown, User, Shield } from "lucide-react";
import { useRoles } from "@/hooks/use-rbac";
import { useMemo } from "react";

interface UsersSidebarProps {
  selectedRole: string | null;
  stats: {
    total: number;
    superAdmins: number;
    regularUsers: number;
  };
  allUsers: any[];
  onRoleChange: (role: string | null) => void;
}

export function UsersSidebar({
  selectedRole,
  stats,
  allUsers,
  onRoleChange,
}: UsersSidebarProps) {
  // 获取所有角色
  const { data: rolesData } = useRoles();
  const allRoles = useMemo(() => rolesData?.list || [], [rolesData?.list]);

  // 动态计算角色统计
  const roleStats = useMemo(() => {
    const roleCount: Record<string, number> = {};

    allUsers.forEach((user) => {
      user.roles.forEach((role: string) => {
        roleCount[role] = (roleCount[role] || 0) + 1;
      });
    });

    return roleCount;
  }, [allUsers]);

  // 构建角色项目列表
  const roleItems = useMemo(() => {
    const items: Array<{
      role: string | null;
      icon: any;
      label: string;
      count: number;
    }> = [{ role: null, icon: Users, label: "全部", count: stats.total }];

    // 添加静态角色
    items.push(
      {
        role: "super_admin",
        icon: Crown,
        label: "超级管理员",
        count: stats.superAdmins,
      },
      { role: "user", icon: User, label: "普通用户", count: stats.regularUsers }
    );

    // 添加动态角色（排除已有的静态角色）
    allRoles.forEach((roleName: string) => {
      if (roleName !== "super_admin" && roleName !== "user") {
        items.push({
          role: roleName,
          icon: Shield,
          label: roleName,
          count: roleStats[roleName] || 0,
        });
      }
    });

    return items;
  }, [stats, allRoles, roleStats]);

  return (
    <div className="w-64 lg:w-72 border-r bg-background hidden md:flex">
      <div className="h-full flex flex-col w-full">
        {/* 角色筛选 */}
        <div className="px-4 py-4">
          <div className="space-y-0.5">
            {roleItems.map(({ role, icon: Icon, label, count }) => (
              <button
                key={role || "all"}
                onClick={() => onRoleChange(role)}
                className={`w-full flex items-center gap-3 px-3 py-2.5 rounded-full text-left transition-colors ${
                  selectedRole === role
                    ? "bg-accent text-accent-foreground font-medium"
                    : "hover:bg-accent/50 text-foreground"
                }`}
              >
                <Icon className="h-5 w-5 flex-shrink-0" />
                <span className="flex-1 text-sm">{label}</span>
                <span className="text-xs text-muted-foreground font-medium">
                  {count}
                </span>
              </button>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
}
