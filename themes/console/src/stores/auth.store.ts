/**
 * 认证状态管理 - Zustand
 * 专注于认证状态和token管理
 */

import { create } from "zustand";
import { persist } from "zustand/middleware";

import { redirect } from "@tanstack/react-router";

import { AUTH_ROUTES, STORAGE_KEYS } from "@/constants";
import type { LoginResponse } from "@/types";

// 认证检查函数
export const checkAuth = (redirectTo: string = AUTH_ROUTES.LOGIN) => {
  const token = localStorage.getItem(STORAGE_KEYS.ACCESS_TOKEN);
  if (!token) {
    throw redirect({ to: redirectTo });
  }
  return true;
};

// 认证状态接口
export interface AuthStoreState {
  // ===== 状态 =====
  isAuthenticated: boolean; // 是否已认证
  loading: boolean; // 认证加载状态

  // ===== 操作方法 =====
  setAuthenticated: (authenticated: boolean) => void; // 设置认证状态
  setLoading: (loading: boolean) => void; // 设置加载状态
  login: (tokens: LoginResponse) => void; // 登录
  logout: () => void; // 登出
  refreshTokens: (accessToken: string, refreshToken: string) => void; // 刷新token后更新状态
  clearAuth: () => void; // 清理认证状态
}

export const useAuthStore = create<AuthStoreState>()(
  persist(
    (set) => ({
      // ===== 初始状态 =====
      // 从 localStorage 读取认证状态（仅在浏览器端）
      isAuthenticated:
        typeof window !== "undefined"
          ? !!localStorage.getItem(STORAGE_KEYS.ACCESS_TOKEN)
          : false,
      loading: false,

      // ===== 操作方法 =====

      // 设置认证状态
      setAuthenticated: (authenticated: boolean) => {
        set({ isAuthenticated: authenticated });
      },

      // 设置加载状态
      setLoading: (loading: boolean) => {
        set({ loading });
      },

      // 登录
      login: (tokens: LoginResponse) => {
        // 保存 tokens 到 localStorage
        localStorage.setItem(STORAGE_KEYS.ACCESS_TOKEN, tokens.access_token);
        localStorage.setItem(STORAGE_KEYS.REFRESH_TOKEN, tokens.refresh_token);

        // 更新认证状态
        set({ isAuthenticated: true, loading: false });
      },

      // 登出
      logout: () => {
        localStorage.removeItem(STORAGE_KEYS.ACCESS_TOKEN);
        localStorage.removeItem(STORAGE_KEYS.REFRESH_TOKEN);

        set({ isAuthenticated: false, loading: false });

        // 清理 Zustand 持久化存储
        localStorage.removeItem(STORAGE_KEYS.AUTH_STORE);

        // 重定向到登录页
        window.location.href = AUTH_ROUTES.LOGIN;
      },

      // 刷新token后更新状态
      refreshTokens: (accessToken: string, refreshToken: string) => {
        localStorage.setItem(STORAGE_KEYS.ACCESS_TOKEN, accessToken);
        localStorage.setItem(STORAGE_KEYS.REFRESH_TOKEN, refreshToken);

        // 保持认证状态
        set({ isAuthenticated: true });
      },

      // 清理认证状态（不重定向）
      clearAuth: () => {
        // 清理 localStorage 中的 tokens
        localStorage.removeItem(STORAGE_KEYS.ACCESS_TOKEN);
        localStorage.removeItem(STORAGE_KEYS.REFRESH_TOKEN);

        // 清理认证状态
        set({ isAuthenticated: false, loading: false });
      },
    }),
    {
      name: STORAGE_KEYS.AUTH_STORE, // localStorage 中的键名
      partialize: (state) => ({
        isAuthenticated: state.isAuthenticated,
      }),
    }
  )
);
