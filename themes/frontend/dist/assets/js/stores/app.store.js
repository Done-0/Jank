/**
 * 应用状态管理
 */

import { STORAGE_KEYS, THEME_VALUES } from "../constants/index.js";

export class AppStore {
  constructor() {
    this.state = {
      currentSection: 'home',
      theme: THEME_VALUES.LIGHT,
      loading: false,
      error: null
    };
    
    this.listeners = [];
    this.loadPersistedState();
  }

  /**
   * 获取当前状态
   */
  getState() {
    return { ...this.state };
  }

  /**
   * 订阅状态变化
   */
  subscribe(listener) {
    this.listeners.push(listener);
    return () => {
      const index = this.listeners.indexOf(listener);
      if (index > -1) {
        this.listeners.splice(index, 1);
      }
    };
  }

  /**
   * 通知所有监听器
   */
  notify() {
    this.listeners.forEach(listener => listener(this.state));
    this.persistState();
  }

  /**
   * 设置当前页面
   */
  setCurrentSection(section) {
    this.state.currentSection = section;
    this.notify();
  }

  /**
   * 切换主题
   */
  toggleTheme() {
    this.state.theme = this.state.theme === THEME_VALUES.LIGHT 
      ? THEME_VALUES.DARK 
      : THEME_VALUES.LIGHT;
    this.applyTheme();
    this.notify();
  }

  /**
   * 设置主题
   */
  setTheme(theme) {
    this.state.theme = theme;
    this.applyTheme();
    this.notify();
  }

  /**
   * 应用主题到DOM
   */
  applyTheme() {
    const body = document.body;
    if (this.state.theme === THEME_VALUES.DARK) {
      body.classList.add('dark');
    } else {
      body.classList.remove('dark');
    }
  }

  /**
   * 设置全局加载状态
   */
  setLoading(loading) {
    this.state.loading = loading;
    this.notify();
  }

  /**
   * 设置全局错误状态
   */
  setError(error) {
    this.state.error = error;
    this.notify();
  }

  /**
   * 加载持久化状态
   */
  loadPersistedState() {
    try {
      const savedTheme = localStorage.getItem(STORAGE_KEYS.THEME);
      if (savedTheme && Object.values(THEME_VALUES).includes(savedTheme)) {
        this.state.theme = savedTheme;
        this.applyTheme();
      }
    } catch (error) {
      console.warn('Failed to load persisted theme:', error);
    }
  }

  /**
   * 持久化状态
   */
  persistState() {
    try {
      localStorage.setItem(STORAGE_KEYS.THEME, this.state.theme);
    } catch (error) {
      console.warn('Failed to persist theme:', error);
    }
  }
}
