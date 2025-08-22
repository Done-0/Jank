/**
 * 文章状态管理
 */

import { STORAGE_KEYS } from "../constants/index.js";

export class PostStore {
  constructor() {
    this.state = {
      posts: [],
      currentPost: null,
      loading: false,
      error: null,
      pagination: {
        currentPage: 1,
        pageSize: 5,
        total: 0
      }
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
   * 设置文章列表
   */
  setPosts(posts) {
    this.state.posts = posts;
    this.notify();
  }

  /**
   * 设置当前文章
   */
  setCurrentPost(post) {
    this.state.currentPost = post;
    this.notify();
  }

  /**
   * 设置加载状态
   */
  setLoading(loading) {
    this.state.loading = loading;
    this.notify();
  }

  /**
   * 设置错误状态
   */
  setError(error) {
    this.state.error = error;
    this.notify();
  }

  /**
   * 设置分页
   */
  setPagination(pagination) {
    this.state.pagination = { ...this.state.pagination, ...pagination };
    this.notify();
  }

  /**
   * 加载持久化状态
   */
  loadPersistedState() {
    try {
      const saved = localStorage.getItem(STORAGE_KEYS.BLOG_STATE);
      if (saved) {
        const parsedState = JSON.parse(saved);
        this.state.pagination.pageSize = 5;
      }
    } catch (error) {
      console.warn('Failed to load persisted state:', error);
    }
  }

  /**
   * 持久化状态
   */
  persistState() {
    try {
      const stateToSave = {
        pageSize: this.state.pagination.pageSize,
      };
      localStorage.setItem(STORAGE_KEYS.BLOG_STATE, JSON.stringify(stateToSave));
    } catch (error) {
      console.warn('Failed to persist state:', error);
    }
  }
}
