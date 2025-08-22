import { Users, Edit, Zap, Settings, Shield, Mail } from "lucide-react";
import {
  USER_ENDPOINTS,
  POST_ENDPOINTS,
  PLUGIN_ENDPOINTS,
  THEME_ENDPOINTS,
  RBAC_ENDPOINTS,
  CATEGORY_ENDPOINTS,
  VERIFICATION_ENDPOINTS,
} from "@/api/endpoints";

/**
 * RBAC 相关常量定义
 */

// ===== RBAC 动作枚举 =====
export const RBAC_ACTION = {
  ALL: "*",
  GET: "GET",
  POST: "POST",
  PUT: "PUT",
  DELETE: "DELETE",
} as const;

// ===== RBAC 查询键 =====
export const RBAC_QUERY_KEYS = {
  CHECK_PERMISSION: "checkPermission",
  USER_ROLES: "userRoles",
  LIST_ROLES: "listRoles",
  LIST_PERMISSIONS: "listPermissions",
  GET_ROLE_PERMISSIONS: "getRolePermissions",
} as const;

// ===== 预定义资源选项 =====
export const RESOURCE_GROUPS = [
  {
    name: "用户管理",
    key: "user",
    icon: Users,
    resources: [
      {
        value: USER_ENDPOINTS.LIST_USERS,
        name: "查看用户列表",
        description: "查看所有用户信息",
      },
      {
        value: USER_ENDPOINTS.GET_PROFILE,
        name: "管理用户资料",
        description: "查看和编辑用户资料",
      },
      {
        value: USER_ENDPOINTS.REGISTER,
        name: "用户注册",
        description: "创建新用户账户",
      },
      {
        value: USER_ENDPOINTS.LOGIN,
        name: "用户登录",
        description: "用户身份验证",
      },
      {
        value: USER_ENDPOINTS.LOGOUT,
        name: "用户登出",
        description: "退出用户会话",
      },
      {
        value: USER_ENDPOINTS.UPDATE,
        name: "更新用户信息",
        description: "修改用户基本信息",
      },
      {
        value: USER_ENDPOINTS.RESET_PASSWORD,
        name: "重置密码",
        description: "重置用户密码",
      },
      {
        value: USER_ENDPOINTS.REFRESH_TOKEN,
        name: "刷新令牌",
        description: "刷新用户访问令牌",
      },
      {
        value: USER_ENDPOINTS.UPDATE_USER_ROLE,
        name: "用户角色管理",
        description: "管理用户角色分配",
      },
      {
        value: "/api/v1/user/*",
        name: "用户管理所有权限",
        description: "对用户模块的全部操作",
      },
    ],
  },
  {
    name: "验证码管理",
    key: "verification",
    icon: Mail,
    resources: [
      {
        value: VERIFICATION_ENDPOINTS.SEND_EMAIL_CODE,
        name: "发送邮箱验证码",
        description: "发送邮箱验证码进行身份验证",
      },
      {
        value: "/api/v1/verification/*",
        name: "验证码管理所有权限",
        description: "对验证码模块的全部操作",
      },
    ],
  },
  {
    name: "文章管理",
    key: "post",
    icon: Edit,
    resources: [
      {
        value: POST_ENDPOINTS.LIST_PUBLISHED_POSTS,
        name: "查看已发布文章",
        description: "浏览所有已发布文章",
      },
      {
        value: POST_ENDPOINTS.LIST_POSTS_BY_STATUS,
        name: "按状态查看文章",
        description: "按状态筛选查看文章",
      },
      {
        value: POST_ENDPOINTS.CREATE_POST,
        name: "创建文章",
        description: "发布新文章",
      },
      {
        value: POST_ENDPOINTS.UPDATE_POST,
        name: "编辑文章",
        description: "修改现有文章",
      },
      {
        value: POST_ENDPOINTS.DELETE_POST,
        name: "删除文章",
        description: "删除指定文章",
      },
      {
        value: POST_ENDPOINTS.GET_POST,
        name: "查看文章详情",
        description: "查看单篇文章详细信息",
      },
      {
        value: "/api/v1/post/*",
        name: "文章管理所有权限",
        description: "对文章模块的全部操作",
      },
    ],
  },
  {
    name: "插件管理",
    key: "plugin",
    icon: Zap,
    resources: [
      {
        value: PLUGIN_ENDPOINTS.LIST_PLUGINS,
        name: "查看插件列表",
        description: "浏览所有插件",
      },
      {
        value: PLUGIN_ENDPOINTS.REGISTER_PLUGIN,
        name: "注册插件",
        description: "安装新插件",
      },
      {
        value: PLUGIN_ENDPOINTS.UNREGISTER_PLUGIN,
        name: "卸载插件",
        description: "删除现有插件",
      },
      {
        value: PLUGIN_ENDPOINTS.EXECUTE_PLUGIN,
        name: "执行插件",
        description: "运行插件功能",
      },
      {
        value: PLUGIN_ENDPOINTS.GET_PLUGIN,
        name: "查看插件详情",
        description: "查看插件详细信息",
      },
      {
        value: "/api/v1/plugin/*",
        name: "插件管理所有权限",
        description: "对插件模块的全部操作",
      },
    ],
  },
  {
    name: "主题管理",
    key: "theme",
    icon: Settings,
    resources: [
      {
        value: THEME_ENDPOINTS.LIST_THEMES,
        name: "查看主题列表",
        description: "浏览所有主题",
      },
      {
        value: THEME_ENDPOINTS.SWITCH_THEME,
        name: "切换主题",
        description: "更改网站主题",
      },
      {
        value: THEME_ENDPOINTS.GET_ACTIVE_THEME,
        name: "获取当前主题",
        description: "查看当前激活主题",
      },
      {
        value: "/api/v1/theme/*",
        name: "主题管理所有权限",
        description: "对主题模块的全部操作",
      },
    ],
  },
  {
    name: "分类管理",
    key: "category",
    icon: Edit,
    resources: [
      {
        value: CATEGORY_ENDPOINTS.LIST_CATEGORIES,
        name: "查看分类列表",
        description: "浏览所有分类",
      },
      {
        value: CATEGORY_ENDPOINTS.CREATE_CATEGORY,
        name: "创建分类",
        description: "新建文章分类",
      },
      {
        value: CATEGORY_ENDPOINTS.UPDATE_CATEGORY,
        name: "编辑分类",
        description: "修改现有分类",
      },
      {
        value: CATEGORY_ENDPOINTS.DELETE_CATEGORY,
        name: "删除分类",
        description: "删除指定分类",
      },
      {
        value: CATEGORY_ENDPOINTS.GET_CATEGORY,
        name: "查看分类详情",
        description: "查看单个分类详细信息",
      },
      {
        value: "/api/v1/category/*",
        name: "分类管理所有权限",
        description: "对分类模块的全部操作",
      },
    ],
  },
  {
    name: "权限管理",
    key: "rbac",
    icon: Shield,
    resources: [
      // 角色管理
      {
        value: RBAC_ENDPOINTS.LIST_ROLES,
        name: "查看角色列表",
        description: "查看系统所有角色",
      },
      {
        value: RBAC_ENDPOINTS.CREATE_ROLE,
        name: "创建角色",
        description: "创建新的系统角色",
      },
      {
        value: RBAC_ENDPOINTS.DELETE_ROLE,
        name: "删除角色",
        description: "删除现有角色",
      },
      {
        value: RBAC_ENDPOINTS.GET_ROLE_PERMISSIONS,
        name: "查看角色权限",
        description: "查看角色拥有的权限",
      },

      // 权限管理
      {
        value: RBAC_ENDPOINTS.LIST_PERMISSIONS,
        name: "查看权限列表",
        description: "查看系统所有权限策略",
      },
      {
        value: RBAC_ENDPOINTS.CREATE_PERMISSION,
        name: "创建权限",
        description: "创建新的权限策略",
      },
      {
        value: RBAC_ENDPOINTS.DELETE_PERMISSION,
        name: "删除权限",
        description: "删除现有权限",
      },
      {
        value: RBAC_ENDPOINTS.ASSIGN_PERMISSION,
        name: "分配权限",
        description: "为角色分配权限",
      },
      {
        value: RBAC_ENDPOINTS.REVOKE_PERMISSION,
        name: "撤销权限",
        description: "撤销角色权限",
      },

      // 用户角色管理
      {
        value: RBAC_ENDPOINTS.ASSIGN_ROLE,
        name: "分配角色",
        description: "为用户分配角色",
      },
      {
        value: RBAC_ENDPOINTS.REVOKE_ROLE,
        name: "撤销角色",
        description: "撤销用户角色",
      },
      {
        value: RBAC_ENDPOINTS.GET_USER_ROLES,
        name: "查看用户角色",
        description: "查看用户拥有的角色",
      },

      // 权限检查
      {
        value: RBAC_ENDPOINTS.CHECK_PERMISSION,
        name: "权限检查",
        description: "检查用户权限",
      },

      {
        value: "/api/v1/rbac/*",
        name: "权限管理所有权限",
        description: "对权限模块的全部操作",
      },
    ],
  },
] as const;

// ===== 类型定义 =====
export type RbacAction = (typeof RBAC_ACTION)[keyof typeof RBAC_ACTION];
export type ResourceGroup = (typeof RESOURCE_GROUPS)[number];
export type ResourceItem = ResourceGroup["resources"][number];
