/**
 * 分类状态管理
 * 参考 console 主题的具体命名规范
 */

export class CategoryStore {
  constructor() {
    this.state = {
      categories: [],
      currentCategory: null,
      loading: false,
      error: null
    };
    
    this.listeners = [];
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
  }

  /**
   * 设置分类列表
   */
  setCategories(categories) {
    this.state.categories = categories;
    this.notify();
  }

  /**
   * 设置当前分类
   */
  setCurrentCategory(category) {
    this.state.currentCategory = category;
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
}
