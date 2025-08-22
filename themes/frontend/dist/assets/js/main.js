/**
 * 博客主应用
 * 整合所有模块，负责应用初始化和路由管理
 */

import { ApiService } from "./services/index.js";
import { AppStore, PostStore, CategoryStore } from "./stores/index.js";
import { PostRenderer, CategoryRenderer, UiRenderer, ArticleRenderer } from "./components/index.js";
import { DomHelpers } from "./utils/index.js";

class BlogApp {
  constructor() {
    // 依赖注入
    this.apiService = new ApiService();
    this.appStore = new AppStore();
    this.postStore = new PostStore();
    this.categoryStore = new CategoryStore();
    this.postRenderer = new PostRenderer();
    this.categoryRenderer = new CategoryRenderer();
    this.uiRenderer = new UiRenderer();
    this.articleRenderer = new ArticleRenderer();
    
    this.init();
  }

  /**
   * 应用初始化
   */
  init() {
    this.articleRenderer.showSidebars();
    this.setupEvents();
    this.setupRouter();
    this.handleInitialRoute();
  }

  /**
   * 设置事件监听
   */
  setupEvents() {
    // 使用事件委托处理所有点击事件
    document.addEventListener('click', (e) => {
      // 主题切换按钮
      if (e.target.closest('[data-theme-toggle]')) {
        this.appStore.toggleTheme();
        return;
      }

      // 文章卡片点击事件
      const postCard = e.target.closest('.post-card');
      if (postCard) {
        const postId = postCard.dataset.postId;
        if (postId) {
          this.viewPost(postId);
        }
        return;
      }

      // 分类按钮点击事件
      const categoryButton = e.target.closest('.category-button');
      if (categoryButton) {
        const categoryId = categoryButton.dataset.categoryId;
        let actualCategoryId = null;
        if (categoryId && categoryId !== 'null' && categoryId !== '') {
          actualCategoryId = categoryId;
        }

        this.filterByCategory(actualCategoryId);
        return;
      }

      // 返回按钮点击事件
      const backButton = e.target.closest('#back-to-home') || e.target.closest('.back-button');
      if (backButton) {
        e.preventDefault();
        e.stopPropagation();
        this.backToHome();
        return;
      }

      // 导航条Logo点击事件
      const navLogo = e.target.closest('.nav-logo');
      if (navLogo) {
        e.preventDefault();
        this.backToHome();
        return;
      }
    });
  }

  /**
   * 设置路由系统
   */
  setupRouter() {
    // 监听浏览器前进后退
    window.addEventListener('popstate', (e) => {
      this.handleRoute();
    });
  }

  /**
   * 处理初始路由
   */
  handleInitialRoute() {
    this.handleRoute();
  }

  /**
   * 路由处理
   */
  handleRoute() {
    const path = window.location.pathname;
    const pathSegments = path.split('/').filter(segment => segment);

    // 根据路径处理不同的路由
    if (pathSegments.length === 0) {
      // 首页: /
      this.loadHomePage();
    } else if (pathSegments[0] === 'post' && pathSegments[1]) {
      // 文章详情页: /post/:id
      const postId = pathSegments[1];
      this.viewPostByUrl(postId);
    } else if (pathSegments[0] === 'category' && pathSegments[1]) {
      // 分类页面: /category/:id
      const categoryId = pathSegments[1];
      this.filterByCategoryByUrl(categoryId);
    } else {
      // 404 页面
      this.show404();
    }
  }

  /**
   * 页面导航
   */
  navigate(section, updateHistory = true) {
    DomHelpers.removeClassFromAll(DomHelpers.getElements('.page-section'), 'active');
    DomHelpers.removeClassFromAll(DomHelpers.getElements('.nav-link'), 'active');
    
    const page = DomHelpers.getElement(`#${section}`);
    const nav = DomHelpers.getElement(`[href="#${section}"]`);
    
    if (page) {
      DomHelpers.addClass(page, 'active');
      if (nav) DomHelpers.addClass(nav, 'active');
      
      if (updateHistory) window.location.hash = section;
      
      this.appStore.setCurrentSection(section);
      this.loadPage(section);
    }
  }

  /**
   * 加载页面内容
   */
  async loadPage(section) {
    this.appStore.setLoading(true);
    
    try {
      switch (section) {
        case 'home':
          await this.loadHomePage();
          break;
        case 'posts':
          await this.loadPosts();
          break;
        case 'categories':
          await this.loadCategories();
          break;
      }
    } catch (error) {
      console.error('Failed to load page:', error);
      this.uiRenderer.renderError(`${section}-content`, '加载失败，请重试');
    } finally {
      this.appStore.setLoading(false);
    }
  }

  /**
   * 加载首页
   */
  async loadHomePage() {
    // 确保显示首页布局，隐藏文章详情页
    this.articleRenderer.hideArticleView();
    this.articleRenderer.showSidebars();
    
    // 重置页面标题
    document.title = 'Blog';
    
    // 加载分类按钮
    await this.loadCategoriesForSidebar();
    
    // 加载文章（默认显示所有文章）
    await this.loadPosts();
    
    // 加载统计数据
    await this.loadStats();
  }

  /**
   * 加载侧边栏分类
   */
  async loadCategoriesForSidebar() {
    const data = await this.apiService.getCategories(1, 100);
    if (data?.list) {
      this.categoryStore.setCategories(data.list);
      this.categoryRenderer.renderCategoryButtons(data.list);
    }
  }

  /**
   * 加载统计信息
   */
  async loadStats() {
    const postsData = await this.apiService.getPublishedPosts(1, 1);
    const categoriesData = await this.apiService.getCategories();
    
    // 更新统计数字
    const totalPostsEl = document.getElementById('total-posts');
    const totalCategoriesEl = document.getElementById('total-categories');
    
    if (totalPostsEl) {
      totalPostsEl.textContent = postsData?.total || 0;
    }
    
    if (totalCategoriesEl) {
      totalCategoriesEl.textContent = categoriesData?.list?.length || 0;
    }
  }

  /**
   * 加载文章页面
   */
  async loadPosts() {
    const postState = this.postStore.getState();
    
    this.uiRenderer.renderLoading('posts-container');
    
    const categoryId = this.categoryStore.getState().currentCategory?.id;
    const data = await this.apiService.getPublishedPosts(
      postState.pagination.currentPage,
      postState.pagination.pageSize,
      categoryId
    );
    
    if (data?.list) {
      this.postStore.setPosts(data.list);
      this.postRenderer.renderPosts(data.list, 'posts-container');
      
      const paginationData = {
        currentPage: data.page_no || 1,
        total: data.total || 0,
        pageSize: data.page_size || 5
      };
      
      this.postStore.setPagination(paginationData);
      this.uiRenderer.renderPagination(
        paginationData,
        'pagination-container',
        (page) => this.goToPage(page)
      );
    } else {
      this.uiRenderer.renderError('posts-container', '无法加载文章列表');
    }
  }

  /**
   * 加载分类页面
   */
  async loadCategories() {
    this.uiRenderer.renderLoading('categories-grid');
    
    const data = await this.apiService.getCategories(1, 50);
    if (data?.list) {
      this.categoryStore.setCategories(data.list);
      this.categoryRenderer.renderCategories(data.list, 'categories-grid');
    } else {
      this.uiRenderer.renderError('categories-grid', '无法加载分类列表');
    }
  }



  /**
   * 翻页
   */
  goToPage(page) {
    const currentPagination = this.postStore.getState().pagination;
    this.postStore.setPagination({ 
      ...currentPagination,
      currentPage: page 
    });
    this.loadPosts();
  }

  /**
   * 查看文章详情（点击卡片时）
   */
  async viewPost(id) {
    // 更新URL并触发路由
    this.navigateToPost(id);
  }

  /**
   * 通过URL查看文章详情
   */
  async viewPostByUrl(id) {
    try {
      // 显示加载状态
      this.appStore.setLoading(true);
      
      // 获取文章详情
      const articleData = await this.apiService.getPost(id);
      
      if (articleData) {
        // 渲染文章详情
        this.articleRenderer.renderArticle(articleData);
        
        // 显示文章详情页
        this.articleRenderer.showArticleView();
        
        // 更新页面标题
        document.title = `${articleData.title} - Blog`;
      } else {
        console.error('Failed to load article:', id);
        this.show404();
      }
    } catch (error) {
      console.error('Error loading article:', error);
      this.show404();
    } finally {
      this.appStore.setLoading(false);
    }
  }

  /**
   * 导航到文章页面
   */
  navigateToPost(id) {
    const url = `/post/${id}`;
    window.history.pushState({ type: 'post', id }, '', url);
    this.handleRoute();
  }

  /**
   * 通过URL过滤分类
   */
  async filterByCategoryByUrl(categoryId) {
    try {
      // 显示首页布局
      this.articleRenderer.hideArticleView();
      this.articleRenderer.showSidebars();
      
      // 加载分类数据
      await this.loadCategoriesForSidebar();
      
      // 过滤文章
      await this.filterByCategory(categoryId);
      
      // 更新页面标题
      const categories = this.categoryStore.getState().categories;
      const category = categories.find(cat => cat.id === categoryId);
      document.title = category ? `${category.name} - Blog` : 'Blog';
    } catch (error) {
      console.error('Error loading category:', error);
      this.show404();
    }
  }

  /**
   * 导航到分类页面
   */
  navigateToCategory(categoryId) {
    const url = categoryId ? `/category/${categoryId}` : '/';
    window.history.pushState({ type: 'category', id: categoryId }, '', url);
    this.handleRoute();
  }

  /**
   * 显示404页面
   */
  show404() {
    this.articleRenderer.hideArticleView();
    this.articleRenderer.showSidebars();
    this.uiRenderer.renderError('posts-container', '页面未找到 (404)');
    document.title = '页面未找到 - Blog';
  }

  /**
   * 返回首页
   */
  backToHome() {
    // 导航到首页URL
    window.history.pushState({}, '', '/');
    this.handleRoute();
  }

  /**
   * 更新分类过滤时的URL
   */
  async filterByCategory(categoryId) {
    // 更新分类状态
    const categories = this.categoryStore.getState().categories;
    const selectedCategory = categories.find(cat => cat.id === categoryId) || null;
    this.categoryStore.setCurrentCategory(selectedCategory);
    
    // 重置分页
    this.postStore.setPagination({ currentPage: 1 });
    
    // 更新分类按钮高亮状态
    this.categoryRenderer.renderCategoryButtons(categories, categoryId);
    
    // 重新加载文章
    this.loadPosts();
    
    // 更新URL（不触发路由，因为我们已经在处理了）
    const url = categoryId ? `/category/${categoryId}` : '/';
    window.history.replaceState({ type: 'category', id: categoryId }, '', url);
    
    // 更新页面标题
    const category = selectedCategory;
    document.title = category ? `${category.name} - Blog` : 'Blog';
  }
}

// 初始化应用
const blogApp = new BlogApp();
