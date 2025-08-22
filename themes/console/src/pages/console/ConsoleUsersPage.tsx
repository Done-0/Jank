/**
 * 用户管理页面
 */

import { useState, useMemo } from "react";

import { UsersSidebar } from "@/components/console/users/UsersSidebar";
import { UsersContent } from "@/components/console/users/UsersContent";

import { useUsers } from "@/hooks/use-user";
import type { UserItem } from "@/types/user";

export function ConsoleUsersPage() {
  // ===== 状态管理 =====
  const [selectedRole, setSelectedRole] = useState<string | null>(null);
  const [searchQuery, setSearchQuery] = useState("");

  // ===== 数据获取 =====
  const { data: usersData, isLoading } = useUsers({
    page_no: 1,
    page_size: 100,
    ...(searchQuery && { keyword: searchQuery }),
  });

  // ===== 计算数据 =====
  const allUsers = useMemo(() => usersData?.list || [], [usersData?.list]);

  // ===== 计算统计数据 =====
  const stats = useMemo(() => {
    return {
      total: allUsers.length,
      superAdmins: allUsers.filter((u: UserItem) =>
        u.roles.includes("super_admin")
      ).length,
      regularUsers: allUsers.filter(
        (u: UserItem) => u.roles.includes("user") || u.roles.length === 0
      ).length,
    };
  }, [allUsers]);

  // ===== 事件处理 =====
  const handleRoleChange = (role: string | null) => setSelectedRole(role);
  const handleSearchChange = (query: string) => setSearchQuery(query);

  // ===== 渲染 =====
  return (
    <div className="flex h-full">
      <div className="hidden md:block">
        <UsersSidebar
          selectedRole={selectedRole}
          stats={stats}
          allUsers={allUsers}
          onRoleChange={handleRoleChange}
        />
      </div>

      <div className="flex-1 flex flex-col h-full">
        <UsersContent
          selectedRole={selectedRole}
          searchQuery={searchQuery}
          allUsers={allUsers}
          isLoading={isLoading}
          onSearchChange={handleSearchChange}
        />
      </div>
    </div>
  );
}
