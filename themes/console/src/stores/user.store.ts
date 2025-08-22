/**
 * 用户信息状态管理 - Zustand
 * 专注于用户资料和相关UI状态
 */

import { create } from "zustand";
import { persist } from "zustand/middleware";

import { STORAGE_KEYS } from "@/constants";
import type { GetProfileResponse } from "@/types";

// 用户状态接口
export interface UserStoreState {
  // ===== 状态 =====
  user: GetProfileResponse | null; // 当前用户信息
  loading: boolean; // 用户信息加载状态

  // ===== 操作方法 =====
  setUser: (user: GetProfileResponse | null) => void; // 设置用户信息
  setLoading: (loading: boolean) => void; // 设置加载状态
  updateUser: (updates: Partial<GetProfileResponse>) => void; // 更新用户部分信息
  clearUser: () => void; // 清理用户信息
}

export const useUserStore = create<UserStoreState>()(
  persist(
    (set, get) => ({
      // ===== 初始状态 =====
      user: null,
      loading: false,

      // ===== 操作方法 =====

      // 设置用户信息
      setUser: (user: GetProfileResponse | null) => {
        set({ user });

        if (user) {
          localStorage.setItem(STORAGE_KEYS.USER_INFO, JSON.stringify(user));
        } else {
          localStorage.removeItem(STORAGE_KEYS.USER_INFO);
        }
      },

      // 设置加载状态
      setLoading: (loading: boolean) => {
        set({ loading });
      },

      // 更新用户部分信息
      updateUser: (updates: Partial<GetProfileResponse>) => {
        const currentUser = get().user;
        if (currentUser) {
          set({
            user: {
              ...currentUser,
              ...updates,
            },
          });
        }
      },

      // 清理用户信息
      clearUser: () => {
        set({ user: null, loading: false });

        // 清理 localStorage 中的用户信息
        localStorage.removeItem(STORAGE_KEYS.USER_INFO);

        // 清理 Zustand 持久化存储
        localStorage.removeItem(STORAGE_KEYS.USER_STORE);
      },
    }),
    {
      name: STORAGE_KEYS.USER_STORE, // localStorage 中的键名
      partialize: (state) => ({
        user: state.user,
      }),
    }
  )
);
